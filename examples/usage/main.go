package main

import (
	"github.com/FishPiOffical/golang-sdk/config"
	"github.com/FishPiOffical/golang-sdk/sdk"
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

	username      = "8888"
	associateName = "888"
)

func main() {

	// 鉴权

	// 通用
	//getUserInfoByUsername()
	//postUsersNames()
	//getUsersEmotions()

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
