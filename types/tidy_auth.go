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

// UserInfo 用户信息
type UserInfo struct {
	UserCity           string      `json:"userCity"`
	UserOnlineFlag     bool        `json:"userOnlineFlag"`
	UserPoint          int         `json:"userPoint"`
	UserAppRole        UserAppRole `json:"userAppRole"`
	UserIntro          string      `json:"userIntro"`
	UserNo             string      `json:"userNo"`
	OnlineMinute       int         `json:"onlineMinute"`
	UserAvatarURL      string      `json:"userAvatarURL"`
	UserNickname       string      `json:"userNickname"`
	OId                string      `json:"oId"`
	UserName           string      `json:"userName"`
	CardBg             string      `json:"cardBg"`
	FollowingUserCount int         `json:"followingUserCount"`
	SysMetal           string      `json:"sysMetal"`
	UserRole           string      `json:"userRole"`
	FollowerCount      int         `json:"followerCount"`
	UserURL            string      `json:"userURL"`
}
