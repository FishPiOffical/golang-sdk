package sdk

import (
	"fishpi-golang-sdk/types"
	"fmt"
)

// GetNotificationCount 获取未读通知数量
func (s *FishPiSDK) GetNotificationCount() (*types.NotificationCount, error) {
	var resp types.NotificationCount
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get("/notifications/unread/count")

	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetNotifications 获取通知列表
func (s *FishPiSDK) GetNotifications(noticeType types.NotificationType) ([]*types.NotificationInfo, error) {
	url := fmt.Sprintf("/api/getNotifications?type=%s", noticeType)

	var resp types.ApiResponse[[]*types.NotificationInfo]
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get(url)

	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Msg)
	}

	return resp.Data, nil
}

// MarkNotificationRead 标记指定类型通知为已读
func (s *FishPiSDK) MarkNotificationRead(noticeType types.NotificationType) error {
	url := fmt.Sprintf("/notifications/make-read/%s", noticeType)

	var resp types.SimpleResponse
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get(url)

	if err != nil {
		return err
	}

	if resp.Code != 0 {
		return fmt.Errorf(resp.Msg)
	}

	return nil
}

// MarkAllNotificationsRead 标记所有通知为已读
func (s *FishPiSDK) MarkAllNotificationsRead() error {
	var resp types.SimpleResponse
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get("/notifications/all-read")

	if err != nil {
		return err
	}

	if resp.Code != 0 {
		return fmt.Errorf(resp.Msg)
	}

	return nil
}
