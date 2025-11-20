package types

// BreezemoonInfo 清风明月信息
type BreezemoonInfo struct {
	OId                            string `json:"oId"`
	BreezemoonContent              string `json:"breezemoonContent"`
	BreezemoonAuthorName           string `json:"breezemoonAuthorName"`
	BreezemoonAuthorThumbnailURL48 string `json:"breezemoonAuthorThumbnailURL48"`

	BreezemoonCreateTime string `json:"breezemoonCreateTime,omitempty"`
	BreezemoonCreated    int64  `json:"breezemoonCreated,omitempty"`
	BreezemoonUpdated    int64  `json:"breezemoonUpdated,omitempty"`
	BreezemoonCity       string `json:"breezemoonCity,omitempty"`
	TimeAgo              string `json:"timeAgo,omitempty"`
}

// GetBreezemoonsResponse 获取清风明月响应
type GetBreezemoonsResponse struct {
	Code        int              `json:"code"`
	Msg         string           `json:"msg,omitempty"`
	Breezemoons []BreezemoonInfo `json:"breezemoons,omitempty"`
}

type BreezemoonPagination struct {
	PaginationPageCount   int   `json:"paginationPageCount"`
	PaginationPageNums    []int `json:"paginationPageNums"`
	PaginationRecordCount int   `json:"paginationRecordCount"`
}

type GetUserBreezemoonsData struct {
	Pagination  BreezemoonPagination `json:"pagination"`
	Breezemoons []*BreezemoonInfo    `json:"breezemoons"`
}
