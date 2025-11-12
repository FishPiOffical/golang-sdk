package sdk

import (
	"bytes"
	"encoding/base32"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"fishpi-golang-sdk/types"
	"github.com/imroc/req/v3"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

// FishPiSDK SDK客户端（新版本）
type FishPiSDK struct {
	apiKey string
	domain string
	client *req.Client
}

// NewSDK 创建新的SDK实例
func NewSDK(apiKey string, domain ...string) *FishPiSDK {
	d := "fishpi.cn"
	if len(domain) > 0 && domain[0] != "" {
		d = domain[0]
	}

	baseURL := fmt.Sprintf("https://%s", d)

	client := req.NewClient().
		SetBaseURL(baseURL).
		SetCommonHeader("User-Agent", UserAgent).
		SetCommonQueryParam("apiKey", apiKey)

	return &FishPiSDK{
		apiKey: apiKey,
		domain: d,
		client: client,
	}
}

// getUserAgent 获取User-Agent
func (s *FishPiSDK) getUserAgent() string {
	return UserAgent
}

// Client 旧版客户端（保持兼容性）
type Client struct {
	configProvider ConfigProvider
	client         *req.Client
}

type Option func(client *Client)

func NewClient(configProvider ConfigProvider, options ...Option) *Client {
	conf := configProvider.Get()

	if conf.BaseUrl == "" {
		conf.BaseUrl = BaseUrl
	}

	reqClient := req.NewClient().
		SetBaseURL(conf.BaseUrl).
		SetCommonHeader("User-Agent", UserAgent).
		SetCommonQueryParam("apiKey", conf.ApiKey)

	client := &Client{
		configProvider: configProvider,
		client:         reqClient,
	}

	for _, option := range options {
		option(client)
	}

	return client
}

// GetConfig 获取配置
func (c *Client) GetConfig() *Config {
	return c.configProvider.Get()
}

// UpdateConfig 更新配置
func (c *Client) UpdateConfig(config *Config) error {
	if err := c.configProvider.Update(config); err != nil {
		return err
	}

	// 更新 HTTP 客户端配置
	c.client.SetBaseURL(config.BaseUrl).
		SetCommonQueryParam("apiKey", config.ApiKey)

	return nil
}

// PostApiGetKey 获取登录key
func (c *Client) PostApiGetKey() error {
	conf := c.configProvider.Get()

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
			return err
		}
		request.MfaCode = passcode
	}

	res, err := c.client.R().
		SetBodyJsonMarshal(request).
		Post("/api/getKey")
	if err != nil {
		return err
	}

	resp := new(types.PostApiGetKeyResponse)
	if err = res.Unmarshal(resp); err != nil {
		return err
	}
	if resp.Code != 0 {
		return errors.New(resp.Msg)
	}

	// 更新配置中的 ApiKey
	conf.ApiKey = resp.Key
	if err = c.configProvider.Update(conf); err != nil {
		return err
	}

	// 更新 HTTP 客户端配置
	c.client.SetCommonQueryParam("apiKey", conf.ApiKey)

	return nil
}

// GetApiUser 获取自己的信息
func (c *Client) GetApiUser() (*types.ApiResponse[*types.UserInfo], error) {
	res, err := c.client.R().Get("/api/user")
	if err != nil {
		return nil, err
	}

	response := new(types.ApiResponse[*types.UserInfo])
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetCaptcha 获取验证码
func (c *Client) GetCaptcha() (http.Header, *bytes.Buffer, error) {
	body := new(bytes.Buffer)
	res, err := c.client.R().
		SetOutput(body).
		Get("/captcha")
	if err != nil {
		return nil, nil, err
	}
	return res.Header, body, nil
}

type PostRegisterRequest struct {
	UserName   string `json:"userName"`
	UserPhone  string `json:"userPhone"`
	InviteCode string `json:"invitecode"`
	Captcha    string `json:"captcha"`
}

// PostRegister 用户注册
func (c *Client) PostRegister(req *PostRegisterRequest) error {
	res, err := c.client.R().
		SetBodyJsonMarshal(req).
		Post("/register")
	if err != nil {
		return err
	}
	slog.Info("请求结果", slog.String("res", res.String()))

	return nil
}

// GetVerify 验证短信验证码是否正确
func (c *Client) GetVerify(code string) error {
	res, err := c.client.R().
		SetQueryParam("code", code).
		Get("/verify")
	if err != nil {
		return err
	}
	slog.Info("请求结果", slog.String("res", res.String()))

	return nil
}
