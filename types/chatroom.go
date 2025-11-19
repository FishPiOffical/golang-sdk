package types

// ChatRoomSource 聊天室消息来源
type ChatRoomSource struct {
	Client  ClientType `json:"client"`
	Version string     `json:"version"`
}

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

// RedPacketInfo 红包信息
type RedPacketInfo struct {
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
	RedPacket RedPacketInfo `json:"redPacket"`
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

// MuteItem 禁言项
type MuteItem struct {
	OId              string `json:"oId"`
	UserName         string `json:"userName"`
	UserNickname     string `json:"userNickname"`
	UserAvatarURL    string `json:"userAvatarURL"`
	MuteTime         int64  `json:"muteTime"`
	MuteEndTime      int64  `json:"muteEndTime"`
	MuteReason       string `json:"muteReason"`
	MuteOperatorName string `json:"muteOperatorName"`
}

// IMusicMessage 音乐消息
type IMusicMessage struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	URL    string `json:"url"`
	Cover  string `json:"cover"`
}

// IWeatherMessage 天气消息
type IWeatherMessage struct {
	City        string `json:"city"`
	Temperature string `json:"temperature"`
	Weather     string `json:"weather"`
	AQI         string `json:"aqi"`
}

// PostChatroomRedPacketOpenRequest 打开聊天室红包请求
type PostChatroomRedPacketOpenRequest struct {
	OId     string      `json:"oId"`
	Gesture GestureType `json:"gesture,omitempty"`
}

// PostChatroomRedPacketOpenResponse 打开聊天室红包响应
type PostChatroomRedPacketOpenResponse struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Who  []string `json:"who"`
	Got  int      `json:"got"`
	Type string   `json:"type"`
}

// PostChatroomSendRequest 发送聊天室消息请求
type PostChatroomSendRequest struct {
	Content string `json:"content"`
	Client  string `json:"client,omitempty"`
}

// ChatroomMoreRequest 聊天室历史消息请求
type ChatroomMoreRequest struct {
	Page int             `json:"page,omitempty"`
	Size int             `json:"size,omitempty"`
	Mode ChatContentType `json:"mode,omitempty"`
}

// ChatroomContextRequest 聊天室上下文消息请求
type ChatroomContextRequest struct {
	OId  string          `json:"oId"`
	Size int             `json:"size,omitempty"`
	Type ChatMessageType `json:"type,omitempty"`
	Mode ChatContentType `json:"mode,omitempty"`
}
