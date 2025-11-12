package sdk

import (
	"fishpi-golang-sdk/types"
	"fmt"
)

// GetArticleComments 获取文章评论列表
func (s *FishPiSDK) GetArticleComments(articleId string, page, size int) ([]*types.CommentInfo, error) {
	if articleId == "" {
		return nil, fmt.Errorf("articleId is required")
	}
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20
	}

	url := fmt.Sprintf("/article/%s/comments?p=%d&size=%d", articleId, page, size)

	var resp types.ApiResponse[[]*types.CommentInfo]
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get(url)

	if err != nil {
		return nil, fmt.Errorf("failed to get article comments: %w", err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("get article comments failed: %s", resp.Msg)
	}

	return resp.Data, nil
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

// GetCloudEmojis 获取云端表情包
func (s *FishPiSDK) GetCloudEmojis() ([]string, error) {
	var resp types.ApiResponse[[]string]
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get("/cloud/get")

	if err != nil {
		return nil, fmt.Errorf("failed to get cloud emojis: %w", err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("get cloud emojis failed: %s", resp.Msg)
	}

	return resp.Data, nil
}

// SyncCloudEmojis 同步云端表情包
func (s *FishPiSDK) SyncCloudEmojis(emojis []string) error {
	var resp types.SimpleResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]interface{}{
			"emojis": emojis,
		}).
		SetSuccessResult(&resp).
		Post("/cloud/sync")

	if err != nil {
		return fmt.Errorf("failed to sync cloud emojis: %w", err)
	}

	if resp.Code != 0 {
		return fmt.Errorf("sync cloud emojis failed: %s", resp.Msg)
	}

	return nil
}
