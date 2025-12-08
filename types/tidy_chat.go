package types

type ChatMessage struct {
	ToId             string `json:"toId"`
	Preview          string `json:"preview"`
	UserSession      string `json:"user_session"`
	SenderAvatar     string `json:"senderAvatar"`
	Markdown         string `json:"markdown"`
	ReceiverAvatar   string `json:"receiverAvatar"`
	OId              string `json:"oId"`
	Time             string `json:"time"`
	FromId           string `json:"fromId"`
	SenderUserName   string `json:"senderUserName"`
	Content          string `json:"content"`
	ReceiverUserName string `json:"receiverUserName"`

	ReceiverOnlineFlag bool `json:"receiverOnlineFlag,omitempty"` // GET /chat/get-list 接口返回
}

type GetChatGetListResponse struct {
	Result int            `json:"result"`
	Data   []*ChatMessage `json:"data"`
	Cached bool           `json:"cached"`
}
