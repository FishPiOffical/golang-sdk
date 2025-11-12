package types

// NotificationCount 通知数量
type NotificationCount struct {
	UserNotifyStatus                   bool `json:"userNotifyStatus"`
	UnreadNotificationCount            int  `json:"unreadNotificationCnt"`
	UnreadReplyNotificationCount       int  `json:"unreadReplyNotificationCnt"`
	UnreadPointNotificationCount       int  `json:"unreadPointNotificationCnt"`
	UnreadAtNotificationCount          int  `json:"unreadAtNotificationCnt"`
	UnreadBroadcastNotificationCount   int  `json:"unreadBroadcastNotificationCnt"`
	UnreadSysAnnounceNotificationCount int  `json:"unreadSysAnnounceNotificationCnt"`
	UnreadNewFollowerNotificationCount int  `json:"unreadNewFollowerNotificationCnt"`
	UnreadFollowingNotificationCount   int  `json:"unreadFollowingNotificationCnt"`
	UnreadCommentedNotificationCount   int  `json:"unreadCommentedNotificationCnt"`
}

// NotificationInfo 通知信息
type NotificationInfo struct {
	OId                string   `json:"oId"`
	DataId             string   `json:"dataId"`
	DataType           DataType `json:"dataType"`
	Description        string   `json:"description"`
	HasRead            bool     `json:"hasRead"`
	CreateTime         int64    `json:"createTime"`
	AuthorName         string   `json:"authorName"`
	AuthorThumbnailURL string   `json:"authorThumbnailURL"`
	ArticleTitle       string   `json:"articleTitle,omitempty"`
	ArticleTags        string   `json:"articleTags,omitempty"`
	ArticleType        int      `json:"articleType,omitempty"`
	ArticlePermalink   string   `json:"articlePermalink,omitempty"`
}

// NotificationList 通知列表
type NotificationList struct {
	Notifications []NotificationInfo `json:"list"`
	Pagination    Pagination         `json:"pagination"`
}

// UserMessage 用户频道消息
type UserMessage struct {
	Type    string      `json:"type"`
	Command string      `json:"command"`
	Data    interface{} `json:"data"`
}

// BreezemoonData 清风明月通知数据
type BreezemoonData struct {
	OId                          string `json:"oId"`
	BreezemoonContent            string `json:"breezemoonContent"`
	BreezemoonAuthorName         string `json:"breezemoonAuthorName"`
	BreezemoonAuthorThumbnailURL string `json:"breezemoonAuthorThumbnailURL"`
	BreezemoonCreateTime         int64  `json:"breezemoonCreateTime"`
	BreezemoonUpdated            int64  `json:"breezemoonUpdated"`
	BreezemoonCity               string `json:"breezemoonCity"`
}

// ArticleData 文章通知数据
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
