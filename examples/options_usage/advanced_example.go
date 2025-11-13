package main

import (
	"log"
	"log/slog"
	"os"

	"fishpi-golang-sdk/sdk"
)

func main1() {
	// 示例1: 只使用日志目录选项
	example1()

	// 示例2: 只使用自定义反序列化器选项
	example2()

	// 示例3: 同时使用两个选项
	example3()
}

// example1 演示如何使用日志目录选项
func example1() {
	log.Println("=== 示例1: 使用日志目录选项 ===")

	config := &sdk.Config{
		ApiKey: "your-api-key-here",
	}
	provider := sdk.NewMemoryConfigProvider(config)

	// 创建SDK实例，启用请求日志
	client := sdk.NewSDK(
		provider,
		sdk.WithLogDir("./logs"), // 所有请求/响应将保存到 ./logs 目录
	)

	// 发起请求，日志将自动保存
	userInfo, err := client.GetApiUser()
	if err != nil {
		log.Printf("获取用户信息失败: %v\n", err)
		return
	}

	log.Printf("用户信息: %+v\n", userInfo)
	log.Println("请求日志已保存到 ./logs 目录")
}

// example2 演示如何使用自定义反序列化器选项
func example2() {
	log.Println("\n=== 示例2: 使用自定义反序列化器选项 ===")

	// 创建自定义logger，只显示WARNING及以上级别的日志
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelWarn,
	}))

	config := &sdk.Config{
		ApiKey: "your-api-key-here",
	}
	provider := sdk.NewMemoryConfigProvider(config)

	// 创建SDK实例，启用字段差异检测
	client := sdk.NewSDK(
		provider,
		sdk.WithCustomUnmarshaler(logger), // 当JSON字段和结构体不匹配时会输出警告
	)

	// 发起请求，如果响应JSON中有结构体未定义的字段，会输出警告日志
	userInfo, err := client.GetApiUser()
	if err != nil {
		log.Printf("获取用户信息失败: %v\n", err)
		return
	}

	log.Printf("用户信息: %+v\n", userInfo)
	log.Println("如果有未匹配的字段，上面会显示警告信息")
}

// example3 演示如何同时使用两个选项
func example3() {
	log.Println("\n=== 示例3: 同时使用两个选项 ===")

	// 创建JSON格式的logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelWarn,
	}))

	config := &sdk.Config{
		ApiKey: "your-api-key-here",
	}
	provider := sdk.NewMemoryConfigProvider(config)

	// 创建SDK实例，同时启用日志和字段差异检测
	client := sdk.NewSDK(
		provider,
		sdk.WithLogDir("./logs"),          // 保存请求日志
		sdk.WithCustomUnmarshaler(logger), // 检测字段差异
	)

	// 发起请求
	userInfo, err := client.GetApiUser()
	if err != nil {
		log.Printf("获取用户信息失败: %v\n", err)
		return
	}

	log.Printf("用户信息: %+v\n", userInfo)
	log.Println("请求日志已保存，字段差异（如有）已输出")
}
