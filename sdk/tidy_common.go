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
