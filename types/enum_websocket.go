//go:generate go-enum --marshal --names --values --ptr --mustparse
package types

// UserChannelCommand 用户频道命令
/*
ENUM(
bzUpdate // 清风明月更新
refreshNotification // 未读消息数通知
chatUnreadCountRefresh // 聊天未读消息数通知
newIdleChatMessage // 新的闲聊消息
warnBroadcast // 警告广播
)
*/
type UserChannelCommand string

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

// ChatroomMsgJsonType 聊天室消息JSON类型
/*
ENUM(
redPacket // 红包
weather // 天气
music // 音乐
)
*/
type ChatroomMsgJsonType string

// ChatroomMsgJsonSubType 聊天室消息JSON子类型
/*
ENUM(
random // 拼手气红包
average // 平分红包
specify // 专属红包
heartbeat // 心跳红包
rockPaperScissors // 猜拳红包
weather // 天气
music // 音乐
)
*/
type ChatroomMsgJsonSubType string
