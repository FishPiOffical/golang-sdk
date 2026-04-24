package sdk

import "github.com/fishpioffical/golang-sdk/types"

// PostArticleReaction 给帖子添加表情
func (s *FishPiSDK) PostArticleReaction(articleId string, groupType types.ReactionGroupType, value types.ReactionEmojiValue) (*types.ApiResponse[*types.PostReactionData], error) {
	response := new(types.ApiResponse[*types.PostReactionData])

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"articleId": articleId,
			"groupType": groupType,
			"value":     value,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/article/reaction")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostCommentReaction 给帖子添加表情
func (s *FishPiSDK) PostCommentReaction(commentId string, groupType types.ReactionGroupType, value types.ReactionEmojiValue) (*types.ApiResponse[*types.PostReactionData], error) {
	response := new(types.ApiResponse[*types.PostReactionData])

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"commentId": commentId,
			"groupType": groupType,
			"value":     value,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/comment/reaction")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostChatRoomReaction 给聊天室消息添加表情
func (s *FishPiSDK) PostChatRoomReaction(oId string, groupType types.ReactionGroupType, value types.ReactionEmojiValue) (*types.ApiResponse[*types.PostReactionData], error) {
	response := new(types.ApiResponse[*types.PostReactionData])

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"oId":       oId,
			"groupType": groupType,
			"value":     value,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/chat-room/reaction")

	if err != nil {
		return nil, err
	}

	return response, nil
}
