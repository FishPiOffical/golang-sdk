package types

// GetChatroomNodeGetResponse 获取聊天室节点响应
type GetChatroomNodeGetResponse struct {
	Code      int                 `json:"code"`
	Msg       string              `json:"msg"`
	Data      string              `json:"data"`
	ApiKey    string              `json:"apiKey"`
	Available []*ChatroomNodeInfo `json:"avaliable"`
}

// ChatroomNodeInfo 聊天室节点信息
type ChatroomNodeInfo struct {
	Name   string `json:"name"`
	Node   string `json:"node"`
	Online int    `json:"online"`
	Weight int    `json:"weight"`
}

// ChatroomMsgData 聊天消息数据
type ChatroomMsgData struct {
	UserOId       int64  `json:"userOId"`
	UserAvatarURL string `json:"userAvatarURL"`
	UserNickname  string `json:"userNickname"`
	SysMetal      string `json:"sysMetal"`
	Client        string `json:"client"`
	Time          string `json:"time"`
	OId           string `json:"oId"`
	UserName      string `json:"userName"`
	Content       string `json:"content"`

	Md               string `json:"md,omitempty"`
	UserHomeURL      string `json:"userHomeURL,omitempty"`
	Type             string `json:"type,omitempty"`
	UserAvatarURL20  string `json:"userAvatarURL20,omitempty"`
	UserAvatarURL210 string `json:"userAvatarURL210,omitempty"`
	UserAvatarURL48  string `json:"userAvatarURL48,omitempty"`
}

// PostChatroomRedPacketOpenResponse 打开聊天室红包响应
type PostChatroomRedPacketOpenResponse struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`

	Recivers []any           `json:"recivers,omitempty"`
	Who      []*RedPacketWho `json:"who,omitempty"`
	Info     *RedPacketInfo  `json:"info,omitempty"`
}

type RedPacketWho struct {
	UserMoney int    `json:"userMoney"`
	Time      string `json:"time"`
	Avatar    string `json:"avatar"`
	UserName  string `json:"userName"`
	UserId    string `json:"userId"`
}

type RedPacketInfo struct {
	Msg              string      `json:"msg"`
	UserAvatarURL    string      `json:"userAvatarURL"`
	UserAvatarURL20  string      `json:"userAvatarURL20"`
	Count            int         `json:"count"`
	UserName         string      `json:"userName"`
	UserAvatarURL210 string      `json:"userAvatarURL210"`
	Got              int         `json:"got"`
	UserAvatarURL48  string      `json:"userAvatarURL48"`
	Gesture          GestureType `json:"gesture,omitempty"`
}
