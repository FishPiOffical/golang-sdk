# FishPi Golang SDK - 快速参考

## 安装与导入

```go
import (
    sdk "fishpi-golang-sdk"
    "fishpi-golang-sdk/types"
)
```

## 初始化

```go
// 方式1: 使用 API Key
fishpi := sdk.NewSDK("your-api-key")

// 方式2: 使用配置文件（兼容旧版）
config := sdk.NewFileConfigProvider("config.toml")
client := sdk.NewClient(config)
```

## 常用类型速查

### 用户
- `types.UserInfo` - 用户信息
- `types.TransferRequest` - 转账请求
- `types.CheckinRequest` - 签到请求

### 文章
- `types.ArticleInfo` - 文章信息
- `types.PostArticleRequest` - 发布文章请求
- `types.ArticleType` - 文章类型枚举

### 评论
- `types.CommentInfo` - 评论信息
- `types.PostCommentRequest` - 发布评论请求

### 清风明月
- `types.BreezemoonInfo` - 清风明月信息
- `types.PostBreezemoonRequest` - 发布请求

### 私聊
- `types.ChatMessage` - 私聊消息
- `types.ChatMessageData` - 消息数据

### 聊天室
- `types.ChatroomMessage` - 聊天室消息
- `types.ChatroomMsgData` - 消息数据
- `types.RedPacketInfo` - 红包信息
- `types.OnlineUser` - 在线用户

### 通知
- `types.NotificationInfo` - 通知信息
- `types.UserMessage` - 用户频道消息

## 枚举常量

### 文章类型
```go
types.ArticleTypeNormal     // 0 - 普通帖子
types.ArticleTypeDiscussion // 1 - 讨论区
types.ArticleTypeCity       // 2 - 同城
types.ArticleTypeQnA        // 3 - 问答
```

### 通知类型
```go
types.NotificationTypePoint       // "point"
types.NotificationTypeCommented   // "commented"
types.NotificationTypeReply       // "reply"
types.NotificationTypeAt          // "at"
types.NotificationTypeFollowing   // "following"
types.NotificationTypeBroadcast   // "broadcast"
types.NotificationTypeSysAnnounce // "sys-announce"
```

### 聊天室消息类型
```go
types.ChatroomMsgTypeOnline          // "online"
types.ChatroomMsgTypeMsg             // "msg"
types.ChatroomMsgTypeRevoke          // "revoke"
types.ChatroomMsgTypeRedPacket       // "redPacket"
types.ChatroomMsgTypeRedPacketStatus // "redPacketStatus"
types.ChatroomMsgTypeDiscussChanged  // "discussChanged"
types.ChatroomMsgTypeCustomMessage   // "customMessage"
types.ChatroomMsgTypeBarrager        // "barrager"
```

### 红包类型
```go
types.ChatroomRedPacketTypeRandom           // "random"
types.ChatroomRedPacketTypeAverage          // "average"
types.ChatroomRedPacketTypeSpecify          // "specify"
types.ChatroomRedPacketTypeHeartbeat        // "heartbeat"
types.ChatroomRedPacketTypeRockPaperScissors // "rockPaperScissors"
```

### 猜拳类型
```go
types.GestureTypeRock     // 0 - 石头
types.GestureTypeScissors // 1 - 剪刀
types.GestureTypePaper    // 2 - 布
```

### 客户端类型
```go
types.ClientTypeGolang
types.ClientTypeWeb
types.ClientTypeMobile
types.ClientTypeWindows
types.ClientTypeMacOS
// ... 更多
```

## WebSocket 快速使用

### 聊天室
```go
ws := fishpi.NewChatroomWebSocket("wss://fishpi.cn/chat-room-channel")
ws.OnMessage(func(msg *types.ChatroomMessage) {
    // 处理消息
})
ws.OnError(func(err error) {
    // 处理错误
})
ws.OnClose(func() {
    // 处理关闭
})
ws.Connect()
defer ws.Close()
```

### 私聊
```go
ws := fishpi.NewPrivateChatWebSocket()
ws.OnMessage(func(msg *types.ChatMessage) {
    // 处理消息
})
ws.Connect()
defer ws.Close()

// 发送消息
ws.SendMessage("userName", "Hello!")
```

### 用户通知
```go
ws := fishpi.NewUserNotificationWebSocket()
ws.OnMessage(func(msg *types.UserMessage) {
    switch msg.Type {
    case "article":
        // 文章通知
    case "comment":
        // 评论通知
    case "at":
        // @ 通知
    }
})
ws.Connect()
defer ws.Close()
```

## 常用 API

### 用户操作
```go
// 获取用户信息
userInfo, _ := client.GetApiUser()

// 签到
resp, _ := client.PostUserCheckin()

// 转账
req := &types.TransferRequest{
    UserName: "targetUser",
    Amount:   100,
    Memo:     "备注",
}
client.PostUserTransfer(req)
```

### 文章操作
```go
// 发布文章
req := &types.PostArticleRequest{
    ArticleTitle:       "标题",
    ArticleContent:     "内容",
    ArticleTags:        "标签1,标签2",
    ArticleCommentable: true,
    ArticleType:        types.ArticleTypeNormal,
}
client.PostArticle(req)

// 获取文章列表
list, _ := client.GetArticleList("", "", 1, 20)
```

### 评论操作
```go
// 发布评论
req := &types.PostCommentRequest{
    ArticleId:      "articleId",
    CommentContent: "评论内容",
}
client.PostComment(req)

// 点赞评论
client.PostCommentVoteUp("commentId")
```

### 清风明月
```go
// 发布清风明月
req := &types.PostBreezemoonRequest{
    BreezemoonContent: "今天天气真好",
}
client.PostBreezemoon(req)

// 获取列表
list, _ := client.GetBreezemoonList(1, 20)
```

## 错误处理

所有 API 调用都返回 `(result, error)`，建议：

```go
result, err := client.SomeAPI()
if err != nil {
    log.Printf("API调用失败: %v", err)
    return
}
// 使用 result
```

## 运行示例

```bash
# 设置环境变量
export FISHPI_API_KEY="your-api-key"

# 运行聊天室示例
go run examples/chatroom_ws/main.go

# 运行私聊示例
go run examples/chat_ws/main.go

# 运行通知示例
go run examples/notification_ws/main.go
```

## 项目结构

```
fishpi-golang-sdk/
├── types/          # 类型定义
├── sdk.go          # SDK主文件
├── websocket.go    # WebSocket功能
├── user.go         # 用户API
├── article.go      # 文章API
├── comment.go      # 评论API
├── breezemoon.go   # 清风明月API
├── enum.go         # 枚举定义
└── examples/       # 示例代码
```

## 依赖

- `github.com/imroc/req/v3` - HTTP客户端
- `github.com/lxzan/gws` - WebSocket客户端
- `github.com/pquerna/otp` - TOTP支持

