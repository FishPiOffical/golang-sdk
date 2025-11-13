package main

import (
	"fishpi-golang-sdk/sdk"
	"fmt"
	"os"
)

func main3() {
	fmt.Println("=== 测试 WithLogDir 修复 ===")

	// 创建配置
	config := &sdk.Config{
		ApiKey: "test-key",
	}
	provider := sdk.NewMemoryConfigProvider(config)

	// 创建SDK - 只测试WithLogDir选项
	fmt.Println("创建SDK with WithLogDir...")
	client := sdk.NewSDK(
		provider,
		sdk.WithLogDir("./test_logs"),
	)

	if client != nil {
		fmt.Println("✅ SDK 创建成功!")

		// 检查目录是否创建
		if _, err := os.Stat("./test_logs"); err == nil {
			fmt.Println("✅ 日志目录已创建: ./test_logs")
		} else {
			fmt.Printf("❌ 日志目录创建失败: %v\n", err)
		}
	} else {
		fmt.Println("❌ SDK 创建失败")
	}

	fmt.Println("\n修复说明:")
	fmt.Println("- 之前使用 SetOutputFile().EnableDump() 会导致 'file already closed' 错误")
	fmt.Println("- 现在使用 EnableDumpToFile(path) 正确地保存日志")
	fmt.Println("- 这样不会干扰响应体的正常读取")
}
