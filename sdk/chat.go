package sdk

import (
	"fishpi-golang-sdk/types"
	"fmt"
)

// GetChatList 获取私聊列表
func (s *FishPiSDK) GetChatList() ([]*types.GetChatListData, error) {
	var resp types.ApiResponse[[]*types.GetChatListData]
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get("/chat/get-list")

	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Msg)
	}

	return resp.Data, nil
}

// GetChatMessages 获取与指定用户的私聊消息
func (s *FishPiSDK) GetChatMessages(toUser string, page, pageSize int) ([]*types.ChatMessageData, error) {
	url := fmt.Sprintf("/chat/get-message?toUser=%s&page=%d&pageSize=%d", toUser, page, pageSize)

	var resp types.ApiResponse[[]*types.ChatMessageData]
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get(url)

	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Msg)
	}

	return resp.Data, nil
}

// SendChatMessage 发送私聊消息
func (s *FishPiSDK) SendChatMessage(toUser, content string) error {
	var resp types.SimpleResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"toUser":  toUser,
			"content": content,
		}).
		SetSuccessResult(&resp).
		Post("/chat/send")

	if err != nil {
		return err
	}

	if resp.Code != 0 {
		return fmt.Errorf(resp.Msg)
	}

	return nil
}

// MarkChatRead 标记私聊消息已读
func (s *FishPiSDK) MarkChatRead(fromUser string) error {
	var resp types.SimpleResponse
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get("/chat/mark-as-read?fromUser=" + fromUser)

	if err != nil {
		return err
	}

	if resp.Code != 0 {
		return fmt.Errorf(resp.Msg)
	}

	return nil
}

// GetChatUnread 获取未读私聊消息
func (s *FishPiSDK) GetChatUnread() (*types.GetChatHasUnreadData, error) {
	var resp types.ApiResponse[*types.GetChatHasUnreadData]
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get("/chat/has-unread")

	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Msg)
	}

	return resp.Data, nil
}
