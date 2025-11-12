//go:generate go-enum --marshal --names --values --ptr --mustparse
package types

// ArticleListType 文章列表类型
/*
ENUM(
// 最新
// 热门
hot
// 精华
good
// 精选
perfect
// 回复
reply
)
*/
type ArticleListType string

// NotificationType 通知类型
/*
ENUM(
point // 积分
commented // 收到的回帖
reply // 收到的回复
at // 提及我的
following // 我关注的
broadcast // 同城
sys-announce // 系统
)
*/
type NotificationType string

// ChatroomMsgType 消息类型
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
rock // 石头
scissors // 剪刀
paper // 布
)
*/
type GestureType int

// VoteType 投票类型
/*
ENUM(
unVote = -1 // 未投票
voted = 0  // 点赞
)
*/
type VoteType int

// ChatContentType 聊天内容类型
/*
ENUM(
md   // Markdown
html // HTML
)
*/
type ChatContentType string

// ChatMessageType 聊天消息查询类型
/*
ENUM(
context // 上下文
before  // 之前
after   // 之后
)
*/
type ChatMessageType int
