package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"fishpi-golang-sdk/types"

	"github.com/lxzan/gws"
)

// WebSocketClient WebSocket客户端接口
type WebSocketClient interface {
	Connect() error
	Close() error
	IsConnected() bool
	SendMessage(data interface{}) error
	OnMessage(callback func(message []byte))
	OnError(callback func(err error))
	OnClose(callback func())
}

// ChatroomWebSocket 聊天室WebSocket连接
type ChatroomWebSocket struct {
	sdk       *FishPiSDK
	conn      *gws.Conn
	endpoint  string
	mu        sync.RWMutex
	connected bool
	ctx       context.Context
	cancel    context.CancelFunc

	// 回调函数
	onMessageCallback func(message *types.ChatroomMessage)
	onErrorCallback   func(err error)
	onCloseCallback   func()
}

// chatroomWebSocketHandler WebSocket事件处理器
type chatroomWebSocketHandler struct {
	ws *ChatroomWebSocket
}

func (h *chatroomWebSocketHandler) OnOpen(socket *gws.Conn) {
	log.Println("Chatroom WebSocket connected")
}

func (h *chatroomWebSocketHandler) OnClose(socket *gws.Conn, err error) {
	h.ws.mu.Lock()
	h.ws.connected = false
	h.ws.mu.Unlock()

	log.Printf("Chatroom WebSocket closed: %v", err)

	if h.ws.onCloseCallback != nil {
		h.ws.onCloseCallback()
	}
}

func (h *chatroomWebSocketHandler) OnPing(socket *gws.Conn, payload []byte) {
	_ = socket.WritePong(payload)
}

func (h *chatroomWebSocketHandler) OnPong(socket *gws.Conn, payload []byte) {}

func (h *chatroomWebSocketHandler) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()

	data := message.Bytes()

	// 解析聊天室消息
	var msg types.ChatroomMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		if h.ws.onErrorCallback != nil {
			h.ws.onErrorCallback(fmt.Errorf("failed to unmarshal chatroom message: %w", err))
		}
		return
	}

	if h.ws.onMessageCallback != nil {
		h.ws.onMessageCallback(&msg)
	}
}

// NewChatroomWebSocket 创建聊天室WebSocket连接
func (s *FishPiSDK) NewChatroomWebSocket(endpoint string) *ChatroomWebSocket {
	ctx, cancel := context.WithCancel(context.Background())

	return &ChatroomWebSocket{
		sdk:      s,
		endpoint: endpoint,
		ctx:      ctx,
		cancel:   cancel,
	}
}

// Connect 连接到聊天室
func (ws *ChatroomWebSocket) Connect() error {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	if ws.connected {
		return fmt.Errorf("already connected")
	}

	// 构建WebSocket URL
	wsURL := fmt.Sprintf("%s?apiKey=%s", ws.endpoint, ws.sdk.GetAPIKey())

	// 创建WebSocket连接
	header := http.Header{}
	header.Set("User-Agent", ws.sdk.GetUserAgent())

	socket, _, err := gws.NewClient(&chatroomWebSocketHandler{ws: ws}, &gws.ClientOption{
		Addr:          wsURL,
		RequestHeader: header,
		PermessageDeflate: gws.PermessageDeflate{
			Enabled:               true,
			ServerContextTakeover: true,
			ClientContextTakeover: true,
		},
	})

	if err != nil {
		return fmt.Errorf("failed to connect to chatroom: %w", err)
	}

	ws.conn = socket
	ws.connected = true

	// 启动读取循环
	go func() {
		socket.ReadLoop()
	}()

	return nil
}

// Close 关闭连接
func (ws *ChatroomWebSocket) Close() error {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	if !ws.connected || ws.conn == nil {
		return nil
	}

	ws.cancel()
	err := ws.conn.WriteClose(1000, []byte("client closing"))
	if err != nil {
		return err
	}

	ws.connected = false
	return nil
}

// IsConnected 检查是否已连接
func (ws *ChatroomWebSocket) IsConnected() bool {
	ws.mu.RLock()
	defer ws.mu.RUnlock()
	return ws.connected
}

// SendMessage 发送消息
func (ws *ChatroomWebSocket) SendMessage(content string) error {
	ws.mu.RLock()
	defer ws.mu.RUnlock()

	if !ws.connected || ws.conn == nil {
		return fmt.Errorf("not connected")
	}

	data, err := json.Marshal(map[string]interface{}{
		"content": content,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	return ws.conn.WriteMessage(gws.OpcodeText, data)
}

// OnMessage 设置消息回调
func (ws *ChatroomWebSocket) OnMessage(callback func(message *types.ChatroomMessage)) {
	ws.onMessageCallback = callback
}

// OnError 设置错误回调
func (ws *ChatroomWebSocket) OnError(callback func(err error)) {
	ws.onErrorCallback = callback
}

// OnClose 设置关闭回调
func (ws *ChatroomWebSocket) OnClose(callback func()) {
	ws.onCloseCallback = callback
}

// GetParser 获取消息解析器
func (ws *ChatroomWebSocket) GetParser() *MessageParser {
	return NewMessageParser()
}

// SendRaw 发送原始数据
func (ws *ChatroomWebSocket) SendRaw(data []byte) error {
	ws.mu.RLock()
	defer ws.mu.RUnlock()

	if !ws.connected || ws.conn == nil {
		return fmt.Errorf("not connected")
	}

	return ws.conn.WriteMessage(gws.OpcodeText, data)
}

// PrivateChatWebSocket 私聊WebSocket连接
type PrivateChatWebSocket struct {
	sdk       *FishPiSDK
	conn      *gws.Conn
	endpoint  string
	mu        sync.RWMutex
	connected bool
	ctx       context.Context
	cancel    context.CancelFunc

	// 回调函数
	onMessageCallback func(message *types.ChatMessage)
	onErrorCallback   func(err error)
	onCloseCallback   func()
}

// privateChatWebSocketHandler WebSocket事件处理器
type privateChatWebSocketHandler struct {
	ws *PrivateChatWebSocket
}

func (h *privateChatWebSocketHandler) OnOpen(socket *gws.Conn) {
	log.Println("Private chat WebSocket connected")
}

func (h *privateChatWebSocketHandler) OnClose(socket *gws.Conn, err error) {
	h.ws.mu.Lock()
	h.ws.connected = false
	h.ws.mu.Unlock()

	log.Printf("Private chat WebSocket closed: %v", err)

	if h.ws.onCloseCallback != nil {
		h.ws.onCloseCallback()
	}
}

func (h *privateChatWebSocketHandler) OnPing(socket *gws.Conn, payload []byte) {
	_ = socket.WritePong(payload)
}

func (h *privateChatWebSocketHandler) OnPong(socket *gws.Conn, payload []byte) {}

func (h *privateChatWebSocketHandler) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()

	data := message.Bytes()

	// 解析私聊消息
	var msg types.ChatMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		if h.ws.onErrorCallback != nil {
			h.ws.onErrorCallback(fmt.Errorf("failed to unmarshal chat message: %w", err))
		}
		return
	}

	if h.ws.onMessageCallback != nil {
		h.ws.onMessageCallback(&msg)
	}
}

// NewPrivateChatWebSocket 创建私聊WebSocket连接
func (s *FishPiSDK) NewPrivateChatWebSocket() *PrivateChatWebSocket {
	ctx, cancel := context.WithCancel(context.Background())

	endpoint := fmt.Sprintf("wss://%s/chat-channel", s.GetConfig().BaseUrl)

	return &PrivateChatWebSocket{
		sdk:      s,
		endpoint: endpoint,
		ctx:      ctx,
		cancel:   cancel,
	}
}

// Connect 连接到私聊
func (ws *PrivateChatWebSocket) Connect() error {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	if ws.connected {
		return fmt.Errorf("already connected")
	}

	// 构建WebSocket URL
	wsURL := fmt.Sprintf("%s?apiKey=%s", ws.endpoint, ws.sdk.GetAPIKey())

	// 创建WebSocket连接
	header := http.Header{}
	header.Set("User-Agent", ws.sdk.GetUserAgent())

	socket, _, err := gws.NewClient(&privateChatWebSocketHandler{ws: ws}, &gws.ClientOption{
		Addr:          wsURL,
		RequestHeader: header,
		PermessageDeflate: gws.PermessageDeflate{
			Enabled:               true,
			ServerContextTakeover: true,
			ClientContextTakeover: true,
		},
	})

	if err != nil {
		return fmt.Errorf("failed to connect to private chat: %w", err)
	}

	ws.conn = socket
	ws.connected = true

	// 启动读取循环
	go func() {
		socket.ReadLoop()
	}()

	return nil
}

// Close 关闭连接
func (ws *PrivateChatWebSocket) Close() error {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	if !ws.connected || ws.conn == nil {
		return nil
	}

	ws.cancel()
	err := ws.conn.WriteClose(1000, []byte("client closing"))
	if err != nil {
		return err
	}

	ws.connected = false
	return nil
}

// IsConnected 检查是否已连接
func (ws *PrivateChatWebSocket) IsConnected() bool {
	ws.mu.RLock()
	defer ws.mu.RUnlock()
	return ws.connected
}

// SendMessage 发送消息
func (ws *PrivateChatWebSocket) SendMessage(toUserName, content string) error {
	ws.mu.RLock()
	defer ws.mu.RUnlock()

	if !ws.connected || ws.conn == nil {
		return fmt.Errorf("not connected")
	}

	data, err := json.Marshal(map[string]interface{}{
		"toUserName": toUserName,
		"content":    content,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	return ws.conn.WriteMessage(gws.OpcodeText, data)
}

// OnMessage 设置消息回调
func (ws *PrivateChatWebSocket) OnMessage(callback func(message *types.ChatMessage)) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.onMessageCallback = callback
}

// OnError 设置错误回调
func (ws *PrivateChatWebSocket) OnError(callback func(err error)) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.onErrorCallback = callback
}

// OnClose 设置关闭回调
func (ws *PrivateChatWebSocket) OnClose(callback func()) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.onCloseCallback = callback
}

// UserNotificationWebSocket 用户通知WebSocket连接
type UserNotificationWebSocket struct {
	sdk       *FishPiSDK
	conn      *gws.Conn
	endpoint  string
	mu        sync.RWMutex
	connected bool
	ctx       context.Context
	cancel    context.CancelFunc

	// 回调函数
	onMessageCallback func(message *types.UserMessage)
	onErrorCallback   func(err error)
	onCloseCallback   func()
}

// userNotificationWebSocketHandler WebSocket事件处理器
type userNotificationWebSocketHandler struct {
	ws *UserNotificationWebSocket
}

func (h *userNotificationWebSocketHandler) OnOpen(socket *gws.Conn) {
	log.Println("User notification WebSocket connected")
}

func (h *userNotificationWebSocketHandler) OnClose(socket *gws.Conn, err error) {
	h.ws.mu.Lock()
	h.ws.connected = false
	h.ws.mu.Unlock()

	log.Printf("User notification WebSocket closed: %v", err)

	if h.ws.onCloseCallback != nil {
		h.ws.onCloseCallback()
	}
}

func (h *userNotificationWebSocketHandler) OnPing(socket *gws.Conn, payload []byte) {
	_ = socket.WritePong(payload)
}

func (h *userNotificationWebSocketHandler) OnPong(socket *gws.Conn, payload []byte) {}

func (h *userNotificationWebSocketHandler) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()

	data := message.Bytes()

	// 解析用户通知消息
	var msg types.UserMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		if h.ws.onErrorCallback != nil {
			h.ws.onErrorCallback(fmt.Errorf("failed to unmarshal user message: %w", err))
		}
		return
	}

	if h.ws.onMessageCallback != nil {
		h.ws.onMessageCallback(&msg)
	}
}

// NewUserNotificationWebSocket 创建用户通知WebSocket连接
func (s *FishPiSDK) NewUserNotificationWebSocket() *UserNotificationWebSocket {
	ctx, cancel := context.WithCancel(context.Background())

	endpoint := fmt.Sprintf("wss://%s/user-channel", s.GetConfig().BaseUrl)

	return &UserNotificationWebSocket{
		sdk:      s,
		endpoint: endpoint,
		ctx:      ctx,
		cancel:   cancel,
	}
}

// Connect 连接到用户通知频道
func (ws *UserNotificationWebSocket) Connect() error {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	if ws.connected {
		return fmt.Errorf("already connected")
	}

	// 构建WebSocket URL
	wsURL := fmt.Sprintf("%s?apiKey=%s", ws.endpoint, ws.sdk.GetAPIKey())

	// 创建WebSocket连接
	header := http.Header{}
	header.Set("User-Agent", ws.sdk.GetUserAgent())

	socket, _, err := gws.NewClient(&userNotificationWebSocketHandler{ws: ws}, &gws.ClientOption{
		Addr:          wsURL,
		RequestHeader: header,
		PermessageDeflate: gws.PermessageDeflate{
			Enabled:               true,
			ServerContextTakeover: true,
			ClientContextTakeover: true,
		},
	})

	if err != nil {
		return fmt.Errorf("failed to connect to user notification: %w", err)
	}

	ws.conn = socket
	ws.connected = true

	// 启动读取循环
	go func() {
		socket.ReadLoop()
	}()

	// 启动心跳
	go ws.heartbeat()

	return nil
}

// heartbeat 发送心跳
func (ws *UserNotificationWebSocket) heartbeat() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ws.ctx.Done():
			return
		case <-ticker.C:
			ws.mu.RLock()
			if ws.connected && ws.conn != nil {
				_ = ws.conn.WritePing([]byte("heartbeat"))
			}
			ws.mu.RUnlock()
		}
	}
}

// Close 关闭连接
func (ws *UserNotificationWebSocket) Close() error {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	if !ws.connected || ws.conn == nil {
		return nil
	}

	ws.cancel()
	err := ws.conn.WriteClose(1000, []byte("client closing"))
	if err != nil {
		return err
	}

	ws.connected = false
	return nil
}

// IsConnected 检查是否已连接
func (ws *UserNotificationWebSocket) IsConnected() bool {
	ws.mu.RLock()
	defer ws.mu.RUnlock()
	return ws.connected
}

// OnMessage 设置消息回调
func (ws *UserNotificationWebSocket) OnMessage(callback func(message *types.UserMessage)) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.onMessageCallback = callback
}

// OnError 设置错误回调
func (ws *UserNotificationWebSocket) OnError(callback func(err error)) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.onErrorCallback = callback
}

// OnClose 设置关闭回调
func (ws *UserNotificationWebSocket) OnClose(callback func()) {
	ws.mu.Lock()
	defer ws.mu.Unlock()
	ws.onCloseCallback = callback
}
