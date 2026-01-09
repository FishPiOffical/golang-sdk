//go:generate go-enum --names --values --ptr --mustparse
package types

// ArticleListType 文章列表类型
/*
ENUM(
// 最新
hot // 热门
good // 精华
perfect // 精选
reply // 回复
)
*/
type ArticleListType string

// ArticlePerfect 是否精选文章
/*
ENUM(
yes=1 // 是
no=0 // 否
)
*/
type ArticlePerfect int

// GetArticleType 获取文章类型
/*
ENUM(
recent // 最近
tag // 标签
domain // 领域
)
*/
type GetArticleType string

// GetArticleOrder 获取文章排序方式
/*
ENUM(
hot // 热门
good // 点赞
reply // 回复
perfect // 精选
)
*/
type GetArticleOrder string

// VoteResult 文章投票结果
/*
ENUM(
cancel=0 // 取消投票
success=-1 // 投票成功
)
*/
type VoteResult int
