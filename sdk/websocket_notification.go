package sdk

import (
	"fmt"
	"strings"

	"github.com/FishPiOffical/golang-sdk/types"
)

// UserNotificationWebSocket 用户通知WebSocket客户端
type UserNotificationWebSocket struct {
	*Client[types.UserChannelMsg]
}

// NewUserNotificationWebSocket 创建用户通知WebSocket连接
func (s *FishPiSDK) NewUserNotificationWebSocket(opts ...ClientOption[types.UserChannelMsg]) *UserNotificationWebSocket {
	baseUrl := strings.Replace(s.GetConfig().BaseUrl, "http", "ws", 1)
	endpoint := fmt.Sprintf("%s/user-channel?apiKey=%s", baseUrl, s.GetAPIKey())

	// 将 endpoint 作为选项传入
	allOpts := append([]ClientOption[types.UserChannelMsg]{
		WithEndpoint[types.UserChannelMsg](endpoint),
	}, opts...)

	client := NewClient(allOpts...)

	return &UserNotificationWebSocket{
		Client: client,
	}
}
