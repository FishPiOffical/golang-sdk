package fishPiSdk

import (
	"bytes"
	"encoding/base32"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/imroc/req/v3"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type Client struct {
	configProvider ConfigProvider

	client *req.Client
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

	request := &PostApiGetKeyRequest{
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

	resp := new(PostApiGetKeyResponse)
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
func (c *Client) GetApiUser() (*ApiResponse[*ApiUserGetData], error) {
	res, err := c.client.R().Get("/api/user")
	if err != nil {
		return nil, err
	}

	response := new(ApiResponse[*ApiUserGetData])
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

type PostRegister2Request struct {
	UserAppRole  string `json:"userAppRole"`
	UserPassword string `json:"userPassword"`
	UserId       string `json:"userId"`
	R            string `json:"r"`
}

// PostRegister2 设置密码等基础信息
func (c *Client) PostRegister2(req *PostRegister2Request) error {
	res, err := c.client.R().
		SetBodyJsonMarshal(req).
		SetQueryParam("r", req.R).
		Post("/register2")
	if err != nil {
		return err
	}
	slog.Info("请求结果", slog.String("res", res.String()))
	return nil
}

// GetChatRevoke 撤回私聊消息
func (c *Client) GetChatRevoke(oId string) (*ResponseResult[any], error) {
	res, err := c.client.R().
		SetQueryParam("oId", oId).
		Get("/chat/revoke")
	if err != nil {
		return nil, err
	}
	response := new(ResponseResult[any])
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetChatMarkAsRead 标记私聊消息为已读
func (c *Client) GetChatMarkAsRead(fromUser string) error {
	res, err := c.client.R().
		SetQueryParam("fromUser", fromUser).
		Get("/chat/mark-as-read")
	if err != nil {
		return err
	}
	response := new(ResponseResult[any])
	if err = res.Unmarshal(&response); err != nil {
		return err
	}
	if response.Result != 0 {
		return errors.New(response.Msg)
	}
	return nil
}

// GetChatGetMessage 获取私聊消息
func (c *Client) GetChatGetMessage(toUser string, page int, pageSize int) (*ApiResponse[[]*GetChatGetMessageData], error) {
	res, err := c.client.R().
		SetQueryParamsAnyType(map[string]interface{}{
			"toUser":   toUser,
			"page":     page,
			"pageSize": pageSize,
		}).
		Get("/chat/get-message")
	if err != nil {
		return nil, err
	}
	response := new(ApiResponse[[]*GetChatGetMessageData])
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetChatHasUnread 获取私聊未读消息
func (c *Client) GetChatHasUnread() (*ResponseResult[[]*GetChatHasUnreadData], error) {
	res, err := c.client.R().
		Get("/chat/has-unread")
	if err != nil {
		return nil, err
	}
	response := new(ResponseResult[[]*GetChatHasUnreadData])
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetChatGetList 获取私聊聊天列表
func (c *Client) GetChatGetList() (*ResponseResult[[]*GetChatGetListData], error) {
	res, err := c.client.R().
		Get("/chat/get-list")
	if err != nil {
		return nil, err
	}
	response := new(ResponseResult[[]*GetChatGetListData])
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetUserChannelLink 获取用户私聊链接
func (c *Client) GetUserChannelLink() string {
	link := c.client.BaseURL
	link = strings.ReplaceAll(link, "http", "ws")
	link += fmt.Sprintf("/user-channel?apiKey=%s", c.configProvider.Get().ApiKey)
	return link
}

// GetChatChannelLink 获取用户私聊链接
func (c *Client) GetChatChannelLink(toUser string) string {
	link := c.client.BaseURL
	link = strings.ReplaceAll(link, "http", "ws")
	link += fmt.Sprintf("/chat-channel?apiKey=%s&toUser=%s", c.configProvider.Get().ApiKey, toUser)
	return link
}

// GetApiGetNotifications 获取通知列表 todo 还有更多参数需要补全
func (c *Client) GetApiGetNotifications(typ NotificationType, page int) ([]*GetApiGetNotificationsData, error) {
	res, err := c.client.R().
		SetQueryParamsAnyType(map[string]interface{}{
			"type": typ,
			"p":    page,
		}).
		Get("/api/getNotifications")
	if err != nil {
		return nil, err
	}
	response := new(ApiResponse[[]*GetApiGetNotificationsData])
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}
	if response.Code != 0 {
		return nil, errors.New(response.Msg)
	}
	return response.Data, nil
}

// GetChatroomNodeGet 获取节点列表
func (c *Client) GetChatroomNodeGet() (*GetChatroomNodeGetResponse, error) {
	res, err := c.client.R().
		Get("/chat-room/node/get")
	if err != nil {
		return nil, err
	}
	response := new(GetChatroomNodeGetResponse)
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}
	if response.Code != 0 {
		return nil, errors.New(response.Msg)
	}
	for _, node := range response.Avaliable {
		node.Node += fmt.Sprintf("?apiKey=%s", c.configProvider.Get().ApiKey)
	}
	return response, nil
}

// PostChatroomRedPacketOpen 打开聊天室红包
func (c *Client) PostChatroomRedPacketOpen(req *PostChatroomRedPacketOpenRequest) (*PostChatroomRedPacketOpenResponse, error) {
	res, err := c.client.R().
		SetBodyJsonMarshal(req).
		Post("/chat-room/red-packet/open")
	if err != nil {
		return nil, err
	}
	response := new(PostChatroomRedPacketOpenResponse)
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}
	return response, nil
}
