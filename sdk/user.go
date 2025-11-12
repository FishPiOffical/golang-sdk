package sdk

import "fishpi-golang-sdk/types"

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

// GetUserLiveness 获取当前用户活跃度
func (s *FishPiSDK) GetUserLiveness() (*types.UserLivenessResponse, error) {
	res, err := s.client.R().Get("/user/liveness")
	if err != nil {
		return nil, err
	}

	var response types.UserLivenessResponse
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetUserCheckedIn 检查是否已签到
func (s *FishPiSDK) GetUserCheckedIn() (*types.UserCheckedInResponse, error) {
	res, err := s.client.R().Get("/user/checkedIn")
	if err != nil {
		return nil, err
	}

	var response types.UserCheckedInResponse
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetIsCollectedLiveness 检查是否已领取昨日活跃奖励
func (s *FishPiSDK) GetIsCollectedLiveness() (*types.IsCollectedLivenessResponse, error) {
	res, err := s.client.R().Get("/api/activity/is-collected-liveness")
	if err != nil {
		return nil, err
	}

	var response types.IsCollectedLivenessResponse
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetYesterdayLivenessReward 领取昨日活跃度奖励
func (s *FishPiSDK) GetYesterdayLivenessReward() (*types.YesterdayLivenessRewardResponse, error) {
	res, err := s.client.R().Get("/activity/yesterday-liveness-reward-api")
	if err != nil {
		return nil, err
	}

	var response types.YesterdayLivenessRewardResponse
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

// GetUserEmotions 获取用户常用表情
func (s *FishPiSDK) GetUserEmotions() (*types.UserEmotionsResponse, error) {
	res, err := s.client.R().Get("/users/emotions")
	if err != nil {
		return nil, err
	}

	var response types.UserEmotionsResponse
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
