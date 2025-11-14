package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/FishPiOffical/golang-sdk/sdk"
)

func main() {
	// 创建自定义logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelWarn,
	}))

	// 从环境变量读取API Key，或使用默认值
	apiKey := os.Getenv("FISHPI_API_KEY")
	if apiKey == "" {
		log.Println("警告: 未设置 FISHPI_API_KEY 环境变量")
		log.Println("使用方式: export FISHPI_API_KEY=your-api-key")
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
		sdk.WithLogDir("./logs"),          // 设置日志目录
		sdk.WithCustomUnmarshaler(logger), // 设置自定义反序列化器
	)

	log.Println("SDK 已创建，日志将保存到 ./logs 目录")

	// 使用SDK
	userInfo, err := client.GetApiUser()
	if err != nil {
		log.Fatalf("Failed to get user info: %v", err)
	}

	log.Printf("User info: %+v", userInfo)
	log.Println("请求日志已保存到 ./logs 目录")
}
