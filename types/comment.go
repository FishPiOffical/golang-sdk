package types

// PostCommentRequest 发布评论请求
type PostCommentRequest struct {
	ArticleId         string `json:"articleId"`
	CommentContent    string `json:"commentContent"`
	CommentAnonymous  bool   `json:"commentAnonymous,omitempty"`
	CommentVisible    int    `json:"commentVisible,omitempty"` // 0 or 1
	CommentOriginalId string `json:"commentOriginalCommentId,omitempty"`
}

// UpdateCommentRequest 更新评论请求
type UpdateCommentRequest struct {
	CommentContent string `json:"commentContent"`
}

// CommentInfo 评论信息
type CommentInfo struct {
	OId                         string `json:"oId"`
	CommentContent              string `json:"commentContent"`
	CommentCreateTime           int64  `json:"commentCreateTime"`
	CommentUpdateTime           int64  `json:"commentUpdateTime"`
	CommentAuthorName           string `json:"commentAuthorName"`
	CommentAuthorThumbnailURL   string `json:"commentAuthorThumbnailURL"`
	CommentAuthorThumbnailURL48 string `json:"commentAuthorThumbnailURL48"`
	CommentSharpURL             string `json:"commentSharpURL"`
	CommentGoodCnt              int    `json:"commentGoodCnt"`
	CommentBadCnt               int    `json:"commentBadCnt"`
	CommentReplyCnt             int    `json:"commentReplyCnt"`
	CommentScore                int    `json:"commentScore"`
	CommentAnonymous            int    `json:"commentAnonymous"`
	CommentVisible              int    `json:"commentVisible"`
	CommentThankCnt             int    `json:"commentThankCnt"`
	CommentOriginalCommentId    string `json:"commentOriginalCommentId"`
}
