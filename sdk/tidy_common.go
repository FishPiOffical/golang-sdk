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

// PostFollowUser 关注用户
func (s *FishPiSDK) PostFollowUser(followingId string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"apiKey":      s.GetAPIKey(),
			"followingId": followingId,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/follow/user")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostUnfollowUser 取消关注用户
func (s *FishPiSDK) PostUnfollowUser(followingId string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"apiKey":      s.GetAPIKey(),
			"followingId": followingId,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/unfollow/user")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetMembership 获取用户VIP信息
func (s *FishPiSDK) GetMembership(userId string) (*types.ApiResponse[*types.MemberShip], error) {
	response := new(types.ApiResponse[*types.MemberShip])

	_, err := s.client.R().
		SetPathParam("userId", userId).
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/api/membership/{userId}")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetLogsMore 获取操作日志
func (s *FishPiSDK) GetLogsMore(page, pageSize int) (*types.ApiResponse[[]*types.LogInfo], error) {
	response := new(types.ApiResponse[[]*types.LogInfo])

	_, err := s.client.R().
		SetQueryParamsAnyType(map[string]any{
			"page":     page,
			"pageSize": pageSize,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/logs/more")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostSettingsAvatar 更新头像
func (s *FishPiSDK) PostSettingsAvatar(avatar string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"userAvatarURL": avatar,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/api/settings/avatar")
	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostSettingsProfiles 更新个人资料
func (s *FishPiSDK) PostSettingsProfiles(profiles *types.PostSettingsProfilesRequest) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(profiles).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/api/settings/profiles")
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetUserUsernamePoint 查询指定用户积分
func (s *FishPiSDK) GetUserUsernamePoint(username string) (*types.ApiResponse[*types.GetUserUsernamePointData], error) {
	response := new(types.ApiResponse[*types.GetUserUsernamePointData])

	_, err := s.client.R().
		SetPathParam("username", username).
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/user/{username}/point")
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetUserUsernameMedal 查询指定用户勋章
func (s *FishPiSDK) GetUserUsernameMedal(username string) (*types.ApiResponse[*types.GetUserUsernameMedalData], error) {
	response := new(types.ApiResponse[*types.GetUserUsernameMedalData])

	_, err := s.client.R().
		SetPathParam("username", username).
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/user/{username}/medal")
	if err != nil {
		return nil, err
	}

	return response, nil
}
