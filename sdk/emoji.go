package sdk

import (
	"github.com/FishPiOffical/golang-sdk/types"
	"fmt"
)

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
