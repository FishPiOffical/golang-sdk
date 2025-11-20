package sdk

import (
	"fmt"

	"github.com/FishPiOffical/golang-sdk/types"
)

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
