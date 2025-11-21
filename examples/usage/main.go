package main

import (
	"github.com/FishPiOffical/golang-sdk/config"
	"github.com/FishPiOffical/golang-sdk/sdk"
	"github.com/FishPiOffical/golang-sdk/types"
	"github.com/duke-git/lancet/v2/convertor"
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

	username        = "8888"
	associateName   = "888"
	reportArticleId = "1702103071389" // https://fishpi.cn/article/1702103071389
	followingId     = "1734578210153" // https://fishpi.cn/member/wordsKing
	messageOId      = "1763542689788"
	uploadFile1     = "../../_tmp/files/IMG_1045.jpg"
	uploadFile2     = "../../_tmp/files/IMG_13069.jpeg"
	editArticleId   = "1763623304114"
)

func main() {

	// 鉴权

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
	//deleteChatroomRevoke()
	//getMessageRaw()
	//postRedPacketSend()
	//postCloudGet()
	//postCloudSync()
	//getSiGuoYa()

	// 图床
	//postUploadFile()

	// 帖子
	//postArticle()
	//putArticle()
	//getArticles()
	//getArticleDetail()
	getUserArticles()

	// 清风明月
	//getBreezemoons()
	//postBreezemoon()
	//getUserBreezemoons()

}

func getUserInfoByUsername() {
	user, err := client.GetUserInfoByUsername(username)
	if err != nil {
		logger.Error("获取用户信息失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("用户信息", slog.Any("user", user.UserNickname))
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

func postCloudGet() {
	resp, err := client.PostCloudGet(types.CloudGameIdEmojis)
	if err != nil {
		logger.Error("获取云游戏资源失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("云游戏资源结果", slog.Any("resp", resp))
}

func postCloudSync() {
	data := "[\"https://fishpi.cn/gen?scale=0.79&txt=%E5%A5%BD%E5%AE%B6%E4%BC%99&url=https://file.fishpi.cn/2022/06/blob-4a5a6682.png&backcolor=FF00FF&fontcolor=00ff00\",\"https://file.fishpi.cn/2022/01/6W4RCDQ1B3E8IXWDD-3b1dcc36.jpg\",\"https://file.fishpi.cn/vditor/dist/images/emoji/huaji.gif\",\"https://file.fishpi.cn/2022/07/4XD412P6Z3NUADURUIF-cf6e5c49.png\",\"https://file.fishpi.cn/2022/06/image-2fa0b65d.png\",\"https://file.fishpi.cn/2021/11/WY3LZN1MAI1DDONT-e2caa6db.gif\",\"https://file.fishpi.cn/2021/11/image-af346e1a.png\",\"https://file.fishpi.cn/2021/11/video1-f50f21b9.gif\",\"https://tse2-mm.cn.bing.net/th/id/OIP-C.FPgxOL4Y3y45okElx5pzSwAAAA?pid=ImgDet&rs=1\",\"https://unv-shield.librian.net/api/unv_shield?scale=0.79&txt=%E6%88%91%E6%98%AF%E5%BA%9F%E7%89%A9&url=https://www.lingmx.com/52pj/images/die.png&backcolor=568289&fontcolor=ffffff\",\"https://file.fishpi.cn/2022/03/606AE90506E645189E8C89DD0EABE688-dbdd8f22.gif\",\"https://file.fishpi.cn/2022/04/B6ARF4PTIB6JZX-d28e0ddf.jpg\",\"https://file.fishpi.cn/2022/07/1D6841B4-62e975d6.jpg\",\"https://unv-shield.librian.net/api/unv_shield?scale=0.79&txt=lsp%E4%B9%8B%E7%8E%8B%E9%9D%9E%E4%BD%A0%E8%8E%AB%E5%B1%9E&url=https://www.lingmx.com/52pj/images/die.png&backcolor=568289&fontcolor=ffffff\",\"https://file.fishpi.cn/2022/07/image-1900c9b7.png\",\"https://file.fishpi.cn/2022/06/image-be82885f.png\",\"https://file.fishpi.cn/2022/06/image-e04d61d9.png\",\"https://file.fishpi.cn/2022/07/image-abc9d3f6.png\",\"https://file.fishpi.cn/2021/12/image-7d8f1284.png\",\"https://file.fishpi.cn/2022/07/DCD128E70C5B10C49E4D66C8568D10EA-697825f2.jpg\",\"https://file.fishpi.cn/2022/03/image-56c0f695.png\",\"https://file.fishpi.cn/2022/08/34F6E4E51347D5BB164156FF4DFE62BA-9c2cec25.jpg\",\"https://file.fishpi.cn/2021/12/Snipaste20211210094758-990d6133.png\",\"https://file.fishpi.cn/2022/08/60be1bcbc1608b95-1d4762b3.png\",\"https://file.fishpi.cn/2022/08/09A3D30F-7bbd8d34.jpg\",\"https://file.fishpi.cn/2022/06/CB3456129E25827BA532135CB7680B64-0336e313.gif\",\"https://unv-shield.librian.net/api/unv_shield?scale=0.79&txt=%E9%AD%85%E9%AD%94&url=https://pic.stackoverflow.wiki/uploadImages/117/28/120/230/2022/09/10/11/50/29fc900f-2dd0-472d-a772-5c03a9b03156.jpg&backcolor=211a1a&fontcolor=d068da\",\"https://file.fishpi.cn/2022/06/adgif156b20ea-68910939.gif\",\"https://file.fishpi.cn/2022/09/image-3911ca4a.png\",\"https://file.fishpi.cn/2022/11/%E5%9B%BE%E7%89%87-1f3a0fc4.png\",\"https://file.fishpi.cn/2022/12/image-dd15392c.png\",\"https://file.fishpi.cn/2022/09/image-cd84a59e.png\",\"https://file.fishpi.cn/2022/09/%E5%9B%BE%E7%89%87-846950bf.png\",\"https://file.fishpi.cn/2022/12/image-a00a8f3f.png\",\"https://file.fishpi.cn/2022/11/nd83w-888eb73e.gif\",\"https://static.dingtalk.com/media/lQHPJxZuEq0iWwXNAUDNAUCwp1JKstRBmeMCtWjqjoC9AA_320_320.gif?bizType=im\",\"https://file.fishpi.cn/2023/04/20200404023812BehxC-ea84b1b9.gif\",\"https://file.fishpi.cn/2022/01/image-3d8ac437.png\",\"https://file.fishpi.cn/2023/04/1J15M-cc5e47f4.gif\",\"https://file.fishpi.cn/2023/06/ne3eZ-b9e96820.gif\",\"https://file.fishpi.cn/2022/08/XBN9CGP0CHLL2MDZA9L-fe17f120.jpg\",\"https://file.fishpi.cn/2023/08/e5cacb49f4ab060451980e55b06be6d-8805a569.jpg\",\"https://file.fishpi.cn/2023/01/image-f12bf156.png\",\"https://file.fishpi.cn/2023/09/image-eac3523d.png\",\"https://file.fishpi.cn/2023/08/image-d54e2f1d.png\",\"https://file.fishpi.cn/2023/09/image-02ce397b.png\",\"https://fishpi.cn/gen?scale=0.79&txt=%E5%85%84%E5%BC%9F%E4%BB%AC%EF%BC%8C%E5%8F%91%E7%82%B9%E6%B6%A9%E5%9B%BE&url=https://file.fishpi.cn/2022/06/blob-4a5a6682.png&backcolor=FF00FF&fontcolor=00ff00\",\"https://file.fishpi.cn/2023/09/image-26a9d52a.png\",\"https://file.fishpi.cn/2023/09/image-d661fbb3.png\",\"https://file.fishpi.cn/2023/04/image-354d2505.png\",\"https://file.fishpi.cn/2021/11/3I3V8NZ1SBZA35N2S7U1-6a567673.png\",\"https://file.fishpi.cn/2022/09/image-5beacdbd.png\",\"https://file.fishpi.cn/2023/08/image-a9baf2e1.png\",\"https://file.fishpi.cn/2023/08/image-36d1b849.png\",\"https://file.fishpi.cn/2023/10/image-a7e13101.png\",\"https://file.fishpi.cn/2023/10/image-63d9fff9.png\",\"https://file.fishpi.cn/2023/10/image-d74b98cc.png\",\"https://file.fishpi.cn/2023/03/AP2R9HZQBHRIYUK-a6695524.jpg\",\"https://file.fishpi.cn/2023/12/image-5917d1f8.png\",\"https://file.fishpi.cn/2023/12/image-db7980ae.png\",\"https://file.fishpi.cn/2023/12/%E5%9B%BE%E7%89%87-3789ec89.png\",\"https://file.fishpi.cn/2024/01/3cc55d75c8fe412fa6d7e12ca7729355-8c1fefde.png\",\"https://file.fishpi.cn/2024/01/20230706034848avns286q-faa5eb9a.jpg\",\"https://file.fishpi.cn/2023/12/image-a6039b53.png\",\"https://file.fishpi.cn/2024/03/image-80eb3cfb.png\",\"https://file.fishpi.cn/2022/01/a9cf8ef6ly1fiecn56l8wj20b50b2glu-12a17eea.jpg\",\"https://file.fishpi.cn/2024/04/6b4cd9ccd007236eda98d7f80f89f97bfc9a6f1b1df5cd96313f7573fe928087ebfc522-4f35ecc4.png\"]"
	resp, err := client.PostCloudSync(types.CloudGameIdEmojis, data)
	if err != nil {
		logger.Error("同步云游戏资源失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("同步云游戏资源结果", slog.Any("resp", resp))
}

func getSiGuoYa() {
	resp, err := client.GetSiGuoYa()
	if err != nil {
		logger.Error("获取思过崖失败", slog.String("error", err.Error()))
		return
	}
	logger.Info("思过崖结果", slog.Any("resp", resp))
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

func postArticle() {
	resp, err := client.PostArticle(&types.PostArticleRequest{
		ArticleTitle:           "AI带来的提升",
		ArticleContent:         "AI已经发展了这么多年，那么AI对你的工作和生活带来了哪些提升呢？  \n> 请详细说明你的实际体验和感受。🍠水贴将会被删除哦！",
		ArticleTags:            "测试,AI,生活",
		ArticleCommentable:     true,
		ArticleNotifyFollowers: false,
		ArticleType:            types.ArticleTypeQna,
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
		ArticleType:            types.ArticleTypeQna,
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
		Order:   convertor.ToPointer(types.GetArticleOrderHot),
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
