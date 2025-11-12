//go:generate go-enum --marshal --names --values --ptr --mustparse
package types

// VoteType 投票类型
/*
ENUM(
unvoted=-1 // 未投票
up=0 // 点赞
down=1 // 点踩
)
*/
type VoteType int
