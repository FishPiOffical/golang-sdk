//go:generate go-enum --names --values --ptr --mustparse
package types

// ReactionEmojiValue 表情值
/*
ENUM(
thumbsup // 👍 赞
plus // ➕1️⃣ +1/同意/我也是
thumbsdown	// 👎	反对 / 不喜欢
check // ✅	对 / 正确 / 通过
cross	// ❌	错 / 否 / 不通过
star	// ⭐	收藏 / 亮眼 / 很棒
heart	// ❤️️	喜欢 / 支持
fire	// 🔥	热门 / 很强
party	// 🎉	庆祝
laugh	// 😂	好笑
wow	// 😮	惊讶
clap	// 👏	鼓掌 / 认可
hundred	// 💯	满分 / 非常认可
rocket	// 🚀	起飞 / 冲
salute	// 🖖	致意 / 打招呼
handshake	// 🤝	握手 / 达成一致
raisedhands	// 🙌	赞成 / 欢呼
mindblown	// 🤯	震惊 / 太强了
thinking	// 🤔	思考 / 存疑
eyes	// 👀	围观 / 关注
cry	// 😢	难过 / 共情
angry	// 😡	生气 / 不满
pray	// 🙏	祈祷 / 拜托 / 感谢
brokenheart	// 💔	破防 / 难受
heartonfire	// ❤️️🔥	上头 / 强烈喜欢
skull	// 💀	笑死 / 绝了
clown	// 🤡	小丑 / 离谱
poop	// 💩	拉胯 / 吐槽
)
*/
type ReactionEmojiValue string

// ReactionGroupType Reaction分组类型
/*
ENUM(
emoji // 表情
)
*/
type ReactionGroupType string

// ReactionTargetType Reaction目标类型
/*
ENUM(
article // 文章
comment // 评论
chat // 聊天室消息
)
*/
type ReactionTargetType string

type PostReactionData struct {
	TargetId            string             `json:"targetId"`
	TargetType          ReactionTargetType `json:"targetType"`
	GroupType           ReactionGroupType  `json:"groupType"`
	CurrentUserReaction string             `json:"currentUserReaction"`
	Summary             []*ReactionSummary `json:"summary"`
}
