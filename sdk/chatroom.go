package sdk

import (
	"fishpi-golang-sdk/types"
	"fmt"
)

// SendChatroomMessage 发送聊天室消息
func (s *FishPiSDK) SendChatroomMessage(content string) error {
	var resp types.SimpleResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"content": content,
		}).
		SetSuccessResult(&resp).
		Post("/chat-room/send")

	if err != nil {
		return err
	}

	if resp.Code != 0 {
		return fmt.Errorf(resp.Msg)
	}

	return nil
}

// GetChatroomHistory 获取聊天室历史消息
func (s *FishPiSDK) GetChatroomHistory(page int, contentType types.ChatContentType) ([]*types.ChatroomMsgData, error) {
	url := fmt.Sprintf("/chat-room/more?page=%d&type=%s", page, contentType)

	var resp types.ApiResponse[[]*types.ChatroomMsgData]
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

// GetChatroomMessage 获取聊天室指定消息上下文
func (s *FishPiSDK) GetChatroomMessage(oId string, mode types.ChatMessageType, size int, contentType types.ChatContentType) ([]*types.ChatroomMsgData, error) {
	url := fmt.Sprintf("/chat-room/getMessage?oId=%s&mode=%d&size=%d&type=%s", oId, mode, size, contentType)

	var resp types.ApiResponse[[]*types.ChatroomMsgData]
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

// RevokeChatroomMessage 撤回聊天室消息
func (s *FishPiSDK) RevokeChatroomMessage(oId string) error {
	var resp types.SimpleResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"oId": oId,
		}).
		SetSuccessResult(&resp).
		Delete("/chat-room/revoke/" + oId)

	if err != nil {
		return err
	}

	if resp.Code != 0 {
		return fmt.Errorf(resp.Msg)
	}

	return nil
}

// OpenRedPacket 打开红包
func (s *FishPiSDK) OpenRedPacket(oId string, gesture *types.GestureType) (*types.PostChatroomRedPacketOpenResponse, error) {
	body := map[string]interface{}{
		"oId": oId,
	}
	if gesture != nil {
		body["gesture"] = *gesture
	}

	var resp *types.PostChatroomRedPacketOpenResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(body).
		SetSuccessResult(&resp).
		Post("/chat-room/red-packet/open")

	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Msg)
	}

	return resp, nil
}
