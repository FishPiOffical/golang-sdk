package types

import "encoding/json"

type UserChannelMsg struct {
	Command UserChannelCommand      `json:"command"`
	Bz      *UserChannelMsgBzUpdate `json:"bz,omitempty"`

	// UserChannelCommandRefreshNotification UserChannelCommandChatUnreadCountRefresh
	Count int `json:"count,omitempty"`

	// UserChannelCommandRefreshNotification UserChannelCommandChatUnreadCountRefresh UserChannelCommandNewIdleChatMessage
	UserId string `json:"userId,omitempty"`

	// UserChannelCommandNewIdleChatMessage
	SendUserName string `json:"sendUserName,omitempty"`
	SendAvatar   string `json:"sendAvatar,omitempty"`
	Preview      string `json:"preview,omitempty"`

	// UserChannelCommandWarnBroadcast
	WarnBroadcastText string `json:"warnBroadcastText,omitempty"`
	Who               string `json:"who,omitempty"`
}

type UserChannelMsgBzUpdate struct {
	OId                            string `json:"oId"`
	BreezemoonContent              string `json:"breezemoonContent"`
	BreezemoonAuthorName           string `json:"breezemoonAuthorName"`
	BreezemoonAuthorThumbnailURL48 string `json:"breezemoonAuthorThumbnailURL48"`
}

type ChatChannelMsg struct {
	ToId             string `json:"toId"`
	Preview          string `json:"preview"`
	UserSession      string `json:"user_session"`
	Markdown         string `json:"markdown"`
	SenderAvatar     string `json:"senderAvatar"`
	ReceiverAvatar   string `json:"receiverAvatar"`
	OId              string `json:"oId"`
	Time             string `json:"time"`
	FromId           string `json:"fromId"`
	SenderUserName   string `json:"senderUserName"`
	Content          string `json:"content"`
	ReceiverUserName string `json:"receiverUserName"`
}

type ChatroomMsg struct {
	Type ChatroomMsgType `json:"type"` // 消息类型

	// 在线消息
	Discussing    string            `json:"discussing"`    // 讨论的话题
	OnlineChatCnt int               `json:"onlineChatCnt"` // 在线人数
	Users         []*OnlineUserInfo `json:"users"`         // 在线用户信息

	// 话题变更
	NewDiscuss string `json:"newDiscuss"` // 新的话题内容

	// 聊天 撤回 红包领取
	OId string `json:"oId"` // 消息ID

	// 聊天消息
	Time                string             `json:"time"`                // 发布时间
	UserOId             int64              `json:"userOId"`             // 用户ID
	UserName            string             `json:"userName"`            // 用户名
	UserNickname        string             `json:"userNickname"`        // 用户昵称
	UserAvatarURL       string             `json:"userAvatarURL"`       // 用户头像
	UserAvatarURL20     string             `json:"userAvatarURL20"`     // 用户头像 20px
	UserAvatarURL48     string             `json:"userAvatarURL48"`     // 用户头像 48px
	UserAvatarURL210    string             `json:"userAvatarURL210"`    // 用户邮箱 210px
	SysMetal            string             `json:"sysMetal"`            // 徽章数据 json字符串
	Content             string             `json:"content"`             // 消息内容 HTML格式 如果是红包则是JSON格式
	Md                  string             `json:"md"`                  // 消息内容 Markdown格式，红包消息无此栏位
	ReactionSummary     []*ReactionSummary `json:"reactionSummary"`     // 表情反应汇总
	CurrentUserReaction string             `json:"currentUserReaction"` // 自己在该信息中的表情反应

	// 红包领取消息
	Count   int    `json:"count"`   // 红包个数
	Got     int    `json:"got"`     // 已领取个数
	WhoGive string `json:"whoGive"` // 发送者用户名
	WhoGot  string `json:"whoGot"`  // 领取者用户名

	// 客户端
	Client string `json:"client"` // 消息客户端

	// 普通消息
	Message string `json:"message"` // 普通消息的消息内容

	// 弹幕消息
	BarrageColor   string `json:"barragerColor"`
	BarrageContent string `json:"barragerContent"`

	// 回应消息
	Summary        []*ReactionSummary `json:"summary,omitempty"`        // 回应汇总
	GroupType      ReactionGroupType  `json:"groupType,omitempty"`      // 回应分组类型
	AuthorUserId   string             `json:"authorUserId,omitempty"`   // 消息作者用户ID
	AuthorReaction string             `json:"authorReaction,omitempty"` // 消息作者的回应
}

func (msg *ChatroomMsg) GetMetalList() ([]*Metal, error) {
	if msg.SysMetal == "" {
		return []*Metal{}, nil
	}
	sysMetal := &SysMetal{}
	if err := json.Unmarshal([]byte(msg.SysMetal), sysMetal); err != nil {
		return nil, err
	}
	return sysMetal.List, nil
}

func (msg *ChatroomMsg) GetJsonInfo() *ChatroomMsgJsonInfo {
	data := new(ChatroomMsgJsonInfo)

	if err := json.Unmarshal([]byte(msg.Content), data); err != nil {
		return nil
	}

	return data
}

type ReactionSummary struct {
	Value       ReactionEmojiValue       `json:"value"`       // 表情值
	Emoji       string                   `json:"emoji"`       // 表情字符
	Count       int                      `json:"count"`       // 该表情当前总数
	Selected    bool                     `json:"selected"`    // 当前登陆用户是否已选中该表情
	Users       []string                 `json:"users"`       // 点赞过该表情的用户显示名列表
	UserDetails []*ReactionSummaryDetail `json:"userDetails"` // 点赞过该表情的用户详情列表
}

type ReactionSummaryDetail struct {
	UserName    string `json:"userName"`
	DisplayName string `json:"displayName"`
	AvatarURL   string `json:"avatarURL"`
}

type OnlineUserInfo struct {
	UserName         string `json:"userName"`         // 用户名
	HomePage         string `json:"homePage"`         // 用户首页
	UserAvatarURL    string `json:"userAvatarURL"`    // 用户头像
	UserAvatarURL20  string `json:"userAvatarURL20"`  // 用户头像 20px
	UserAvatarURL48  string `json:"userAvatarURL48"`  // 用户头像 48px
	UserAvatarURL210 string `json:"userAvatarURL210"` // 用户邮箱 210px
}

type SysMetal struct {
	List []*Metal `json:"list"` // 徽章列表数据
}

// ChatroomMsgJsonInfo json的数据结构
type ChatroomMsgJsonInfo struct {
	// 消息类型 redPacket-红包 weather-天气 music-音乐
	MsgType ChatroomMsgJsonType `json:"msgType"`
	/*
		红包类型 random(拼手气红包), average(平分红包)，specify(专属红包)，heartbeat(心跳红包)，rockPaperScissors(猜拳红包)
		天气类型 weather
		音乐类型 music
	*/
	Type ChatroomMsgJsonSubType `json:"type"`

	Msg      string        `json:"msg"`      // 红包祝福语
	Recivers string        `json:"recivers"` // 红包接收者用户名，专属红包有效
	SenderId string        `json:"senderId"` // 发送者id
	Money    int           `json:"money"`    // 红包积分
	Count    int           `json:"count"`    // 红包个数
	Got      int           `json:"got"`      // 已领取个数
	Who      []interface{} `json:"who"`      // 已领取者信息

	Date        string `json:"date"`        // 日期
	St          string `json:"st"`          // 一句话
	Min         string `json:"min"`         // 最低温度
	T           string `json:"t"`           // 城市
	Max         string `json:"max"`         // 最高温度
	WeatherCode string `json:"weatherCode"` // 天气

	CoverURL string `json:"coverURL"` // 封面链接
	From     string `json:"from"`     // 来源
	Source   string `json:"source"`   // 音乐链接
	Title    string `json:"title"`    // 音乐名
}

type ArticleChannelMsg struct {
	ArticleId string             `json:"articleId"`
	Type      ArticleChannelType `json:"type"`

	// ArticleChannelTypeArticleHeat 相关参数
	Operation ArticleChannelOperation `json:"operation,omitempty"`

	// ArticleChannelTypeArticleReaction ArticleChannelTypeCommentReaction 相关参数
	TargetId      string             `json:"targetId,omitempty"`      // 目标对象oId
	TargetType    ReactionTargetType `json:"targetType,omitempty"`    // 目标对象类型
	GroupType     ReactionGroupType  `json:"groupType,omitempty"`     // 表情分组类型
	Summary       []*ReactionSummary `json:"summary,omitempty"`       // 表情反应汇总
	ActorUserId   string             `json:"actorUserId,omitempty"`   // 触发操作的用户ID
	ActorReaction string             `json:"actorReaction,omitempty"` // 触发操作的用户的表情反应

	// ArticleChannelTypeComment 相关参数
	CommentNice               bool              `json:"commentNice,omitempty"`
	CommentCreateTimeStr      string            `json:"commentCreateTimeStr,omitempty"`
	ReactionSummary           []interface{}     `json:"reactionSummary,omitempty"`
	CommentAuthorId           string            `json:"commentAuthorId,omitempty"`
	CommentUA                 string            `json:"commentUA,omitempty"`
	CommentScore              int               `json:"commentScore,omitempty"`
	CommentCreateTime         int64             `json:"commentCreateTime,omitempty"`
	CommentVote               int               `json:"commentVote,omitempty"`
	CommentRevisionCount      int               `json:"commentRevisionCount,omitempty"`
	TimeAgo                   string            `json:"timeAgo,omitempty"`
	CommentOriginalCommentId  string            `json:"commentOriginalCommentId,omitempty"`
	SysMetal                  []*Metal          `json:"sysMetal,omitempty"`
	CurrentUserReaction       string            `json:"currentUserReaction,omitempty"`
	CommentGoodCnt            int               `json:"commentGoodCnt,omitempty"`
	CommentVisible            int               `json:"commentVisible,omitempty"`
	CommentOnArticleId        string            `json:"commentOnArticleId,omitempty"`
	RewardedCnt               int               `json:"rewardedCnt,omitempty"`
	CommentThankLabel         string            `json:"commentThankLabel,omitempty"`
	CommentSharpURL           string            `json:"commentSharpURL,omitempty"`
	CommentAnonymous          int               `json:"commentAnonymous,omitempty"`
	CmtTpl                    string            `json:"cmtTpl,omitempty"`
	CommentIP                 string            `json:"commentIP,omitempty"`
	CommentReplyCnt           int               `json:"commentReplyCnt,omitempty"`
	OId                       string            `json:"oId,omitempty"`
	CommentContent            string            `json:"commentContent,omitempty"`
	Article                   *ArticleInfo      `json:"article,omitempty"`
	CommentStatus             int               `json:"commentStatus,omitempty"`
	Commenter                 *ArticleCommenter `json:"commenter,omitempty"`
	CommentAuthorName         string            `json:"commentAuthorName,omitempty"`
	CommentThankCnt           int               `json:"commentThankCnt,omitempty"`
	CommentBadCnt             int               `json:"commentBadCnt,omitempty"`
	Rewarded                  bool              `json:"rewarded,omitempty"`
	CommentId                 string            `json:"commentId,omitempty"`
	CommentAuthorThumbnailURL string            `json:"commentAuthorThumbnailURL,omitempty"`
	CommentAudioURL           string            `json:"commentAudioURL,omitempty"`
	CommentQnAOffered         int               `json:"commentQnAOffered,omitempty"`
	CommentAuthorNickName     string            `json:"commentAuthorNickName,omitempty"`
}

type T struct {
	CommentNice              bool          `json:"commentNice"`
	CommentCreateTimeStr     string        `json:"commentCreateTimeStr"`
	ReactionSummary          []interface{} `json:"reactionSummary"`
	CommentAuthorId          string        `json:"commentAuthorId"`
	CommentUA                string        `json:"commentUA"`
	CommentScore             int           `json:"commentScore"`
	CommentCreateTime        int64         `json:"commentCreateTime"`
	CommentVote              int           `json:"commentVote"`
	Type                     string        `json:"type"`
	CommentRevisionCount     int           `json:"commentRevisionCount"`
	TimeAgo                  string        `json:"timeAgo"`
	CommentOriginalCommentId string        `json:"commentOriginalCommentId"`
	SysMetal                 []struct {
		Data        string `json:"data"`
		Name        string `json:"name"`
		Description string `json:"description"`
		ExpireDate  string `json:"expireDate"`
		Id          string `json:"id"`
		Attr        string `json:"attr"`
		Type        string `json:"type"`
		Enabled     bool   `json:"enabled"`
		Order       int    `json:"order"`
	} `json:"sysMetal"`
	CurrentUserReaction string `json:"currentUserReaction"`
	CommentGoodCnt      int    `json:"commentGoodCnt"`
	CommentVisible      int    `json:"commentVisible"`
	CommentOnArticleId  string `json:"commentOnArticleId"`
	RewardedCnt         int    `json:"rewardedCnt"`
	CommentThankLabel   string `json:"commentThankLabel"`
	CommentSharpURL     string `json:"commentSharpURL"`
	CommentAnonymous    int    `json:"commentAnonymous"`
	CmtTpl              string `json:"cmtTpl"`
	CommentIP           string `json:"commentIP"`
	CommentReplyCnt     int    `json:"commentReplyCnt"`
	ArticleId           string `json:"articleId"`
	OId                 string `json:"oId"`
	CommentContent      string `json:"commentContent"`
	Article             struct {
		ArticleShowInList      int     `json:"articleShowInList"`
		ArticleCreateTime      int64   `json:"articleCreateTime"`
		ArticleIP              string  `json:"articleIP"`
		ArticleEditorType      int     `json:"articleEditorType"`
		ArticleRandomDouble    float64 `json:"articleRandomDouble"`
		ArticleAuthorId        string  `json:"articleAuthorId"`
		ArticleBadCnt          int     `json:"articleBadCnt"`
		ArticleRewardPoint     int     `json:"articleRewardPoint"`
		ArticleLatestCmtTime   int64   `json:"articleLatestCmtTime"`
		ArticleGoodCnt         int     `json:"articleGoodCnt"`
		ArticleQnAOfferPoint   int     `json:"articleQnAOfferPoint"`
		ArticleStatement       int     `json:"articleStatement"`
		ArticleType            int     `json:"articleType"`
		Offered                bool    `json:"offered"`
		ArticleViewCount       int     `json:"articleViewCount"`
		ArticleCommentable     bool    `json:"articleCommentable"`
		ArticleWatchCnt        int     `json:"articleWatchCnt"`
		ArticleContent         string  `json:"articleContent"`
		ArticleUA              string  `json:"articleUA"`
		ArticleAudioURL        string  `json:"articleAudioURL"`
		ArticleCommentCount    int     `json:"articleCommentCount"`
		ArticleImg1URL         string  `json:"articleImg1URL"`
		ArticlePushOrder       int     `json:"articlePushOrder"`
		ArticleCollectCnt      int     `json:"articleCollectCnt"`
		ArticleTitle           string  `json:"articleTitle"`
		ArticleLatestCmterName string  `json:"articleLatestCmterName"`
		ArticleAnonymousView   int     `json:"articleAnonymousView"`
		ArticleTags            string  `json:"articleTags"`
		OId                    string  `json:"oId"`
		ArticleStick           int     `json:"articleStick"`
		ArticleAnonymous       int     `json:"articleAnonymous"`
		ArticleThankCnt        int     `json:"articleThankCnt"`
		ArticleRewardContent   string  `json:"articleRewardContent"`
		RedditScore            int     `json:"redditScore"`
		ArticleUpdateTime      int64   `json:"articleUpdateTime"`
		ArticleStatus          int     `json:"articleStatus"`
		ArticlePerfect         int     `json:"articlePerfect"`
		ArticlePermalink       string  `json:"articlePermalink"`
		ArticleCity            string  `json:"articleCity"`
	} `json:"article"`
	CommentStatus int `json:"commentStatus"`
	Commenter     struct {
		UserQQ                        string `json:"userQQ"`
		UserOnlineFlag                bool   `json:"userOnlineFlag"`
		UserPhone                     string `json:"userPhone"`
		OnlineMinute                  int    `json:"onlineMinute"`
		UserPointStatus               int    `json:"userPointStatus"`
		UserLatestLoginIP             string `json:"userLatestLoginIP"`
		UserFollowerStatus            int    `json:"userFollowerStatus"`
		UserGuideStep                 int    `json:"userGuideStep"`
		UserOnlineStatus              int    `json:"userOnlineStatus"`
		UserCurrentCheckinStreakStart int    `json:"userCurrentCheckinStreakStart"`
		ChatRoomPictureStatus         int    `json:"chatRoomPictureStatus"`
		UserTags                      string `json:"userTags"`
		UserCommentStatus             int    `json:"userCommentStatus"`
		UserTimezone                  string `json:"userTimezone"`
		UserURL                       string `json:"userURL"`
		UserForwardPageStatus         int    `json:"userForwardPageStatus"`
		UserUAStatus                  int    `json:"userUAStatus"`
		UserIndexRedirectURL          string `json:"userIndexRedirectURL"`
		UserLatestArticleTime         int64  `json:"userLatestArticleTime"`
		UserTagCount                  int    `json:"userTagCount"`
		UserNickname                  string `json:"userNickname"`
		UserListViewMode              int    `json:"userListViewMode"`
		UserLongestCheckinStreak      int    `json:"userLongestCheckinStreak"`
		UserAvatarType                int    `json:"userAvatarType"`
		UserSubMailSendTime           int64  `json:"userSubMailSendTime"`
		UserUpdateTime                int64  `json:"userUpdateTime"`
		UserSubMailStatus             int    `json:"userSubMailStatus"`
		UserJoinPointRank             int    `json:"userJoinPointRank"`
		UserLatestLoginTime           int64  `json:"userLatestLoginTime"`
		UserCity                      string `json:"userCity"`
		UserPassword                  string `json:"userPassword"`
		UserAppRole                   int    `json:"userAppRole"`
		UserAvatarViewMode            int    `json:"userAvatarViewMode"`
		UserStatus                    int    `json:"userStatus"`
		UserLongestCheckinStreakEnd   int    `json:"userLongestCheckinStreakEnd"`
		UserWatchingArticleStatus     int    `json:"userWatchingArticleStatus"`
		UserLatestCmtTime             int64  `json:"userLatestCmtTime"`
		UserProvince                  string `json:"userProvince"`
		UserCurrentCheckinStreak      int    `json:"userCurrentCheckinStreak"`
		UserNo                        int    `json:"userNo"`
		UserAvatarURL                 string `json:"userAvatarURL"`
		UserFollowingTagStatus        int    `json:"userFollowingTagStatus"`
		UserLanguage                  string `json:"userLanguage"`
		UserJoinUsedPointRank         int    `json:"userJoinUsedPointRank"`
		UserCurrentCheckinStreakEnd   int    `json:"userCurrentCheckinStreakEnd"`
		UserFollowingArticleStatus    int    `json:"userFollowingArticleStatus"`
		UserKeyboardShortcutsStatus   int    `json:"userKeyboardShortcutsStatus"`
		UserReplyWatchArticleStatus   int    `json:"userReplyWatchArticleStatus"`
		UserCountry                   string `json:"userCountry"`
		UserCommentViewMode           int    `json:"userCommentViewMode"`
		UserBreezemoonStatus          int    `json:"userBreezemoonStatus"`
		UserCheckinTime               int64  `json:"userCheckinTime"`
		Secret2Fa                     string `json:"secret2fa"`
		UserEmail                     string `json:"userEmail"`
		UserUsedPoint                 int    `json:"userUsedPoint"`
		UserArticleStatus             int    `json:"userArticleStatus"`
		UserPoint                     int    `json:"userPoint"`
		UserCommentCount              int    `json:"userCommentCount"`
		UserIntro                     string `json:"userIntro"`
		UserMobileSkin                string `json:"userMobileSkin"`
		UserListPageSize              int    `json:"userListPageSize"`
		OId                           string `json:"oId"`
		UserName                      string `json:"userName"`
		UserGeoStatus                 int    `json:"userGeoStatus"`
		UserLongestCheckinStreakStart int    `json:"userLongestCheckinStreakStart"`
		UserSkin                      string `json:"userSkin"`
		UserNotifyStatus              int    `json:"userNotifyStatus"`
		UserFollowingUserStatus       int    `json:"userFollowingUserStatus"`
		UserArticleCount              int    `json:"userArticleCount"`
		Mbti                          string `json:"mbti"`
		UserRole                      string `json:"userRole"`
	} `json:"commenter"`
	CommentAuthorName         string `json:"commentAuthorName"`
	CommentThankCnt           int    `json:"commentThankCnt"`
	CommentBadCnt             int    `json:"commentBadCnt"`
	Rewarded                  bool   `json:"rewarded"`
	CommentId                 string `json:"commentId"`
	CommentAuthorThumbnailURL string `json:"commentAuthorThumbnailURL"`
	CommentAudioURL           string `json:"commentAudioURL"`
	CommentQnAOffered         int    `json:"commentQnAOffered"`
	CommentAuthorNickName     string `json:"commentAuthorNickName"`
}
