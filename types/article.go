package types

// ArticleDetail 文章详情（包含评论等）
type ArticleDetail struct {
	ArticleInfo
	ArticleComments     []CommentInfo `json:"articleComments"`
	ArticleNiceComments []CommentInfo `json:"articleNiceComments"`
	Pagination          Pagination    `json:"pagination"`
}

// Pagination 分页信息
type Pagination struct {
	PageCount   int `json:"paginationPageCount"`
	PageNum     int `json:"paginationPageNum"`
	PageSize    int `json:"paginationPageSize"`
	RecordCount int `json:"paginationRecordCount"`
	WindowSize  int `json:"paginationWindowSize"`
	FirstPage   int `json:"paginationFirstPage"`
	LastPage    int `json:"paginationLastPage"`
	NextPage    int `json:"paginationNextPage"`
	PrevPage    int `json:"paginationPreviousPage"`
}

// ArticleReward 文章打赏
type ArticleReward struct {
	ArticleId string `json:"articleId"`
}

// ArticleHeat 文章热度消息
type ArticleHeat struct {
	ArticleHeat string `json:"articleHeat"`
}
