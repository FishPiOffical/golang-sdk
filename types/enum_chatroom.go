//go:generate go-enum --names --values --ptr --mustparse
package types

// ChatroomMsgType 聊天室消息类型
/*
ENUM(
online // 在线
discussChanged // 话题变更
revoke // 撤回
msg // 聊天
redPacket // 红包
redPacketStatus // 红包领取
customMessage // 进入离开聊天室消息
barrager // 弹幕
)
*/
type ChatroomMsgType string

// ChatroomRedPacketType 聊天室红包类型
/*
ENUM(
random // 拼手气红包
average // 平分红包
specify // 专属红包
heartbeat // 心跳红包
rockPaperScissors // 猜拳红包
)
*/
type ChatroomRedPacketType string

// GestureType 猜拳类型
/*
ENUM(
rock=0 // 石头
scissors=1 // 剪刀
paper=2 // 布
)
*/
type GestureType int

// ChatContentType 聊天内容类型
/*
ENUM(
md // Markdown
html // HTML
)
*/
type ChatContentType string

// ChatMessageType 聊天消息查询类型
/*
ENUM(
context=0 // 上下文
before=1 // 之前
after=2 // 之后
)
*/
type ChatMessageType int

// CloudGameId 云游戏ID
/*
ENUM(
emojis // 表情包
)
*/
type CloudGameId string
