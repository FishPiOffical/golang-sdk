package sdk

import (
	"github.com/FishPiOffical/golang-sdk/types"
)

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

// GetNotifications 获取通知列表
func (s *FishPiSDK) GetNotifications(noticeType types.NotificationType, page int) (*types.ApiResponse[[]*types.NotificationInfo], error) {
	response := new(types.ApiResponse[[]*types.NotificationInfo])
	if page <= 0 {
		page = 1
	}

	_, err := s.client.R().
		SetQueryParamsAnyType(map[string]any{
			"type": noticeType,
			"p":    page,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/api/getNotifications")

	if err != nil {
		return nil, err
	}

	return response, nil
}
