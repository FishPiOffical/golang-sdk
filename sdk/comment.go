package sdk

import (
	"github.com/FishPiOffical/golang-sdk/types"
	"fmt"
)

// ThankComment 感谢评论
func (s *FishPiSDK) ThankComment(commentId string) error {
	if commentId == "" {
		return fmt.Errorf("commentId is required")
	}

	var resp types.SimpleResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"commentId": commentId,
		}).
		SetSuccessResult(&resp).
		Post("/comment/thank")

	if err != nil {
		return fmt.Errorf("failed to thank comment: %w", err)
	}

	if resp.Code != 0 {
		return fmt.Errorf("thank comment failed: %s", resp.Msg)
	}

	return nil
}

// RemoveComment 删除评论
func (s *FishPiSDK) RemoveComment(commentId string) error {
	if commentId == "" {
		return fmt.Errorf("commentId is required")
	}

	var resp types.SimpleResponse
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Post("/comment/" + commentId + "/remove")

	if err != nil {
		return fmt.Errorf("failed to remove comment: %w", err)
	}

	if resp.Code != 0 {
		return fmt.Errorf("remove comment failed: %s", resp.Msg)
	}

	return nil
}
