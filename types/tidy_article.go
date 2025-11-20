package types

import "fmt"

// PostArticleRequest 发布文章请求
type PostArticleRequest struct {
	ArticleTitle           string            `json:"articleTitle"`                   // 帖子标题
	ArticleContent         string            `json:"articleContent"`                 // 帖子内容
	ArticleTags            string            `json:"articleTags"`                    // 帖子标签
	ArticleCommentable     bool              `json:"articleCommentable"`             // 是否允许评论
	ArticleNotifyFollowers bool              `json:"articleNotifyFollowers"`         // 是否通知帖子关注着
	ArticleType            ArticleType       `json:"articleType"`                    // 帖子类型
	ArticleShowInList      ArticleShowInList `json:"articleShowInList"`              // 是否在列表展示
	ArticleRewardContent   *string           `json:"articleRewardContent,omitempty"` // 打赏内容
	ArticleRewardPoint     *int              `json:"articleRewardPoint,omitempty"`   // 打赏积分
	ArticleAnonymous       *bool             `json:"articleAnonymous,omitempty"`     // 是否匿名
	ArticleQnAOfferPoint   *int              `json:"articleQnAOfferPoint,omitempty"` // 提问悬赏积分
}

type PostArticleResponse struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg,omitempty"`
	ArticleId string `json:"articleId,omitempty"`
}

type GetArticlesRequest struct {
	Type    GetArticleType
	Keyword string
	Order   *GetArticleOrder
	Page    int
	Size    int
}

func (request *GetArticlesRequest) ToPath() string {
	path := fmt.Sprintf("/%s", request.Type)

	if request.Keyword != "" {
		path += fmt.Sprintf("/%s", request.Keyword)
	}

	if request.Order != nil {
		path += fmt.Sprintf("/%s", *request.Order)
	}

	return path
}

type T struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data struct {
		Articles []struct {
			ArticleCreateTime            string  `json:"articleCreateTime"`
			ArticleRewardPoint           int     `json:"articleRewardPoint"`
			ArticleLatestCmtTime         string  `json:"articleLatestCmtTime"`
			ArticleThumbnailURL          string  `json:"articleThumbnailURL"`
			ArticleStatement             int     `json:"articleStatement"`
			ArticleAuthorName            string  `json:"articleAuthorName"`
			ArticleType                  int     `json:"articleType"`
			ArticleCreateTimeStr         string  `json:"articleCreateTimeStr"`
			ArticleViewCount             int     `json:"articleViewCount"`
			ArticleCommentable           bool    `json:"articleCommentable"`
			ArticleAuthorThumbnailURL20  string  `json:"articleAuthorThumbnailURL20"`
			ArticlePreviewContent        string  `json:"articlePreviewContent"`
			ArticleCommentCount          int     `json:"articleCommentCount"`
			ArticleLatestCmterName       string  `json:"articleLatestCmterName"`
			ArticleAnonymousView         int     `json:"articleAnonymousView"`
			CmtTimeAgo                   string  `json:"cmtTimeAgo"`
			ArticleLatestCmtTimeStr      string  `json:"articleLatestCmtTimeStr"`
			ArticleViewCntDisplayFormat  string  `json:"articleViewCntDisplayFormat"`
			ArticleHeat                  int     `json:"articleHeat"`
			ArticlePerfect               int     `json:"articlePerfect"`
			ArticleAuthorThumbnailURL210 string  `json:"articleAuthorThumbnailURL210"`
			ArticlePermalink             string  `json:"articlePermalink"`
			ArticleCity                  string  `json:"articleCity"`
			ArticleShowInList            int     `json:"articleShowInList"`
			ArticleEditorType            int     `json:"articleEditorType"`
			ArticleRandomDouble          float64 `json:"articleRandomDouble"`
			ArticleAuthorId              string  `json:"articleAuthorId"`
			ArticleBadCnt                int     `json:"articleBadCnt"`
			ArticleGoodCnt               int     `json:"articleGoodCnt"`
			ArticleQnAOfferPoint         int     `json:"articleQnAOfferPoint"`
			ArticleStickRemains          int     `json:"articleStickRemains"`
			TimeAgo                      string  `json:"timeAgo"`
			ArticleUpdateTimeStr         string  `json:"articleUpdateTimeStr"`
			Offered                      bool    `json:"offered"`
			ArticleWatchCnt              int     `json:"articleWatchCnt"`
			ArticleTitleEmoj             string  `json:"articleTitleEmoj"`
			ArticleTitleEmojUnicode      string  `json:"articleTitleEmojUnicode"`
			ArticleAudioURL              string  `json:"articleAudioURL"`
			ArticleAuthorThumbnailURL48  string  `json:"articleAuthorThumbnailURL48"`
			ArticleImg1URL               string  `json:"articleImg1URL"`
			ArticlePushOrder             int     `json:"articlePushOrder"`
			ArticleCollectCnt            int     `json:"articleCollectCnt"`
			ArticleTitle                 string  `json:"articleTitle"`
			ArticleTags                  string  `json:"articleTags"`
			OId                          string  `json:"oId"`
			ArticleStick                 int     `json:"articleStick"`
			ArticleAnonymous             int     `json:"articleAnonymous"`
			ArticleThankCnt              int     `json:"articleThankCnt"`
			ArticleRewardContent         string  `json:"articleRewardContent"`
			RedditScore                  float64 `json:"redditScore"`
			ArticleUpdateTime            string  `json:"articleUpdateTime"`
			ArticleStatus                int     `json:"articleStatus"`
		} `json:"articles"`
		Pagination struct {
			PaginationPageCount int   `json:"paginationPageCount"`
			PaginationPageNums  []int `json:"paginationPageNums"`
		} `json:"pagination"`
		Domain struct {
			DomainTitle    string `json:"domainTitle"`
			DomainIconPath string `json:"domainIconPath"`
			DomainType     string `json:"domainType"`
			DomainTags     []struct {
				TagShowSideAd     int     `json:"tagShowSideAd"`
				TagIconPath       string  `json:"tagIconPath"`
				TagStatus         int     `json:"tagStatus"`
				TagBadCnt         int     `json:"tagBadCnt"`
				TagRandomDouble   float64 `json:"tagRandomDouble"`
				TagTitle          string  `json:"tagTitle"`
				OId               string  `json:"oId"`
				TagURI            string  `json:"tagURI"`
				TagAd             string  `json:"tagAd"`
				TagGoodCnt        int     `json:"tagGoodCnt"`
				TagCSS            string  `json:"tagCSS"`
				TagCommentCount   int     `json:"tagCommentCount"`
				TagFollowerCount  int     `json:"tagFollowerCount"`
				TagSeoTitle       string  `json:"tagSeoTitle"`
				TagLinkCount      int     `json:"tagLinkCount"`
				TagSeoDesc        string  `json:"tagSeoDesc"`
				TagReferenceCount int     `json:"tagReferenceCount"`
				TagSeoKeywords    string  `json:"tagSeoKeywords"`
				TagDescription    string  `json:"tagDescription"`
			} `json:"domainTags"`
			DomainSeoTitle    string `json:"domainSeoTitle"`
			DomainStatus      int    `json:"domainStatus"`
			OId               string `json:"oId"`
			DomainURI         string `json:"domainURI"`
			DomainSeoKeywords string `json:"domainSeoKeywords"`
			DomainDescription string `json:"domainDescription"`
			DomainSort        int    `json:"domainSort"`
			DomainNav         int    `json:"domainNav"`
			DomainCSS         string `json:"domainCSS"`
			DomainTagCnt      int    `json:"domainTagCnt"`
			DomainSeoDesc     string `json:"domainSeoDesc"`
		} `json:"domain"`
	} `json:"data"`
}

type ArticleParticipant struct {
	ArticleParticipantURL          string `json:"articleParticipantURL"`
	CommentId                      string `json:"commentId"`
	OId                            string `json:"oId"`
	ArticleParticipantName         string `json:"articleParticipantName"`
	ArticleParticipantThumbnailURL string `json:"articleParticipantThumbnailURL"`
}

type ArticleAuthor struct {
	UserOnlineFlag              bool   `json:"userOnlineFlag"`
	OnlineMinute                int    `json:"onlineMinute"`
	UserPointStatus             int    `json:"userPointStatus"`
	UserFollowerStatus          int    `json:"userFollowerStatus"`
	UserGuideStep               int    `json:"userGuideStep"`
	UserOnlineStatus            int    `json:"userOnlineStatus"`
	ChatRoomPictureStatus       int    `json:"chatRoomPictureStatus"`
	UserTags                    string `json:"userTags"`
	UserCommentStatus           int    `json:"userCommentStatus"`
	UserTimezone                string `json:"userTimezone"`
	UserURL                     string `json:"userURL"`
	UserForwardPageStatus       int    `json:"userForwardPageStatus"`
	UserUAStatus                int    `json:"userUAStatus"`
	UserIndexRedirectURL        string `json:"userIndexRedirectURL"`
	UserLatestArticleTime       int64  `json:"userLatestArticleTime"`
	UserTagCount                int    `json:"userTagCount"`
	UserNickname                string `json:"userNickname"`
	UserListViewMode            int    `json:"userListViewMode"`
	UserAvatarType              int    `json:"userAvatarType"`
	UserSubMailStatus           int    `json:"userSubMailStatus"`
	UserJoinPointRank           int    `json:"userJoinPointRank"`
	UserAppRole                 int    `json:"userAppRole"`
	UserAvatarViewMode          int    `json:"userAvatarViewMode"`
	UserStatus                  int    `json:"userStatus"`
	UserWatchingArticleStatus   int    `json:"userWatchingArticleStatus"`
	UserProvince                string `json:"userProvince"`
	UserNo                      int    `json:"userNo"`
	UserAvatarURL               string `json:"userAvatarURL"`
	UserFollowingTagStatus      int    `json:"userFollowingTagStatus"`
	UserLanguage                string `json:"userLanguage"`
	UserJoinUsedPointRank       int    `json:"userJoinUsedPointRank"`
	UserFollowingArticleStatus  int    `json:"userFollowingArticleStatus"`
	UserKeyboardShortcutsStatus int    `json:"userKeyboardShortcutsStatus"`
	UserReplyWatchArticleStatus int    `json:"userReplyWatchArticleStatus"`
	UserCommentViewMode         int    `json:"userCommentViewMode"`
	UserBreezemoonStatus        int    `json:"userBreezemoonStatus"`
	UserUsedPoint               int    `json:"userUsedPoint"`
	UserArticleStatus           int    `json:"userArticleStatus"`
	UserPoint                   int    `json:"userPoint"`
	UserCommentCount            int    `json:"userCommentCount"`
	UserIntro                   string `json:"userIntro"`
	UserMobileSkin              string `json:"userMobileSkin"`
	UserListPageSize            int    `json:"userListPageSize"`
	OId                         string `json:"oId"`
	UserName                    string `json:"userName"`
	UserGeoStatus               int    `json:"userGeoStatus"`
	UserSkin                    string `json:"userSkin"`
	UserNotifyStatus            int    `json:"userNotifyStatus"`
	UserFollowingUserStatus     int    `json:"userFollowingUserStatus"`
	UserArticleCount            int    `json:"userArticleCount"`
	Mbti                        string `json:"mbti"`
	UserRole                    string `json:"userRole"`
}

// ArticleInfo 文章信息
type ArticleInfo struct {
	ArticleShowInList            int                   `json:"articleShowInList"`
	ArticleCreateTime            string                `json:"articleCreateTime"`
	ArticleAuthorId              string                `json:"articleAuthorId"`
	ArticleBadCnt                int                   `json:"articleBadCnt"`
	ArticleParticipants          []*ArticleParticipant `json:"articleParticipants"`
	ArticleLatestCmtTime         string                `json:"articleLatestCmtTime"`
	ArticleGoodCnt               int                   `json:"articleGoodCnt"`
	ArticleQnAOfferPoint         int                   `json:"articleQnAOfferPoint"`
	ArticleThumbnailURL          string                `json:"articleThumbnailURL"`
	ArticleStickRemains          int                   `json:"articleStickRemains"`
	TimeAgo                      string                `json:"timeAgo"`
	ArticleUpdateTimeStr         string                `json:"articleUpdateTimeStr"`
	ArticleAuthorName            string                `json:"articleAuthorName"`
	ArticleType                  int                   `json:"articleType"`
	Offered                      bool                  `json:"offered"`
	ArticleCreateTimeStr         string                `json:"articleCreateTimeStr"`
	ArticleViewCount             int                   `json:"articleViewCount"`
	ArticleAuthorThumbnailURL20  string                `json:"articleAuthorThumbnailURL20"`
	ArticleWatchCnt              int                   `json:"articleWatchCnt"`
	ArticlePreviewContent        string                `json:"articlePreviewContent"`
	ArticleTitleEmoj             string                `json:"articleTitleEmoj"`
	ArticleTitleEmojUnicode      string                `json:"articleTitleEmojUnicode"`
	ArticleAuthorThumbnailURL48  string                `json:"articleAuthorThumbnailURL48"`
	ArticleCommentCount          int                   `json:"articleCommentCount"`
	ArticleCollectCnt            int                   `json:"articleCollectCnt"`
	ArticleTitle                 string                `json:"articleTitle"`
	ArticleLatestCmterName       string                `json:"articleLatestCmterName"`
	ArticleTags                  string                `json:"articleTags"`
	OId                          string                `json:"oId"`
	CmtTimeAgo                   string                `json:"cmtTimeAgo"`
	ArticleStick                 int                   `json:"articleStick"`
	ArticleTagObjs               []*ArticleTag         `json:"articleTagObjs"`
	ArticleLatestCmtTimeStr      string                `json:"articleLatestCmtTimeStr"`
	ArticleAnonymous             int                   `json:"articleAnonymous"`
	ArticleViewCntDisplayFormat  string                `json:"articleViewCntDisplayFormat,omitempty"`
	ArticleThankCnt              int                   `json:"articleThankCnt"`
	ArticleUpdateTime            string                `json:"articleUpdateTime"`
	ArticleStatus                int                   `json:"articleStatus"`
	ArticleHeat                  int                   `json:"articleHeat"`
	ArticlePerfect               int                   `json:"articlePerfect"`
	ArticleAuthorThumbnailURL210 string                `json:"articleAuthorThumbnailURL210"`
	ArticlePermalink             string                `json:"articlePermalink"`
	ArticleAuthor                *ArticleAuthor        `json:"articleAuthor"`
	ArticleRewardPoint           int                   `json:"articleRewardPoint"`
	ArticleStatement             int                   `json:"articleStatement"`
	ArticleCommentable           bool                  `json:"articleCommentable"`
	ArticleAnonymousView         int                   `json:"articleAnonymousView"`
	ArticleCity                  string                `json:"articleCity"`
	ArticleEditorType            int                   `json:"articleEditorType"`
	ArticleRandomDouble          float64               `json:"articleRandomDouble"`
	ArticleAudioURL              string                `json:"articleAudioURL"`
	ArticleImg1URL               string                `json:"articleImg1URL"`
	ArticlePushOrder             int                   `json:"articlePushOrder"`
	ArticleRewardContent         string                `json:"articleRewardContent"`
	RedditScore                  float64               `json:"redditScore"`
}

// ArticlePagination 分页信息
type ArticlePagination struct {
	PaginationPageCount int   `json:"paginationPageCount"`
	PaginationPageNums  []int `json:"paginationPageNums"`
}

type ArticleTag struct {
	TagShowSideAd      int           `json:"tagShowSideAd"`
	TagIconPath        string        `json:"tagIconPath"`
	TagStatus          int           `json:"tagStatus"`
	TagBadCnt          int           `json:"tagBadCnt"`
	TagRandomDouble    float64       `json:"tagRandomDouble"`
	TagTitle           string        `json:"tagTitle"`
	IsReserved         bool          `json:"isReserved"`
	OId                string        `json:"oId"`
	TagURI             string        `json:"tagURI"`
	TagAd              string        `json:"tagAd"`
	TagGoodCnt         int           `json:"tagGoodCnt"`
	TagCSS             string        `json:"tagCSS"`
	TagCommentCount    int           `json:"tagCommentCount"`
	TagDescriptionText string        `json:"tagDescriptionText"`
	TagFollowerCount   int           `json:"tagFollowerCount"`
	TagRelatedTags     []*ArticleTag `json:"tagRelatedTags,omitempty"`
	TagDomains         []interface{} `json:"tagDomains"`
	TagSeoTitle        string        `json:"tagSeoTitle"`
	TagLinkCount       int           `json:"tagLinkCount"`
	TagSeoDesc         string        `json:"tagSeoDesc"`
	TagReferenceCount  int           `json:"tagReferenceCount"`
	TagSeoKeywords     string        `json:"tagSeoKeywords"`
	TagDescription     string        `json:"tagDescription"`
}

type ArticleDomain struct {
	DomainTitle       string        `json:"domainTitle"`
	DomainIconPath    string        `json:"domainIconPath"`
	DomainType        string        `json:"domainType"`
	DomainTags        []*ArticleTag `json:"domainTags"`
	DomainSeoTitle    string        `json:"domainSeoTitle"`
	DomainStatus      int           `json:"domainStatus"`
	OId               string        `json:"oId"`
	DomainURI         string        `json:"domainURI"`
	DomainSeoKeywords string        `json:"domainSeoKeywords"`
	DomainDescription string        `json:"domainDescription"`
	DomainSort        int           `json:"domainSort"`
	DomainNav         int           `json:"domainNav"`
	DomainCSS         string        `json:"domainCSS"`
	DomainTagCnt      int           `json:"domainTagCnt"`
	DomainSeoDesc     string        `json:"domainSeoDesc"`
}

// ArticleList 文章列表
type ArticleList struct {
	Articles   []*ArticleInfo    `json:"articles"`
	Pagination ArticlePagination `json:"pagination,omitempty"`
	Tag        ArticleTag        `json:"tag,omitempty"`
	Domain     ArticleDomain     `json:"domain,omitempty"`
}
