package sdk

import (
	"github.com/FishPiOffical/golang-sdk/types"
	"fmt"
)

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
