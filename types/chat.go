package types

// ChatMessageData 私聊消息数据
type ChatMessageData struct {
	OId              string `json:"oId"`
	Content          string `json:"content"`
	Markdown         string `json:"markdown"`
	Preview          string `json:"preview"`
	FromId           string `json:"fromId"`
	ToId             string `json:"toId"`
	FromUserName     string `json:"fromUserName"`
	ToUserName       string `json:"toUserName"`
	SenderUserName   string `json:"senderUserName"`
	ReceiverUserName string `json:"receiverUserName"`
	SenderAvatar     string `json:"senderAvatar"`
	ReceiverAvatar   string `json:"receiverAvatar"`
	Time             int64  `json:"time"`
	UserSession      string `json:"user_session"`
}

// ChatMessage 私聊消息
type ChatMessage struct {
	Type string          `json:"type"`
	Data ChatMessageData `json:"data"`
}
