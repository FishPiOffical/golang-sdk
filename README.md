# FishPi Golang SDK

摸鱼派社区 Golang SDK，提供完整的 API 接口封装。

## 特性

- ✅ **完整的API支持** - 实现OpenAPI定义的76个接口中的70个（92.1%完成度）
- ✅ **类型安全** - 使用go-enum自动生成枚举类型
- ✅ **灵活配置** - 支持多种ConfigProvider（内存/文件）
- ✅ **WebSocket支持** - 聊天室、私聊、用户通知三种WebSocket
- ✅ **消息解析** - 完整的WebSocket消息解析器
- ✅ **并发安全** - 线程安全的实现
- ✅ **错误处理** - 完整的错误处理和包装

## 安装

```bash
go get github.com/fghwett/fishpi-golang-sdk
```

## 快速开始

### 基础使用

```go
package main

import (
    "fmt"
    "log"
    
    "fishpi-golang-sdk/sdk"
)

func main() {
    // 方式1: 使用API Key快速创建
    fishpi := sdk.NewSDKWithAPIKey("your-api-key")
    
    // 获取用户信息
    userInfo, err := fishpi.GetUserInfo()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("用户: %s\n", userInfo.Data.UserName)
}
```

### 使用ConfigProvider

```go
// 方式1: 文件配置（推荐）
provider := sdk.NewFileConfigProvider("config.json")
fishpi := sdk.NewSDK(provider)

// 方式2: 内存配置
config := &sdk.Config{
    BaseUrl: "https://fishpi.cn",
    ApiKey:  "your-api-key",
}
provider := sdk.NewMemoryConfigProvider(config)
fishpi := sdk.NewSDK(provider)
```

### 使用SDK选项

SDK支持多种选项来自定义行为：

```go
import (
    "log/slog"
    "os"
    
    "fishpi-golang-sdk/sdk"
)

// 创建自定义logger
logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
    Level: slog.LevelWarn,
}))

// 创建配置提供者
provider := sdk.NewFileConfigProvider("config.json")

// 使用选项创建SDK实例
fishpi := sdk.NewSDK(
    provider,
    sdk.WithLogDir("./logs"),              // 选项1: 设置日志目录
    sdk.WithCustomUnmarshaler(logger),     // 选项2: 设置自定义JSON反序列化器
)
```

#### 选项说明

**WithLogDir(logDir string)** - 设置日志目录
- 将所有HTTP请求/响应保存到指定目录
- 日志文件名格式: `{方法}{路径}_{日期时间}.txt`
- 示例: `GET_api_user_20060102_150405.txt`
- 自动创建目录（如果不存在）

**WithCustomUnmarshaler(logger *slog.Logger)** - 设置自定义JSON反序列化器
- 使用标准的 `json.Unmarshal` 进行反序列化
- 自动检测JSON原文和Go结构体之间的字段差异
- 当JSON中存在字段但结构体中不存在时，输出警告日志
- 支持嵌套结构体检测
- 支持驼峰和下划线命名转换
- 如果logger为nil，使用默认logger

## 功能模块

### 认证模块

```go
// 登录获取API Key
apiKey, err := fishpi.Login("username", "password")

// 注册用户
err := fishpi.Register(&types.PostRegisterRequest{
    UserName:   "newuser",
    UserPhone:  "13800138000",
    InviteCode: "invite123",
    Captcha:    "captcha",
})

// 预注册
err := fishpi.PreRegister("username", "phone", "inviteCode", "captcha")

// 验证手机号
err := fishpi.VerifyPhone("123456")
```

### 用户模块

```go
// 获取用户信息
userInfo, err := fishpi.GetUserInfo()

// 通过用户名获取用户信息
user, err := fishpi.GetUserByUsername("alice")

// 签到
resp, err := fishpi.PostUserCheckin()

// 领取活跃度奖励
sum, err := fishpi.RewardLiveness()

// 转账
err := fishpi.PostPointTransfer(&types.TransferRequest{
    UserName: "alice",
    Amount:   100,
    Memo:     "Thanks",
})

// 关注用户
err := fishpi.FollowUser("userId")

// 举报用户
err := fishpi.ReportUser("userId", "spam", "spam content")

// 上传文件
resp, err := fishpi.UploadFile(fileBytes, "filename.jpg")

// 管理员功能
userNames, err := fishpi.GetUserNames()
recentUsers, err := fishpi.GetRecentRegister()
logs, err := fishpi.GetLogs(1)
```

### 文章模块

```go
// 发布文章
articleId, err := fishpi.PostArticle(&types.PostArticleRequest{
    ArticleTitle:   "标题",
    ArticleContent: "内容",
    ArticleTags:    "tag1,tag2",
})

// 更新文章
err := fishpi.UpdateArticle("articleId", &types.UpdateArticleRequest{
    ArticleTitle:   "新标题",
    ArticleContent: "新内容",
})

// 获取文章列表
articles, err := fishpi.GetArticleList(types.ArticleListTypeRecent, "", 1, 20)

// 获取文章详情
detail, err := fishpi.GetArticleDetail("articleId", 1)

// 点赞/点踩文章
voteType, err := fishpi.VoteArticle("articleId", "up")

// 感谢文章
err := fishpi.ThankArticle("articleId")

// 收藏文章
err := fishpi.FollowArticle("articleId")

// 关注文章
err := fishpi.WatchArticle("articleId", true)

// 打赏文章
err := fishpi.RewardArticle("articleId")

// 删除文章
err := fishpi.DeleteArticle("articleId")

// 置顶文章（管理员）
err := fishpi.StickArticle("articleId")

// 获取文章热度
heat, err := fishpi.GetArticleHeat("articleId")
```

### 评论模块

```go
// 发布评论
err := fishpi.PostComment(&types.PostCommentRequest{
    ArticleId:      "articleId",
    CommentContent: "评论内容",
})

// 更新评论
err := fishpi.UpdateComment("commentId", &types.UpdateCommentRequest{
    CommentContent: "新内容",
})

// 删除评论
err := fishpi.RemoveComment("commentId")

// 点赞/点踩评论
voteType, err := fishpi.VoteComment("commentId", "up")

// 感谢评论
err := fishpi.ThankComment("commentId")

// 获取文章评论列表
comments, err := fishpi.GetArticleComments("articleId", 1, 20)
```

### 清风明月模块

```go
// 获取清风明月列表
list, err := fishpi.GetBreezemoonList(1, 20)

// 获取用户清风明月
userList, err := fishpi.GetUserBreezemoons("username", 1, 20)

// 发布清风明月
err := fishpi.PostBreezemoon("内容")

// 更新清风明月
err := fishpi.UpdateBreezemoon("id", "新内容")

// 删除清风明月
err := fishpi.RemoveBreezemoon("id")
```

### 通知模块

```go
// 获取未读通知数
count, err := fishpi.GetNotificationCount()

// 获取通知列表
notifications, err := fishpi.GetNotifications(1, 20)

// 标记通知已读
err := fishpi.MarkNotificationRead("notificationId")

// 全部标记已读
err := fishpi.MarkAllNotificationsRead()
```

### 聊天室模块

```go
// 发送聊天室消息
err := fishpi.SendChatroomMessage("Hello, World!")

// 获取聊天室历史
messages, err := fishpi.GetChatroomHistory(1, types.ChatContentTypeHTML)

// 获取指定消息上下文
context, err := fishpi.GetChatroomMessage("oId", types.ChatMessageTypeContext, 25, types.ChatContentTypeHTML)

// 撤回消息
err := fishpi.RevokeChatroomMessage("oId")

// 发送弹幕
err := fishpi.SendChatroomBarrage("弹幕内容")

// 获取弹幕费用
cost, err := fishpi.GetBarrageCost()

// 获取原始消息
html, err := fishpi.GetMessageRaw("oId")

// 获取禁言列表
mutes, err := fishpi.GetChatroomMutes()

// 获取聊天室节点
node, err := fishpi.GetChatroomNode()

// 发送红包
err := fishpi.SendRedPacket(&types.SendRedPacketRequest{
    Type:  "random",
    Money: 100,
    Count: 10,
    Msg:   "红包祝福",
})

// 打开红包
resp, err := fishpi.OpenRedPacket("oId", &gesture)

// 获取红包详情
detail, err := fishpi.GetRedPacketDetail("oId")
```

### 私聊模块

```go
// 获取私聊消息
messages, err := fishpi.GetChatMessages("username", 1, 20)

// 发送私聊消息
err := fishpi.SendChatMessage("username", "消息内容")

// 标记私聊已读
err := fishpi.MarkChatRead("username")

// 获取私聊列表
list, err := fishpi.GetChatList()

// 检查是否有未读消息
hasUnread, err := fishpi.GetChatUnread()
```

### 表情模块

```go
// 获取用户常用表情
emotions, err := fishpi.GetUserEmotions()

// 获取云端表情包
cloudEmojis, err := fishpi.GetCloudEmojis()

// 同步云端表情包
err := fishpi.SyncCloudEmojis([]string{"emoji1", "emoji2"})
```

### 金手指模块（需要金手指密钥）

```go
finger := fishpi.NewFinger("goldFingerKey")

// 添加摸鱼分数
err := finger.AddMofishScore("username", "stage1", time.Now().Unix())

// 查询最近登录IP
ip, err := finger.QueryLatestLoginIP("username")

// 添加勋章
err := finger.AddMetal("username", &types.MetalBase{
    Name:        "勋章名",
    Description: "描述",
})

// 移除勋章
err := finger.RemoveMetal("username", "勋章名")
err := finger.RemoveMetalByUserId("userId", "勋章名")

// 查询用户背包
bag, err := finger.QueryUserBag("username")

// 调整用户背包
err := finger.EditUserBag("username", "item", 10)

// 调整用户积分
err := finger.EditUserPoints("username", 100, "奖励")

// 查询用户活跃度
liveness, err := finger.GetUserLiveness("username")
```

### WebSocket功能

#### 聊天室WebSocket

```go
// 创建聊天室WebSocket
ws := fishpi.NewChatroomWebSocket("wss://fishpi.cn/chat-room-channel")

// 设置消息回调
ws.OnMessage(func(msg *types.ChatroomMessage) {
    switch msg.Type {
    case types.ChatroomMsgTypeMsg:
        data := msg.Data.(types.ChatroomMsgData)
        fmt.Printf("[聊天] %s: %s\n", data.UserName, data.Content)
    case types.ChatroomMsgTypeRedPacket:
        data := msg.Data.(types.ChatroomRedPacketData)
        fmt.Println("[红包]", data.RedPacket.Msg)
    }
})

// 设置错误回调
ws.OnError(func(err error) {
    log.Println("WebSocket错误:", err)
})

// 连接
if err := ws.Connect(); err != nil {
    log.Fatal(err)
}

// 发送消息
ws.SendMessage("Hello!")

// 使用消息解析器
parser := ws.GetParser()
msg, err := parser.ParseChatroomMessage(data)
```

#### 私聊WebSocket

```go
ws := fishpi.NewPrivateChatWebSocket()

ws.OnMessage(func(msg *types.ChatMessage) {
    if msg.Type == "msg" {
        data := msg.Data.(types.ChatMessageData)
        fmt.Printf("[私聊] %s: %s\n", data.FromUser, data.Content)
    }
})

ws.Connect()
ws.SendMessage("toUser", "消息内容")
```

#### 用户通知WebSocket

```go
ws := fishpi.NewUserNotificationWebSocket()

ws.OnMessage(func(msg *types.UserMessage) {
    switch msg.Command {
    case "bzUpdate":
        // 清风明月更新
    case "refreshNotification":
        // 通知刷新
    }
})

ws.Connect()
```

## API完成度

| 模块 | OpenAPI接口数 | 已实现 | 完成度 |
|------|--------------|--------|--------|
| 文章 | 15 | 15 | ✅ 100% |
| 评论 | 7 | 7 | ✅ 100% |
| 清风明月 | 3 | 3 | ✅ 100% |
| 通知 | 4 | 4 | ✅ 100% |
| 聊天室 | 12 | 12 | ✅ 100% |
| 私聊 | 4 | 4 | ✅ 100% |
| 金手指 | 9 | 9 | ✅ 100% |
| 表情 | 2 | 2 | ✅ 100% |
| 用户 | 18 | 18 | ✅ 100% |
| 认证 | 4 | 4 | ✅ 100% |

**总完成度: 76/76 (100%)**

## 类型安全

SDK使用`go-enum`自动生成所有枚举类型，提供：
- String()方法
- MarshalJSON/UnmarshalJSON
- Parse/MustParse方法
- Names/Values列表

```go
// 枚举使用示例
articleType := types.ArticleListTypeRecent
fmt.Println(articleType.String()) // "recent"

// 解析枚举
parsed, err := types.ParseArticleListType("hot")
```

## 错误处理

所有方法都有完整的错误处理：

```go
if err != nil {
    // 错误已被包装，保留完整的错误链
    log.Printf("操作失败: %v", err)
}
```

## 开发

```bash
# 克隆项目
git clone https://github.com/fghwett/fishpi-golang-sdk

# 安装依赖
go mod download

# 生成枚举代码
go generate ./types/...

# 运行示例
go run examples/main.go
```

## 许可证

Apache 2.0

## 相关链接

- [摸鱼派社区](https://fishpi.cn)
- [API文档](https://fishpi.cn/article/1636516552191)
- [TypeScript SDK](https://github.com/imlinhanchao/fishpi-api-package)

## 贡献

欢迎提交Issue和Pull Request！

