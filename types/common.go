package types

// ResponseResult 特殊响应结构(使用sc字段而非code)
type ResponseResult[T any] struct {
	Result int    `json:"sc"`
	Msg    string `json:"msg"`
	Data   T      `json:"data"`
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
