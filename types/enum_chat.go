//go:generate go-enum --marshal --names --values --ptr --mustparse
package types

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

// GestureType 猜拳类型
/*
ENUM(
rock=0 // 石头
scissors=1 // 剪刀
paper=2 // 布
)
*/
type GestureType int
