package types

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
