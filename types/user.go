package types

// PostRegisterRequest 用户注册请求
type PostRegisterRequest struct {
	UserName   string `json:"userName"`
	UserPhone  string `json:"userPhone"`
	InviteCode string `json:"invitecode"`
	Captcha    string `json:"captcha"`
}
