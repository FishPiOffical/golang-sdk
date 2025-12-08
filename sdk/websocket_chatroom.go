package sdk

import "github.com/FishPiOffical/golang-sdk/types"

// ChatroomWebSocket 聊天室WebSocket客户端
type ChatroomWebSocket struct {
	*Client[types.ChatroomMsg]
}

// NewChatroomWebSocket 创建聊天室WebSocket连接
func (s *FishPiSDK) NewChatroomWebSocket(endpoint string, opts ...ClientOption[types.ChatroomMsg]) *ChatroomWebSocket {
	// 将 endpoint 作为选项传入
	allOpts := append([]ClientOption[types.ChatroomMsg]{
		WithEndpoint[types.ChatroomMsg](endpoint),
	}, opts...)

	client := NewClient(allOpts...)

	return &ChatroomWebSocket{
		Client: client,
	}
}
