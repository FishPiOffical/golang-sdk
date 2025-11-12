package fishPiSdk

import "time"

// ApiResponse 通用API响应结构(带Data字段)
type ApiResponse[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

// SimpleResponse 简单响应结构(无Data字段)
type SimpleResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// ResponseResult 特殊响应结构(使用sc字段而非code)
type ResponseResult[T any] struct {
	Result int    `json:"sc"`
	Msg    string `json:"msg"`
	Data   T      `json:"data"`
}

// PostApiGetKeyRequest 获取ApiKey请求
type PostApiGetKeyRequest struct {
	NameOrEmail  string `json:"nameOrEmail"`
	UserPassword string `json:"userPassword"`
	MfaCode      string `json:"mfaCode,omitempty"`
}

// PostApiGetKeyResponse 获取ApiKey响应
type PostApiGetKeyResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Key  string `json:"Key"`
}

// ApiUserGetData 获取用户信息数据
type ApiUserGetData struct {
	UserNo               int       `json:"userNo"`
	UserName             string    `json:"userName"`
	UserNickname         string    `json:"userNickname"`
	UserAvatarURL        string    `json:"userAvatarURL"`
	UserURL              string    `json:"userURL"`
	UserIntro            string    `json:"userIntro"`
	UserPoint            int       `json:"userPoint"`
	UserAppRole          int       `json:"userAppRole"`
	UserRole             string    `json:"userRole"`
	OnlineMinute         int       `json:"onlineMinute"`
	UserCity             string    `json:"userCity"`
	SysMetal             string    `json:"sysMetal"`
	CardBg               string    `json:"cardBg"`
	FollowingUserCount   int       `json:"followingUserCount"`
	FollowerCount        int       `json:"followerCount"`
	CanFollow            string    `json:"canFollow"`
	OnlineFlag           bool      `json:"onlineFlag"`
	UserCreateTime       time.Time `json:"userCreateTime"`
	UserCreateTimeStr    string    `json:"userCreateTimeStr"`
	UserUpdateTime       time.Time `json:"userUpdateTime"`
	UserUpdateTimeStr    string    `json:"userUpdateTimeStr"`
	UserCheckedIn        bool      `json:"checkedIn"`
	CheckedInToday       bool      `json:"checkedInToday"`
	CurrentCheckedInDays int       `json:"currentCheckedInDays"`
	LongestCheckedInDays int       `json:"longestCheckedInDays"`
}

// GetChatGetMessageData 获取私聊消息数据
type GetChatGetMessageData struct {
	OId          string `json:"oId"`
	Content      string `json:"content"`
	FromUserName string `json:"fromUserName"`
	ToUserName   string `json:"toUserName"`
	Time         int64  `json:"time"`
}

// GetChatHasUnreadData 获取私聊未读消息数据
type GetChatHasUnreadData struct {
	FromUserName      string `json:"fromUserName"`
	FromUserNickname  string `json:"fromUserNickname"`
	FromUserAvatarURL string `json:"fromUserAvatarURL"`
	FromUserURL       string `json:"fromUserURL"`
	UnreadCount       int    `json:"unreadChatCnt"`
}

// GetChatGetListData 获取私聊列表数据
type GetChatGetListData struct {
	UserName    string `json:"userName"`
	Nickname    string `json:"nickname"`
	AvatarURL   string `json:"avatarURL"`
	UnreadCount int    `json:"unreadChatCnt"`
	LastMsg     string `json:"lastMsg"`
	LastMsgTime int64  `json:"lastMsgTime"`
}

// GetApiGetNotificationsData 获取通知列表数据
type GetApiGetNotificationsData struct {
	DataId             string `json:"dataId"`
	DataType           int    `json:"dataType"`
	Description        string `json:"description"`
	HasRead            bool   `json:"hasRead"`
	CreateTime         int64  `json:"createTime"`
	AuthorName         string `json:"authorName"`
	AuthorThumbnailURL string `json:"authorThumbnailURL"`
}

// GetChatroomNodeGetResponse 获取聊天室节点响应
type GetChatroomNodeGetResponse struct {
	Code      int                 `json:"code"`
	Msg       string              `json:"msg"`
	Avaliable []*ChatroomNodeInfo `json:"avaliable"`
}

// ChatroomNodeInfo 聊天室节点信息
type ChatroomNodeInfo struct {
	Name   string `json:"name"`
	Node   string `json:"node"`
	Online int    `json:"online"`
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

// ChatroomMessage 聊天室消息
type ChatroomMessage struct {
	Type ChatroomMsgType `json:"type"`
	Data interface{}     `json:"data"`
	Time int64           `json:"time"`
}

// ChatroomOnlineData 在线用户数据
type ChatroomOnlineData struct {
	Users []OnlineUser `json:"users"`
}

// OnlineUser 在线用户
type OnlineUser struct {
	UserName      string `json:"userName"`
	UserNickname  string `json:"userNickname"`
	UserAvatarURL string `json:"userAvatarURL"`
	HomepageURL   string `json:"homepageURL"`
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

// ChatroomRedPacketData 红包消息数据
type ChatroomRedPacketData struct {
	ChatroomMsgData
	RedPacket RedPacketInfo `json:"redPacket"`
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

// ChatMessage 私聊消息
type ChatMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// ChatMessageData 私聊消息数据
type ChatMessageData struct {
	OId          string `json:"oId"`
	Content      string `json:"content"`
	FromUserName string `json:"fromUserName"`
	ToUserName   string `json:"toUserName"`
	Time         int64  `json:"time"`
}

// UserMessage 用户频道消息
type UserMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// BreezemoonData 清风明月数据
type BreezemoonData struct {
	OId                          string `json:"oId"`
	BreezemoonContent            string `json:"breezemoonContent"`
	BreezemoonAuthorName         string `json:"breezemoonAuthorName"`
	BreezemoonAuthorThumbnailURL string `json:"breezemoonAuthorThumbnailURL"`
	BreezemoonCreateTime         int64  `json:"breezemoonCreateTime"`
	BreezemoonUpdated            int64  `json:"breezemoonUpdated"`
	BreezemoonCity               string `json:"breezemoonCity"`
}

// ArticleData 文章数据
type ArticleData struct {
	OId                       string `json:"oId"`
	ArticleTitle              string `json:"articleTitle"`
	ArticleContent            string `json:"articleContent"`
	ArticleTags               string `json:"articleTags"`
	ArticleAuthorName         string `json:"articleAuthorName"`
	ArticleAuthorThumbnailURL string `json:"articleAuthorThumbnailURL"`
	ArticleCreateTime         int64  `json:"articleCreateTime"`
	ArticleUpdateTime         int64  `json:"articleUpdateTime"`
	ArticleViewCount          int    `json:"articleViewCnt"`
	ArticleCommentCount       int    `json:"articleCommentCnt"`
	ArticleThankCount         int    `json:"articleThankCnt"`
	ArticleGoodCount          int    `json:"articleGoodCnt"`
	ArticleBadCount           int    `json:"articleBadCnt"`
	ArticleCollectCount       int    `json:"articleCollectCnt"`
	ArticleWatchCount         int    `json:"articleWatchCnt"`
	ArticlePermalink          string `json:"articlePermalink"`
}

// PostArticleRequest 发布文章请求
type PostArticleRequest struct {
	ArticleTitle         string `json:"articleTitle"`
	ArticleContent       string `json:"articleContent"`
	ArticleTags          string `json:"articleTags"`
	ArticleCommentable   bool   `json:"articleCommentable"`
	ArticleType          int    `json:"articleType"`
	ArticleRewardContent string `json:"articleRewardContent,omitempty"`
	ArticleRewardPoint   int    `json:"articleRewardPoint,omitempty"`
	ArticleQnAOfferPoint int    `json:"articleQnAOfferPoint,omitempty"`
}

// PostBreezemoonRequest 发布清风明月请求
type PostBreezemoonRequest struct {
	BreezemoonContent string `json:"breezemoonContent"`
}

// PostCommentRequest 发布评论请求
type PostCommentRequest struct {
	ArticleId      string `json:"articleId"`
	CommentContent string `json:"commentContent"`
}

// TransferRequest 转账请求
type TransferRequest struct {
	UserName string `json:"userName"`
	Amount   int    `json:"amount"`
	Memo     string `json:"memo"`
}
