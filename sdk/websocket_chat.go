package sdk

import (
	"fmt"
	"strings"

	"github.com/FishPiOffical/golang-sdk/types"
)

// PrivateChatWebSocket 私聊WebSocket客户端
type PrivateChatWebSocket struct {
	*Client[types.ChatChannelMsg]
}

// NewPrivateChatWebSocket 创建私聊WebSocket连接
func (s *FishPiSDK) NewPrivateChatWebSocket(toUserId string, opts ...ClientOption[types.ChatChannelMsg]) *PrivateChatWebSocket {
	baseUrl := strings.Replace(s.GetConfig().BaseUrl, "http", "ws", 1)

	endpoint := fmt.Sprintf("%s/chat-channel?apiKey=%s&toUser=%s", baseUrl, s.GetAPIKey(), toUserId)

	// 将 endpoint 作为选项传入
	allOpts := append([]ClientOption[types.ChatChannelMsg]{
		WithEndpoint[types.ChatChannelMsg](endpoint),
	}, opts...)

	client := NewClient(allOpts...)

	return &PrivateChatWebSocket{
		Client: client,
	}
}

// SendMessage 发送私聊消息
func (ws *PrivateChatWebSocket) SendMessage(content string) error {
	return ws.SendRaw([]byte(content))
}
