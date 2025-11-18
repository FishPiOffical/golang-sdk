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
