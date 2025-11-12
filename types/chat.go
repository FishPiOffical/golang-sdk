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

// ChatRevoke 私聊撤回消息
type ChatRevoke struct {
	Type string `json:"type"` // "revoke"
	Data string `json:"data"` // message oId
}

// GetChatListData 获取私聊列表数据
type GetChatListData struct {
	UserName    string `json:"userName"`
	Nickname    string `json:"nickname"`
	AvatarURL   string `json:"avatarURL"`
	UnreadCount int    `json:"unreadChatCnt"`
	LastMsg     string `json:"lastMsg"`
	LastMsgTime int64  `json:"lastMsgTime"`
}

// GetChatHasUnreadData 获取私聊未读消息数据
type GetChatHasUnreadData struct {
	FromUserName      string `json:"fromUserName"`
	FromUserNickname  string `json:"fromUserNickname"`
	FromUserAvatarURL string `json:"fromUserAvatarURL"`
	FromUserURL       string `json:"fromUserURL"`
	UnreadCount       int    `json:"unreadChatCnt"`
}

// PostChatSendRequest 发送私聊请求
type PostChatSendRequest struct {
	ToUserName string `json:"toUserName"`
	Content    string `json:"content"`
}

// ChatQuery 私聊查询参数
type ChatQuery struct {
	Page     int  `json:"page,omitempty"`
	Size     int  `json:"size,omitempty"`
	AutoRead bool `json:"autoRead,omitempty"`
}

// ChatListItem 私聊列表项
type ChatListItem struct {
	UserName      string `json:"userName"`
	UserNickname  string `json:"userNickname"`
	UserAvatarURL string `json:"userAvatarURL"`
	LastMessage   string `json:"lastMessage"`
	LastTime      int64  `json:"lastTime"`
	UnreadCount   int    `json:"unreadCount"`
}
