//go:generate go-enum --marshal --names --values --ptr --mustparse
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

// ArticleType 文章类型
/*
ENUM(
normal=0 // 普通帖子
discussion=1 // 讨论区
city=2 // 同城
qna=3 // 问答
)
*/
type ArticleType int
