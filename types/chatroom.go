package types

// ClientType 客户端类型
type ClientType string

const (
	ClientTypeWeb         ClientType = "Web"
	ClientTypePC          ClientType = "PC"
	ClientTypeMobile      ClientType = "Mobile"
	ClientTypeWindows     ClientType = "Windows"
	ClientTypeMacOS       ClientType = "macOS"
	ClientTypeLinux       ClientType = "Linux"
	ClientTypeIOS         ClientType = "iOS"
	ClientTypeAndroid     ClientType = "Android"
	ClientTypeIDEA        ClientType = "IDEA"
	ClientTypeChrome      ClientType = "Chrome"
	ClientTypeEdge        ClientType = "Edge"
	ClientTypeVSCode      ClientType = "VSCode"
	ClientTypePython      ClientType = "Python"
	ClientTypeGolang      ClientType = "Golang"
	ClientTypeRust        ClientType = "Rust"
	ClientTypeHarmony     ClientType = "Harmony"
	ClientTypeCLI         ClientType = "CLI"
	ClientTypeBird        ClientType = "Bird"
	ClientTypeIceNet      ClientType = "IceNet"
	ClientTypeElvesOnline ClientType = "ElvesOnline"
	ClientTypeOther       ClientType = "Other"
)

// ChatRoomSource 聊天室消息来源
type ChatRoomSource struct {
	Client  ClientType `json:"client"`
	Version string     `json:"version"`
}

// ChatContentType 历史消息内容类型
type ChatContentType string

const (
	ChatContentTypeMarkdown ChatContentType = "md"
	ChatContentTypeHTML     ChatContentType = "html"
)

// ChatMessageType 历史消息查询类型
type ChatMessageType int

const (
	ChatMessageTypeContext ChatMessageType = 0 // 前后消息
	ChatMessageTypeBefore  ChatMessageType = 1 // 前面的消息
	ChatMessageTypeAfter   ChatMessageType = 2 // 后面的消息
)

// ChatroomMsgType 聊天室消息类型
type ChatroomMsgType string

const (
	ChatroomMsgTypeOnline          ChatroomMsgType = "online"
	ChatroomMsgTypeDiscussChanged  ChatroomMsgType = "discussChanged"
	ChatroomMsgTypeRevoke          ChatroomMsgType = "revoke"
	ChatroomMsgTypeMsg             ChatroomMsgType = "msg"
	ChatroomMsgTypeRedPacket       ChatroomMsgType = "redPacket"
	ChatroomMsgTypeRedPacketStatus ChatroomMsgType = "redPacketStatus"
	ChatroomMsgTypeCustomMessage   ChatroomMsgType = "customMessage"
	ChatroomMsgTypeBarrager        ChatroomMsgType = "barrager"
)

// ChatroomRedPacketType 聊天室红包类型
type ChatroomRedPacketType string

const (
	ChatroomRedPacketTypeRandom            ChatroomRedPacketType = "random"
	ChatroomRedPacketTypeAverage           ChatroomRedPacketType = "average"
	ChatroomRedPacketTypeSpecify           ChatroomRedPacketType = "specify"
	ChatroomRedPacketTypeHeartbeat         ChatroomRedPacketType = "heartbeat"
	ChatroomRedPacketTypeRockPaperScissors ChatroomRedPacketType = "rockPaperScissors"
)

// GestureType 猜拳类型
type GestureType int

const (
	GestureTypeRock     GestureType = 0 // 石头
	GestureTypeScissors GestureType = 1 // 剪刀
	GestureTypePaper    GestureType = 2 // 布
)

// ChatroomNodeInfo 聊天室节点信息
type ChatroomNodeInfo struct {
	Name   string `json:"name"`
	Node   string `json:"node"`
	Online int    `json:"online"`
}

// GetChatroomNodeGetResponse 获取聊天室节点响应
type GetChatroomNodeGetResponse struct {
	Code      int                 `json:"code"`
	Msg       string              `json:"msg"`
	Available []*ChatroomNodeInfo `json:"avaliable"`
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

// ChatroomMsgData 聊天消息数据
type ChatroomMsgData struct {
	OId           string `json:"oId"`
	UserName      string `json:"userName"`
	UserNickname  string `json:"userNickname"`
	UserAvatarURL string `json:"userAvatarURL"`
	Time          int64  `json:"time"`
	Content       string `json:"content"`
	Md            string `json:"md"`
	UserHomeURL   string `json:"userHomeURL"`
	SysMetal      string `json:"sysMetal"`
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
