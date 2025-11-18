package sdk

import "github.com/FishPiOffical/golang-sdk/types"

// GetNotificationCount 获取未读通知数量
func (s *FishPiSDK) GetNotificationCount() (*types.GetNotificationCountResponse, error) {
	response := new(types.GetNotificationCountResponse)
	_, err := s.client.R().
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/notifications/unread/count")

	if err != nil {
		return nil, err
	}

	return response, nil
}
