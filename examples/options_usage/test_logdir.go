package main

import (
	"log"
	"os"

	"github.com/FishPiOffical/golang-sdk/sdk"
)

func main2() {
	// 只测试 WithLogDir 选项
	log.Println("=== 测试 WithLogDir 选项 ===")

	config := &sdk.Config{
		ApiKey: os.Getenv("FISHPI_API_KEY"), // 从环境变量读取
	}

	if config.ApiKey == "" {
		log.Println("警告: 未设置 FISHPI_API_KEY 环境变量")
		log.Println("使用方式: export FISHPI_API_KEY=your-api-key")
		log.Println("跳过实际API调用测试")

		// 仅测试SDK创建
		provider := sdk.NewMemoryConfigProvider(config)
		client := sdk.NewSDK(
			provider,
			sdk.WithLogDir("./logs"),
		)

		if client != nil {
			log.Println("✅ SDK 创建成功（使用 WithLogDir）")
			log.Println("✅ 日志目录已创建: ./logs")
		}
		return
	}

	provider := sdk.NewMemoryConfigProvider(config)

	// 使用日志目录选项创建SDK
	client := sdk.NewSDK(
		provider,
		sdk.WithLogDir("./logs"),
	)

	// 使用SDK
	log.Println("正在调用 GetApiUser()...")
	userInfo, err := client.GetApiUser()
	if err != nil {
		log.Fatalf("❌ 调用失败: %v", err)
	}

	log.Printf("✅ 调用成功!")
	log.Printf("用户信息: %+v", userInfo)
	log.Printf("日志已保存到 ./logs 目录")
}
