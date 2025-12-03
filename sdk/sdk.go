package sdk

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/FishPiOffical/golang-sdk/config"
	"github.com/FishPiOffical/golang-sdk/types"

	"github.com/imroc/req/v3"
)

// FishPiSDK SDK客户端
type FishPiSDK struct {
	configProvider config.Provider
	client         *req.Client
	logDir         string
	logger         *slog.Logger
}

// NewSDK 创建新的SDK实例（使用ConfigProvider）
func NewSDK(configProvider config.Provider, options ...Option) *FishPiSDK {
	conf := configProvider.Get()

	if conf.BaseUrl == "" {
		conf.BaseUrl = BaseUrl
	}

	if conf.UserAgent == "" {
		conf.UserAgent = UserAgent
	}

	reqClient := req.NewClient().
		SetBaseURL(conf.BaseUrl).
		SetCommonHeader("User-Agent", conf.UserAgent).
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

	conf := &config.Config{
		BaseUrl: fmt.Sprintf("https://%s", d),
		ApiKey:  apiKey,
	}

	provider := config.NewMemoryConfigProvider(conf)
	return NewSDK(provider)
}

// GetConfig 获取配置
func (s *FishPiSDK) GetConfig() *config.Config {
	return s.configProvider.Get()
}

// UpdateConfig 更新配置
func (s *FishPiSDK) UpdateConfig(config *config.Config) error {
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
	return s.client.Headers.Get("User-Agent")
}

// GetAPIKey 获取当前API Key
func (s *FishPiSDK) GetAPIKey() string {
	return s.configProvider.Get().ApiKey
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
