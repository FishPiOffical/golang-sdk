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

// UploadFileResponse 上传文件响应
type UploadFileResponse struct {
	SuccMap map[string]string `json:"succMap"` // 成功上传的文件URL映射
}

// SendRedPacketRequest 发送红包请求
type SendRedPacketRequest struct {
	Type     string `json:"type"`               // average, random, specify, heartbeat, rockPaperScissors
	Money    int    `json:"money"`              // 总金额（单位：积分）
	Count    int    `json:"count"`              // 红包个数
	Msg      string `json:"msg"`                // 红包祝福语
	Recivers string `json:"recivers,omitempty"` // 指定领取人（专属红包）
	Gesture  int    `json:"gesture,omitempty"`  // 猜拳类型（猜拳红包）
}

// RedPacketDetail 红包详情
type RedPacketDetail struct {
	OId     string   `json:"oId"`
	Type    string   `json:"type"`
	Money   int      `json:"money"`
	Count   int      `json:"count"`
	Msg     string   `json:"msg"`
	Got     int      `json:"got"`
	WhoGot  []string `json:"whoGot"`
	WhoMiss []string `json:"whoMiss"`
}
