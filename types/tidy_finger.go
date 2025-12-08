package types

type LatestLoginIp struct {
	UserLatestLoginIp string `json:"userLatestLoginIp,omitempty"` // 用户最近登录IP
	UserId            string `json:"userId,omitempty"`            // 用户ID
}

type PostUserLivenessResponse struct {
	Code     int     `json:"code,omitempty"`
	Msg      string  `json:"msg,omitempty"`
	Liveness float64 `json:"liveness,omitempty"`
}

type PostYesterdayLivenessRewardApiResponse struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Sum  int    `json:"sum,omitempty"`
}
