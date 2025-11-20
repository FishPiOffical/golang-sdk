package types

// PostBreezemoonRequest 发布清风明月请求
type PostBreezemoonRequest struct {
	BreezemoonContent string `json:"breezemoonContent"`
}

// UpdateBreezemoonRequest 更新清风明月请求
type UpdateBreezemoonRequest struct {
	BreezemoonContent string `json:"breezemoonContent"`
}

// BreezemoonList 清风明月列表
type BreezemoonList struct {
	Breezemoons []BreezemoonInfo `json:"breezemoons"`
	Pagination  Pagination       `json:"pagination"`
}
