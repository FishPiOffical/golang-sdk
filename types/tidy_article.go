package types

// PostArticleRequest 发布文章请求
type PostArticleRequest struct {
	ArticleTitle           string            `json:"articleTitle"`                   // 帖子标题
	ArticleContent         string            `json:"articleContent"`                 // 帖子内容
	ArticleTags            string            `json:"articleTags"`                    // 帖子标签
	ArticleCommentable     bool              `json:"articleCommentable"`             // 是否允许评论
	ArticleNotifyFollowers bool              `json:"articleNotifyFollowers"`         // 是否通知帖子关注着
	ArticleType            ArticleType       `json:"articleType"`                    // 帖子类型
	ArticleShowInList      ArticleShowInList `json:"articleShowInList"`              // 是否在列表展示
	ArticleRewardContent   *string           `json:"articleRewardContent,omitempty"` // 打赏内容
	ArticleRewardPoint     *int              `json:"articleRewardPoint,omitempty"`   // 打赏积分
	ArticleAnonymous       *bool             `json:"articleAnonymous,omitempty"`     // 是否匿名
	ArticleQnAOfferPoint   *int              `json:"articleQnAOfferPoint,omitempty"` // 提问悬赏积分
}

type PostArticleResponse struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg,omitempty"`
	ArticleId string `json:"articleId,omitempty"`
}
