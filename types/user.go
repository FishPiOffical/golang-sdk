package types

// PostApiGetKeyRequest 获取ApiKey请求
type PostApiGetKeyRequest struct {
	NameOrEmail  string `json:"nameOrEmail"`
	UserPassword string `json:"userPassword"`
	MfaCode      string `json:"mfaCode,omitempty"`
}

// PostApiGetKeyResponse 获取ApiKey响应
type PostApiGetKeyResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Key  string `json:"Key"`
}

// CheckinRequest 签到请求
type CheckinRequest struct{}

// CheckinResponse 签到响应
type CheckinResponse struct {
	Sum               int  `json:"sum"`
	YesterdayLiveness int  `json:"yesterdayLiveness"`
	CheckedIn         bool `json:"checkedIn"`
}

// LivenessReward 活跃度奖励信息
type LivenessReward struct {
	Sum int `json:"sum"`
}

// UserInfoResponse 用户信息响应
type UserInfoResponse struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data UserInfo `json:"data"`
}

// PostRegisterRequest 用户注册请求
type PostRegisterRequest struct {
	UserName   string `json:"userName"`
	UserPhone  string `json:"userPhone"`
	InviteCode string `json:"invitecode"`
	Captcha    string `json:"captcha"`
}
