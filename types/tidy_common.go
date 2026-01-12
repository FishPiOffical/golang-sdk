//go:generate go-enum --marshal --names --values --ptr --mustparse
package types

import "encoding/json"

// SimpleResponse 简单响应结构(无Data字段)
type SimpleResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
}

// ApiResponse 通用API响应结构(带Data字段)
type ApiResponse[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data T      `json:"data"`
}

// ResultResponse 通用API响应结构(带Data字段)
type ResultResponse[T any] struct {
	Result int    `json:"result"`
	Msg    string `json:"msg,omitempty"`
	Data   T      `json:"data,omitempty"`
	Cached bool   `json:"cached,omitempty"` // /chat/get-list /chat/has-unread 接口返回
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
	UserRole           UserRole    `json:"userRole"`
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

// UserRole 用户角色 新手、成员、超级会员、纪律委员、OP、管理员
/*
ENUM(
new=新手
vip=会员
svip=超级会员
police=纪律委员
op=OP
admin=管理员
)
*/
type UserRole string

func (user *User) GetMetalList() ([]*Metal, error) {
	if user.SysMetal == "" {
		return []*Metal{}, nil
	}
	sysMetal := &SysMetal{}
	if err := json.Unmarshal([]byte(user.SysMetal), sysMetal); err != nil {
		return nil, err
	}
	return sysMetal.List, nil
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

type RecentRegUser struct {
	UserName     string `json:"userName"`
	UserNickname string `json:"userNickname"`
}

// PostPointTransferRequest 转账请求
type PostPointTransferRequest struct {
	UserName string `json:"userName"`
	Amount   int    `json:"amount"`
	Memo     string `json:"memo"`
}

type LogInfo struct {
	Key1   string      `json:"key1"`   // 2025-12-08 17:57:15
	Key2   string      `json:"key2"`   // 49.232.59.*
	Data   string      `json:"data"`   // 用户: aaa12123, 积分: 128, 备注: 你被烟花烫伤(4 次)!
	Public bool        `json:"public"` // true
	Key3   string      `json:"key3"`   // 扣除积分
	OId    string      `json:"oId"`    // 1765187835408
	Type   LogInfoType `json:"type"`   // simple
}

type MemberShip struct {
	ConfigJson string `json:"configJson"`
	CreatedAt  int64  `json:"createdAt"`
	OId        string `json:"oId"`
	State      int    `json:"state"`
	UserId     string `json:"userId"`
	LvCode     string `json:"lvCode"`
	ExpiresAt  int64  `json:"expiresAt"`
	UpdatedAt  int64  `json:"updatedAt"`
}

func (membership *MemberShip) GetConfig() (*MemberShipConfig, error) {
	config := &MemberShipConfig{}

	if err := json.Unmarshal([]byte(membership.ConfigJson), config); err != nil {
		return nil, err
	}

	return config, nil
}

type MemberShipConfig struct {
	JointVip    bool   `json:"jointVip"`
	Color       string `json:"color"`
	Underline   bool   `json:"underline"`
	Metal       bool   `json:"metal"`
	AutoCheckin int    `json:"autoCheckin"`
	Bold        bool   `json:"bold"`
}

type PostSettingsProfilesRequest struct {
	UserNickname string `json:"userNickname,omitempty"` // 昵称
	UserTags     string `json:"userTags,omitempty"`     // 用户标签，多个标签用逗号分隔
	UserURL      string `json:"userURL,omitempty"`      // 个人主页URL
	UserIntro    string `json:"userIntro,omitempty"`    // 个人简介
	Mbti         string `json:"mbti,omitempty"`         // MBTI性格类型（例如：ENFP）todo 类型枚举
}

type GetUserUsernamePointData struct {
	UserPoint int    `json:"userPoint"` // 用户积分
	UserName  string `json:"userName"`  // 用户名
}

type GetUserUsernameMedalData struct {
	List []*GetUserUsernameMedalInfo `json:"list"` // 勋章列表
}

type GetUserUsernameMedalInfo struct {
	Data        string `json:"data"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ExpireDate  string `json:"expireDate"`
	Id          string `json:"id"`
	Attr        string `json:"attr"`
	Type        string `json:"type"`
	Enabled     bool   `json:"enabled"`
	Order       int    `json:"order"`
}
