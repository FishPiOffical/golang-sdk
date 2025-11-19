//go:generate go-enum --names --values --ptr --mustparse
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
article = 0						// 文章
comment = 1						// 评论
at = 2							// @
commented = 3					// 被评论
followingUser = 4 				// 关注用户
pointCharge = 5					// 积分 - 充值
pointTransfer = 6				// 积分 - 转账
pointArticleReward = 7			// 积分 - 文章打赏
pointCommentThank = 8 			// 积分 - 评论感谢
broadcast = 9					// 同城广播
pointExchange = 10 				// 积分 - 交易
abusePointDeduct = 11 			// 积分 - 违规扣除
pointArticleThank = 12 			// 积分 - 文章被感谢
reply = 13						// 回复
invitecodeUsed = 14 			// 使用邀请码
sysAnnounceArticle = 15 		// 系统公告 - 文章
sysAnnounceNewUser = 16 		// 系统公告 - 新用户
newFollower = 17 				// 新的关注者
invitationLinkUsed = 18 		// 邀请链接
sysAnnounceRoleChanged = 19 	// 系统通知 - 角色变化
followingArticleUpdate = 20 	// 关注的文章更新
followingArticleComment = 21	// 关注的文章评论
pointPerfectArticle = 22 		// 积分 - 文章优选
articleNewFollower = 23 		// 文章新的被关注者
articleNewWatcher = 24 		// 文章新的关注者
commentVoteUp = 25 			// 评论点赞
commentVoteDown = 26 			// 评论点踩
articleVoteUp = 27 			// 文章被点赞
articleVoteDown = 28 			// 文章被点踩
pointCommentAccept = 33 		// 积分 - 评论被接受
pointReportHandled = 36 		// 积分 - 举报处理
chatRoomAt = 38 				// 聊天室 @
redPacket = 39 				// 专属红包提醒
)
*/
type DataType int
