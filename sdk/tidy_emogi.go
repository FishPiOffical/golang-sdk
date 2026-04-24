package sdk

import "github.com/fishpioffical/golang-sdk/types"

// GetEmojiGroups 获取表情包分组列表
func (s *FishPiSDK) GetEmojiGroups() (*types.ApiResponse[[]*types.EmojiGroup], error) {
	response := new(types.ApiResponse[[]*types.EmojiGroup])

	_, err := s.client.R().
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/api/emoji/groups")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetEmojiGroupEmojis 获取分组内表情
func (s *FishPiSDK) GetEmojiGroupEmojis(groupId string) (*types.ApiResponse[[]*types.Emoji], error) {
	response := new(types.ApiResponse[[]*types.Emoji])

	_, err := s.client.R().
		SetSuccessResult(response).
		SetErrorResult(response).
		SetQueryParamsAnyType(map[string]any{
			"groupId": groupId,
		}).
		Get("/api/emoji/group/emojis")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostEmojiUpload 上传URL到“全部”分组
func (s *FishPiSDK) PostEmojiUpload(imageUrl string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetSuccessResult(response).
		SetErrorResult(response).
		SetBodyJsonMarshal(map[string]any{
			"url": imageUrl,
			//"apiKey": s.GetConfig().ApiKey,
		}).
		Post("/api/emoji/upload")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostEmojiGroupCreate 创建分组
func (s *FishPiSDK) PostEmojiGroupCreate(name string, sort int) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetSuccessResult(response).
		SetErrorResult(response).
		SetBodyJsonMarshal(map[string]any{
			"name": name,
			"sort": sort,
		}).
		Post("/api/emoji/group/create")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostEmojiGroupUpdate 更新分组
func (s *FishPiSDK) PostEmojiGroupUpdate(groupId string, name string, sort int) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetSuccessResult(response).
		SetErrorResult(response).
		SetBodyJsonMarshal(map[string]any{
			"groupId": groupId,
			"name":    name,
			"sort":    sort,
		}).
		Post("/api/emoji/group/update")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostEmojiGroupDelete 删除分组
func (s *FishPiSDK) PostEmojiGroupDelete(groupId string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetSuccessResult(response).
		SetErrorResult(response).
		SetBodyJsonMarshal(map[string]any{
			"groupId": groupId,
		}).
		Post("/api/emoji/group/delete")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostEmojiGroupAddEmoji 分组添加已有表情
func (s *FishPiSDK) PostEmojiGroupAddEmoji(groupId string, emojiId string, sort int, name string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetSuccessResult(response).
		SetErrorResult(response).
		SetBodyJsonMarshal(map[string]any{
			"groupId": groupId,
			"emojiId": emojiId,
			"sort":    sort,
			"name":    name,
		}).
		Post("/api/emoji/group/add-emoji")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostEmojiGroupAddUrlEmoji 分组添加URL表情
func (s *FishPiSDK) PostEmojiGroupAddUrlEmoji(groupId string, url string, sort int, name string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetSuccessResult(response).
		SetErrorResult(response).
		SetBodyJsonMarshal(map[string]any{
			"groupId": groupId,
			"url":     url,
			"sort":    sort,
			"name":    name,
		}).
		Post("/api/emoji/group/add-url-emoji")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostEmojiGroupRemoveEmoji 从分组移除表情
func (s *FishPiSDK) PostEmojiGroupRemoveEmoji(groupId string, emojiId string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetSuccessResult(response).
		SetErrorResult(response).
		SetBodyJsonMarshal(map[string]any{
			"groupId": groupId,
			"emojiId": emojiId,
		}).
		Post("/api/emoji/group/remove-emoji")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostEmojiEmojiUpdate 更新表情项（重命名/排序）
func (s *FishPiSDK) PostEmojiEmojiUpdate(oId string, groupId string, name string, sort int) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetSuccessResult(response).
		SetErrorResult(response).
		SetBodyJsonMarshal(map[string]any{
			"oId":     oId,
			"groupId": groupId,
			"name":    name,
			"sort":    sort,
		}).
		Post("/api/emoji/emoji/update")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostEmojiEmojiMigrate 迁移旧表情到V2
func (s *FishPiSDK) PostEmojiEmojiMigrate() (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/api/emoji/emoji/migrate")

	if err != nil {
		return nil, err
	}

	return response, nil
}
