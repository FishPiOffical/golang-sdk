package sdk

import "github.com/FishPiOffical/golang-sdk/types"

// GetChatMessage 获取用户私聊历史消息
func (s *FishPiSDK) GetChatMessage(toUser string, page, pageSize int) (*types.ResultResponse[[]*types.ChatMessage], error) {
	response := new(types.ResultResponse[[]*types.ChatMessage])

	_, err := s.client.R().
		SetQueryParamsAnyType(map[string]any{
			"toUser":   toUser,
			"page":     page,
			"pageSize": pageSize,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/chat/get-message")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetChatMarkAsRead 标记用户消息已读
func (s *FishPiSDK) GetChatMarkAsRead(fromUser string) (*types.ResultResponse[any], error) {
	response := new(types.ResultResponse[any])

	_, err := s.client.R().
		SetQueryParamsAnyType(map[string]any{
			"fromUser": fromUser,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/chat/mark-as-read")

	if err != nil {
		return nil, err
	}

	return response, nil
}
