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
