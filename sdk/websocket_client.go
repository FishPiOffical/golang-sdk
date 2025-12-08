package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/lxzan/gws"
)

// ReconnectStrategy 重连策略接口
type ReconnectStrategy interface {
	NextDelay(attempt int) time.Duration
}

// FixedDelayStrategy 固定延迟重连策略
type FixedDelayStrategy struct {
	Delay time.Duration
}

// NextDelay 返回固定的重连延迟
func (s *FixedDelayStrategy) NextDelay(_ int) time.Duration {
	return s.Delay
}

// ExponentialBackoffStrategy 指数退避重连策略
type ExponentialBackoffStrategy struct {
	BaseDelay  time.Duration // 基础延迟时间
	MaxDelay   time.Duration // 最大延迟时间
	Multiplier float64       // 指数倍率
}

// NextDelay 计算指数退避延迟
func (s *ExponentialBackoffStrategy) NextDelay(attempt int) time.Duration {
	if attempt <= 0 {
		return s.BaseDelay
	}

	delay := float64(s.BaseDelay) * math.Pow(s.Multiplier, float64(attempt-1))
	if delay > float64(s.MaxDelay) {
		return s.MaxDelay
	}
	return time.Duration(delay)
}

// Client 泛型WebSocket客户端
type Client[T any] struct {
	endpoint  string
	userAgent string

	// 连接相关
	connMu    sync.RWMutex
	conn      *gws.Conn
	stateMu   sync.RWMutex
	connected atomic.Bool

	// 上下文控制
	ctx    context.Context
	cancel context.CancelFunc

	// 配置
	logger               *slog.Logger
	autoReconnect        bool
	reconnectStrategy    ReconnectStrategy
	maxReconnectAttempts int
	heartbeatInterval    time.Duration
	heartbeatFunc        func() []byte
	onReconnectFailed    func(attempts int, err error)

	// 回调函数
	callbackMu        sync.RWMutex
	onOpenCallback    func(client *Client[T])
	onMessageCallback func(message *T)
	onErrorCallback   func(err error)
	onCloseCallback   func()

	// 消息解析函数
	messageParser func(data []byte) (*T, error)

	// 重连状态
	reconnectMu       sync.Mutex
	reconnecting      bool
	reconnectAttempts int
}

// ClientOption 客户端选项函数
type ClientOption[T any] func(*Client[T])

// WithContext 设置自定义上下文
func WithContext[T any](ctx context.Context) ClientOption[T] {
	return func(c *Client[T]) {
		if c.cancel != nil {
			c.cancel()
		}
		c.ctx, c.cancel = context.WithCancel(ctx)
	}
}

// WithLogger 设置日志组件
func WithLogger[T any](logger *slog.Logger) ClientOption[T] {
	return func(c *Client[T]) {
		c.logger = logger
	}
}

// WithAutoReconnect 设置是否自动重连
func WithAutoReconnect[T any](enabled bool) ClientOption[T] {
	return func(c *Client[T]) {
		c.autoReconnect = enabled
	}
}

// WithReconnectStrategy 设置重连策略
func WithReconnectStrategy[T any](strategy ReconnectStrategy) ClientOption[T] {
	return func(c *Client[T]) {
		c.reconnectStrategy = strategy
	}
}

// WithMaxReconnectAttempts 设置最大重连次数（0表示无限重连）
func WithMaxReconnectAttempts[T any](max int) ClientOption[T] {
	return func(c *Client[T]) {
		c.maxReconnectAttempts = max
	}
}

// WithHeartbeat 设置心跳配置
func WithHeartbeat[T any](interval time.Duration, messageFunc func() []byte) ClientOption[T] {
	return func(c *Client[T]) {
		c.heartbeatInterval = interval
		c.heartbeatFunc = messageFunc
	}
}

// WithReconnectFailedCallback 设置重连失败回调
func WithReconnectFailedCallback[T any](fn func(attempts int, err error)) ClientOption[T] {
	return func(c *Client[T]) {
		c.onReconnectFailed = fn
	}
}

// WithEndpoint 设置 WebSocket 端点
func WithEndpoint[T any](endpoint string) ClientOption[T] {
	return func(c *Client[T]) {
		c.endpoint = endpoint
	}
}

// WithWSUserAgent 设置 WebSocket User-Agent
func WithWSUserAgent[T any](userAgent string) ClientOption[T] {
	return func(c *Client[T]) {
		c.userAgent = userAgent
	}
}

// WithMessageParser 设置自定义消息解析器
func WithMessageParser[T any](parser func(data []byte) (*T, error)) ClientOption[T] {
	return func(c *Client[T]) {
		c.messageParser = parser
	}
}

// defaultMessageParser 默认消息解析器，使用 json.Unmarshal
func defaultMessageParser[T any](data []byte) (*T, error) {
	var msg T
	if err := json.Unmarshal(data, &msg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal message: %w", err)
	}
	return &msg, nil
}

// NewClient 创建泛型WebSocket客户端（可被其他包调用）
func NewClient[T any](opts ...ClientOption[T]) *Client[T] {
	ctx, cancel := context.WithCancel(context.Background())

	client := &Client[T]{
		ctx:                  ctx,
		cancel:               cancel,
		logger:               slog.Default(),
		userAgent:            UserAgent, // 默认使用 UserAgent 常量
		autoReconnect:        true,      // 默认自动重连
		reconnectStrategy:    &ExponentialBackoffStrategy{BaseDelay: 1 * time.Second, MaxDelay: 60 * time.Second, Multiplier: 2.0},
		maxReconnectAttempts: 0,                       // 默认无限重连
		messageParser:        defaultMessageParser[T], // 默认使用 JSON 解析
	}

	// 应用选项
	for _, opt := range opts {
		opt(client)
	}

	return client
}

// ApplyOptions 应用选项（用于创建后动态配置）
func (c *Client[T]) ApplyOptions(opts ...ClientOption[T]) {
	for _, opt := range opts {
		opt(c)
	}
}

// Connect 连接到WebSocket服务器
func (c *Client[T]) Connect() error {
	c.stateMu.Lock()
	if c.connected.Load() {
		c.stateMu.Unlock()
		return fmt.Errorf("already connected")
	}
	c.stateMu.Unlock()

	err := c.dial()
	if err != nil {
		c.logger.Error("failed to connect", slog.String("endpoint", c.endpoint), slog.Any("error", err))

		// 如果启用自动重连，启动重连循环
		if c.autoReconnect {
			go c.reconnectLoop()
		}
		return err
	}

	c.logger.Info("websocket connected", slog.String("endpoint", c.endpoint))
	return nil
}

// dial 执行实际的连接操作
func (c *Client[T]) dial() error {
	// 直接使用 endpoint
	wsURL := c.endpoint

	// 创建WebSocket连接
	header := http.Header{}
	header.Set("User-Agent", c.userAgent)

	socket, _, err := gws.NewClient(c.createHandler(), &gws.ClientOption{
		Addr:          wsURL,
		RequestHeader: header,
		PermessageDeflate: gws.PermessageDeflate{
			Enabled:               true,
			ServerContextTakeover: true,
			ClientContextTakeover: true,
		},
	})

	if err != nil {
		return err
	}

	c.connMu.Lock()
	c.conn = socket
	c.connMu.Unlock()

	c.connected.Store(true)

	// 启动读取循环
	go func() {
		socket.ReadLoop()
	}()

	// 启动心跳（如果配置了）
	if c.heartbeatInterval > 0 && c.heartbeatFunc != nil {
		go c.startHeartbeat()
	}

	return nil
}

// createHandler 创建WebSocket事件处理器
func (c *Client[T]) createHandler() gws.Event {
	return &genericWebSocketHandler[T]{client: c}
}

// reconnectLoop 重连循环
func (c *Client[T]) reconnectLoop() {
	c.reconnectMu.Lock()
	if c.reconnecting {
		c.reconnectMu.Unlock()
		return
	}
	c.reconnecting = true
	c.reconnectAttempts = 0
	c.reconnectMu.Unlock()

	defer func() {
		c.reconnectMu.Lock()
		c.reconnecting = false
		c.reconnectMu.Unlock()
	}()

	for {
		select {
		case <-c.ctx.Done():
			c.logger.Debug("reconnect loop stopped: context cancelled")
			return
		default:
		}

		c.reconnectMu.Lock()
		c.reconnectAttempts++
		attempt := c.reconnectAttempts
		c.reconnectMu.Unlock()

		// 检查是否达到最大重连次数
		if c.maxReconnectAttempts > 0 && attempt > c.maxReconnectAttempts {
			c.logger.Warn("max reconnect attempts reached",
				slog.Int("attempts", attempt),
				slog.Int("max", c.maxReconnectAttempts))

			if c.onReconnectFailed != nil {
				c.onReconnectFailed(attempt-1, fmt.Errorf("max reconnect attempts reached"))
			}
			return
		}

		// 计算延迟
		delay := c.reconnectStrategy.NextDelay(attempt)
		c.logger.Info("attempting to reconnect",
			slog.Int("attempt", attempt),
			slog.Duration("delay", delay))

		// 等待延迟
		select {
		case <-c.ctx.Done():
			return
		case <-time.After(delay):
		}

		// 尝试重连
		err := c.dial()
		if err == nil {
			c.logger.Info("reconnected successfully", slog.Int("after_attempts", attempt))
			c.reconnectMu.Lock()
			c.reconnectAttempts = 0
			c.reconnectMu.Unlock()
			return
		}

		c.logger.Warn("reconnect failed",
			slog.Int("attempt", attempt),
			slog.Any("error", err))
	}
}

// Close 关闭连接
func (c *Client[T]) Close() error {
	c.stateMu.Lock()
	defer c.stateMu.Unlock()

	if !c.connected.Load() {
		return nil
	}

	// 取消上下文（停止心跳和重连）
	c.cancel()

	c.connMu.Lock()
	conn := c.conn
	c.connMu.Unlock()

	if conn != nil {
		err := conn.WriteClose(1000, []byte("client closing"))
		if err != nil {
			c.logger.Warn("error writing close frame", slog.Any("error", err))
		}
	}

	c.connected.Store(false)
	c.logger.Debug("websocket closed")

	return nil
}

// IsConnected 检查是否已连接
func (c *Client[T]) IsConnected() bool {
	return c.connected.Load()
}

// SendRaw 发送原始字节数据
func (c *Client[T]) SendRaw(data []byte) error {
	c.connMu.RLock()
	conn := c.conn
	connected := c.connected.Load()
	c.connMu.RUnlock()

	if !connected || conn == nil {
		return fmt.Errorf("not connected")
	}

	return conn.WriteMessage(gws.OpcodeText, data)
}

// SendJSON 发送JSON数据
func (c *Client[T]) SendJSON(v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("failed to marshal json: %w", err)
	}
	return c.SendRaw(data)
}

// OnOpen 设置连接打开回调
func (c *Client[T]) OnOpen(callback func(client *Client[T])) {
	c.callbackMu.Lock()
	defer c.callbackMu.Unlock()
	c.onOpenCallback = callback
}

// OnMessage 设置消息回调
func (c *Client[T]) OnMessage(callback func(message *T)) {
	c.callbackMu.Lock()
	defer c.callbackMu.Unlock()
	c.onMessageCallback = callback
}

// OnError 设置错误回调
func (c *Client[T]) OnError(callback func(err error)) {
	c.callbackMu.Lock()
	defer c.callbackMu.Unlock()
	c.onErrorCallback = callback
}

// OnClose 设置关闭回调
func (c *Client[T]) OnClose(callback func()) {
	c.callbackMu.Lock()
	defer c.callbackMu.Unlock()
	c.onCloseCallback = callback
}

// startHeartbeat 启动心跳
func (c *Client[T]) startHeartbeat() {
	ticker := time.NewTicker(c.heartbeatInterval)
	defer ticker.Stop()

	c.logger.Debug("heartbeat started", slog.Duration("interval", c.heartbeatInterval))

	for {
		select {
		case <-c.ctx.Done():
			c.logger.Debug("heartbeat stopped: context cancelled")
			return
		case <-ticker.C:
			if !c.connected.Load() {
				continue
			}

			msg := c.heartbeatFunc()
			if msg == nil {
				continue
			}

			c.connMu.RLock()
			conn := c.conn
			c.connMu.RUnlock()

			if conn != nil {
				err := conn.WriteMessage(gws.OpcodeText, msg)
				if err != nil {
					c.logger.Warn("failed to send heartbeat", slog.Any("error", err))
				} else {
					c.logger.Debug("heartbeat sent")
				}
			}
		}
	}
}

// genericWebSocketHandler 泛型WebSocket事件处理器
type genericWebSocketHandler[T any] struct {
	client *Client[T]
}

func (h *genericWebSocketHandler[T]) OnOpen(_ *gws.Conn) {
	h.client.logger.Debug("websocket connection opened")

	h.client.callbackMu.RLock()
	callback := h.client.onOpenCallback
	h.client.callbackMu.RUnlock()

	if callback != nil {
		callback(h.client)
	}
}

func (h *genericWebSocketHandler[T]) OnClose(_ *gws.Conn, err error) {
	h.client.connected.Store(false)

	if err != nil {
		h.client.logger.Info("websocket connection closed", slog.Any("error", err))
	} else {
		h.client.logger.Debug("websocket connection closed")
	}

	h.client.callbackMu.RLock()
	callback := h.client.onCloseCallback
	h.client.callbackMu.RUnlock()

	if callback != nil {
		callback()
	}

	// 如果启用自动重连且不是主动关闭
	if h.client.autoReconnect && h.client.ctx.Err() == nil {
		go h.client.reconnectLoop()
	}
}

func (h *genericWebSocketHandler[T]) OnPing(socket *gws.Conn, payload []byte) {
	_ = socket.WritePong(payload)
}

func (h *genericWebSocketHandler[T]) OnPong(_ *gws.Conn, _ []byte) {
	h.client.logger.Debug("received pong")
}

func (h *genericWebSocketHandler[T]) OnMessage(_ *gws.Conn, message *gws.Message) {
	defer func() {
		if err := message.Close(); err != nil {
			h.client.logger.Debug("failed to close message", slog.Any("error", err))
		}
	}()

	data := message.Bytes()

	// 使用消息解析器解析消息
	msg, err := h.client.messageParser(data)
	if err != nil {
		h.client.logger.Error("failed to parse message", slog.Any("error", err))

		h.client.callbackMu.RLock()
		errorCallback := h.client.onErrorCallback
		h.client.callbackMu.RUnlock()

		if errorCallback != nil {
			errorCallback(fmt.Errorf("failed to parse message: %w", err))
		}
		return
	}

	h.client.callbackMu.RLock()
	messageCallback := h.client.onMessageCallback
	h.client.callbackMu.RUnlock()

	if messageCallback != nil {
		messageCallback(msg)
	}
}
