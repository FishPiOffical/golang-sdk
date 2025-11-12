// Package fishpi_golang_sdk 提供类型别名以保持向后兼容
// 建议使用 fishpi-golang-sdk/types 和 fishpi-golang-sdk/sdk 包
package fishpi_golang_sdk

import "fishpi-golang-sdk/types"

// 为了保持向后兼容性，创建类型别名

// ApiResponse 通用API响应结构(带Data字段)
type ApiResponse[T any] = types.ApiResponse[T]

// SimpleResponse 简单响应结构(无Data字段)
type SimpleResponse = types.SimpleResponse

// ResponseResult 特殊响应结构(使用sc字段而非code)
type ResponseResult[T any] = types.ResponseResult[T]

// UserInfo 用户信息
type UserInfo = types.UserInfo

// ArticleInfo 文章信息
type ArticleInfo = types.ArticleInfo

// ArticleList 文章列表
type ArticleList = types.ArticleList

// PostArticleRequest 发布文章请求
type PostArticleRequest = types.PostArticleRequest

// CommentInfo 评论信息
type CommentInfo = types.CommentInfo

// PostCommentRequest 发布评论请求
type PostCommentRequest = types.PostCommentRequest

// BreezemoonInfo 清风明月信息
type BreezemoonInfo = types.BreezemoonInfo

// BreezemoonList 清风明月列表
type BreezemoonList = types.BreezemoonList

// PostBreezemoonRequest 发布清风明月请求
type PostBreezemoonRequest = types.PostBreezemoonRequest

// ChatMessage 私聊消息
type ChatMessage = types.ChatMessage

// ChatMessageData 私聊消息数据
type ChatMessageData = types.ChatMessageData

// ChatroomMessage 聊天室消息
type ChatroomMessage = types.ChatroomMessage

// NotificationInfo 通知信息
type NotificationInfo = types.NotificationInfo

// Pagination 分页信息
type Pagination = types.Pagination
