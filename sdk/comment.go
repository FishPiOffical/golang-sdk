package sdk

import (
	"github.com/FishPiOffical/golang-sdk/types"
	"fmt"
)

// UpdateComment 更新评论
func (s *FishPiSDK) UpdateComment(commentId string, req *types.UpdateCommentRequest) error {
	if commentId == "" {
		return fmt.Errorf("commentId is required")
	}
	if req.CommentContent == "" {
		return fmt.Errorf("comment content is required")
	}

	var resp types.SimpleResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(req).
		SetSuccessResult(&resp).
		Put("/comment/" + commentId)

	if err != nil {
		return fmt.Errorf("failed to update comment: %w", err)
	}

	if resp.Code != 0 {
		return fmt.Errorf("update comment failed: %s", resp.Msg)
	}

	return nil
}

// VoteComment 评论投票
func (s *FishPiSDK) VoteComment(commentId string, voteType string) (types.VoteType, error) {
	if commentId == "" {
		return types.VoteTypeUnvoted, fmt.Errorf("commentId is required")
	}
	if voteType != "up" && voteType != "down" {
		return types.VoteTypeUnvoted, fmt.Errorf("voteType must be 'up' or 'down'")
	}

	url := fmt.Sprintf("/vote/%s/comment", voteType)

	var resp types.ApiResponse[map[string]types.VoteType]
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"dataId": commentId,
		}).
		SetSuccessResult(&resp).
		Post(url)

	if err != nil {
		return types.VoteTypeUnvoted, fmt.Errorf("failed to vote comment: %w", err)
	}

	if resp.Code != 0 {
		return types.VoteTypeUnvoted, fmt.Errorf("vote comment failed: %s", resp.Msg)
	}

	if resp.Data == nil {
		return types.VoteTypeUnvoted, fmt.Errorf("vote response data is nil")
	}

	voteResult, ok := resp.Data["type"]
	if !ok {
		return types.VoteTypeUnvoted, fmt.Errorf("vote type not found in response")
	}

	return voteResult, nil
}

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
