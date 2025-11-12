package types

// ApiResponse 通用API响应结构(带Data字段)
type ApiResponse[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

// SimpleResponse 简单响应结构(无Data字段)
type SimpleResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// ResponseResult 特殊响应结构(使用sc字段而非code)
type ResponseResult[T any] struct {
	Result int    `json:"sc"`
	Msg    string `json:"msg"`
	Data   T      `json:"data"`
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

// Metal 勋章信息
type Metal struct {
	Name            string `json:"name"`
	Attr            string `json:"attr"`
	Description     string `json:"description"`
	Data            string `json:"data"`
	BackgroundImage string `json:"backgroundImage"`
	Enabled         bool   `json:"enabled"`
}
