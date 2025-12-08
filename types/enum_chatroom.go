//go:generate go-enum --names --values --ptr --mustparse
package types

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

// ChatContentType 聊天内容类型
/*
ENUM(
md // Markdown
html // HTML
)
*/
type ChatContentType string
