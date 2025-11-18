package types

import "time"

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
	OId                  string      `json:"oId"`
	UserNo               string      `json:"userNo"`
	UserName             string      `json:"userName"`
	UserNickname         string      `json:"userNickname"`
	UserAvatarURL        string      `json:"userAvatarURL"`
	UserAvatarURL210     string      `json:"userAvatarURL210"`
	UserAvatarURL48      string      `json:"userAvatarURL48"`
	UserURL              string      `json:"userURL"`
	UserIntro            string      `json:"userIntro"`
	UserPoint            int         `json:"userPoint"`
	UserAppRole          UserAppRole `json:"userAppRole"`
	UserRole             string      `json:"userRole"`
	OnlineMinute         int         `json:"onlineMinute"`
	UserCity             string      `json:"userCity"`
	SysMetal             string      `json:"sysMetal"`
	CardBg               string      `json:"cardBg"`
	FollowingUserCount   int         `json:"followingUserCount"`
	FollowerCount        int         `json:"followerCount"`
	CanFollow            CanFollow   `json:"canFollow"`
	UserOnlineFlag       bool        `json:"userOnlineFlag"`
	UserCreateTime       time.Time   `json:"userCreateTime"`
	UserCreateTimeStr    string      `json:"userCreateTimeStr"`
	UserUpdateTime       time.Time   `json:"userUpdateTime"`
	UserUpdateTimeStr    string      `json:"userUpdateTimeStr"`
	UserCheckedIn        bool        `json:"checkedIn"`
	CheckedInToday       bool        `json:"checkedInToday"`
	CurrentCheckedInDays int         `json:"currentCheckedInDays"`
	LongestCheckedInDays int         `json:"longestCheckedInDays"`
	MBTI                 string      `json:"mbti"`
}

// TransferRequest 转账请求
type TransferRequest struct {
	UserName string `json:"userName"`
	Amount   int    `json:"amount"`
	Memo     string `json:"memo"`
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

// IsCollectedLivenessResponse 是否已领取昨日活跃奖励响应
type IsCollectedLivenessResponse struct {
	Code                               int    `json:"code"`
	Msg                                string `json:"msg"`
	IsCollectedYesterdayLivenessReward bool   `json:"isCollectedYesterdayLivenessReward"`
}

// PostRegisterRequest 用户注册请求
type PostRegisterRequest struct {
	UserName   string `json:"userName"`
	UserPhone  string `json:"userPhone"`
	InviteCode string `json:"invitecode"`
	Captcha    string `json:"captcha"`
}
