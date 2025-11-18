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
