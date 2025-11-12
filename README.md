# FishPi Golang SDK

FishPi社区的Golang SDK，支持完整的API和WebSocket功能。

## 功能特性

- ✅ 完整的类型定义，按模块分类
- ✅ 用户管理（登录、注册、签到、转账等）
- ✅ 文章管理（发布、更新、列表查询等）
- ✅ 评论管理（发布、更新、投票等）
- ✅ 清风明月管理
- ✅ 私聊功能（HTTP + WebSocket）
- ✅ 聊天室功能（HTTP + WebSocket）
- ✅ 通知系统（HTTP + WebSocket）
- ✅ 红包功能
- ✅ 枚举类型支持

## 安装

```bash
go get github.com/yourusername/fishpi-golang-sdk
```

## 项目结构

```
fishpi-golang-sdk/
├── types/              # 类型定义（按模块分类）
│   ├── common.go      # 通用类型
│   ├── user.go        # 用户相关类型
│   ├── article.go     # 文章相关类型
│   ├── comment.go     # 评论相关类型
│   ├── breezemoon.go  # 清风明月相关类型
│   ├── chat.go        # 私聊相关类型
│   ├── chatroom.go    # 聊天室相关类型
│   └── notice.go      # 通知相关类型
├── sdk.go             # SDK主文件
├── user.go            # 用户相关API
├── article.go         # 文章相关API
├── comment.go         # 评论相关API
├── breezemoon.go      # 清风明月相关API
├── websocket.go       # WebSocket功能
├── enum.go            # 枚举定义
└── consts.go          # 常量定义
```

## 快速开始

### 基本使用

```go
package main

import (
    "fmt"
    sdk "fishpi-golang-sdk"
    "fishpi-golang-sdk/types"
)

func main() {
    // 使用配置文件创建客户端（旧版方式）
    config := sdk.NewFileConfigProvider("config.toml")
    client := sdk.NewClient(config)
    
    // 或者直接使用API Key创建SDK实例（新版方式）
    fishpi := sdk.NewSDK("your-api-key")
    
    // 获取用户信息
    userInfo, err := client.GetApiUser()
    if err != nil {
        panic(err)
    }
    fmt.Printf("用户: %s\n", userInfo.Data.UserName)
}
```

### 文章操作

```go
// 发布文章
req := &types.PostArticleRequest{
    ArticleTitle:       "测试文章",
    ArticleContent:     "这是一篇测试文章",
    ArticleTags:        "测试,Golang",
    ArticleCommentable: true,
    ArticleType:        types.ArticleTypeNormal,
}

resp, err := client.PostArticle(req)
if err != nil {
    panic(err)
}
fmt.Printf("文章ID: %s\n", resp.ArticleId)

// 获取文章列表
articles, err := client.GetArticleList("", "", 1, 20)
if err != nil {
    panic(err)
}
for _, article := range articles.Data.Articles {
    fmt.Printf("%s - %s\n", article.ArticleTitle, article.ArticleAuthorName)
}
```

### 清风明月操作

```go
// 发布清风明月
req := &types.PostBreezemoonRequest{
    BreezemoonContent: "今天天气真好！",
}

err := client.PostBreezemoon(req)
if err != nil {
    panic(err)
}

// 获取清风明月列表
list, err := client.GetBreezemoonList(1, 20)
if err != nil {
    panic(err)
}
```

### WebSocket - 聊天室

```go
package main

import (
    "fmt"
    "log"
    "time"
    
    sdk "fishpi-golang-sdk"
    "fishpi-golang-sdk/types"
)

func main() {
    // 创建SDK实例
    fishpi := sdk.NewSDK("your-api-key")
    
    // 获取聊天室节点
    nodes, err := fishpi.GetChatroomNodes()
    if err != nil {
        panic(err)
    }
    
    // 创建聊天室WebSocket连接
    ws := fishpi.NewChatroomWebSocket(nodes.Available[0].Node)
    
    // 设置消息回调
    ws.OnMessage(func(msg *types.ChatroomMessage) {
        switch msg.Type {
        case types.ChatroomMsgTypeMsg:
            // 处理聊天消息
            data := msg.Data.(map[string]interface{})
            fmt.Printf("[%s] %s\n", data["userName"], data["content"])
            
        case types.ChatroomMsgTypeOnline:
            // 处理在线用户
            fmt.Println("在线用户更新")
            
        case types.ChatroomMsgTypeRedPacket:
            // 处理红包消息
            fmt.Println("收到红包！")
        }
    })
    
    // 设置错误回调
    ws.OnError(func(err error) {
        log.Printf("WebSocket错误: %v\n", err)
    })
    
    // 设置关闭回调
    ws.OnClose(func() {
        log.Println("WebSocket已关闭")
    })
    
    // 连接
    if err := ws.Connect(); err != nil {
        panic(err)
    }
    
    // 发送消息
    time.Sleep(2 * time.Second)
    ws.SendMessage("大家好！")
    
    // 保持连接
    select {}
}
```

### WebSocket - 私聊

```go
package main

import (
    "fmt"
    "log"
    
    sdk "fishpi-golang-sdk"
    "fishpi-golang-sdk/types"
)

func main() {
    fishpi := sdk.NewSDK("your-api-key")
    
    // 创建私聊WebSocket连接
    ws := fishpi.NewPrivateChatWebSocket()
    
    // 设置消息回调
    ws.OnMessage(func(msg *types.ChatMessage) {
        fmt.Printf("收到私聊: [%s] %s\n", 
            msg.Data.FromUserName, 
            msg.Data.Content)
    })
    
    ws.OnError(func(err error) {
        log.Printf("错误: %v\n", err)
    })
    
    // 连接
    if err := ws.Connect(); err != nil {
        panic(err)
    }
    
    // 发送私聊消息
    ws.SendMessage("targetUser", "你好！")
    
    select {}
}
```

### WebSocket - 用户通知

```go
package main

import (
    "fmt"
    "log"
    
    sdk "fishpi-golang-sdk"
    "fishpi-golang-sdk/types"
)

func main() {
    fishpi := sdk.NewSDK("your-api-key")
    
    // 创建用户通知WebSocket连接
    ws := fishpi.NewUserNotificationWebSocket()
    
    // 设置消息回调
    ws.OnMessage(func(msg *types.UserMessage) {
        switch msg.Type {
        case "article":
            // 文章通知
            fmt.Println("收到文章通知")
        case "breezemoon":
            // 清风明月通知
            fmt.Println("收到清风明月通知")
        case "comment":
            // 评论通知
            fmt.Println("收到评论通知")
        }
    })
    
    // 连接
    if err := ws.Connect(); err != nil {
        panic(err)
    }
    
    select {}
}
```

## 枚举类型

SDK使用go-enum生成了完整的枚举类型支持：

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
types.ChatroomMsgTypeDiscussChanged  // 话题变更
types.ChatroomMsgTypeRevoke          // 撤回
types.ChatroomMsgTypeMsg             // 聊天
types.ChatroomMsgTypeRedPacket       // 红包
types.ChatroomMsgTypeRedPacketStatus // 红包领取
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

### 文章类型
```go
types.ArticleTypeNormal     // 普通帖子
types.ArticleTypeDiscussion // 讨论区
types.ArticleTypeCity       // 同城
types.ArticleTypeQnA        // 问答
```

### 客户端类型
```go
types.ClientTypeGolang   // Golang客户端
types.ClientTypeWeb      // 网页端
types.ClientTypeMobile   // 移动端
// ... 更多客户端类型
```

## 高级功能

### 用户签到

```go
resp, err := client.PostUserCheckin()
if err != nil {
    panic(err)
}
fmt.Printf("签到成功，获得%d积分\n", resp.Sum)
```

### 转账

```go
req := &types.TransferRequest{
    UserName: "targetUser",
    Amount:   100,
    Memo:     "转账备注",
}

err := client.PostUserTransfer(req)
if err != nil {
    panic(err)
}
```

### 评论投票

```go
// 点赞评论
err := client.PostCommentVoteUp("commentId")

// 点踩评论
err := client.PostCommentVoteDown("commentId")
```

### 打开红包

```go
req := &types.PostChatroomRedPacketOpenRequest{
    OId: "redPacketId",
    Gesture: types.GestureTypeRock, // 猜拳红包需要
}

resp, err := client.PostChatroomRedPacketOpen(req)
if err != nil {
    panic(err)
}
fmt.Printf("抢到%d积分\n", resp.Got)
```

## 类型系统

所有的结构体都按照模块分类在`types`包中：

- `types.UserInfo` - 用户信息
- `types.ArticleInfo` - 文章信息
- `types.CommentInfo` - 评论信息
- `types.BreezemoonInfo` - 清风明月信息
- `types.ChatMessage` - 私聊消息
- `types.ChatroomMessage` - 聊天室消息
- `types.NotificationInfo` - 通知信息
- 更多...

## 依赖

- `github.com/imroc/req/v3` - HTTP客户端
- `github.com/lxzan/gws` - WebSocket客户端
- `github.com/pquerna/otp` - TOTP支持

## 开发

### 生成枚举代码

```bash
go generate ./...
```

### 构建

```bash
go build ./...
```

### 测试

```bash
go test ./...
```

## License

MIT License

## 贡献

欢迎提交Issue和Pull Request！

## 参考

- [FishPi社区](https://fishpi.cn)
- [TypeScript SDK](https://github.com/imlinhanchao/fishpi-js)

