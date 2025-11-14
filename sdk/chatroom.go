package sdk

import (
	"github.com/FishPiOffical/golang-sdk/types"
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

// GetChatroomNode 获取聊天室节点信息
func (s *FishPiSDK) GetChatroomNode() (*types.GetChatroomNodeGetResponse, error) {
	var resp types.GetChatroomNodeGetResponse
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get("/chat-room/node/get")

	if err != nil {
		return nil, fmt.Errorf("failed to get chatroom node: %w", err)
	}

	return &resp, nil
}

// GetChatroomMutes 获取聊天室禁言列表
func (s *FishPiSDK) GetChatroomMutes() ([]*types.MuteItem, error) {
	var resp types.ApiResponse[[]*types.MuteItem]
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get("/chat-room/mutes")

	if err != nil {
		return nil, fmt.Errorf("failed to get mutes list: %w", err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("get mutes list failed: %s", resp.Msg)
	}

	return resp.Data, nil
}

// SendChatroomBarrage 发送弹幕
func (s *FishPiSDK) SendChatroomBarrage(content string) error {
	if content == "" {
		return fmt.Errorf("content is required")
	}

	var resp types.SimpleResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"content": content,
		}).
		SetSuccessResult(&resp).
		Post("/chat-room/barrager")

	if err != nil {
		return fmt.Errorf("failed to send barrage: %w", err)
	}

	if resp.Code != 0 {
		return fmt.Errorf("send barrage failed: %s", resp.Msg)
	}

	return nil
}

// SendRedPacket 发送红包
func (s *FishPiSDK) SendRedPacket(req *types.SendRedPacketRequest) error {
	if req == nil {
		return fmt.Errorf("request is required")
	}

	var resp types.SimpleResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(req).
		SetSuccessResult(&resp).
		Post("/chat-room/red-packet/send")

	if err != nil {
		return fmt.Errorf("failed to send red packet: %w", err)
	}

	if resp.Code != 0 {
		return fmt.Errorf("send red packet failed: %s", resp.Msg)
	}

	return nil
}

// GetRedPacketDetail 获取红包详情
func (s *FishPiSDK) GetRedPacketDetail(oId string) (*types.RedPacketDetail, error) {
	if oId == "" {
		return nil, fmt.Errorf("oId is required")
	}

	var resp types.ApiResponse[*types.RedPacketDetail]
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get("/chat-room/red-packet/" + oId)

	if err != nil {
		return nil, fmt.Errorf("failed to get red packet detail: %w", err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("get red packet detail failed: %s", resp.Msg)
	}

	return resp.Data, nil
}

// GetBarrageCost 获取弹幕费用
func (s *FishPiSDK) GetBarrageCost() (int, error) {
	var resp types.ApiResponse[map[string]int]
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get("/chat-room/barrage/cost")

	if err != nil {
		return 0, fmt.Errorf("failed to get barrage cost: %w", err)
	}

	if resp.Code != 0 {
		return 0, fmt.Errorf("get barrage cost failed: %s", resp.Msg)
	}

	if resp.Data != nil {
		if cost, ok := resp.Data["cost"]; ok {
			return cost, nil
		}
	}

	return 0, nil
}

// GetMessageRaw 获取原始消息（HTML格式）
func (s *FishPiSDK) GetMessageRaw(oId string) (string, error) {
	if oId == "" {
		return "", fmt.Errorf("oId is required")
	}

	var resp types.ApiResponse[map[string]string]
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get("/cr/raw/" + oId)

	if err != nil {
		return "", fmt.Errorf("failed to get message raw: %w", err)
	}

	if resp.Code != 0 {
		return "", fmt.Errorf("get message raw failed: %s", resp.Msg)
	}

	if resp.Data != nil {
		if raw, ok := resp.Data["html"]; ok {
			return raw, nil
		}
	}

	return "", nil
}
