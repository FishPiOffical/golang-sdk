package types

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
