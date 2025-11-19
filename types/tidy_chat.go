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
	OId              string `json:"oId"`
	UserName         string `json:"userName"`
	UserNickname     string `json:"userNickname"`
	UserAvatarURL    string `json:"userAvatarURL"`
	Time             string `json:"time"`
	Content          string `json:"content"`
	Md               string `json:"md"`
	UserHomeURL      string `json:"userHomeURL"`
	SysMetal         string `json:"sysMetal"`
	UserOId          int64  `json:"userOId"`
	Type             string `json:"type"`
	UserAvatarURL20  string `json:"userAvatarURL20"`
	Client           string `json:"client"`
	UserAvatarURL210 string `json:"userAvatarURL210"`
	UserAvatarURL48  string `json:"userAvatarURL48"`
}
