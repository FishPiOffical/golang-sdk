package sdk

import (
	"bytes"
	"encoding/base32"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"fishpi-golang-sdk/types"

	"github.com/imroc/req/v3"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

// FishPiSDK SDK客户端
type FishPiSDK struct {
	configProvider ConfigProvider
	client         *req.Client
	logDir         string
	logger         *slog.Logger
}

// Option SDK选项函数
type Option func(sdk *FishPiSDK)

// WithLogDir 设置日志目录，将请求日志保存到指定目录
func WithLogDir(logDir string) Option {
	return func(sdk *FishPiSDK) {
		sdk.logDir = logDir
		// 确保目录存在
		if err := os.MkdirAll(logDir, 0755); err != nil {
			slog.Error("failed to create log directory", slog.String("dir", logDir), slog.Any("error", err))
			return
		}

		sdk.client.OnBeforeRequest(func(client *req.Client, req *req.Request) error {
			// 生成日志文件名: 方法+路径+日期.txt
			method := req.Method
			path := strings.ReplaceAll(req.RawURL, "/", "_")
			if path == "" {
				path = "root"
			}
			date := time.Now().Format("20060102_150405")
			filename := fmt.Sprintf("%s%s_%s.txt", method, path, date)
			logPath := filepath.Join(sdk.logDir, filename)

			// 使用 SetDumpOptions 而不是 SetOutputFile，避免干扰响应读取
			req.EnableDumpToFile(logPath)
			return nil
		})
	}
}

// WithCustomUnmarshaler 设置自定义JSON反序列化方法，会检测字段差异并记录日志
func WithCustomUnmarshaler(logger *slog.Logger) Option {
	return func(sdk *FishPiSDK) {
		if logger == nil {
			logger = slog.Default()
		}
		sdk.logger = logger

		sdk.client.SetJsonUnmarshal(func(data []byte, v interface{}) error {
			// 先使用标准库进行反序列化
			if err := json.Unmarshal(data, v); err != nil {
				return err
			}

			// 检测字段差异
			detectFieldDifferences(data, v, logger)
			return nil
		})
	}
}

// detectFieldDifferences 检测JSON原文和结构体之间的字段差异
func detectFieldDifferences(data []byte, v interface{}, logger *slog.Logger) {
	// 解析原始JSON为map
	var rawData map[string]interface{}
	if err := json.Unmarshal(data, &rawData); err != nil {
		// 如果不是对象类型，跳过检测
		return
	}

	// 获取结构体的所有字段
	structFields := getStructFields(v)

	// 查找原文中存在但结构体中不存在的字段
	for key := range rawData {
		if !containsField(structFields, key) {
			logger.Warn("JSON field not found in struct",
				slog.String("field", key),
				slog.String("type", fmt.Sprintf("%T", v)),
				slog.Any("value", rawData[key]))
		}
	}
}

// getStructFields 获取结构体的所有字段名（包括json tag）
func getStructFields(v interface{}) map[string]bool {
	fields := make(map[string]bool)

	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fields
	}

	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		// 添加字段名
		fields[field.Name] = true

		// 添加json tag名
		if jsonTag := field.Tag.Get("json"); jsonTag != "" {
			tagName := strings.Split(jsonTag, ",")[0]
			if tagName != "" && tagName != "-" {
				fields[tagName] = true
			}
		}

		// 如果是嵌套结构体，递归处理
		if field.Type.Kind() == reflect.Struct || (field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct) {
			fieldVal := val.Field(i)
			if field.Type.Kind() == reflect.Ptr {
				if !fieldVal.IsNil() {
					nestedFields := getStructFields(fieldVal.Interface())
					for k := range nestedFields {
						fields[k] = true
					}
				}
			} else {
				nestedFields := getStructFields(fieldVal.Addr().Interface())
				for k := range nestedFields {
					fields[k] = true
				}
			}
		}
	}

	return fields
}

// containsField 检查字段是否存在于字段集合中
func containsField(fields map[string]bool, key string) bool {
	// 直接匹配
	if fields[key] {
		return true
	}

	// 尝试驼峰和下划线转换
	camelKey := toCamelCase(key)
	if fields[camelKey] {
		return true
	}

	snakeKey := toSnakeCase(key)
	if fields[snakeKey] {
		return true
	}

	return false
}

// toCamelCase 将下划线命名转换为驼峰命名
func toCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}

// toSnakeCase 将驼峰命名转换为下划线命名
func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}

// NewSDK 创建新的SDK实例（使用ConfigProvider）
func NewSDK(configProvider ConfigProvider, options ...Option) *FishPiSDK {
	conf := configProvider.Get()

	if conf.BaseUrl == "" {
		conf.BaseUrl = BaseUrl
	}

	reqClient := req.NewClient().
		SetBaseURL(conf.BaseUrl).
		SetCommonHeader("User-Agent", UserAgent).
		SetCommonQueryParam("apiKey", conf.ApiKey)

	sdk := &FishPiSDK{
		configProvider: configProvider,
		client:         reqClient,
	}

	// 应用选项
	for _, option := range options {
		option(sdk)
	}

	return sdk
}

// NewSDKWithAPIKey 使用API Key快速创建SDK实例（便捷方法）
func NewSDKWithAPIKey(apiKey string, domain ...string) *FishPiSDK {
	d := "fishpi.cn"
	if len(domain) > 0 && domain[0] != "" {
		d = domain[0]
	}

	config := &Config{
		BaseUrl: fmt.Sprintf("https://%s", d),
		ApiKey:  apiKey,
	}

	provider := NewMemoryConfigProvider(config)
	return NewSDK(provider)
}

// GetConfig 获取配置
func (s *FishPiSDK) GetConfig() *Config {
	return s.configProvider.Get()
}

// UpdateConfig 更新配置
func (s *FishPiSDK) UpdateConfig(config *Config) error {
	if err := s.configProvider.Update(config); err != nil {
		return fmt.Errorf("failed to update config: %w", err)
	}

	// 更新 HTTP 客户端配置
	s.client.SetBaseURL(config.BaseUrl).
		SetCommonQueryParam("apiKey", config.ApiKey)

	return nil
}

// GetUserAgent 获取User-Agent
func (s *FishPiSDK) GetUserAgent() string {
	return UserAgent
}

// GetAPIKey 获取当前API Key
func (s *FishPiSDK) GetAPIKey() string {
	return s.configProvider.Get().ApiKey
}

// PostApiGetKey 获取登录key
func (s *FishPiSDK) PostApiGetKey() error {
	conf := s.configProvider.Get()

	if conf.Username == "" || conf.Password == "" {
		return fmt.Errorf("username and password are required")
	}

	request := &types.PostApiGetKeyRequest{
		NameOrEmail:  conf.Username,
		UserPassword: conf.Password,
	}
	if conf.MfaCode != "" {
		request.MfaCode = conf.MfaCode
	}
	if conf.Totp != "" {
		secret := base32.StdEncoding.EncodeToString([]byte(conf.Totp))
		passcode, err := totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
			Period:    30,
			Skew:      1,
			Digits:    otp.DigitsSix,
			Algorithm: otp.AlgorithmSHA512,
		})
		if err != nil {
			return fmt.Errorf("failed to generate TOTP: %w", err)
		}
		request.MfaCode = passcode
	}

	res, err := s.client.R().
		SetBodyJsonMarshal(request).
		Post("/api/getKey")
	if err != nil {
		return fmt.Errorf("failed to get api key: %w", err)
	}

	resp := new(types.PostApiGetKeyResponse)
	if err = res.Unmarshal(resp); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}
	if resp.Code != 0 {
		return fmt.Errorf("get api key failed: %s", resp.Msg)
	}

	// 更新配置中的 ApiKey
	conf.ApiKey = resp.Key
	if err = s.configProvider.Update(conf); err != nil {
		return fmt.Errorf("failed to save api key: %w", err)
	}

	// 更新 HTTP 客户端配置
	s.client.SetCommonQueryParam("apiKey", conf.ApiKey)

	return nil
}

// GetApiUser 获取自己的信息
func (s *FishPiSDK) GetApiUser() (*types.ApiResponse[*types.UserInfo], error) {
	res, err := s.client.R().Get("/api/user")
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}

	response := new(types.ApiResponse[*types.UserInfo])
	if err = res.Unmarshal(&response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response, nil
}

// GetCaptcha 获取验证码
func (s *FishPiSDK) GetCaptcha() (http.Header, *bytes.Buffer, error) {
	body := new(bytes.Buffer)
	res, err := s.client.R().
		SetOutput(body).
		Get("/captcha")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get captcha: %w", err)
	}
	return res.Header, body, nil
}

// Register 用户注册
func (s *FishPiSDK) Register(req *types.PostRegisterRequest) error {
	if req == nil {
		return fmt.Errorf("request is required")
	}
	if req.UserName == "" {
		return fmt.Errorf("userName is required")
	}

	res, err := s.client.R().
		SetBodyJsonMarshal(req).
		Post("/register")
	if err != nil {
		return fmt.Errorf("failed to register: %w", err)
	}
	slog.Info("请求结果", slog.String("res", res.String()))

	return nil
}

// VerifyPhone 验证短信验证码是否正确
func (s *FishPiSDK) VerifyPhone(code string) error {
	if code == "" {
		return fmt.Errorf("code is required")
	}

	res, err := s.client.R().
		SetQueryParam("code", code).
		Get("/verify")
	if err != nil {
		return fmt.Errorf("failed to verify code: %w", err)
	}
	slog.Info("验证结果", slog.String("res", res.String()))

	return nil
}

// PreRegister 预注册
func (s *FishPiSDK) PreRegister(userName, userPhone, inviteCode, captcha string) error {
	if userName == "" {
		return fmt.Errorf("userName is required")
	}

	var resp types.SimpleResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"userName":   userName,
			"userPhone":  userPhone,
			"invitecode": inviteCode,
			"captcha":    captcha,
		}).
		SetSuccessResult(&resp).
		Post("/register/preregister")

	if err != nil {
		return fmt.Errorf("failed to pre-register: %w", err)
	}

	if resp.Code != 0 {
		return fmt.Errorf("pre-register failed: %s", resp.Msg)
	}

	return nil
}
