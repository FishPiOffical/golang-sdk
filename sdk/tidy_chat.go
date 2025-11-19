package sdk

import (
	"github.com/FishPiOffical/golang-sdk/types"
)

// GetChatroomBarragePrice 获取弹幕价格
func (s *FishPiSDK) GetChatroomBarragePrice() (*types.ApiResponse[string], error) {
	response := new(types.ApiResponse[string])
	_, err := s.client.R().
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/chat-room/barrager/get")

	if err != nil {
		return nil, err
	}

	return response, nil
}
