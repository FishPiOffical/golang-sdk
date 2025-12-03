package types

// ChatroomMessage 聊天室消息
type ChatroomMessage struct {
	Type ChatroomMsgType `json:"type"`
	Data interface{}     `json:"data"`
	Time int64           `json:"time"`
}

// OnlineUser 在线用户
type OnlineUser struct {
	UserName      string `json:"userName"`
	UserNickname  string `json:"userNickname"`
	UserAvatarURL string `json:"userAvatarURL"`
	HomepageURL   string `json:"homepageURL"`
}

// ChatroomOnlineData 在线用户数据
type ChatroomOnlineData struct {
	Users []OnlineUser `json:"users"`
}

// ChatroomDiscussData 话题变更数据
type ChatroomDiscussData struct {
	Discuss string `json:"newDiscuss"`
}

// ChatroomRevokeData 撤回消息数据
type ChatroomRevokeData struct {
	OId string `json:"oId"`
}

// RedPacketInfo1 红包信息
type RedPacketInfo1 struct {
	OId     string                `json:"oId"`
	Type    ChatroomRedPacketType `json:"type"`
	Count   int                   `json:"count"`
	Money   int                   `json:"money"`
	Msg     string                `json:"msg"`
	Got     int                   `json:"got"`
	WhoGot  []string              `json:"whoGot"`
	WhoMiss []string              `json:"whoMiss"`
	Gesture *GestureType          `json:"gesture,omitempty"`
}

// ChatroomRedPacketData 红包消息数据
type ChatroomRedPacketData struct {
	ChatroomMsgData
	RedPacket RedPacketInfo1 `json:"redPacket"`
}

// ChatroomRedPacketStatusData 红包领取状态数据
type ChatroomRedPacketStatusData struct {
	OId     string   `json:"oId"`
	Count   int      `json:"count"`
	Got     int      `json:"got"`
	WhoGot  []string `json:"whoGot"`
	WhoMiss []string `json:"whoMiss"`
}

// ChatroomCustomData 自定义消息数据
type ChatroomCustomData struct {
	Content string `json:"content"`
}

// ChatroomBarragerData 弹幕消息数据
type ChatroomBarragerData struct {
	UserName      string `json:"userName"`
	UserNickname  string `json:"userNickname"`
	UserAvatarURL string `json:"userAvatarURL"`
	Content       string `json:"content"`
}
