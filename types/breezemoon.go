package types

// PostBreezemoonRequest 发布清风明月请求
type PostBreezemoonRequest struct {
	BreezemoonContent string `json:"breezemoonContent"`
}

// UpdateBreezemoonRequest 更新清风明月请求
type UpdateBreezemoonRequest struct {
	BreezemoonContent string `json:"breezemoonContent"`
}

// BreezemoonInfo 清风明月信息
type BreezemoonInfo struct {
	OId                            string `json:"oId"`
	BreezemoonContent              string `json:"breezemoonContent"`
	BreezemoonAuthorName           string `json:"breezemoonAuthorName"`
	BreezemoonAuthorThumbnailURL   string `json:"breezemoonAuthorThumbnailURL"`
	BreezemoonAuthorThumbnailURL48 string `json:"breezemoonAuthorThumbnailURL48"`
	BreezemoonCreateTime           int64  `json:"breezemoonCreateTime"`
	BreezemoonUpdated              int64  `json:"breezemoonUpdated"`
	BreezemoonCity                 string `json:"breezemoonCity"`
	TimeAgo                        string `json:"timeAgo"`
}

// BreezemoonList 清风明月列表
type BreezemoonList struct {
	Breezemoons []BreezemoonInfo `json:"breezemoons"`
	Pagination  Pagination       `json:"pagination"`
}
