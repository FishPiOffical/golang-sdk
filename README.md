# FishPi Golang SDK

摸鱼派社区 Golang SDK，提供完整的 API 接口封装。

## 特性

- ✅ **完整的API支持** - 实现OpenAPI定义的76个接口中的76个（100%完成度）
- ✅ **类型安全** - 使用go-enum自动生成枚举类型
- ✅ **灵活配置** - 支持多种ConfigProvider（内存/文件）
- ✅ **WebSocket支持** - 泛型架构，支持聊天室、私聊、用户通知
- ✅ **自动重连** - 可配置重连策略（固定延迟/指数退避）
- ✅ **心跳机制** - 可选的自定义心跳配置
- ✅ **消息解析** - 完整的WebSocket消息解析器
- ✅ **并发安全** - 细粒度锁优化，线程安全的实现
- ✅ **结构化日志** - 使用slog.Logger，支持自定义日志级别
- ✅ **错误处理** - 完整的错误处理和包装

## 目录

- [接口](#接口)
- [安装](#安装)
- [快速开始](#快速开始)
- [WebSocket](#websocket)
- [类型安全](#类型安全)
- [开发](#开发)
- [许可证](#许可证)
- [相关链接](#相关链接)
- [贡献](#贡献)

## 接口
- [ ] 鉴权
    - [x] 获取apiKey
    - [x] 查询用户信息
    - [ ] 注册用户
        - [ ] 预注册账号 POST /register
        - [ ] 验证手机验证码 GET /verify
        - [ ] 注册账号 POST /register2
- [ ] OpenID 接入
    - [ ] 获取授权链接
    - [ ] 签名校验
    - [ ] 获取用户信息
- [ ] 新增
    - [x] 获取用户VIP信息 通用 GET /api/membership/{userId}
    - [x] 获取操作日志 通用 GET /logs/more
    - [x] 修改头像 通用 POST /api/settings/avatar
    - [x] 修改信息 通用 POST /api/settings/profiles
- [ ] 杂项
    - [ ] 勋章链接生成
    - [ ] 客户端版本解析
- [ ] 通用
    - [x] 通过API累计用户的在线时间 WS
    - [x] 查询成员信息
    - [x] 用户名联想
    - [x] 用户常用表情
    - [x] 获取活跃度
    - [x] 获取签到状态
    - [x] 领取昨日活跃奖励
    - [x] 查询在昨日奖励领取状态
    - [x] 举报
    - [x] 查询最近注册的20个用户
    - [x] 转账
    - [x] 关注用户
    - [x] 取关用户
- [x] 通知
    - [x] 通知计数
    - [x] 通知详情
    - [x] 批量已读类型的通知
    - [x] 已读所有消息
- [ ] 聊天室
    - [x] 获取发送弹幕的价格
    - [x] 连接聊天室 WS
    - [x] 聊天室地址API
    - [x] 聊天历史消息
        - [x] 通过聊天消息的oId获取前后消息
    - [x] 发送消息
        - [x] 弹幕
        - [x] 红包
    - [x] 撤回消息
    - [x] 获取消息markdown
    - [x] 打开红包
    - [x] 获取表情包
        - [x] 默认表情包
    - [x] 同步表情包
    - [x] 获取禁言中的成员列表（思过崖）
- [ ] 图床
    - [x] 上传图片
    - [ ] 限制
- [ ] 帖子
    - [x] 发帖
    - [x] 更新帖子
    - [x] 帖子列表
        - [x] 特别注意
        - [x] 最近
        - [x] 按标签
        - [x] 按领域
    - [x] 获取指定帖子
    - [x] 获取指定用户的帖子列表
    - [x] 给文章点赞
    - [x] 感谢文章
    - [x] 获取帖子的评论列表
    - [x] 评论/回复
    - [x] 更新评论
    - [x] 给评论点赞
    - [x] 感谢评论
    - [x] 删除评论
    - [x] 获取帖子当前正在阅读的人数
        - [x] /article-channel WSS
    - [x] 收藏帖子
    - [x] 取消收藏帖子
    - [x] 关注帖子
    - [x] 取消关注帖子
    - [x] 打赏帖子
    - [x] 获取帖子的Markdown原文
- [x] 清风明月
    - [x] 获取清风明月列表
    - [x] 发布清风明月
    - [x] 获取指定用户的清风明月列表
- [ ] 私信
    - [x] 消息通知 /user-channel WSS
    - [x] 获取用户私聊历史消息
    - [x] 标记用户消息已读
    - [x] 获取私聊用户列表以及第一条消息
    - [x] 获取未读消息
    - [x] 撤回私聊消息
- [ ] 敏感操作
    - [ ] 永久注销删除用户
- [ ] 金手指
    - [ ] 注意
    - [x] 上传摸鱼大闯关关卡数据
    - [x] 查询用户最近登录的IP地址
    - [x] 添加勋章
    - [x] 移除勋章
    - [x] 移除勋章（通过userId）
    - [x] 查询用户背包
    - [x] 调整用户背包
    - [x] 调整用户积分
    - [x] 获取用户活跃度
    - [x] 领取指定用户的昨日活跃奖励

## 安装

```bash
go get github.com/FishPiOffical/golang-sdk
```

## 快速开始

### 基础使用

```go
// 使用 API Key 创建 SDK 实例
fishpi := sdk.NewSDKWithAPIKey("your-api-key")

// 获取用户信息
userInfo, err := fishpi.GetUserInfo()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("用户: %s\n", userInfo.Data.UserName)
```

### 配置方式

```go
// 方式1: API Key（推荐）
fishpi := sdk.NewSDKWithAPIKey("your-api-key")

// 方式2: 文件配置
provider := sdk.NewFileConfigProvider("config.json")
fishpi := sdk.NewSDK(provider)

// 方式3: 内存配置
provider := sdk.NewMemoryConfigProvider(&sdk.Config{
    BaseUrl: "https://fishpi.cn",
    ApiKey:  "your-api-key",
})
fishpi := sdk.NewSDK(provider)
```

### SDK 选项

```go
// 启用请求日志
fishpi := sdk.NewSDK(provider, 
    sdk.WithLogDir("./logs"),
)

// 自定义 JSON 解析器（用于调试）
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
fishpi := sdk.NewSDK(provider,
    sdk.WithCustomUnmarshaler(logger),
)
```

### 完整示例

查看 `examples/usage/main.go` 获取完整的使用示例，包括：

- **配置管理**: 使用 YAML 配置文件
- **日志配置**: 使用 devslog 美化日志输出
- **API 调用**: 用户信息、文章、清风明月等功能
- **WebSocket**: 聊天室、私聊、通知的实时通信
- **错误处理**: 完整的错误处理示例

运行示例：

1. 增加本地配置 `_tmp/config.yaml`

```yaml
baseUrl: https://fishpi.cn
userAgent: your-user-agent-here

apiKey: your-api-key-here
username: your-username-here
password: your-password-here
passwordMd5: your-md5-password-here
mfaCode: your-mfa-code-here
totp: your-totp-secret-here

pointGoldFingerKey: "your-point-gold-finger-key-here"
livenessGoldFingerKey: "your-liveness-gold-finger-key-here"
gameGoldFingerKey: "your-game-gold-finger-key-here"
queryGoldFingerKey: "your-query-gold-finger-key-here"
metalGoldFingerKey: "your-metal-gold-finger-key-here"
itemGoldFingerKey: "your-item-gold-finger-key-here"

```

2. 运行代码

```bash
cd examples/usage
go run main.go
```

> **⚠️ 特别说明**
> 
> 1. 示例代码需要配置 `config.yaml` 文件，包含 `apiKey` 和 `baseUrl`
> 2. WebSocket 连接需要有效的 API Key
> 3. 部分功能（如金手指）需要特殊权限
> 4. 建议先在摸鱼派社区获取 API Key：https://fishpi.cn/settings/account

## WebSocket

基于泛型的 WebSocket 客户端，支持自动重连、心跳、自定义日志等企业级特性。

**核心特性**：
- 🔧 **泛型架构** - 类型安全的消息处理
- 🔄 **自动重连** - 默认启用，支持指数退避/固定延迟策略  
- 💓 **心跳机制** - 可配置心跳间隔和自定义消息
- 🔒 **并发安全** - 细粒度锁优化
- ⚙️ **函数式选项** - 灵活配置
- 📝 **结构化日志** - 使用 slog.Logger

**快速开始**：

```go
// 1. 聊天室（自动重连）
ws := fishpi.NewChatroomWebSocket("wss://fishpi.cn/chat-room-channel?apiKey=xxx")
ws.OnMessage(func(msg *types.ChatroomMessage) {
    fmt.Printf("收到消息: %s\n", msg.Type)
})
ws.Connect()
ws.SendMessage("Hello!")

// 2. 私聊
ws := fishpi.NewPrivateChatWebSocket()
ws.OnMessage(func(msg *types.ChatMessage) {
    fmt.Printf("[私聊] %s\n", msg.Data.Content)
})
ws.Connect()

// 3. 用户通知（带心跳）
ws := fishpi.NewUserNotificationWebSocket(
    sdk.WithHeartbeat[types.UserMessage](30*time.Second, func() []byte {
        return []byte(`{"type":"ping"}`)
    }),
)
ws.Connect()
```

**高级配置**：

```go
ws := fishpi.NewChatroomWebSocket("wss://...",
    // 重连策略
    sdk.WithReconnectStrategy[types.ChatroomMessage](&sdk.ExponentialBackoffStrategy{
        BaseDelay:  1 * time.Second,
        MaxDelay:   60 * time.Second,
        Multiplier: 2.0,
    }),
    // 最大重连次数（0=无限）
    sdk.WithMaxReconnectAttempts[types.ChatroomMessage](10),
    // 重连失败回调
    sdk.WithReconnectFailedCallback[types.ChatroomMessage](func(attempts int, err error) {
        log.Printf("重连失败: %v", err)
    }),
    // 自定义日志
    sdk.WithLogger[types.ChatroomMessage](customLogger),
)
```

> 📖 **详细 API 文档**: 查看 [examples/usage/main.go](./examples/usage/main.go) 获取所有 API 的实际使用示例

## 类型安全

SDK 使用 `go-enum` 自动生成所有枚举类型：

```go
// 枚举使用
articleType := types.ArticleListTypeRecent
fmt.Println(articleType.String()) // "recent"

// 解析枚举
parsed, err := types.ParseArticleListType("hot")
```

## 开发

```bash
# 克隆项目
git clone https://github.com/FishPiOffical/golang-sdk

# 安装依赖
go mod download

# 生成枚举代码
go generate ./types/...
```

## 许可证

Apache 2.0

## 相关链接

- [摸鱼派社区](https://fishpi.cn)
- [API文档 v2.1.8](https://fishpi.cn/article/1636516552191)
- [TypeScript SDK 2025.12.17](https://github.com/FishPiOffical/fishpi.js/commit/5e213dd8b5ad4f163152de45eac05f2e5a33db59)

## 贡献

欢迎提交Issue和Pull Request！
