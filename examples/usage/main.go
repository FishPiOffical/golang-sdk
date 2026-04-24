package main

import (
	"encoding/json"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/duke-git/lancet/v2/convertor"
	"github.com/fishpioffical/golang-sdk/config"
	"github.com/fishpioffical/golang-sdk/sdk"
	"github.com/fishpioffical/golang-sdk/types"
	"github.com/golang-cz/devslog"

	"log/slog"
	"os"
)

var (
	client *sdk.FishPiSDK
	logger *slog.Logger
)

func init() {
	opts := &devslog.Options{
		HandlerOptions: &slog.HandlerOptions{
			AddSource: false,
			Level:     slog.LevelDebug,
		},
		TimeFormat:        "[15:04:05]",
		NewLineAfterLog:   true,
		DebugColor:        devslog.Magenta,
		StringerFormatter: true,
	}

	logger = slog.New(devslog.NewHandler(os.Stdout, opts))
	slog.SetDefault(logger)

	provider := config.NewFileYamlProvider(configPath)

	// 使用选项创建SDK
	client = sdk.NewSDK(
		provider,
		sdk.WithLogDir(logPath),                   // 设置日志目录
		sdk.WithCustomUnmarshaler(slog.Default()), // 设置自定义反序列化器
	)
}

const (
	configPath = "../../_tmp/config.yaml"
	logPath    = "../../_tmp/logs/"

	username                  = "8888"
	associateName             = "888"
	reportArticleId           = "1702103071389" // https://fishpi.cn/article/1702103071389
	followingId               = "1734578210153" // https://fishpi.cn/member/wordsKing
	messageOId                = "1763542689788"
	uploadFile1               = "../../_tmp/files/IMG_1045.jpg"
	uploadFile2               = "../../_tmp/files/IMG_13069.jpeg"
	editArticleId             = "1763623304114" // https://fishpi.cn/article/1763623304114
	editArticleCommentId      = "1763629146604" // 念
	editArticleFirstCommentId = "1763624701751" // 厌子
	editArticleReplyCommentId = "1764215313804" // 回复念
	otherArticleId            = "1630569106133"
	botUserName               = "its21f"
	chatMessageOId            = "1765184305408"
	userIdYui                 = "1630488635229"
	userIdIWPZ                = "1637917131504"                                    // 和平鸽
	userIdDRDA                = "1678416418912"                                    // 加辣
	userId8888                = "1656984017362"                                    // 开摆
	avatarUrl                 = "https://file.fishpi.cn/2022/08/blob-fbe21682.png" // 开摆的头像

	realm    = "https://test.fishpi.cn"
	returnTo = "https://test.fishpi.cn/return_to"

	medalIdAdmin    = "0"
	medalIdYearOne  = "1"  // 摸鱼派1岁啦
	medalIdFans16   = "5"  // 摸鱼派粉丝
	medalIdTest     = "62" // 测试徽章
	medalNameSister = "小姐姐认证"
)

func main() {

	// 鉴权
	//postApiGetKey()
	//getApiUser()

	// 通用
	//getUserInfoByUsername()
	//postUsersNames()
	//getUsersEmotions()
	//getUserLiveness()
	//getUserCheckedIn()
	//getYesterdayLivenessReward()
	//getIsCollectedLiveness()
	//postReport()
	//getUserRecentReg()
	//postPointTransfer()
	//postFollowUser()
	//postUnfollowUser()
	//getMembership()
	//getLogsMore()
	//postSettingsAvatar()
	//postSettingsProfiles()
	//getUserUsernamePoint()
	//getUserUsernameMedal()

	// 通知
	//getNotificationCount()
	//getNotificationsPoint()
	//getNotificationsCommented()
	//getNotificationsReply()
	//getNotificationsAt()
	//getNotificationsFollowing()
	//getNotificationsBroadcast()
	//getNotificationsSysAnnounce()
	//getNotificationsMarkRead()
	//getNotificationsAllRead()

	// 聊天室
	//getChatroomBarragePrice()
	//getChatroomNode()
	//getChatroomMore()
	//getChatroomMessage()
	//postChatroomSend()
	//sendBarrage()
	//sendRedPacket()
	//deleteChatroomRevoke()
	//getMessageRaw()
	//postRedPacketSend()
	//postCloudGet()
	//postCloudSync()
	//getDefaultEmojis()
	//getSiGuoYa()

	// 表情包V2
	//getEmojiGroups()
	//getEmojiGroupEmojis()
	//postEmojiUpload()
	//postEmojiGroupCreate()
	//postEmojiGroupUpdate()
	//postEmojiGroupDelete()
	//postEmojiGroupAddEmoji()
	//postEmojiGroupAddUrlEmoji()
	//postEmojiGroupRemoveEmoji()
	//postEmojiEmojiUpdate()
	//postEmojiEmojiMigrate()

	// 图床
	//postUploadFile()

	// 帖子
	//postArticle()
	//putArticle()
	//getArticles()
	//getArticleDetail()
	//getUserArticles()
	//postVoteUpArticle()
	//postArticleThank()
	//getArticleComments()
	//postComment()
	//putComment()
	//postVoteUpComment()
	//postCommentThank()
	//postCommentRemove()
	//postArticleHeat()
	//postFollowArticle()
	//postUnfollowArticle()
	//postFollowArticleWatch()
	//postUnfollowArticleWatch()
	//postArticleReward()
	//getArticleMd()

	// 清风明月
	//getBreezemoons()
	//postBreezemoon()
	//getUserBreezemoons()

	// 私信
	//getChatMessage()
	//getChatMarkAsRead()
	//getChatGetList()
	//getChatHasUnread()
	//getChatRevoke()

	// 金手指
	//postMofishScore()
	//postQueryLatestLoginIp()
	//postGiveMetal()
	//postRemoveMetal()
	//postRemoveMetalByUserId()
	//postUserQueryItems()
	//postUserEditItems()
	//postUserEditPoints()
	//postUserLiveness()
	//postYesterdayLivenessReward()
	//postUserEditNotification()

	// websocket
	//userChannelWebsocket() // 通知
	//chatChannelWebsocket() // 私聊
	//chatroomWebsocket() // 聊天室
	//articleChannelWebsocket() // 文章热度通知

	// openid
	//getUserInfoById()
	//getOpenIdUrl()
	//runOpenIdServer()

	// medal
	//postMedalMyList()
	//postMedalMyReorder()
	//postMedalMyDisplay()
	//postMedalUserList()
	//postMedalAdminList()
	//postMedalAdminSearch()
	//postMedalAdminDetail()
	//postMedalAdminDelete()
	//postMedalAdminEdit()
	//postMedalAdminCreate()
	//postMedalAdminGrant()
	//postMedalAdminRevoke()
	//postMedalAdminOwners()
	//getMedalUrl()
	//postMedalAdminUserMedals()

	// reaction
	//postArticleReaction()
	//postCommentReaction()
	postChatRoomReaction()

}

func getUserInfoByUsername() {
	user, err := client.GetUserInfoByUsername("drda")
	if err != nil {
		logger.Error("获取用户信息失败", slog.String("error", err.Error()))
		return
	}
	metalList, MetalErr := user.GetMetalList()
	logger.Info("用户信息",
		slog.Any("user", user.UserNickname),
		slog.Any("metalList", metalList),
		slog.Any("MetalErr", MetalErr),
	)
}

func postUsersNames() {
	resp, err := client.PostUsersNames(associateName)
	if err != nil {
		logger.Error("用户名联想失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("用户名联想结果", slog.Any("resp", resp))
}

func getUsersEmotions() {
	resp, err := client.GetUsersEmotions()
	if err != nil {
		logger.Error("获取用户常用表情失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("用户常用表情结果", slog.Any("resp", resp))
}

func getUserLiveness() {
	resp, err := client.GetUserLiveness()
	if err != nil {
		logger.Error("获取用户活跃度失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("用户活跃度结果", slog.Any("resp", resp))
}

func getUserCheckedIn() {
	resp, err := client.GetUserCheckedIn()
	if err != nil {
		logger.Error("检查用户是否签到失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("用户签到结果", slog.Any("resp", resp))
}

func getYesterdayLivenessReward() {
	resp, err := client.GetYesterdayLivenessReward()
	if err != nil {
		logger.Error("获取昨日活跃度奖励失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("昨日活跃度奖励结果", slog.Any("resp", resp))
}

func getIsCollectedLiveness() {
	resp, err := client.GetIsCollectedLiveness()
	if err != nil {
		logger.Error("检查是否已领取昨日活跃度奖励失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("是否已领取昨日活跃度奖励结果", slog.Any("resp", resp))
}

func postReport() {
	resp, err := client.PostReport(reportArticleId, types.ReportDataTypeArticle, types.ReportTypeOther, "接口测试举报")
	if err != nil {
		logger.Error("举报用户失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("举报用户结果", slog.Any("resp", resp))
}

func getUserRecentReg() {
	resp, err := client.GetUserRecentReg()
	if err != nil {
		logger.Error("获取用户最近注册信息失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("用户最近注册信息结果", slog.Any("resp", resp))
}

func postPointTransfer() {
	resp, err := client.PostPointTransfer(&types.PostPointTransferRequest{
		UserName: username,
		Amount:   2,
		Memo:     "接口测试转账",
	})
	if err != nil {
		logger.Error("用户积分转账失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("用户积分转账结果", slog.Any("resp", resp))
}

func postFollowUser() {
	resp, err := client.PostFollowUser(followingId)
	if err != nil {
		logger.Error("关注用户失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("关注用户结果", slog.Any("resp", resp))
}

func postUnfollowUser() {
	resp, err := client.PostUnfollowUser(followingId)
	if err != nil {
		logger.Error("取消关注用户失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("取消关注用户结果", slog.Any("resp", resp))
}

func getMembership() {
	resp, err := client.GetMembership(userIdDRDA)
	if err != nil {
		logger.Error("获取用户VIP信息失败", slog.String("error", err.Error()))
		return
	}
	conf, e := resp.Data.GetConfig()
	logger.Info("用户VIP信息结果", slog.Any("resp", resp), slog.Any("config", conf), slog.Any("config err", e))
}

func getLogsMore() {
	resp, err := client.GetLogsMore(1, 1)
	if err != nil {
		logger.Error("获取操作日志失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("操作日志结果", slog.Any("resp", resp))
}

func postSettingsAvatar() {
	resp, err := client.PostSettingsAvatar(avatarUrl)
	if err != nil {
		logger.Error("更新用户头像失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("更新用户头像结果", slog.Any("resp", resp))
}

func postSettingsProfiles() {
	resp, err := client.PostSettingsProfiles(&types.PostSettingsProfilesRequest{
		UserNickname: "开摆",
		UserTags:     "开摆,哇咔咔",
		UserIntro:    "我就是我，不一样的烟火。",
		UserURL:      "https://www.zqcnc.cn",
		Mbti:         "ISFJ-T",
	})
	if err != nil {
		logger.Error("更新用户资料失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("更新用户资料结果", slog.Any("resp", resp))
}

func getUserUsernamePoint() {
	resp, err := client.GetUserUsernamePoint(username)
	if err != nil {
		logger.Error("查询指定用户积分失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("指定用户积分结果", slog.Any("resp", resp))
}

func getUserUsernameMedal() {
	resp, err := client.GetUserUsernameMedal(username)
	if err != nil {
		logger.Error("查询指定用户勋章失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("指定用户勋章结果", slog.Any("resp", resp))
}

func getNotificationCount() {
	resp, err := client.GetNotificationCount()
	if err != nil {
		logger.Error("获取未读通知数量失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("未读通知数量结果", slog.Any("resp", resp))
}

func getNotificationsPoint() {
	resp, err := client.GetNotifications(types.NotificationTypePoint, 1)
	if err != nil {
		logger.Error("获取积分通知列表失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("积分通知列表结果", slog.Any("resp", resp))
}

func getNotificationsCommented() {
	resp, err := client.GetNotifications(types.NotificationTypeCommented, 1)
	if err != nil {
		logger.Error("获取评论通知列表失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("评论通知列表结果", slog.Any("resp", resp))
}

func getNotificationsReply() {
	resp, err := client.GetNotifications(types.NotificationTypeReply, 1)
	if err != nil {
		logger.Error("获取回复通知列表失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("回复通知列表结果", slog.Any("resp", resp))
}

func getNotificationsAt() {
	resp, err := client.GetNotifications(types.NotificationTypeAt, 1)
	if err != nil {
		logger.Error("获取@通知列表失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("@通知列表结果", slog.Any("resp", resp))
}

func getNotificationsFollowing() {
	resp, err := client.GetNotifications(types.NotificationTypeFollowing, 1)
	if err != nil {
		logger.Error("获取关注通知列表失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("关注通知列表结果", slog.Any("resp", resp))
}

func getNotificationsBroadcast() {
	resp, err := client.GetNotifications(types.NotificationTypeBroadcast, 1)
	if err != nil {
		logger.Error("获取系统通知列表失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("系统通知列表结果", slog.Any("resp", resp))
}

func getNotificationsSysAnnounce() {
	resp, err := client.GetNotifications(types.NotificationTypeSysAnnounce, 1)
	if err != nil {
		logger.Error("获取公告通知列表失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("公告通知列表结果", slog.Any("resp", resp))
}

func getNotificationsMarkRead() {
	resp, err := client.GetNotificationsMarkRead(types.NotificationTypePoint)
	if err != nil {
		logger.Error("标记指定类型通知为已读失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("标记指定类型通知为已读结果", slog.Any("resp", resp))
}

func getNotificationsAllRead() {
	resp, err := client.GetNotificationsAllRead()
	if err != nil {
		logger.Error("标记所有通知为已读失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("标记所有通知为已读结果", slog.Any("resp", resp))
}

func getChatroomBarragePrice() {
	resp, err := client.GetChatroomBarragePrice()
	if err != nil {
		logger.Error("获取弹幕价格失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("弹幕价格结果", slog.Any("resp", resp))
}

func getChatroomNode() {
	resp, err := client.GetChatroomNode()
	if err != nil {
		logger.Error("获取聊天室节点信息失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("聊天室节点信息结果", slog.Any("resp", resp))
}

func getChatroomMore() {
	resp, err := client.GetChatroomMore(1)
	if err != nil {
		logger.Error("获取聊天室历史消息失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("聊天室历史消息结果", slog.Any("resp", resp.Msg))
}

func getChatroomMessage() {
	resp, err := client.GetChatroomMessage(messageOId, 10, types.ChatMessageTypeContext)
	if err != nil {
		logger.Error("获取聊天室指定消息上下文失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("聊天室指定消息上下文结果", slog.Any("resp", resp.Msg))
}

func postChatroomSend() {
	resp, err := client.PostChatroomSend("🎵 你在烦恼什么呢")
	if err != nil {
		logger.Error("发送聊天室消息失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("发送聊天室消息结果", slog.Any("resp", resp))
}

func sendBarrage() {
	resp, err := client.SendBarrage("还有人吗", "#66ccff")
	if err != nil {
		logger.Error("发送弹幕消息失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("发送弹幕消息结果", slog.Any("resp", resp))
}

func sendRedPacket() {
	resp, err := client.SendRedPacket(&types.RedPacket{
		Type:  types.ChatroomRedPacketTypeRandom,
		Money: 32,
		Count: 32,
		Msg:   "让我看看谁还在啊",
	})
	if err != nil {
		logger.Error("发送红包消息失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("发送红包消息结果", slog.Any("resp", resp))
}

func deleteChatroomRevoke() {
	resp, err := client.DeleteChatroomRevoke("1763545820979")
	if err != nil {
		logger.Error("撤回聊天室消息失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("撤回聊天室消息结果", slog.Any("resp", resp))
}

func getMessageRaw() {
	resp, err := client.GetMessageRaw(messageOId)
	if err != nil {
		logger.Error("获取消息原始数据失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("消息原始数据结果", slog.Any("resp", resp))
}

func postRedPacketSend() {
	resp, err := client.PostRedPacketOpen("1763607318962", types.GestureTypePaper)
	if err != nil {
		logger.Error("打开红包失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("打开红包结果", slog.Any("resp", resp))
}

//func postCloudGet() {
//	resp, err := client.PostCloudGet(types.CloudGameIdEmojis)
//	if err != nil {
//		logger.Error("获取云游戏资源失败", slog.String("error", err.Error()))
//		return
//	}
//	logger.Info("云游戏资源结果", slog.Any("resp", resp))
//}

//func postCloudSync() {
//	data := "[\"https://fishpi.cn/gen?scale=0.79&txt=%E5%A5%BD%E5%AE%B6%E4%BC%99&url=https://file.fishpi.cn/2022/06/blob-4a5a6682.png&backcolor=FF00FF&fontcolor=00ff00\",\"https://file.fishpi.cn/2022/01/6W4RCDQ1B3E8IXWDD-3b1dcc36.jpg\",\"https://file.fishpi.cn/vditor/dist/images/emoji/huaji.gif\",\"https://file.fishpi.cn/2022/07/4XD412P6Z3NUADURUIF-cf6e5c49.png\",\"https://file.fishpi.cn/2022/06/image-2fa0b65d.png\",\"https://file.fishpi.cn/2021/11/WY3LZN1MAI1DDONT-e2caa6db.gif\",\"https://file.fishpi.cn/2021/11/image-af346e1a.png\",\"https://file.fishpi.cn/2021/11/video1-f50f21b9.gif\",\"https://tse2-mm.cn.bing.net/th/id/OIP-C.FPgxOL4Y3y45okElx5pzSwAAAA?pid=ImgDet&rs=1\",\"https://unv-shield.librian.net/api/unv_shield?scale=0.79&txt=%E6%88%91%E6%98%AF%E5%BA%9F%E7%89%A9&url=https://www.lingmx.com/52pj/images/die.png&backcolor=568289&fontcolor=ffffff\",\"https://file.fishpi.cn/2022/03/606AE90506E645189E8C89DD0EABE688-dbdd8f22.gif\",\"https://file.fishpi.cn/2022/04/B6ARF4PTIB6JZX-d28e0ddf.jpg\",\"https://file.fishpi.cn/2022/07/1D6841B4-62e975d6.jpg\",\"https://unv-shield.librian.net/api/unv_shield?scale=0.79&txt=lsp%E4%B9%8B%E7%8E%8B%E9%9D%9E%E4%BD%A0%E8%8E%AB%E5%B1%9E&url=https://www.lingmx.com/52pj/images/die.png&backcolor=568289&fontcolor=ffffff\",\"https://file.fishpi.cn/2022/07/image-1900c9b7.png\",\"https://file.fishpi.cn/2022/06/image-be82885f.png\",\"https://file.fishpi.cn/2022/06/image-e04d61d9.png\",\"https://file.fishpi.cn/2022/07/image-abc9d3f6.png\",\"https://file.fishpi.cn/2021/12/image-7d8f1284.png\",\"https://file.fishpi.cn/2022/07/DCD128E70C5B10C49E4D66C8568D10EA-697825f2.jpg\",\"https://file.fishpi.cn/2022/03/image-56c0f695.png\",\"https://file.fishpi.cn/2022/08/34F6E4E51347D5BB164156FF4DFE62BA-9c2cec25.jpg\",\"https://file.fishpi.cn/2021/12/Snipaste20211210094758-990d6133.png\",\"https://file.fishpi.cn/2022/08/60be1bcbc1608b95-1d4762b3.png\",\"https://file.fishpi.cn/2022/08/09A3D30F-7bbd8d34.jpg\",\"https://file.fishpi.cn/2022/06/CB3456129E25827BA532135CB7680B64-0336e313.gif\",\"https://unv-shield.librian.net/api/unv_shield?scale=0.79&txt=%E9%AD%85%E9%AD%94&url=https://pic.stackoverflow.wiki/uploadImages/117/28/120/230/2022/09/10/11/50/29fc900f-2dd0-472d-a772-5c03a9b03156.jpg&backcolor=211a1a&fontcolor=d068da\",\"https://file.fishpi.cn/2022/06/adgif156b20ea-68910939.gif\",\"https://file.fishpi.cn/2022/09/image-3911ca4a.png\",\"https://file.fishpi.cn/2022/11/%E5%9B%BE%E7%89%87-1f3a0fc4.png\",\"https://file.fishpi.cn/2022/12/image-dd15392c.png\",\"https://file.fishpi.cn/2022/09/image-cd84a59e.png\",\"https://file.fishpi.cn/2022/09/%E5%9B%BE%E7%89%87-846950bf.png\",\"https://file.fishpi.cn/2022/12/image-a00a8f3f.png\",\"https://file.fishpi.cn/2022/11/nd83w-888eb73e.gif\",\"https://static.dingtalk.com/media/lQHPJxZuEq0iWwXNAUDNAUCwp1JKstRBmeMCtWjqjoC9AA_320_320.gif?bizType=im\",\"https://file.fishpi.cn/2023/04/20200404023812BehxC-ea84b1b9.gif\",\"https://file.fishpi.cn/2022/01/image-3d8ac437.png\",\"https://file.fishpi.cn/2023/04/1J15M-cc5e47f4.gif\",\"https://file.fishpi.cn/2023/06/ne3eZ-b9e96820.gif\",\"https://file.fishpi.cn/2022/08/XBN9CGP0CHLL2MDZA9L-fe17f120.jpg\",\"https://file.fishpi.cn/2023/08/e5cacb49f4ab060451980e55b06be6d-8805a569.jpg\",\"https://file.fishpi.cn/2023/01/image-f12bf156.png\",\"https://file.fishpi.cn/2023/09/image-eac3523d.png\",\"https://file.fishpi.cn/2023/08/image-d54e2f1d.png\",\"https://file.fishpi.cn/2023/09/image-02ce397b.png\",\"https://fishpi.cn/gen?scale=0.79&txt=%E5%85%84%E5%BC%9F%E4%BB%AC%EF%BC%8C%E5%8F%91%E7%82%B9%E6%B6%A9%E5%9B%BE&url=https://file.fishpi.cn/2022/06/blob-4a5a6682.png&backcolor=FF00FF&fontcolor=00ff00\",\"https://file.fishpi.cn/2023/09/image-26a9d52a.png\",\"https://file.fishpi.cn/2023/09/image-d661fbb3.png\",\"https://file.fishpi.cn/2023/04/image-354d2505.png\",\"https://file.fishpi.cn/2021/11/3I3V8NZ1SBZA35N2S7U1-6a567673.png\",\"https://file.fishpi.cn/2022/09/image-5beacdbd.png\",\"https://file.fishpi.cn/2023/08/image-a9baf2e1.png\",\"https://file.fishpi.cn/2023/08/image-36d1b849.png\",\"https://file.fishpi.cn/2023/10/image-a7e13101.png\",\"https://file.fishpi.cn/2023/10/image-63d9fff9.png\",\"https://file.fishpi.cn/2023/10/image-d74b98cc.png\",\"https://file.fishpi.cn/2023/03/AP2R9HZQBHRIYUK-a6695524.jpg\",\"https://file.fishpi.cn/2023/12/image-5917d1f8.png\",\"https://file.fishpi.cn/2023/12/image-db7980ae.png\",\"https://file.fishpi.cn/2023/12/%E5%9B%BE%E7%89%87-3789ec89.png\",\"https://file.fishpi.cn/2024/01/3cc55d75c8fe412fa6d7e12ca7729355-8c1fefde.png\",\"https://file.fishpi.cn/2024/01/20230706034848avns286q-faa5eb9a.jpg\",\"https://file.fishpi.cn/2023/12/image-a6039b53.png\",\"https://file.fishpi.cn/2024/03/image-80eb3cfb.png\",\"https://file.fishpi.cn/2022/01/a9cf8ef6ly1fiecn56l8wj20b50b2glu-12a17eea.jpg\",\"https://file.fishpi.cn/2024/04/6b4cd9ccd007236eda98d7f80f89f97bfc9a6f1b1df5cd96313f7573fe928087ebfc522-4f35ecc4.png\"]"
//	resp, err := client.PostCloudSync(types.CloudGameIdEmojis, data)
//	if err != nil {
//		logger.Error("同步云游戏资源失败", slog.String("error", err.Error()))
//		return
//	}
//	logger.Info("同步云游戏资源结果", slog.Any("resp", resp))
//}

func getDefaultEmojis() {
	emojis := client.GetDefaultEmojis()
	logger.Info("默认表情列表结果", slog.Any("emojis", emojis))
}

func getSiGuoYa() {
	resp, err := client.GetSiGuoYa()
	if err != nil {
		logger.Error("获取思过崖失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("思过崖结果", slog.Any("resp", resp))
}

func getEmojiGroups() {
	resp, err := client.GetEmojiGroups()
	if err != nil {
		logger.Error("获取表情包分组列表失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("表情包分组列表结果", slog.Any("resp", resp))
}

func getEmojiGroupEmojis() {
	resp, err := client.GetEmojiGroupEmojis("1777001844008")
	if err != nil {
		logger.Error("获取表情包分组表情列表失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("表情包分组表情列表结果", slog.Any("resp", resp))
}

func postEmojiUpload() {
	resp, err := client.PostEmojiUpload("https://www.aweoo.com/upload/aweoo-favicon.png")
	if err != nil {
		logger.Error("上传表情失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("上传表情结果", slog.Any("resp", resp))
}

func postEmojiGroupCreate() {
	resp, err := client.PostEmojiGroupCreate("接口测试表情包分组", 1)
	if err != nil {
		logger.Error("创建表情包分组失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("创建表情包分组结果", slog.Any("resp", resp))
}

func postEmojiGroupUpdate() {
	resp, err := client.PostEmojiGroupUpdate("1777001844008", "接口测试表情包分组（已修改）", 1)
	if err != nil {
		logger.Error("修改表情包分组失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("修改表情包分组结果", slog.Any("resp", resp))
}

func postEmojiGroupDelete() {
	resp, err := client.PostEmojiGroupDelete("1777001844008")
	if err != nil {
		logger.Error("删除表情包分组失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("删除表情包分组结果", slog.Any("resp", resp))
}

func postEmojiGroupAddEmoji() {
	resp, err := client.PostEmojiGroupAddEmoji("1777001844008", "1769942904605", 1, "滑稽")
	if err != nil {
		logger.Error("添加表情到表情包分组失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("添加表情到表情包分组结果", slog.Any("resp", resp))
}

func postEmojiGroupAddUrlEmoji() {
	resp, err := client.PostEmojiGroupAddUrlEmoji("1777001844008", "https://www.aweoo.com/upload/aweoo-favicon.png", 1, "Favicon捏")
	if err != nil {
		logger.Error("添加URL表情到表情包分组失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("添加URL表情到表情包分组结果", slog.Any("resp", resp))
}

func postEmojiGroupRemoveEmoji() {
	resp, err := client.PostEmojiGroupRemoveEmoji("1777001844008", "1777002223602")
	if err != nil {
		logger.Error("从表情包分组移除表情失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("从表情包分组移除表情结果", slog.Any("resp", resp))
}

func postEmojiEmojiUpdate() {
	resp, err := client.PostEmojiEmojiUpdate("1777002063214", "1777001844008", "滑稽（已修改）", 1)
	if err != nil {
		logger.Error("修改表情失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("修改表情结果", slog.Any("resp", resp))
}

func postEmojiEmojiMigrate() {
	resp, err := client.PostEmojiEmojiMigrate()
	if err != nil {
		logger.Error("表情迁移失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("表情迁移结果", slog.Any("resp", resp))
}

func postUploadFile() {
	resp, err := client.PostUploadFile(
		[]string{
			uploadFile1,
			uploadFile2,
		})
	if err != nil {
		logger.Error("上传文件失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("上传文件结果", slog.Any("resp", resp))
}

func getBreezemoons() {
	resp, err := client.GetBreezemoons(1, 20)
	if err != nil {
		logger.Error("获取清风明月列表失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("清风明月列表结果", slog.Any("resp", resp))
}

func postBreezemoon() {
	resp, err := client.PostBreezemoon("怎么还没到下班点")
	if err != nil {
		logger.Error("发送清风明月失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("发送清风明月结果", slog.Any("resp", resp))
}

func getUserBreezemoons() {
	resp, err := client.GetUserBreezemoons(username, 1, 20)
	if err != nil {
		logger.Error("获取指定用户的清风明月列表失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("指定用户的清风明月列表结果", slog.Any("resp", resp))
}

func getChatMessage() {
	resp, err := client.GetChatMessage(botUserName, 1, 2)
	if err != nil {
		logger.Error("获取私信聊天消息失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("私信聊天消息结果", slog.Any("resp", resp))
}

func getChatMarkAsRead() {
	resp, err := client.GetChatMarkAsRead(botUserName)
	if err != nil {
		logger.Error("标记私信聊天消息已读失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("标记私信聊天消息已读结果", slog.Any("resp", resp))
}

func getChatGetList() {
	resp, err := client.GetChatGetList()
	if err != nil {
		logger.Error("获取私信聊天用户列表失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("私信聊天用户列表结果", slog.Any("resp", resp))
}

func getChatHasUnread() {
	resp, err := client.GetChatHasUnread()
	if err != nil {
		logger.Error("获取私信聊天未读消息失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("私信聊天未读消息结果", slog.Any("resp", resp))
}

func getChatRevoke() {
	resp, err := client.GetChatRevoke(chatMessageOId)
	if err != nil {
		logger.Error("撤回私信聊天消息失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("撤回私信聊天消息结果", slog.Any("resp", resp))
}

func postArticle() {
	resp, err := client.PostArticle(&types.PostArticleRequest{
		ArticleTitle:           "AI带来的提升",
		ArticleContent:         "AI已经发展了这么多年，那么AI对你的工作和生活带来了哪些提升呢？  \n> 请详细说明你的实际体验和感受。🍠水贴将会被删除哦！",
		ArticleTags:            "测试,AI,生活",
		ArticleCommentable:     true,
		ArticleNotifyFollowers: false,
		ArticleType:            types.ArticleTypeQuestion,
		ArticleShowInList:      types.ArticleShowInListNo,
		ArticleRewardContent:   convertor.ToPointer("感谢您的支持！"),
		ArticleRewardPoint:     convertor.ToPointer(5),
		ArticleAnonymous:       convertor.ToPointer(false),
		ArticleQnAOfferPoint:   convertor.ToPointer(5),
	})
	if err != nil {
		logger.Error("发布文章失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("发布文章结果", slog.Any("resp", resp))
}

func putArticle() {
	resp, err := client.PutArticle(editArticleId, &types.PostArticleRequest{
		ArticleTitle:           "【测试文章】AI带来的提升",
		ArticleContent:         "AI已经发展了这么多年，那么AI对你的工作和生活带来了哪些提升呢？  \n> 请详细说明你的实际体验和感受。🍠水贴将会被删除哦！",
		ArticleTags:            "测试,AI,生活",
		ArticleCommentable:     true,
		ArticleNotifyFollowers: false,
		ArticleType:            types.ArticleTypeQuestion,
		ArticleShowInList:      types.ArticleShowInListNo,
		ArticleRewardContent:   convertor.ToPointer("感谢您的支持！"),
		ArticleRewardPoint:     convertor.ToPointer(5),
		ArticleAnonymous:       convertor.ToPointer(false),
		ArticleQnAOfferPoint:   convertor.ToPointer(5),
	})
	if err != nil {
		logger.Error("更新文章失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("更新文章结果", slog.Any("resp", resp))
}

func getArticles() {
	resp, err := client.GetArticles(&types.GetArticlesRequest{
		Type:    types.GetArticleTypeRecent,
		Keyword: "",
		Order:   convertor.ToPointer(types.GetArticleOrderLong),
		Page:    1,
		Size:    10,
	})
	if err != nil {
		logger.Error("获取文章列表失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("文章列表结果", slog.Any("resp", resp.Code))
}

func getArticleDetail() {
	resp, err := client.GetArticleDetail(editArticleId)
	if err != nil {
		logger.Error("获取文章详情失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("文章详情结果", slog.Any("resp", resp.Msg))
}

func getUserArticles() {
	resp, err := client.GetUserArticles(username, 1, 10)
	if err != nil {
		logger.Error("获取用户文章列表失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("用户文章列表结果", slog.Any("resp", resp.Msg))
}

func postVoteUpArticle() {
	resp, err := client.PostVoteUpArticle(otherArticleId)
	if err != nil {
		logger.Error("点赞文章失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("点赞文章结果", slog.Any("resp", resp))
}

func postArticleThank() {
	resp, err := client.PostArticleThank(otherArticleId)
	if err != nil {
		logger.Error("感谢文章失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("感谢文章结果", slog.Any("resp", resp))
}

func getArticleComments() {
	resp, err := client.GetArticleComments(otherArticleId, 1, 10)
	if err != nil {
		logger.Error("获取文章评论列表失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("文章评论列表结果", slog.Any("resp", resp.Msg))
}

func postComment() {
	resp, err := client.PostComment(&types.PostCommentRequest{
		ArticleId:         editArticleId,
		CommentContent:    "在做测试，没点到隐藏文章",
		CommentVisible:    convertor.ToPointer(true),
		CommentAnonymous:  convertor.ToPointer(false),
		CommentOriginalId: convertor.ToPointer(editArticleCommentId),
	})
	if err != nil {
		logger.Error("发布评论失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("发布评论结果", slog.Any("resp", resp))
}

func putComment() {
	resp, err := client.PutComment(editArticleReplyCommentId, &types.PutCommentRequest{
		ArticleId:        editArticleId,
		CommentContent:   "在做测试，没点到隐藏文章，被你们发现了",
		CommentVisible:   true,
		CommentAnonymous: true,
	})
	if err != nil {
		logger.Error("更新评论失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("更新评论结果", slog.Any("resp", resp))
}

func postVoteUpComment() {
	resp, err := client.PostVoteUpComment(editArticleCommentId)
	if err != nil {
		logger.Error("点赞评论失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("点赞评论结果", slog.Any("resp", resp))
}

func postCommentThank() {
	resp, err := client.PostCommentThank(editArticleFirstCommentId)
	if err != nil {
		logger.Error("感谢评论失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("感谢评论结果", slog.Any("resp", resp))
}

func postCommentRemove() {
	resp, err := client.PostCommentRemove(editArticleReplyCommentId)
	if err != nil {
		logger.Error("删除评论失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("删除评论结果", slog.Any("resp", resp))
}

func postArticleHeat() {
	resp, err := client.PostArticleHeat(editArticleId)
	if err != nil {
		logger.Error("文章升温失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("文章升温结果", slog.Any("resp", resp))
}

func postFollowArticle() {
	resp, err := client.PostFollowArticle(otherArticleId)
	if err != nil {
		logger.Error("收藏文章失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("收藏文章结果", slog.Any("resp", resp))
}

func postUnfollowArticle() {
	resp, err := client.PostUnfollowArticle(otherArticleId)
	if err != nil {
		logger.Error("收藏文章失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("取消收藏文章结果", slog.Any("resp", resp))
}

func postFollowArticleWatch() {
	resp, err := client.PostFollowArticleWatch(otherArticleId)
	if err != nil {
		logger.Error("关注文章失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("关注文章结果", slog.Any("resp", resp))
}

func postUnfollowArticleWatch() {
	resp, err := client.PostUnfollowArticleWatch(otherArticleId)
	if err != nil {
		logger.Error("关注文章失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("关注文章结果", slog.Any("resp", resp))
}

func postArticleReward() {
	resp, err := client.PostArticleReward(otherArticleId)
	if err != nil {
		logger.Error("打赏文章失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("打赏文章结果", slog.Any("resp", resp))
}

func getArticleMd() {
	content, err := client.GetArticleMd(editArticleId)
	if err != nil {
		logger.Error("获取帖子的Markdown原文失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("获取帖子的Markdown原文成功", slog.String("content", convertor.ToString(content)))
}

func postMofishScore() {
	resp, err := client.PostMoFishScore(username, 10, time.Now().UnixMilli())
	if err != nil {
		logger.Error("提交摸鱼分数失败", slog.Any("error", err))
		return
	}
	logger.Info("提交摸鱼分数结果", slog.Any("resp", resp))
}

func postQueryLatestLoginIp() {
	resp, err := client.PostQueryLatestLoginIp(username)
	if err != nil {
		logger.Error("查询用户最新登录IP失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("查询用户最新登录IP结果", slog.Any("resp", resp))
}

//func postGiveMetal() {
//	resp, err := client.PostGiveMetal("AziAzi", &types.Metal{
//		Name:        "最受欢迎设计师S1",
//		Description: "摸鱼派五周年徽章共创计划 N0.1",
//		Attr:        "url=https://file.fishpi.cn/2025/12/7cab0213437e841e8680120e612a9ea3-331d0a76.png&txt=最受欢迎设计师S1&scale=0.79&backcolor=cccccc,7d7b7b&way=top-left",
//		Data:        "",
//	})
//	if err != nil {
//		logger.Error("赠送徽章失败", slog.String("error", err.Error()))
//		return
//	}
//	logger.Info("赠送徽章结果", slog.Any("resp", resp))
//}
//
//func postRemoveMetal() {
//	resp, err := client.PostRemoveMetal(username, "最受欢迎设计师S1测试徽章")
//	if err != nil {
//		logger.Error("移除徽章失败", slog.String("error", err.Error()))
//		return
//	}
//	logger.Info("移除徽章结果", slog.Any("resp", resp))
//}
//
//func postRemoveMetalByUserId() {
//	resp, err := client.PostRemoveMetalByUserId("userOId", "最受欢迎设计师S1测试徽章")
//	if err != nil {
//		logger.Error("通过用户ID移除徽章失败", slog.String("error", err.Error()))
//		return
//	}
//	logger.Info("通过用户ID移除徽章结果", slog.Any("resp", resp))
//}

func postUserQueryItems() {
	resp, err := client.PostUserQueryItems(username)
	if err != nil {
		logger.Error("查询用户物品失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("查询用户物品结果", slog.Any("resp", resp))
}

func postUserEditItems() {
	resp, err := client.PostUserEditItems(username, types.ItemTypeSysCheckinRemain, 0)
	if err != nil {
		logger.Error("编辑用户物品失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("编辑用户物品结果", slog.Any("resp", resp))
}

func postUserEditPoints() {
	resp, err := client.PostUserEditPoints(username, -1, "开摆接口测试调整积分")
	if err != nil {
		logger.Error("编辑用户积分失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("编辑用户积分结果", slog.Any("resp", resp))
}

func postUserLiveness() {
	resp, err := client.PostUserLiveness(username)
	if err != nil {
		logger.Error("获取用户活跃度失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("获取用户活跃度结果", slog.Any("resp", resp))
}

func postYesterdayLivenessReward() {
	resp, err := client.PostYesterdayLivenessReward(username)
	if err != nil {
		logger.Error("领取用户昨日活跃度奖励失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("领取用户昨日活跃度奖励结果", slog.Any("resp", resp))
}

func postUserEditNotification() {
	resp, err := client.PostUserEditNotification(username, "这是一条来自接口测试的通知内容")
	if err != nil {
		logger.Error("发送指定用户通知失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("发送指定用户通知结果", slog.Any("resp", resp))
}

func postApiGetKey() {
	if err := client.PostApiGetKey(); err != nil {
		logger.Error("获取API密钥失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("获取API密钥成功")
}

func getApiUser() {
	resp, err := client.GetApiUser()
	if err != nil {
		logger.Error("获取API用户信息失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("API用户信息", slog.Any("resp", resp))
}

func userChannelWebsocket() {
	ws := client.NewUserNotificationWebSocket(
		sdk.WithHeartbeat[types.UserChannelMsg](30*time.Second, func() []byte {
			// 可以返回自定义心跳消息
			// 这里返回一个简单的ping消息
			msg, _ := json.Marshal(map[string]string{"type": "ping"})
			slog.Info("发送心跳消息", slog.String("message", string(msg)))
			return msg
		}),
	)

	// 设置连接成功回调
	ws.OnOpen(func(c *sdk.Client[types.UserChannelMsg]) {
		slog.Debug("[系统] WebSocket连接已打开")
	})

	// 设置消息回调
	ws.OnMessage(func(msg *types.UserChannelMsg) {
		switch msg.Command {
		case types.UserChannelCommandBzUpdate:
			slog.Info("[通知] 收到业务数据更新通知", slog.Any("msg", msg))
		case types.UserChannelCommandChatUnreadCountRefresh:
			slog.Info("[通知] 收到聊天未读消息数刷新通知", slog.Any("msg", msg))
		case types.UserChannelCommandRefreshNotification:
			slog.Info("[通知] 收到通知刷新请求", slog.Any("msg", msg))
		case types.UserChannelCommandNewIdleChatMessage:
			slog.Info("[通知] 收到新的闲聊消息通知", slog.Any("msg", msg))
		case types.UserChannelCommandWarnBroadcast:
			slog.Info("[通知] 收到系统广播通知", slog.Any("msg", msg))
		default:
			slog.Warn("[通知] 收到未知类型的消息", slog.Any("msg", msg))
		}
	})

	// 设置错误回调
	ws.OnError(func(err error) {
		slog.Error("[系统] WebSocket连接发生错误", slog.Any("error", err))
	})

	// 设置关闭回调
	ws.OnClose(func() {
		slog.Warn("[系统] WebSocket连接已关闭")
	})

	// 连接
	slog.Debug("正在连接到用户通知服务...")
	if err := ws.Connect(); err != nil {
		slog.Error("连接失败", slog.Any("error", err))
	}
	slog.Debug("已连接到用户通知服务！")

	// 设置信号处理，优雅退出
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 保持程序运行
	slog.Debug("用户通知服务已就绪，按 Ctrl+C 退出...")
	<-sigChan

	// 关闭连接
	slog.Debug("正在断开连接...")
	if err := ws.Close(); err != nil {
		slog.Error("关闭连接失败", slog.Any("error", err))
	}
	slog.Debug("已断开连接")
}

func chatChannelWebsocket() {

	ws := client.NewPrivateChatWebSocket(
		botUserName,
		// 可选：禁用自动重连
		// sdk.WithAutoReconnect[types.ChatMessage](false),
	)

	// 设置开始回调
	ws.OnOpen(func(client *sdk.Client[types.ChatChannelMsg]) {
		slog.Info("[系统] WebSocket连接已启动")
	})

	// 设置消息回调
	ws.OnMessage(func(msg *types.ChatChannelMsg) {
		slog.Info("[私聊消息] 来自用户 %s 的消息: %s", msg.SenderUserName, msg.Content)
	})

	// 设置错误回调
	ws.OnError(func(err error) {
		slog.Error("[错误] WebSocket错误", slog.Any("error", err))
	})

	// 设置关闭回调
	ws.OnClose(func() {
		slog.Warn("[系统] WebSocket连接已关闭")
	})

	// 连接
	slog.Info("正在连接到私聊服务...")
	if err := ws.Connect(); err != nil {
		slog.Error("连接失败", slog.Any("error", err))
		os.Exit(1)
	}
	slog.Info("已连接到私聊服务！")

	//示例：发送私聊消息
	if err := ws.SendMessage("帮助"); err != nil {
		slog.Error("发送消息失败", slog.Any("error", err))
	} else {
		slog.Info("已发送私聊消息：帮助")
	}

	// 设置信号处理，优雅退出
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 保持程序运行
	slog.Info("私聊服务已就绪，按 Ctrl+C 退出...")
	<-sigChan

	// 关闭连接
	slog.Info("正在断开连接...")
	if err := ws.Close(); err != nil {
		slog.Error("关闭连接失败", slog.Any("error", err))
	}
	slog.Info("已断开连接")
}

func chatroomWebsocket() {

	resp, err := client.GetChatroomNode()
	if err != nil {
		slog.Error("获取聊天室节点信息失败", slog.Any("error", err))
		return
	}
	if resp.Code != 0 {
		slog.Error("获取聊天室节点信息失败", slog.Any("resp", resp))
		return
	}
	slog.Info("聊天室节点信息结果", slog.Any("resp", resp))

	wsEndpoint := resp.Data

	// 使用函数式选项配置WebSocket客户端
	ws := client.NewChatroomWebSocket(
		wsEndpoint,
		// 可选：配置日志级别
		sdk.WithLogger[types.ChatroomMsg](slog.Default()),
		// 可选：配置重连策略（默认已启用指数退避）
		// sdk.WithReconnectStrategy[types.ChatroomMessage](&sdk.ExponentialBackoffStrategy{
		// 	BaseDelay:  2 * time.Second,
		// 	MaxDelay:   30 * time.Second,
		// 	Multiplier: 1.5,
		// }),
		// 可选：配置最大重连次数（0表示无限重连）
		// sdk.WithMaxReconnectAttempts[types.ChatroomMessage](10),
		// 可选：配置重连失败回调
		sdk.WithReconnectFailedCallback[types.ChatroomMsg](func(attempts int, err error) {
			slog.Error("重连失败", slog.Int("失败次数", attempts), slog.Any("错误", err))
		}),
	)

	// 设置消息回调
	ws.OnMessage(func(msg *types.ChatroomMsg) {
		metals, metalErr := msg.GetMetalList()
		slog.Info("[聊天室消息] 来自用户的消息",
			slog.Any("type", msg.Type),
			slog.String("name", msg.UserName),
			slog.String("nickname", msg.UserNickname),
			slog.Any("metals", metals),
			slog.Any("metalErr", metalErr),
			slog.Any("content", msg.Md),
			slog.Any("jsonInfo", msg.GetJsonInfo()),
			slog.Bool("jsonInfoNil", msg.GetJsonInfo() == nil),
		)
	})

	// 设置错误回调
	ws.OnError(func(err error) {
		slog.Error("[错误] WebSocket错误", slog.Any("error", err))
	})

	// 设置关闭回调
	ws.OnClose(func() {
		slog.Debug("[系统] WebSocket连接已关闭")
	})

	// 连接到聊天室
	slog.Debug("正在连接到聊天室...")
	if err = ws.Connect(); err != nil {
		slog.Error("连接失败", slog.Any("error", err))
		return
	}
	slog.Debug("已连接到聊天室！")

	// 等待一下确保连接建立
	time.Sleep(2 * time.Second)

	//// 发送一条消息
	//fmt.Println("\n发送消息: 大家好，我是Golang SDK！")
	//if err := ws.SendMessage("大家好，我是Golang SDK！"); err != nil {
	//	log.Printf("发送消息失败: %v", err)
	//}

	// 设置信号处理，优雅退出
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 保持程序运行
	slog.Debug("聊天室已就绪，按 Ctrl+C 退出...")
	<-sigChan

	// 关闭连接
	slog.Debug("正在断开连接...")
	if err = ws.Close(); err != nil {
		slog.Error("关闭连接失败", slog.Any("error", err))
	}
	slog.Debug("已断开连接")
}

func articleChannelWebsocket() {

	// 使用函数式选项配置WebSocket客户端
	ws := client.NewArticleChannelWebSocket(
		editArticleId,
		types.ArticleTypeThought,
		// 可选：配置日志级别
		sdk.WithLogger[types.ArticleChannelMsg](slog.Default()),
		// 可选：配置重连失败回调
		sdk.WithReconnectFailedCallback[types.ArticleChannelMsg](func(attempts int, err error) {
			slog.Error("重连失败", slog.Int("失败次数", attempts), slog.Any("错误", err))
		}),
	)

	// 设置消息回调
	ws.OnMessage(func(msg *types.ArticleChannelMsg) {
		slog.Info("[文章频道消息] 来自用户的消息",
			slog.Any("msg", msg),
		)
	})

	// 设置错误回调
	ws.OnError(func(err error) {
		slog.Error("[错误] WebSocket错误", slog.Any("error", err))
	})

	// 设置关闭回调
	ws.OnClose(func() {
		slog.Debug("[系统] WebSocket连接已关闭")
	})

	// 连接到文章频道
	slog.Debug("正在连接到文章频道...")
	if err := ws.Connect(); err != nil {
		slog.Error("连接失败", slog.Any("error", err))
		return
	}
	slog.Debug("已连接到文章频道！")

	// 等待一下确保连接建立
	time.Sleep(2 * time.Second)

	// 设置信号处理，优雅退出
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 保持程序运行
	slog.Debug("文章频道已就绪，按 Ctrl+C 退出...")
	<-sigChan

	// 关闭连接
	slog.Debug("正在断开连接...")
	if err := ws.Close(); err != nil {
		slog.Error("关闭连接失败", slog.Any("error", err))
	}
	slog.Debug("已断开连接")
}

func getUserInfoById() {
	resp, err := client.GetUserInfoById(userIdYui)
	if err != nil {
		logger.Error("通过用户ID获取用户信息失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("通过用户ID获取用户信息结果", slog.Any("resp", resp))
}

func getOpenIdUrl() {
	link := client.GetOpenIdUrl(realm, returnTo)

	logger.Info("获取OpenID链接结果", slog.String("link", link))
}

func runOpenIdServer() {
	const baseAddr = "https://test.aweoo.top"
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		addr := client.GetOpenIdUrl(baseAddr, baseAddr+"/openid/callback")
		logger.Info("重定向到OpenID登录页面", slog.String("addr", addr))
		http.Redirect(writer, request, addr, http.StatusFound)
	})
	http.HandleFunc("/openid/callback", func(w http.ResponseWriter, r *http.Request) {

		var err error

		q := r.URL.Query()

		query := make(map[string]string)
		for key := range q {
			query[key] = q.Get(key)
		}

		var openId *string
		if openId, err = client.PostOpenIdVerify(query); err != nil {
			logger.Error("验证OpenID失败", slog.String("error", err.Error()))
			http.Error(w, "验证OpenID失败", http.StatusBadRequest)
		}
		logger.Info("OpenID验证成功", slog.String("openId", convertor.ToString(openId)))

		resp := new(types.ApiResponse[*types.GetUserInfoByIdData])
		if resp, err = client.GetUserInfoById(convertor.ToString(openId)); err != nil {
			logger.Error("通过OpenID获取用户信息失败", slog.String("error", err.Error()))
			http.Error(w, "通过OpenID获取用户信息失败", http.StatusBadRequest)
			return
		}
		if resp.Code != 0 {
			logger.Error("通过OpenID获取用户信息失败", slog.Any("resp", resp))
			http.Error(w, "通过OpenID获取用户信息失败", http.StatusBadRequest)
			return
		}
		logger.Info("获取到用户信息", slog.Any("user", resp.Data))

		// 响应成功页面
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte("<html><body><h1>登录成功！</h1><p>欢迎，" + resp.Data.UserNickname + "！</p></body></html>"))
	})
	logger.Info("启动OpenID回调服务器，监听端口6666")
	if err := http.ListenAndServe(":6666", nil); err != nil {
		logger.Error("启动OpenID回调服务器失败", slog.String("error", err.Error()))
	}
}

func postMedalMyList() {
	resp, err := client.PostMedalMyList()
	if err != nil {
		logger.Error("获取我的徽章列表失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("获取我的徽章列表结果", slog.Any("resp", resp))
}

func postMedalMyReorder() {
	resp, err := client.PostMedalMyReorder(medalIdYearOne, types.MedalReorderDirectionUp)
	if err != nil {
		logger.Error("重新排序我的徽章失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("重新排序我的徽章结果", slog.Any("resp", resp))
}

func postMedalMyDisplay() {
	resp, err := client.PostMedalMyDisplay(medalIdFans16, false)
	if err != nil {
		logger.Error("设置我的徽章显示状态失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("设置我的徽章显示状态结果", slog.Any("resp", resp))
}

func postMedalUserList() {
	resp, err := client.PostMedalUserList(&types.PostMedalUserListRequest{
		//UserId: userIdYui,
		UserName: username,
	})
	if err != nil {
		logger.Error("获取用户徽章列表失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("获取用户徽章列表结果", slog.Any("resp", resp))
}

func postMedalAdminList() {
	resp, err := client.PostMedalAdminList(1, 100)
	if err != nil {
		logger.Error("获取所有徽章列表失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("获取所有徽章列表结果", slog.Any("resp", resp))
}

func postMedalAdminSearch() {
	resp, err := client.PostMedalAdminSearch("开摆")
	if err != nil {
		logger.Error("搜索徽章失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("搜索徽章结果", slog.Any("resp", resp))
}

func postMedalAdminDetail() {
	resp, err := client.PostMedalAdminDetail(medalIdTest)
	if err != nil {
		logger.Error("获取徽章详情失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("获取徽章详情结果", slog.Any("resp", resp))
}

func postMedalAdminDelete() {
	resp, err := client.PostMedalAdminDelete(medalIdTest)
	if err != nil {
		logger.Error("删除徽章失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("删除徽章结果", slog.Any("resp", resp))
}

func postMedalAdminEdit() {
	resp, err := client.PostMedalAdminEdit(
		medalIdTest,
		"开摆的测试勋章",
		types.MedalTypeNormal,
		"这是一个用于测试的开摆勋章，请勿赠送。尊贵的{var1}{var2}，感谢您{var3}天来的支持！更新啦～",
		"url=https://file.fishpi.cn/2025/06/image-6e223db8.png&backcolor=000000&fontcolor=ffffff",
	)
	if err != nil {
		logger.Error("编辑徽章失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("编辑徽章结果", slog.Any("resp", resp))
}

func postMedalAdminCreate() {
	resp, err := client.PostMedalAdminCreate(
		"开摆的测试勋章",
		types.MedalTypeNormal,
		"这是一个用于测试的开摆勋章，请勿赠送。尊贵的{var1}{var2}，感谢您{var3}天来的支持！",
		"url=https://file.fishpi.cn/2025/06/image-6e223db8.png&backcolor=000000&fontcolor=ffffff",
	)
	if err != nil {
		logger.Error("创建徽章失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("创建徽章结果", slog.Any("resp", resp))
}

func postMedalAdminGrant() {
	resp, err := client.PostMedalAdminGrant(
		userId8888,
		medalIdTest,
		time.Now().Add(time.Hour*24*30).UnixMilli(),
		"开摆;开发者;9999",
	)
	if err != nil {
		logger.Error("管理员赠送徽章失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("管理员赠送徽章结果", slog.Any("resp", resp))
}

func postMedalAdminRevoke() {
	resp, err := client.PostMedalAdminRevoke(
		userId8888,
		medalIdTest,
	)
	if err != nil {
		logger.Error("管理员收回徽章失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("管理员收回徽章结果", slog.Any("resp", resp))
}

func postMedalAdminOwners() {
	resp, err := client.PostMedalAdminOwners(medalIdTest, 1, 100)
	if err != nil {
		logger.Error("获取徽章拥有者列表失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("获取徽章拥有者列表结果", slog.Any("resp", resp))
}

func getMedalUrl() {
	link := client.GetMedalUrl(&types.GetMedalUrlRequest{
		//MedalId: medalIdYearOne,
		MedalName: medalNameSister,
	})
	logger.Info("获取徽章图片链接结果", slog.String("link", link))
}

func postMedalAdminUserMedals() {
	resp, err := client.PostMedalAdminUserMedals(&types.PostMedalAdminUserMedalsRequest{
		//UserId: userIdYui,
		UserName: username,
	})
	if err != nil {
		logger.Error("获取用户所有徽章列表失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("获取用户所有徽章列表结果", slog.Any("resp", resp))
}

func postArticleReaction() {
	resp, err := client.PostArticleReaction(editArticleId, types.ReactionGroupTypeEmoji, types.ReactionEmojiValueThinking)
	if err != nil {
		logger.Error("文章表态失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("文章表态结果", slog.Any("resp", resp))
}

func postCommentReaction() {
	resp, err := client.PostCommentReaction(editArticleCommentId, types.ReactionGroupTypeEmoji, types.ReactionEmojiValueThinking)
	if err != nil {
		logger.Error("评论表态失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("评论表态结果", slog.Any("resp", resp))
}

func postChatRoomReaction() {
	resp, err := client.PostChatRoomReaction("1777014957669", types.ReactionGroupTypeEmoji, types.ReactionEmojiValueThinking)
	if err != nil {
		logger.Error("聊天室表态失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("聊天室表态结果", slog.Any("resp", resp))
}
