package sdk

import (
	"github.com/FishPiOffical/golang-sdk/types"
	"fmt"
)

// GetBreezemoonList 获取清风明月列表
func (s *FishPiSDK) GetBreezemoonList(page, size int) ([]*types.BreezemoonInfo, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20
	}

	url := fmt.Sprintf("/api/breezemoons?p=%d&size=%d", page, size)

	var resp types.ApiResponse[map[string][]*types.BreezemoonInfo]
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get(url)

	if err != nil {
		return nil, fmt.Errorf("failed to get breezemoon list: %w", err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("get breezemoon list failed: %s", resp.Msg)
	}

	if resp.Data != nil {
		if list, ok := resp.Data["breezemoons"]; ok {
			return list, nil
		}
	}

	return nil, nil
}

// GetUserBreezemoons 获取用户清风明月列表
func (s *FishPiSDK) GetUserBreezemoons(userName string, page, size int) (*types.BreezemoonList, error) {
	if userName == "" {
		return nil, fmt.Errorf("userName is required")
	}
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20
	}

	url := fmt.Sprintf("/api/user/%s/breezemoons?p=%d&size=%d", userName, page, size)

	var resp types.ApiResponse[*types.BreezemoonList]
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get(url)

	if err != nil {
		return nil, fmt.Errorf("failed to get user breezemoons: %w", err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("get user breezemoons failed: %s", resp.Msg)
	}

	return resp.Data, nil
}

// PostBreezemoon 发送清风明月
func (s *FishPiSDK) PostBreezemoon(content string) error {
	if content == "" {
		return fmt.Errorf("content is required")
	}

	var resp types.SimpleResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"breezemoonContent": content,
		}).
		SetSuccessResult(&resp).
		Post("/breezemoon")

	if err != nil {
		return fmt.Errorf("failed to send breezemoon: %w", err)
	}

	if resp.Code != 0 {
		return fmt.Errorf("send breezemoon failed: %s", resp.Msg)
	}

	return nil
}

// UpdateBreezemoon 更新清风明月
func (s *FishPiSDK) UpdateBreezemoon(id, content string) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	if content == "" {
		return fmt.Errorf("content is required")
	}

	var resp types.SimpleResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"breezemoonContent": content,
		}).
		SetSuccessResult(&resp).
		Put("/breezemoon/" + id)

	if err != nil {
		return fmt.Errorf("failed to update breezemoon: %w", err)
	}

	if resp.Code != 0 {
		return fmt.Errorf("update breezemoon failed: %s", resp.Msg)
	}

	return nil
}

// RemoveBreezemoon 删除清风明月
func (s *FishPiSDK) RemoveBreezemoon(id string) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}

	var resp types.SimpleResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"id": id,
		}).
		SetSuccessResult(&resp).
		Delete("/breezemoon/" + id)

	if err != nil {
		return fmt.Errorf("failed to remove breezemoon: %w", err)
	}

	if resp.Code != 0 {
		return fmt.Errorf("remove breezemoon failed: %s", resp.Msg)
	}

	return nil
}
