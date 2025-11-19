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

// GetChatroomNode 获取聊天室节点信息
func (s *FishPiSDK) GetChatroomNode() (*types.GetChatroomNodeGetResponse, error) {
	response := new(types.GetChatroomNodeGetResponse)
	_, err := s.client.R().
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/chat-room/node/get")

	if err != nil {
		return nil, err
	}

	return response, nil
}
