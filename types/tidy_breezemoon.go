package types

// BreezemoonInfo 清风明月信息
type BreezemoonInfo struct {
	OId                            string `json:"oId"`
	BreezemoonContent              string `json:"breezemoonContent"`
	BreezemoonAuthorName           string `json:"breezemoonAuthorName"`
	BreezemoonAuthorThumbnailURL48 string `json:"breezemoonAuthorThumbnailURL48"`
	BreezemoonCreateTime           string `json:"breezemoonCreateTime"`
	BreezemoonCreated              int64  `json:"breezemoonCreated"`
	BreezemoonUpdated              int64  `json:"breezemoonUpdated"`
	BreezemoonCity                 string `json:"breezemoonCity"`
	TimeAgo                        string `json:"timeAgo"`
}

// GetBreezemoonsResponse 获取清风明月响应
type GetBreezemoonsResponse struct {
	Code        int              `json:"code"`
	Msg         string           `json:"msg,omitempty"`
	Breezemoons []BreezemoonInfo `json:"breezemoons,omitempty"`
}
