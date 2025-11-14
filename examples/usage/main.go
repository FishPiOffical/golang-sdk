package main

import (
	"github.com/FishPiOffical/golang-sdk/sdk"
	"github.com/golang-cz/devslog"

	"log/slog"
	"os"
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

	logger := slog.New(devslog.NewHandler(os.Stdout, opts))
	slog.SetDefault(logger)
}

func main() {

	logger := slog.Default()

	// 从环境变量读取API Key，或使用默认值
	apiKey := os.Getenv("FISHPI_API_KEY")
	if apiKey == "" {
		logger.Warn("警告: 未设置 FISHPI_API_KEY 环境变量")
		logger.Warn("使用方式: export FISHPI_API_KEY=your-api-key")
		apiKey = "your-api-key-here"
	}

	// 创建配置
	config := &sdk.Config{
		ApiKey: apiKey,
	}
	provider := sdk.NewMemoryConfigProvider(config)

	// 使用选项创建SDK
	client := sdk.NewSDK(
		provider,
		sdk.WithLogDir("../../_tmp/logs/"),        // 设置日志目录
		sdk.WithCustomUnmarshaler(slog.Default()), // 设置自定义反序列化器
	)

	user, err := client.GetApiUser()
	if err != nil {
		logger.Error("获取用户信息失败", slog.String("error", err.Error()))
		return
	}
	if user.Code != 0 {
		logger.Error("获取用户信息失败", slog.Any("resp", user))
		return
	}
	logger.Info("用户信息", slog.Any("user", user.Data.UserNickname))
}
