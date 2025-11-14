package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	sdk "github.com/FishPiOffical/golang-sdk/sdk"
	"github.com/FishPiOffical/golang-sdk/types"
)

func main() {
	// 从环境变量获取API Key
	apiKey := os.Getenv("FISHPI_API_KEY")
	if apiKey == "" {
		log.Fatal("请设置环境变量 FISHPI_API_KEY")
	}

	// 创建SDK实例
	fishpi := sdk.NewSDKWithAPIKey(apiKey)

	// 创建用户通知WebSocket连接
	ws := fishpi.NewUserNotificationWebSocket()

	// 设置消息回调
	ws.OnMessage(func(msg *types.UserMessage) {
		switch msg.Type {
		case "article":
			// 文章通知
			fmt.Println("\n[通知] 收到文章通知")
			// 可以解析msg.Data获取具体信息

		case "breezemoon":
			// 清风明月通知
			fmt.Println("\n[通知] 收到清风明月通知")

		case "comment":
			// 评论通知
			fmt.Println("\n[通知] 收到评论通知")

		case "at":
			// @ 通知
			fmt.Println("\n[通知] 有人@了你")

		case "following":
			// 关注通知
			fmt.Println("\n[通知] 你关注的用户有新动态")

		case "point":
			// 积分通知
			fmt.Println("\n[通知] 积分变动通知")

		default:
			fmt.Printf("\n[通知] 未知类型: %s\n", msg.Type)
		}
	})

	// 设置错误回调
	ws.OnError(func(err error) {
		log.Printf("\n[错误] WebSocket错误: %v\n", err)
	})

	// 设置关闭回调
	ws.OnClose(func() {
		log.Println("\n[系统] WebSocket连接已关闭")
	})

	// 连接
	fmt.Println("正在连接到用户通知服务...")
	if err := ws.Connect(); err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	fmt.Println("已连接到用户通知服务！")

	// 设置信号处理，优雅退出
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 保持程序运行
	fmt.Println("\n用户通知服务已就绪，按 Ctrl+C 退出...")
	<-sigChan

	// 关闭连接
	fmt.Println("\n正在断开连接...")
	if err := ws.Close(); err != nil {
		log.Printf("关闭连接失败: %v", err)
	}
	fmt.Println("已断开连接")
}
