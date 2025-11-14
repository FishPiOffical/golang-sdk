package sdk

import (
	"github.com/FishPiOffical/golang-sdk/types"
	"fmt"
)

// GetChatMessages 获取私聊消息
func (s *FishPiSDK) GetChatMessages(toUser string, page, pageSize int) ([]*types.ChatMessageData, error) {
	if toUser == "" {
		return nil, fmt.Errorf("toUser is required")
	}
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	url := fmt.Sprintf("/chat/get-message?toUser=%s&page=%d&pageSize=%d", toUser, page, pageSize)

	var resp types.ApiResponse[[]*types.ChatMessageData]
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get(url)

	if err != nil {
		return nil, fmt.Errorf("failed to get chat messages: %w", err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("get chat messages failed: %s", resp.Msg)
	}

	return resp.Data, nil
}

// SendChatMessage 发送私聊消息
func (s *FishPiSDK) SendChatMessage(toUser, content string) error {
	if toUser == "" {
		return fmt.Errorf("toUser is required")
	}
	if content == "" {
		return fmt.Errorf("content is required")
	}

	var resp types.SimpleResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"toUser":  toUser,
			"content": content,
		}).
		SetSuccessResult(&resp).
		Post("/chat/send")

	if err != nil {
		return fmt.Errorf("failed to send chat message: %w", err)
	}

	if resp.Code != 0 {
		return fmt.Errorf("send chat message failed: %s", resp.Msg)
	}

	return nil
}

// MarkChatRead 标记私聊已读
func (s *FishPiSDK) MarkChatRead(fromUser string) error {
	if fromUser == "" {
		return fmt.Errorf("fromUser is required")
	}

	url := fmt.Sprintf("/chat/mark-as-read?fromUser=%s", fromUser)

	var resp types.SimpleResponse
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get(url)

	if err != nil {
		return fmt.Errorf("failed to mark chat as read: %w", err)
	}

	if resp.Code != 0 {
		return fmt.Errorf("mark chat as read failed: %s", resp.Msg)
	}

	return nil
}

// GetChatList 获取私聊列表
func (s *FishPiSDK) GetChatList() ([]*types.ChatListItem, error) {
	var resp types.ApiResponse[[]*types.ChatListItem]
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get("/chat/list")

	if err != nil {
		return nil, fmt.Errorf("failed to get chat list: %w", err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("get chat list failed: %s", resp.Msg)
	}

	return resp.Data, nil
}

// GetChatUnread 检查是否有未读私聊
func (s *FishPiSDK) GetChatUnread() (bool, error) {
	var resp types.ApiResponse[map[string]bool]
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get("/chat/has-unread")

	if err != nil {
		return false, fmt.Errorf("failed to check unread chat: %w", err)
	}

	if resp.Code != 0 {
		return false, fmt.Errorf("check unread chat failed: %s", resp.Msg)
	}

	if resp.Data != nil {
		if hasUnread, ok := resp.Data["hasUnread"]; ok {
			return hasUnread, nil
		}
	}

	return false, nil
}
