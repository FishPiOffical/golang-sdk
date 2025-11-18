package sdk

import (
	"fmt"

	"github.com/FishPiOffical/golang-sdk/types"
)

// GetUserInfo 获取当前用户信息
func (s *FishPiSDK) GetUserInfo() (*types.UserInfoResponse, error) {
	res, err := s.client.R().Get("/api/user")
	if err != nil {
		return nil, err
	}

	var response types.UserInfoResponse
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// PostPointTransfer 转账
func (s *FishPiSDK) PostPointTransfer(req *types.TransferRequest) (*types.SimpleResponse, error) {
	res, err := s.client.R().
		SetBodyJsonMarshal(req).
		Post("/point/transfer")
	if err != nil {
		return nil, err
	}

	var response types.SimpleResponse
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// PostUserCheckin 用户签到
func (s *FishPiSDK) PostUserCheckin() (*types.CheckinResponse, error) {
	var resp types.ApiResponse[*types.CheckinResponse]
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get("/activity/daily-checkin")

	if err != nil {
		return nil, fmt.Errorf("failed to checkin: %w", err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("checkin failed: %s", resp.Msg)
	}

	return resp.Data, nil
}

// RewardLiveness 领取昨日活跃度奖励
func (s *FishPiSDK) RewardLiveness() (int, error) {
	resp, err := s.GetYesterdayLivenessReward()
	if err != nil {
		return 0, err
	}

	return resp.Sum, nil
}

// IsCheckIn 检查是否已签到
func (s *FishPiSDK) IsCheckIn() (bool, error) {
	resp, err := s.GetUserCheckedIn()
	if err != nil {
		return false, err
	}

	return resp.CheckedIn, nil
}

// IsCollectedLiveness 检查是否已领取昨日活跃度奖励
func (s *FishPiSDK) IsCollectedLiveness() (bool, error) {
	resp, err := s.GetIsCollectedLiveness()
	if err != nil {
		return false, err
	}

	return resp.IsCollectedYesterdayLivenessReward, nil
}

// GetLiveness 获取当前活跃度值
func (s *FishPiSDK) GetLiveness() (float64, error) {
	resp, err := s.GetUserLiveness()
	if err != nil {
		return 0, err
	}

	return resp.Liveness, nil
}

// GetUserByUsername 通过用户名获取用户信息
func (s *FishPiSDK) GetUserByUsername(username string) (*types.UserInfo, error) {
	if username == "" {
		return nil, fmt.Errorf("username is required")
	}

	var resp types.ApiResponse[*types.UserInfo]
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get("/user/" + username)

	if err != nil {
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("get user by username failed: %s", resp.Msg)
	}

	return resp.Data, nil
}

// FollowUser 关注/取消关注用户
func (s *FishPiSDK) FollowUser(followingId string) error {
	if followingId == "" {
		return fmt.Errorf("followingId is required")
	}

	var resp types.SimpleResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"followingId": followingId,
		}).
		SetSuccessResult(&resp).
		Post("/follow")

	if err != nil {
		return fmt.Errorf("failed to follow user: %w", err)
	}

	if resp.Code != 0 {
		return fmt.Errorf("follow user failed: %s", resp.Msg)
	}

	return nil
}

// UploadFile 上传文件
func (s *FishPiSDK) UploadFile(file []byte, fileName string) (*types.UploadFileResponse, error) {
	if len(file) == 0 {
		return nil, fmt.Errorf("file is required")
	}
	if fileName == "" {
		return nil, fmt.Errorf("fileName is required")
	}

	var resp types.ApiResponse[*types.UploadFileResponse]
	_, err := s.client.R().
		SetFileBytes("file[]", fileName, file).
		SetSuccessResult(&resp).
		Post("/upload")

	if err != nil {
		return nil, fmt.Errorf("failed to upload file: %w", err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("upload file failed: %s", resp.Msg)
	}

	return resp.Data, nil
}

// GetUserNames 获取用户名列表（管理员功能）
func (s *FishPiSDK) GetUserNames() ([]string, error) {
	var resp types.ApiResponse[[]string]
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get("/user/names")

	if err != nil {
		return nil, fmt.Errorf("failed to get user names: %w", err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("get user names failed: %s", resp.Msg)
	}

	return resp.Data, nil
}

// GetRecentRegister 获取最近注册用户（管理员功能）
func (s *FishPiSDK) GetRecentRegister() ([]*types.UserInfo, error) {
	var resp types.ApiResponse[[]*types.UserInfo]
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get("/user/recentRegister")

	if err != nil {
		return nil, fmt.Errorf("failed to get recent register: %w", err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("get recent register failed: %s", resp.Msg)
	}

	return resp.Data, nil
}

// GetLogs 获取日志（管理员功能）
func (s *FishPiSDK) GetLogs(page int) ([]map[string]interface{}, error) {
	if page < 1 {
		page = 1
	}

	var resp types.ApiResponse[[]map[string]interface{}]
	_, err := s.client.R().
		SetQueryParam("p", fmt.Sprintf("%d", page)).
		SetSuccessResult(&resp).
		Get("/logs")

	if err != nil {
		return nil, fmt.Errorf("failed to get logs: %w", err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("get logs failed: %s", resp.Msg)
	}

	return resp.Data, nil
}
