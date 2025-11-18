package main

import (
	"github.com/FishPiOffical/golang-sdk/config"
	"github.com/FishPiOffical/golang-sdk/sdk"
	"github.com/FishPiOffical/golang-sdk/types"
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
