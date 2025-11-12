// Package fishpi_golang_sdk FishPi社区Golang SDK
package fishpi_golang_sdk

import (
	"fishpi-golang-sdk/sdk"
	"fishpi-golang-sdk/types"
)

// SDK别名，方便使用
type (
	// FishPiSDK SDK客户端
	FishPiSDK = sdk.FishPiSDK

	// Client 旧版客户端
	Client = sdk.Client

	// Config 配置
	Config = sdk.Config

	// ConfigProvider 配置提供者
	ConfigProvider = sdk.ConfigProvider
)

// 类型别名
type (
	// UserInfo 用户信息
	UserInfo = types.UserInfo

	// ArticleInfo 文章信息
	ArticleInfo = types.ArticleInfo

	// CommentInfo 评论信息
	CommentInfo = types.CommentInfo

	// BreezemoonInfo 清风明月信息
	BreezemoonInfo = types.BreezemoonInfo

	// ChatMessage 私聊消息
	ChatMessage = types.ChatMessage

	// ChatroomMessage 聊天室消息
	ChatroomMessage = types.ChatroomMessage

	// NotificationInfo 通知信息
	NotificationInfo = types.NotificationInfo
)

// 构造函数别名
var (
	// NewSDK 创建SDK实例
	NewSDK = sdk.NewSDK

	// NewClient 创建旧版Client实例
	NewClient = sdk.NewClient

	// NewFileConfigProvider 创建文件配置提供者
	NewFileConfigProvider = sdk.NewFileConfigProvider
)
