package types

// NotificationType 通知类型
type NotificationType string

const (
	NotificationTypePoint       NotificationType = "point"
	NotificationTypeCommented   NotificationType = "commented"
	NotificationTypeReply       NotificationType = "reply"
	NotificationTypeAt          NotificationType = "at"
	NotificationTypeFollowing   NotificationType = "following"
	NotificationTypeBroadcast   NotificationType = "broadcast"
	NotificationTypeSysAnnounce NotificationType = "sys-announce"
)

// DataType 数据类型
type DataType int

const (
	DataTypeArticle                   DataType = 0
	DataTypeComment                   DataType = 1
	DataTypeAt                        DataType = 2
	DataTypeCommented                 DataType = 3
	DataTypeFollowingUser             DataType = 4
	DataTypePointCharge               DataType = 5
	DataTypePointTransfer             DataType = 6
	DataTypePointArticleReward        DataType = 7
	DataTypePointArticleThank         DataType = 8
	DataTypePointCommentThank         DataType = 9
	DataTypePointInviteRegister       DataType = 10
	DataTypePointInviteJoinActivation DataType = 11
	DataTypePointPerfectArticle       DataType = 12
	DataTypePointCommentAccept        DataType = 13
	DataTypePointFollowing            DataType = 14
	DataTypePointFollower             DataType = 15
	DataTypePointExchange             DataType = 16
	DataTypePointAbuseDeduct          DataType = 17
	DataTypePointStickArticle         DataType = 18
	DataTypePointRedPacket            DataType = 19
	DataTypePointTopArticle           DataType = 20
	DataTypePointReport               DataType = 21
	DataTypePointRepublishArticle     DataType = 22
	DataTypePointSellerBuyerAdvance   DataType = 23
	DataTypePointCommentVote          DataType = 24
	DataTypePointArticleVote          DataType = 25
	DataTypePointAppeal               DataType = 26
	DataTypePointPurchaseInvitecode   DataType = 27
	DataTypePointForward              DataType = 28
)

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
	Type string      `json:"type"`
	Data interface{} `json:"data"`
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
