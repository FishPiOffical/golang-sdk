//go:generate go-enum --marshal --names --values --ptr --mustparse
package types

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

// DataType 数据类型
/*
ENUM(
article=0 // 文章
comment=1 // 评论
at=2 // @
commented=3 // 被评论
followingUser=4 // 关注用户
pointCharge=5 // 积分充值
pointTransfer=6 // 积分转账
pointArticleReward=7 // 文章打赏
pointArticleThank=8 // 文章感谢
pointCommentThank=9 // 评论感谢
pointInviteRegister=10 // 邀请注册
pointInviteJoinActivation=11 // 邀请加入激活
pointPerfectArticle=12 // 完美文章
pointCommentAccept=13 // 评论采纳
pointFollowing=14 // 关注
pointFollower=15 // 被关注
pointExchange=16 // 积分兑换
pointAbuseDeduct=17 // 滥用扣除
pointStickArticle=18 // 置顶文章
pointRedPacket=19 // 红包
pointTopArticle=20 // 顶帖
pointReport=21 // 举报
pointRepublishArticle=22 // 重新发布文章
pointSellerBuyerAdvance=23 // 卖家买家预付
pointCommentVote=24 // 评论投票
pointArticleVote=25 // 文章投票
pointAppeal=26 // 申诉
pointPurchaseInvitecode=27 // 购买邀请码
pointForward=28 // 转发
)
*/
type DataType int
