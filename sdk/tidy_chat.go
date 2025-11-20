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

// GetChatroomMore 获取聊天室历史消息
func (s *FishPiSDK) GetChatroomMore(page int) (*types.ApiResponse[[]*types.ChatroomMsgData], error) {
	response := new(types.ApiResponse[[]*types.ChatroomMsgData])

	_, err := s.client.R().
		SetQueryParamsAnyType(map[string]any{
			"page": page,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/chat-room/more")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetChatroomMessage 获取聊天室指定消息上下文
func (s *FishPiSDK) GetChatroomMessage(oId string, size int, mode types.ChatMessageType) (*types.ApiResponse[[]*types.ChatroomMsgData], error) {
	response := new(types.ApiResponse[[]*types.ChatroomMsgData])

	_, err := s.client.R().
		SetQueryParamsAnyType(map[string]any{
			"oId":  oId,
			"size": size,
			"mode": int(mode),
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/chat-room/getMessage")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostChatroomSend 发送聊天室消息
func (s *FishPiSDK) PostChatroomSend(content string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"content": content,
			"client":  VERSION,
			"apiKey":  s.GetAPIKey(),
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/chat-room/send")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// DeleteChatroomRevoke 撤回聊天室消息
func (s *FishPiSDK) DeleteChatroomRevoke(oId string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"apiKey": s.GetAPIKey(),
		}).
		SetPathParam("oId", oId).
		SetSuccessResult(response).
		SetErrorResult(response).
		Delete("/chat-room/revoke/{oId}")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetMessageRaw 获取原始消息（HTML格式）
func (s *FishPiSDK) GetMessageRaw(oId string) (*string, error) {

	res, err := s.client.R().
		SetPathParam("oId", oId).
		Get("/cr/raw/{oId}")

	if err != nil {
		return nil, err
	}

	result := res.String()

	return &result, nil
}

// PostRedPacketOpen 打开红包
func (s *FishPiSDK) PostRedPacketOpen(oId string, gesture types.GestureType) (*types.PostChatroomRedPacketOpenResponse, error) {
	response := new(types.PostChatroomRedPacketOpenResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"oId":     oId,
			"gesture": int(gesture),
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/chat-room/red-packet/open")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostCloudGet 从云获取指定Key内容
func (s *FishPiSDK) PostCloudGet(gameId types.CloudGameId) (*types.ApiResponse[string], error) {
	response := new(types.ApiResponse[string])

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"gameId": gameId,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/api/cloud/get")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostCloudSync 同步指定key数据到云
func (s *FishPiSDK) PostCloudSync(gameId types.CloudGameId, data string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"gameId": gameId,
			"data":   data,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/api/cloud/sync")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetSiGuoYa 获取思过崖用户
func (s *FishPiSDK) GetSiGuoYa() (*types.ApiResponse[[]*types.MuteUser], error) {
	response := new(types.ApiResponse[[]*types.MuteUser])

	_, err := s.client.R().
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/chat-room/si-guo-list")

	if err != nil {
		return nil, err
	}

	return response, nil
}
