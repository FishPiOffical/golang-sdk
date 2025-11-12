package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	sdk "fishpi-golang-sdk/sdk"
	"fishpi-golang-sdk/types"
)

func main() {
	// 从环境变量获取API Key
	apiKey := os.Getenv("FISHPI_API_KEY")
	if apiKey == "" {
		log.Fatal("请设置环境变量 FISHPI_API_KEY")
	}

	// 创建SDK实例
	fishpi := sdk.NewSDK(apiKey)

	// 创建聊天室WebSocket连接
	// 这里直接使用默认节点，实际使用时应该先获取节点列表
	wsEndpoint := "wss://fishpi.cn/chat-room-channel"
	ws := fishpi.NewChatroomWebSocket(wsEndpoint)

	// 设置消息回调
	ws.OnMessage(func(msg *types.ChatroomMessage) {
		switch msg.Type {
		case "msg":
			// 处理聊天消息
			if data, ok := msg.Data.(map[string]interface{}); ok {
				userName := data["userName"]
				content := data["content"]
				fmt.Printf("\n[聊天] %s: %s\n", userName, content)
			}

		case "online":
			// 处理在线用户
			fmt.Println("\n[系统] 在线用户列表更新")

		case "discussChanged":
			// 处理话题变更
			if data, ok := msg.Data.(map[string]interface{}); ok {
				discuss := data["newDiscuss"]
				fmt.Printf("\n[系统] 话题变更: %s\n", discuss)
			}

		case "revoke":
			// 处理消息撤回
			fmt.Println("\n[系统] 消息已撤回")

		case "redPacket":
			// 处理红包消息
			fmt.Println("\n[红包] 收到红包！快去抢！")

		case "redPacketStatus":
			// 处理红包领取状态
			fmt.Println("\n[红包] 红包领取状态更新")

		case "customMessage":
			// 处理自定义消息（进入/离开聊天室）
			if data, ok := msg.Data.(map[string]interface{}); ok {
				content := data["content"]
				fmt.Printf("\n[系统] %s\n", content)
			}

		case "barrager":
			// 处理弹幕消息
			if data, ok := msg.Data.(map[string]interface{}); ok {
				userName := data["userName"]
				content := data["content"]
				fmt.Printf("\n[弹幕] %s: %s\n", userName, content)
			}

		default:
			fmt.Printf("\n[未知消息类型] %s\n", msg.Type)
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

	// 连接到聊天室
	fmt.Println("正在连接到聊天室...")
	if err := ws.Connect(); err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	fmt.Println("已连接到聊天室！")

	// 等待一下确保连接建立
	time.Sleep(2 * time.Second)

	// 发送一条消息
	fmt.Println("\n发送消息: 大家好，我是Golang SDK！")
	if err := ws.SendMessage("大家好，我是Golang SDK！"); err != nil {
		log.Printf("发送消息失败: %v", err)
	}

	// 设置信号处理，优雅退出
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 保持程序运行
	fmt.Println("\n聊天室已就绪，按 Ctrl+C 退出...")
	<-sigChan

	// 关闭连接
	fmt.Println("\n正在断开连接...")
	if err := ws.Close(); err != nil {
		log.Printf("关闭连接失败: %v", err)
	}
	fmt.Println("已断开连接")
}
