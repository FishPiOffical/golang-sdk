package types

import "encoding/json"

type UserChannelMsg struct {
	Command UserChannelCommand      `json:"command"`
	Bz      *UserChannelMsgBzUpdate `json:"bz,omitempty"`

	// UserChannelCommandRefreshNotification UserChannelCommandChatUnreadCountRefresh
	Count int `json:"count,omitempty"`

	// UserChannelCommandRefreshNotification UserChannelCommandChatUnreadCountRefresh UserChannelCommandNewIdleChatMessage
	UserId string `json:"userId,omitempty"`

	// UserChannelCommandNewIdleChatMessage
	SendUserName string `json:"sendUserName,omitempty"`
	SendAvatar   string `json:"sendAvatar,omitempty"`
	Preview      string `json:"preview,omitempty"`

	// UserChannelCommandWarnBroadcast
	WarnBroadcastText string `json:"warnBroadcastText,omitempty"`
	Who               string `json:"who,omitempty"`
}

type UserChannelMsgBzUpdate struct {
	OId                            string `json:"oId"`
	BreezemoonContent              string `json:"breezemoonContent"`
	BreezemoonAuthorName           string `json:"breezemoonAuthorName"`
	BreezemoonAuthorThumbnailURL48 string `json:"breezemoonAuthorThumbnailURL48"`
}

type ChatChannelMsg struct {
	ToId             string `json:"toId"`
	Preview          string `json:"preview"`
	UserSession      string `json:"user_session"`
	Markdown         string `json:"markdown"`
	SenderAvatar     string `json:"senderAvatar"`
	ReceiverAvatar   string `json:"receiverAvatar"`
	OId              string `json:"oId"`
	Time             string `json:"time"`
	FromId           string `json:"fromId"`
	SenderUserName   string `json:"senderUserName"`
	Content          string `json:"content"`
	ReceiverUserName string `json:"receiverUserName"`
}

type ChatroomMsg struct {
	Type ChatroomMsgType `json:"type"` // 消息类型

	// 在线消息
	Discussing    string            `json:"discussing"`    // 讨论的话题
	OnlineChatCnt int               `json:"onlineChatCnt"` // 在线人数
	Users         []*OnlineUserInfo `json:"users"`         // 在线用户信息

	// 话题变更
	NewDiscuss string `json:"newDiscuss"` // 新的话题内容

	// 聊天 撤回 红包领取
	OId string `json:"oId"` // 消息ID

	// 聊天消息
	Time             string `json:"time"`             // 发布时间
	UserName         string `json:"userName"`         // 用户名
	UserNickname     string `json:"userNickname"`     // 用户昵称
	UserAvatarURL    string `json:"userAvatarURL"`    // 用户头像
	UserAvatarURL20  string `json:"userAvatarURL20"`  // 用户头像 20px
	UserAvatarURL48  string `json:"userAvatarURL48"`  // 用户头像 48px
	UserAvatarURL210 string `json:"userAvatarURL210"` // 用户邮箱 210px
	SysMetal         string `json:"sysMetal"`         // 徽章数据 json字符串
	Content          string `json:"content"`          // 消息内容 HTML格式 如果是红包则是JSON格式
	Md               string `json:"md"`               // 消息内容 Markdown格式，红包消息无此栏位

	// 红包领取消息
	Count   int    `json:"count"`   // 红包个数
	Got     int    `json:"got"`     // 已领取个数
	WhoGive string `json:"whoGive"` // 发送者用户名
	WhoGot  string `json:"whoGot"`  // 领取者用户名

	// 客户端
	Client string `json:"client"` // 消息客户端

	// 普通消息
	Message string `json:"message"` // 普通消息的消息内容

	// 弹幕消息
	BarrageColor   string `json:"barragerColor"`
	BarrageContent string `json:"barragerContent"`
}

func (msg *ChatroomMsg) GetMetalList() ([]*Metal, error) {
	if msg.SysMetal == "" {
		return []*Metal{}, nil
	}
	sysMetal := &SysMetal{}
	if err := json.Unmarshal([]byte(msg.SysMetal), sysMetal); err != nil {
		return nil, err
	}
	return sysMetal.List, nil
}

func (msg *ChatroomMsg) GetJsonInfo() *ChatroomMsgJsonInfo {
	data := new(ChatroomMsgJsonInfo)

	if err := json.Unmarshal([]byte(msg.Content), data); err != nil {
		return nil
	}

	return data
}

type OnlineUserInfo struct {
	UserName         string `json:"userName"`         // 用户名
	HomePage         string `json:"homePage"`         // 用户首页
	UserAvatarURL    string `json:"userAvatarURL"`    // 用户头像
	UserAvatarURL20  string `json:"userAvatarURL20"`  // 用户头像 20px
	UserAvatarURL48  string `json:"userAvatarURL48"`  // 用户头像 48px
	UserAvatarURL210 string `json:"userAvatarURL210"` // 用户邮箱 210px
}

type SysMetal struct {
	List []*Metal `json:"list"` // 徽章列表数据
}

// ChatroomMsgJsonInfo json的数据结构
type ChatroomMsgJsonInfo struct {
	// 消息类型 redPacket-红包 weather-天气 music-音乐
	MsgType ChatroomMsgJsonType `json:"msgType"`
	/*
		红包类型 random(拼手气红包), average(平分红包)，specify(专属红包)，heartbeat(心跳红包)，rockPaperScissors(猜拳红包)
		天气类型 weather
		音乐类型 music
	*/
	Type ChatroomMsgJsonSubType `json:"type"`

	Msg      string        `json:"msg"`      // 红包祝福语
	Recivers string        `json:"recivers"` // 红包接收者用户名，专属红包有效
	SenderId string        `json:"senderId"` // 发送者id
	Money    int           `json:"money"`    // 红包积分
	Count    int           `json:"count"`    // 红包个数
	Got      int           `json:"got"`      // 已领取个数
	Who      []interface{} `json:"who"`      // 已领取者信息

	Date        string `json:"date"`        // 日期
	St          string `json:"st"`          // 一句话
	Min         string `json:"min"`         // 最低温度
	T           string `json:"t"`           // 城市
	Max         string `json:"max"`         // 最高温度
	WeatherCode string `json:"weatherCode"` // 天气

	CoverURL string `json:"coverURL"` // 封面链接
	From     string `json:"from"`     // 来源
	Source   string `json:"source"`   // 音乐链接
	Title    string `json:"title"`    // 音乐名
}
