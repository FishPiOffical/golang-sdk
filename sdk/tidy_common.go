package sdk

import "github.com/FishPiOffical/golang-sdk/types"

// GetUserInfoByUsername 根据用户名获取用户信息
func (s *FishPiSDK) GetUserInfoByUsername(username string) (*types.User, error) {
	response := new(types.User)

	_, err := s.client.R().
		SetPathParam("username", username).
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/user/{username}")
	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostUsersNames 用户名联想
func (s *FishPiSDK) PostUsersNames(name string) (*types.ApiResponse[[]*types.AssociateUser], error) {
	response := new(types.ApiResponse[[]*types.AssociateUser])

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"name": name,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/users/names")
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetUsersEmotions 获取用户常用表情
func (s *FishPiSDK) GetUsersEmotions() (*types.ApiResponse[[]map[string]any], error) {
	response := new(types.ApiResponse[[]map[string]any])

	_, err := s.client.R().
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/users/emotions")
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetUserLiveness 获取当前用户活跃度响应
func (s *FishPiSDK) GetUserLiveness() (*types.UserLivenessResponse, error) {
	response := new(types.UserLivenessResponse)

	_, err := s.client.R().
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/user/liveness")
	if err != nil {
		return nil, err
	}

	return response, nil
}
