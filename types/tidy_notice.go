package types

// GetNotificationCountResponse 通知数量
type GetNotificationCountResponse struct {
	Code                               int `json:"code"`
	UserNotifyStatus                   int `json:"userNotifyStatus"`
	UnreadNotificationCount            int `json:"unreadNotificationCnt"`
	UnreadReplyNotificationCount       int `json:"unreadReplyNotificationCnt"`
	UnreadPointNotificationCount       int `json:"unreadPointNotificationCnt"`
	UnreadAtNotificationCount          int `json:"unreadAtNotificationCnt"`
	UnreadBroadcastNotificationCount   int `json:"unreadBroadcastNotificationCnt"`
	UnreadSysAnnounceNotificationCount int `json:"unreadSysAnnounceNotificationCnt"`
	UnreadNewFollowerNotificationCount int `json:"unreadNewFollowerNotificationCnt"`
	UnreadFollowingNotificationCount   int `json:"unreadFollowingNotificationCnt"`
	UnreadCommentedNotificationCount   int `json:"unreadCommentedNotificationCnt"`
}

// NotificationInfo 通知信息
type NotificationInfo struct {
	HasRead bool `json:"hasRead"`

	// NotificationTypePoint NotificationTypeAt NotificationTypeSysAnnounce NotificationTypeFollowing
	CreateTime string `json:"createTime,omitempty"`

	// NotificationTypePoint NotificationTypeSysAnnounce
	Description string `json:"description,omitempty"`

	// NotificationTypeCommented NotificationTypeReply
	CommentAuthorName         string         `json:"commentAuthorName,omitempty"`
	CommentAuthorThumbnailURL string         `json:"commentAuthorThumbnailURL,omitempty"`
	CommentCreateTime         string         `json:"commentCreateTime,omitempty"`
	CommentSharpURL           string         `json:"commentSharpURL,omitempty"`
	CommentContent            string         `json:"commentContent,omitempty"`
	CommentArticleType        ArticleType    `json:"commentArticleType,omitempty"`
	CommentArticleTitle       string         `json:"commentArticleTitle,omitempty"`
	CommentArticlePerfect     ArticlePerfect `json:"commentArticlePerfect,omitempty"`

	// NotificationTypeAt
	UserName      string `json:"userName,omitempty"`
	UserAvatarURL string `json:"userAvatarURL,omitempty"`
	Deleted       bool   `json:"deleted,omitempty"`
	Content       string `json:"content,omitempty"`

	// NotificationTypeFollowing
	ArticleTitle        string         `json:"articleTitle,omitempty"`
	IsComment           bool           `json:"isComment,omitempty"`
	ArticleTags         string         `json:"articleTags,omitempty"`
	Url                 string         `json:"url,omitempty"`
	ArticleType         ArticleType    `json:"articleType,omitempty"`
	AuthorName          string         `json:"authorName,omitempty"`
	ArticlePerfect      ArticlePerfect `json:"articlePerfect,omitempty"`
	ThumbnailURL        string         `json:"thumbnailURL,omitempty"`
	ArticleCommentCount int            `json:"articleCommentCnt,omitempty"`

	// NotificationTypeSysAnnounce
	OId      string   `json:"oId,omitempty"`
	UserId   string   `json:"userId,omitempty"`
	DataId   string   `json:"dataId,omitempty"`
	DataType DataType `json:"dataType,omitempty"`
}
