package types

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
