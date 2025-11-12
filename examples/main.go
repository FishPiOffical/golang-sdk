package main

import (
	fishpi "fishpi-golang-sdk"
	"fmt"
	"log"
)

func main() {
	// 基础使用示例
	basicExample()
}
func basicExample() {
	fmt.Println("=== 基础使用示例 ===")
	// 创建配置
	config := &fishpi.Config{
		BaseUrl:  fishpi.BaseUrl,
		Username: "your_username",
		Password: "your_password",
	}
	// 创建配置提供者
	provider := fishpi.NewMemoryConfigProvider(config)
	// 创建客户端
	client := fishpi.NewClient(provider)
	// 登录获取 API Key
	fmt.Println("正在登录...")
	if err := client.PostApiGetKey(); err != nil {
		log.Fatal(err)
	}
	// 获取用户信息
	fmt.Println("获取用户信息...")
	userResp, err := client.GetUserInfo()
	if err != nil {
		log.Fatal(err)
	}
	if userResp.Code != 0 {
		log.Fatalf("获取用户信息失败: %s", userResp.Msg)
	}
	fmt.Printf("用户名: %s\n", userResp.Data.UserName)
	fmt.Printf("昵称: %s\n", userResp.Data.UserNickname)
	fmt.Printf("积分: %d\n", userResp.Data.UserPoint)
}
