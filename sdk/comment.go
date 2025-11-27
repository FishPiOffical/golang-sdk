package sdk

import (
	"github.com/FishPiOffical/golang-sdk/types"
	"fmt"
)

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
