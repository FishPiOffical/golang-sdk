package types

type ChatMessage struct {
	// GET /chat/has-unread 接口返回字段
	ToId             string `json:"toId"`
	UserSession      string `json:"user_session"`
	OId              string `json:"oId"`
	FromId           string `json:"fromId"`
	SenderUserName   string `json:"senderUserName"`
	ReceiverUserName string `json:"receiverUserName"`

	// 更多通用字段
	Preview        string `json:"preview,omitempty"`
	SenderAvatar   string `json:"senderAvatar,omitempty"`
	Markdown       string `json:"markdown,omitempty"`
	ReceiverAvatar string `json:"receiverAvatar,omitempty"`
	Time           string `json:"time,omitempty"`
	Content        string `json:"content,omitempty"`

	ReceiverOnlineFlag bool `json:"receiverOnlineFlag,omitempty"` // GET /chat/get-list 接口返回
}

type GetChatGetListResponse struct {
	Result int            `json:"result"`
	Data   []*ChatMessage `json:"data"`
	Cached bool           `json:"cached"`
}
