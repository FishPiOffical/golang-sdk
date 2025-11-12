package fishPiSdk

import "fishpi-golang-sdk/types"

// 为了保持向后兼容性，创建类型别名

// ApiResponse 通用API响应结构(带Data字段)
type ApiResponse[T any] = types.ApiResponse[T]

// SimpleResponse 简单响应结构(无Data字段)
type SimpleResponse = types.SimpleResponse

// ResponseResult 特殊响应结构(使用sc字段而非code)
type ResponseResult[T any] = types.ResponseResult[T]

// PostApiGetKeyRequest 获取ApiKey请求
type PostApiGetKeyRequest = types.PostApiGetKeyRequest

// PostApiGetKeyResponse 获取ApiKey响应
type PostApiGetKeyResponse = types.PostApiGetKeyResponse

// ApiUserGetData 获取用户信息数据（别名）
type ApiUserGetData = types.UserInfo

// GetChatGetMessageData 获取私聊消息数据（已废弃，使用types.ChatMessageData）
type GetChatGetMessageData = types.ChatMessageData

// GetChatHasUnreadData 获取私聊未读消息数据
type GetChatHasUnreadData = types.GetChatHasUnreadData

// GetChatGetListData 获取私聊列表数据
type GetChatGetListData = types.GetChatListData

// GetApiGetNotificationsData 获取通知列表数据
type GetApiGetNotificationsData = types.NotificationInfo

// GetChatroomNodeGetResponse 获取聊天室节点响应
type GetChatroomNodeGetResponse = types.GetChatroomNodeGetResponse

// ChatroomNodeInfo 聊天室节点信息
type ChatroomNodeInfo = types.ChatroomNodeInfo

// PostChatroomRedPacketOpenRequest 打开聊天室红包请求
type PostChatroomRedPacketOpenRequest = types.PostChatroomRedPacketOpenRequest

// PostChatroomRedPacketOpenResponse 打开聊天室红包响应
type PostChatroomRedPacketOpenResponse = types.PostChatroomRedPacketOpenResponse

// ChatroomMessage 聊天室消息
type ChatroomMessage = types.ChatroomMessage

// ChatroomOnlineData 在线用户数据
type ChatroomOnlineData = types.ChatroomOnlineData

// OnlineUser 在线用户
type OnlineUser = types.OnlineUser

// ChatroomDiscussData 话题变更数据
type ChatroomDiscussData = types.ChatroomDiscussData

// ChatroomRevokeData 撤回消息数据
type ChatroomRevokeData = types.ChatroomRevokeData

// ChatroomMsgData 聊天消息数据
type ChatroomMsgData = types.ChatroomMsgData

// ChatroomRedPacketData 红包消息数据
type ChatroomRedPacketData = types.ChatroomRedPacketData

// RedPacketInfo 红包信息
type RedPacketInfo = types.RedPacketInfo

// ChatroomRedPacketStatusData 红包领取状态数据
type ChatroomRedPacketStatusData = types.ChatroomRedPacketStatusData

// ChatroomCustomData 自定义消息数据
type ChatroomCustomData = types.ChatroomCustomData

// ChatroomBarragerData 弹幕消息数据
type ChatroomBarragerData = types.ChatroomBarragerData

// ChatMessage 私聊消息
type ChatMessage = types.ChatMessage

// ChatMessageData 私聊消息数据
type ChatMessageData = types.ChatMessageData

// UserMessage 用户频道消息
type UserMessage = types.UserMessage

// BreezemoonData 清风明月数据
type BreezemoonData = types.BreezemoonData

// ArticleData 文章数据
type ArticleData = types.ArticleData

// PostArticleRequest 发布文章请求
type PostArticleRequest = types.PostArticleRequest

// PostBreezemoonRequest 发布清风明月请求
type PostBreezemoonRequest = types.PostBreezemoonRequest

// PostCommentRequest 发布评论请求
type PostCommentRequest = types.PostCommentRequest

// TransferRequest 转账请求
type TransferRequest = types.TransferRequest
