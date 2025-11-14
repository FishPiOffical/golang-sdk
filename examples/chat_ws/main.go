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

	// 创建私聊WebSocket连接
	ws := fishpi.NewPrivateChatWebSocket()

	// 设置消息回调
	ws.OnMessage(func(msg *types.ChatMessage) {
		if msg.Type == "msg" {
			fmt.Printf("\n[私聊] %s: %s\n",
				msg.Data.SenderUserName,
				msg.Data.Content)
		} else if msg.Type == "revoke" {
			fmt.Printf("\n[系统] 消息已撤回: %s\n", msg.Data.OId)
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
	fmt.Println("正在连接到私聊服务...")
	if err := ws.Connect(); err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	fmt.Println("已连接到私聊服务！")

	// 示例：发送私聊消息
	// ws.SendMessage("targetUserName", "你好！")

	// 设置信号处理，优雅退出
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 保持程序运行
	fmt.Println("\n私聊服务已就绪，按 Ctrl+C 退出...")
	<-sigChan

	// 关闭连接
	fmt.Println("\n正在断开连接...")
	if err := ws.Close(); err != nil {
		log.Printf("关闭连接失败: %v", err)
	}
	fmt.Println("已断开连接")
}
