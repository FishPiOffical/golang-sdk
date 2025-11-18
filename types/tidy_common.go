package types

// SimpleResponse 简单响应结构(无Data字段)
type SimpleResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// ApiResponse 通用API响应结构(带Data字段)
type ApiResponse[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

// User 其他用户的信息
type User struct {
	OId                string      `json:"oId"`
	UserNo             string      `json:"userNo"`
	UserName           string      `json:"userName"`
	UserNickname       string      `json:"userNickname"`
	UserAvatarURL      string      `json:"userAvatarURL"`
	UserAvatarURL20    string      `json:"userAvatarURL20"`
	UserAvatarURL210   string      `json:"userAvatarURL210"`
	UserAvatarURL48    string      `json:"userAvatarURL48"`
	UserURL            string      `json:"userURL"`
	UserIntro          string      `json:"userIntro"`
	UserPoint          int         `json:"userPoint"`
	UserAppRole        UserAppRole `json:"userAppRole"`
	UserRole           string      `json:"userRole"`
	OnlineMinute       int         `json:"onlineMinute"`
	UserCity           string      `json:"userCity"`
	SysMetal           string      `json:"sysMetal"`
	CardBg             string      `json:"cardBg"`
	FollowingUserCount int         `json:"followingUserCount"`
	FollowerCount      int         `json:"followerCount"`
	CanFollow          CanFollow   `json:"canFollow"`
	UserOnlineFlag     bool        `json:"userOnlineFlag"`
	MBTI               string      `json:"mbti"`
	AllMetalOwned      string      `json:"allMetalOwned"`
}

// AssociateUser 联想的用户信息
type AssociateUser struct {
	UserName          string `json:"userName"`
	UserAvatarURL     string `json:"userAvatarURL"`
	UserAvatarURL20   string `json:"userAvatarURL20"`
	UserAvatarURL48   string `json:"userAvatarURL48"`
	UserAvatarURL210  string `json:"userAvatarURL210"`
	UserNameLowerCase string `json:"userNameLowerCase"`
}

// UserLivenessResponse 用户活跃度响应
type UserLivenessResponse struct {
	Code     int     `json:"code"`
	Msg      string  `json:"msg"`
	Liveness float64 `json:"liveness"`
}

// UserCheckedInResponse 用户签到状态响应
type UserCheckedInResponse struct {
	CheckedIn bool `json:"checkedIn"`
}

// YesterdayLivenessRewardResponse 领取昨日活跃奖励响应
type YesterdayLivenessRewardResponse struct {
	Sum int `json:"sum"`
}

// IsCollectedLivenessResponse 是否已领取昨日活跃奖励响应
type IsCollectedLivenessResponse struct {
	IsCollectedYesterdayLivenessReward bool `json:"isCollectedYesterdayLivenessReward"`
}
