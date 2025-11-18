package sdk

import (
	"github.com/FishPiOffical/golang-sdk/types"
)

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

// GetUserCheckedIn 检查是否已签到
func (s *FishPiSDK) GetUserCheckedIn() (*types.UserCheckedInResponse, error) {
	response := new(types.UserCheckedInResponse)

	_, err := s.client.R().
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/user/checkedIn")
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetYesterdayLivenessReward 领取昨日活跃度奖励
func (s *FishPiSDK) GetYesterdayLivenessReward() (*types.YesterdayLivenessRewardResponse, error) {
	response := new(types.YesterdayLivenessRewardResponse)

	_, err := s.client.R().
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/activity/yesterday-liveness-reward-api")
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetIsCollectedLiveness 检查是否已领取昨日活跃奖励
func (s *FishPiSDK) GetIsCollectedLiveness() (*types.IsCollectedLivenessResponse, error) {
	response := new(types.IsCollectedLivenessResponse)

	_, err := s.client.R().
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/api/activity/is-collected-liveness")
	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostReport 举报用户
func (s *FishPiSDK) PostReport(dataId string, dataType types.ReportDataType, reportType types.ReportType, memo string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetFormDataAnyType(map[string]any{
			"apiKey":         s.GetAPIKey(),
			"reportDataId":   dataId,
			"reportDataType": dataType,
			"reportType":     reportType,
			"reportMemo":     memo,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/report")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetUserRecentReg 举报用户
func (s *FishPiSDK) GetUserRecentReg() (*types.ApiResponse[[]*types.RecentRegUser], error) {
	response := new(types.ApiResponse[[]*types.RecentRegUser])

	_, err := s.client.R().
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/api/user/recentReg")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostPointTransfer 转账
func (s *FishPiSDK) PostPointTransfer(req *types.PostPointTransferRequest) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(req).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/point/transfer")
	if err != nil {
		return nil, err
	}

	return response, nil
}
