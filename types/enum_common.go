//go:generate go-enum --marshal --names --values --ptr --mustparse
package types

// ReportDataType 举报数据类型
/*
ENUM(
article = 0 // 文章
comment = 1 // 评论
user = 2 // 用户
chatroom = 3 // 聊天消息
)
*/
type ReportDataType int

// ReportType 举报类型
/*
ENUM(
advertise = 0 	// 垃圾广告
porn = 1 		// 色情
violate = 2 	// 违规
infringement = 3 // 侵权
attacks = 4 	// 人身攻击
impersonate = 5 // 冒充他人账号
advertisingAccount = 6 // 垃圾广告账号
leakPrivacy = 7 // 违规泄露个人信息
other = 49 		// 其它
)
*/
type ReportType int

// LogInfoType 日志信息类型
/*
ENUM(
simple // 基础日志
)
*/
type LogInfoType string
