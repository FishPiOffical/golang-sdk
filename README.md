# FishPi Golang SDK

[![Go Version](https://img.shields.io/badge/Go-%3E%3D1.20-blue)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

FishPi 社区的 Golang SDK，提供完整的 API 和 WebSocket 功能支持。

## ✨ 功能特性

- ✅ **完整的类型系统** - 所有结构体按模块分类，类型安全
- ✅ **枚举类型支持** - 使用 go-enum 生成的完整枚举类型
- ✅ **用户管理** - 登录、注册、签到、转账等
- ✅ **文章管理** - 发布、更新、查询、投票、感谢等
- ✅ **评论管理** - 发布、更新、投票、感谢、删除等
- ✅ **清风明月** - 发布、更新、查询、删除等
- ✅ **私聊功能** - HTTP API + WebSocket 实时通信
- ✅ **聊天室功能** - HTTP API + WebSocket 实时通信
- ✅ **通知系统** - HTTP API + WebSocket 实时推送
- ✅ **红包功能** - 查询、打开、领取等
- ✅ **金手指API** - 游戏数据、勋章管理等

## 📦 安装

```bash
go get github.com/yourusername/fishpi-golang-sdk
```

## 🏗️ 项目结构

```
fishpi-golang-sdk/
├── sdk/                # SDK实现
│   ├── sdk.go         # SDK主文件
│   ├── article.go     # 文章API
│   ├── breezemoon.go  # 清风明月API
│   ├── chat.go        # 私聊API
│   ├── chatroom.go    # 聊天室API
│   ├── comment.go     # 评论API
│   ├── user.go        # 用户API
│   ├── notice.go      # 通知API
│   ├── finger.go      # 金手指API
│   └── websocket.go   # WebSocket功能
├── types/              # 类型定义
│   ├── common.go      # 通用类型
│   ├── user.go        # 用户类型
│   ├── article.go     # 文章类型
│   ├── comment.go     # 评论类型
│   ├── breezemoon.go  # 清风明月类型
│   ├── chat.go        # 私聊类型
│   ├── chatroom.go    # 聊天室类型
│   ├── notice.go      # 通知类型
│   ├── finger.go      # 金手指类型
│   └── enum.go        # 枚举定义
└── examples/           # 示例代码
    ├── chatroom_ws/   # 聊天室WebSocket示例
    ├── chat_ws/       # 私聊WebSocket示例
    └── notification_ws/ # 通知WebSocket示例
```

## 🚀 快速开始

### 基本使用

```go
package main

import (
    "fmt"
    "fishpi-golang-sdk/sdk"
    "fishpi-golang-sdk/types"
)

func main() {
    // 创建SDK实例
    client := sdk.NewSDK("your-api-key")
    
    // 获取用户信息
    userInfo, err := client.GetApiUser()
    if err != nil {
        panic(err)
    }
    fmt.Printf("用户: %s\n", userInfo.Data.UserName)
}
```

### 用户操作

```go
// 签到
resp, err := client.PostUserCheckin()
if err != nil {
    panic(err)
}
fmt.Printf("签到成功，获得 %d 积分\n", resp.Sum)

// 转账
err = client.PostUserTransfer(&types.TransferRequest{
    UserName: "targetUser",
    Amount:   100,
    Memo:     "转账备注",
})

// 领取昨日活跃度奖励
reward, err := client.RewardLiveness()

// 获取常用表情
emotions, err := client.GetUserEmotions()
```

### 文章操作

```go
// 发布文章
articleId, err := client.PostArticle(&types.PostArticleRequest{
    ArticleTitle:       "测试文章",
    ArticleContent:     "这是一篇测试文章的内容",
    ArticleTags:        "测试,Golang",
    ArticleCommentable: true,
    ArticleType:        types.ArticleTypeNormal,
})

// 更新文章
err = client.UpdateArticle(articleId, &types.UpdateArticleRequest{
    ArticleTitle:   "更新后的标题",
    ArticleContent: "更新后的内容",
    ArticleTags:    "测试,Golang,更新",
})

// 获取文章列表
articles, err := client.GetArticleList(types.ArticleListType, "", 1, 20)

// 获取文章详情
detail, err := client.GetArticleDetail(articleId, 1)

// 点赞文章
voteType, err := client.VoteArticle(articleId, "up")

// 感谢文章
err = client.ThankArticle(articleId)

// 关注/取消关注文章
err = client.WatchArticle(articleId, true)
```

### 评论操作

```go
// 发布评论
err := client.PostComment(&types.PostCommentRequest{
    ArticleId:      articleId,
    CommentContent: "这是一条评论",
})

// 更新评论
err = client.UpdateComment(commentId, &types.UpdateCommentRequest{
    CommentContent: "更新后的评论内容",
})

// 点赞评论
voteType, err := client.VoteComment(commentId, "up")

// 感谢评论
err = client.ThankComment(commentId)

// 删除评论
err = client.RemoveComment(commentId)
```

### 清风明月操作

```go
// 发布清风明月
err := client.PostBreezemoon(&types.PostBreezemoonRequest{
    BreezemoonContent: "今天天气真好！",
})

// 更新清风明月
err = client.UpdateBreezemoon(breezemoonId, &types.UpdateBreezemoonRequest{
    BreezemoonContent: "更新后的内容",
})

// 获取清风明月列表
list, err := client.GetBreezemoonList(1, 20)

// 获取用户清风明月列表
userList, err := client.GetUserBreezemoons("userName", 1, 20)

// 删除清风明月
err = client.RemoveBreezemoon(breezemoonId)
```

### 私聊操作

```go
// 获取私聊列表
chatList, err := client.GetChatList()

// 获取与指定用户的私聊消息
messages, err := client.GetChatMessages("userName", 1, 20)

// 发送私聊消息
err = client.SendChatMessage("userName", "你好！")

// 标记消息已读
err = client.MarkChatRead("userName")

// 获取未读消息
unread, err := client.GetChatUnread()
```

### 聊天室操作

```go
// 发送聊天室消息
err := client.SendChatroomMessage("大家好！")

// 获取聊天室历史消息
history, err := client.GetChatroomHistory(1, types.ChatContentTypeHTML)

// 获取指定消息上下文
context, err := client.GetChatroomMessage(
    messageId, 
    types.ChatMessageTypeContext, 
    25, 
    types.ChatContentTypeHTML,
)

// 撤回消息
err = client.RevokeChatroomMessage(messageId)

// 打开红包
result, err := client.OpenRedPacket(redPacketId, nil)

// 打开猜拳红包
gesture := types.GestureTypeRock
result, err := client.OpenRedPacket(redPacketId, &gesture)
```

### 通知操作

```go
// 获取未读通知数量
count, err := client.GetNotificationCount()

// 获取通知列表
notifications, err := client.GetNotifications(types.NotificationTypeAt)

// 标记通知已读
err = client.MarkNotificationRead(types.NotificationTypeAt)

// 标记所有通知已读
err = client.MarkAllNotificationsRead()
```

### 金手指API

```go
// 创建金手指实例
finger := client.NewFinger("gold-finger-key")

// 上传摸鱼大闯关分数
err := finger.AddMofishScore("userName", "stage1", time.Now().UnixMilli())

// 查询用户最近登录IP
ip, err := finger.QueryLatestLoginIP("userName")

// 添加勋章
err = finger.AddMetal("userName", &types.MetalBase{
    Name:        "测试勋章",
    Attr:        []string{"attr1", "attr2"},
    Description: "这是一个测试勋章",
})

// 移除勋章
err = finger.RemoveMetal("userName", "测试勋章")
```

## 🔌 WebSocket 功能

### 聊天室 WebSocket

```go
package main

import (
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"
    
    "fishpi-golang-sdk/sdk"
    "fishpi-golang-sdk/types"
)

func main() {
    client := sdk.NewSDK(os.Getenv("FISHPI_API_KEY"))
    
    // 创建聊天室WebSocket连接
    ws := client.NewChatroomWebSocket("wss://fishpi.cn/chat-room-channel")
    
    // 设置消息回调
    ws.OnMessage(func(msg *types.ChatroomMessage) {
        switch msg.Type {
        case "msg":
            data := msg.Data.(map[string]interface{})
            fmt.Printf("[聊天] %s: %s\n", data["userName"], data["content"])
            
        case "online":
            fmt.Println("[系统] 在线用户更新")
            
        case "redPacket":
            fmt.Println("[红包] 收到红包！")
            
        case "discussChanged":
            data := msg.Data.(map[string]interface{})
            fmt.Printf("[系统] 话题变更: %s\n", data["newDiscuss"])
        }
    })
    
    // 设置错误回调
    ws.OnError(func(err error) {
        log.Printf("错误: %v\n", err)
    })
    
    // 连接
    if err := ws.Connect(); err != nil {
        panic(err)
    }
    defer ws.Close()
    
    // 发送消息
    ws.SendMessage("大家好！")
    
    // 保持连接
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    <-sigChan
}
```

### 私聊 WebSocket

```go
func main() {
    client := sdk.NewSDK(os.Getenv("FISHPI_API_KEY"))
    
    // 创建私聊WebSocket连接
    ws := client.NewPrivateChatWebSocket()
    
    // 设置消息回调
    ws.OnMessage(func(msg *types.ChatMessage) {
        if msg.Type == "msg" {
            fmt.Printf("[私聊] %s: %s\n", 
                msg.Data.SenderUserName, 
                msg.Data.Content)
        }
    })
    
    // 连接
    if err := ws.Connect(); err != nil {
        panic(err)
    }
    defer ws.Close()
    
    // 发送消息
    ws.SendMessage("targetUser", "你好！")
    
    // 保持连接
    select {}
}
```

### 用户通知 WebSocket

```go
func main() {
    client := sdk.NewSDK(os.Getenv("FISHPI_API_KEY"))
    
    // 创建通知WebSocket连接
    ws := client.NewUserNotificationWebSocket()
    
    // 设置消息回调
    ws.OnMessage(func(msg *types.UserMessage) {
        switch msg.Type {
        case "article":
            fmt.Println("[通知] 收到文章通知")
        case "comment":
            fmt.Println("[通知] 收到评论通知")
        case "at":
            fmt.Println("[通知] 有人@了你")
        case "following":
            fmt.Println("[通知] 关注的用户有新动态")
        }
    })
    
    // 连接
    if err := ws.Connect(); err != nil {
        panic(err)
    }
    defer ws.Close()
    
    // 保持连接
    select {}
}
```

## 📚 枚举类型

### 文章列表类型
```go
types.ArticleListType      // 最新
types.ArticleListTypeHot   // 热门
types.ArticleListTypeGood  // 精华
types.ArticleListTypePerfect // 精选
types.ArticleListTypeReply // 回复
```

### 文章类型
```go
types.ArticleTypeNormal     // 普通帖子
types.ArticleTypeDiscussion // 讨论区
types.ArticleTypeCity       // 同城
types.ArticleTypeQnA        // 问答
```

### 通知类型
```go
types.NotificationTypePoint       // 积分
types.NotificationTypeCommented   // 收到的回帖
types.NotificationTypeReply       // 收到的回复
types.NotificationTypeAt          // 提及我的
types.NotificationTypeFollowing   // 我关注的
types.NotificationTypeBroadcast   // 同城
types.NotificationTypeSysAnnounce // 系统
```

### 聊天室消息类型
```go
types.ChatroomMsgTypeOnline          // 在线
types.ChatroomMsgTypeMsg             // 聊天
types.ChatroomMsgTypeRevoke          // 撤回
types.ChatroomMsgTypeRedPacket       // 红包
types.ChatroomMsgTypeRedPacketStatus // 红包领取
types.ChatroomMsgTypeDiscussChanged  // 话题变更
types.ChatroomMsgTypeCustomMessage   // 进入离开聊天室消息
types.ChatroomMsgTypeBarrager        // 弹幕
```

### 红包类型
```go
types.ChatroomRedPacketTypeRandom           // 拼手气红包
types.ChatroomRedPacketTypeAverage          // 平分红包
types.ChatroomRedPacketTypeSpecify          // 专属红包
types.ChatroomRedPacketTypeHeartbeat        // 心跳红包
types.ChatroomRedPacketTypeRockPaperScissors // 猜拳红包
```

### 猜拳类型
```go
types.GestureTypeRock     // 石头
types.GestureTypeScissors // 剪刀
types.GestureTypePaper    // 布
```

### 投票类型
```go
types.VoteType(-1) // 未投票
types.VoteType(0)  // 点赞
types.VoteType(1)  // 点踩
```

### 聊天内容类型
```go
types.ChatContentTypeMd   // Markdown
types.ChatContentTypeHtml // HTML
```

### 聊天消息查询类型
```go
types.ChatMessageTypeContext // 上下文
types.ChatMessageTypeBefore  // 之前
types.ChatMessageTypeAfter   // 之后
```

### 客户端类型
```go
types.ClientTypeGolang   // Golang客户端
types.ClientTypeWeb      // 网页端
types.ClientTypeMobile   // 移动端
types.ClientTypeWindows  // Windows客户端
types.ClientTypeMacOS    // macOS客户端
// ... 更多客户端类型
```

## 🔧 高级功能

### 自定义域名

```go
// 使用自定义域名
client := sdk.NewSDK("api-key", "custom.fishpi.cn")
```

### 旧版Client（配置文件方式）

```go
// 使用配置文件
config := sdk.NewFileConfigProvider("config.toml")
client := sdk.NewClient(config)

// 获取API Key
err := client.PostApiGetKey()
```

## 📖 完整API列表

### 用户API
- `GetApiUser()` - 获取用户信息
- `PostUserCheckin()` - 用户签到
- `GetUserLiveness()` - 获取活跃度
- `IsCheckIn()` - 检查是否已签到
- `IsCollectedLiveness()` - 检查是否已领取昨日活跃奖励
- `RewardLiveness()` - 领取昨日活跃奖励
- `PostUserTransfer()` - 转账
- `GetUserEmotions()` - 获取常用表情

### 文章API
- `PostArticle()` - 发布文章
- `UpdateArticle()` - 更新文章
- `GetArticleList()` - 获取文章列表
- `GetUserArticles()` - 获取用户文章列表
- `GetArticleDetail()` - 获取文章详情
- `VoteArticle()` - 文章投票
- `ThankArticle()` - 感谢文章
- `WatchArticle()` - 关注文章

### 评论API
- `PostComment()` - 发布评论
- `UpdateComment()` - 更新评论
- `VoteComment()` - 评论投票
- `ThankComment()` - 感谢评论
- `RemoveComment()` - 删除评论

### 清风明月API
- `PostBreezemoon()` - 发布清风明月
- `UpdateBreezemoon()` - 更新清风明月
- `GetBreezemoonList()` - 获取清风明月列表
- `GetUserBreezemoons()` - 获取用户清风明月列表
- `RemoveBreezemoon()` - 删除清风明月

### 私聊API
- `GetChatList()` - 获取私聊列表
- `GetChatMessages()` - 获取私聊消息
- `SendChatMessage()` - 发送私聊消息
- `MarkChatRead()` - 标记消息已读
- `GetChatUnread()` - 获取未读消息

### 聊天室API
- `SendChatroomMessage()` - 发送聊天室消息
- `GetChatroomHistory()` - 获取聊天室历史消息
- `GetChatroomMessage()` - 获取指定消息上下文
- `RevokeChatroomMessage()` - 撤回消息
- `OpenRedPacket()` - 打开红包

### 通知API
- `GetNotificationCount()` - 获取未读通知数量
- `GetNotifications()` - 获取通知列表
- `MarkNotificationRead()` - 标记通知已读
- `MarkAllNotificationsRead()` - 标记所有通知已读

### 金手指API
- `NewFinger()` - 创建金手指实例
- `AddMofishScore()` - 上传摸鱼大闯关分数
- `QueryLatestLoginIP()` - 查询用户最近登录IP
- `AddMetal()` - 添加勋章
- `RemoveMetal()` - 移除勋章

### WebSocket API
- `NewChatroomWebSocket()` - 创建聊天室WebSocket连接
- `NewPrivateChatWebSocket()` - 创建私聊WebSocket连接
- `NewUserNotificationWebSocket()` - 创建通知WebSocket连接

## 🛠️ 开发

### 生成枚举代码

```bash
cd types
go generate
```

### 构建

```bash
go build ./...
```

### 运行示例

```bash
# 设置API Key
export FISHPI_API_KEY="your-api-key"

# 运行聊天室示例
go run examples/chatroom_ws/main.go

# 运行私聊示例
go run examples/chat_ws/main.go

# 运行通知示例
go run examples/notification_ws/main.go
```

## 📦 依赖

- `github.com/imroc/req/v3` - HTTP客户端
- `github.com/lxzan/gws` - WebSocket客户端
- `github.com/pquerna/otp` - TOTP支持

## 📄 License

MIT License

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📞 联系方式

- FishPi 社区: https://fishpi.cn
- TypeScript SDK: https://github.com/imlinhanchao/fishpi-js

## 🎯 功能对比

相比 TypeScript SDK，本 Golang SDK 提供了：

- ✅ 完全对等的类型系统
- ✅ 完全对等的API接口
- ✅ 完全对等的WebSocket功能
- ✅ 更强的类型安全性
- ✅ 更好的性能表现
- ✅ 完整的枚举类型支持
- ✅ 清晰的模块划分

## 📊 项目统计

- **类型文件**: 11 个
- **SDK模块**: 10 个
- **示例程序**: 3 个
- **支持的API**: 50+ 个
- **枚举类型**: 10+ 种
- **总代码行数**: 3000+ 行

