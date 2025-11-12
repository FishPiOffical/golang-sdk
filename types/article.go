package types

// PostArticleRequest 发布文章请求
type PostArticleRequest struct {
	ArticleTitle           string      `json:"articleTitle"`
	ArticleContent         string      `json:"articleContent"`
	ArticleTags            string      `json:"articleTags"`
	ArticleCommentable     bool        `json:"articleCommentable"`
	ArticleType            ArticleType `json:"articleType"`
	ArticleRewardContent   string      `json:"articleRewardContent,omitempty"`
	ArticleRewardPoint     int         `json:"articleRewardPoint,omitempty"`
	ArticleQnAOfferPoint   int         `json:"articleQnAOfferPoint,omitempty"`
	ArticleNotifyFollowers bool        `json:"articleNotifyFollowers,omitempty"`
	ArticleShowInList      int         `json:"articleShowInList,omitempty"` // 0 or 1
	ArticleAnonymous       bool        `json:"articleAnonymous,omitempty"`
}

// UpdateArticleRequest 更新文章请求
type UpdateArticleRequest struct {
	ArticleTitle         string      `json:"articleTitle"`
	ArticleContent       string      `json:"articleContent"`
	ArticleTags          string      `json:"articleTags"`
	ArticleCommentable   bool        `json:"articleCommentable"`
	ArticleType          ArticleType `json:"articleType"`
	ArticleRewardContent string      `json:"articleRewardContent,omitempty"`
	ArticleRewardPoint   int         `json:"articleRewardPoint,omitempty"`
}

// ArticleInfo 文章信息
type ArticleInfo struct {
	OId                         string `json:"oId"`
	ArticleTitle                string `json:"articleTitle"`
	ArticleContent              string `json:"articleContent"`
	ArticleContentHTML          string `json:"articleContentHTML"`
	ArticleTags                 string `json:"articleTags"`
	ArticleAuthorName           string `json:"articleAuthorName"`
	ArticleAuthorThumbnailURL   string `json:"articleAuthorThumbnailURL"`
	ArticleAuthorThumbnailURL48 string `json:"articleAuthorThumbnailURL48"`
	ArticleCreateTime           int64  `json:"articleCreateTime"`
	ArticleUpdateTime           int64  `json:"articleUpdateTime"`
	ArticleViewCount            int    `json:"articleViewCnt"`
	ArticleCommentCount         int    `json:"articleCommentCnt"`
	ArticleThankCount           int    `json:"articleThankCnt"`
	ArticleGoodCount            int    `json:"articleGoodCnt"`
	ArticleBadCount             int    `json:"articleBadCnt"`
	ArticleCollectCount         int    `json:"articleCollectCnt"`
	ArticleWatchCount           int    `json:"articleWatchCnt"`
	ArticlePermalink            string `json:"articlePermalink"`
	ArticlePerfectCount         int    `json:"articlePerfectCnt"`
	ArticleAnonymous            int    `json:"articleAnonymous"`
	ArticleType                 int    `json:"articleType"`
	ArticleStick                int64  `json:"articleStick"`
	ArticleQnAOfferPoint        int    `json:"articleQnAOfferPoint"`
	Offered                     bool   `json:"offered"`
	ArticlePreviewContent       string `json:"articlePreviewContent"`
	ArticleToC                  string `json:"articleToC"`
	ArticleAudioURL             string `json:"articleAudioURL"`
}

// ArticleDetail 文章详情（包含评论等）
type ArticleDetail struct {
	ArticleInfo
	ArticleComments     []CommentInfo `json:"articleComments"`
	ArticleNiceComments []CommentInfo `json:"articleNiceComments"`
	Pagination          Pagination    `json:"pagination"`
}

// ArticleList 文章列表
type ArticleList struct {
	Articles   []ArticleInfo `json:"articles"`
	Pagination Pagination    `json:"pagination"`
}

// ArticleReward 文章打赏
type ArticleReward struct {
	ArticleId string `json:"articleId"`
}

// ArticleHeat 文章热度消息
type ArticleHeat struct {
	ArticleHeat string `json:"articleHeat"`
}
