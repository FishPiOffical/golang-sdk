package sdk

import (
	"github.com/FishPiOffical/golang-sdk/types"
)

// PostMoFishScore 上传摸鱼大闯关关卡数据 未验证
func (s *FishPiSDK) PostMoFishScore(userName string, stage, time int64) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]interface{}{
			"goldFingerKey": s.GetConfig().GameGoldFingerKey,
			"userName":      userName,
			"stage":         stage,
			"time":          time,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/api/games/mofish/score")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostQueryLatestLoginIp 查询用户最近登录的IP地址
func (s *FishPiSDK) PostQueryLatestLoginIp(userName string) (*types.ApiResponse[*types.LatestLoginIp], error) {
	response := new(types.ApiResponse[*types.LatestLoginIp])
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"goldFingerKey": s.GetConfig().QueryGoldFingerKey,
			"userName":      userName,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/user/query/latest-login-ip")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostGiveMetal 添加勋章
// Deprecated: 改用 PostMedalAdminGrant
func (s *FishPiSDK) PostGiveMetal(userName string, metal *types.Metal) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]interface{}{
			"goldFingerKey": s.GetConfig().MetalGoldFingerKey,
			"userName":      userName,
			"name":          metal.Name,
			"description":   metal.Description,
			"attr":          metal.Attr,
			"data":          metal.Data,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/user/edit/give-metal")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostRemoveMetal 移除勋章
// Deprecated: 改用 PostMedalAdminRevoke
func (s *FishPiSDK) PostRemoveMetal(userName, metalName string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"goldFingerKey": s.GetConfig().MetalGoldFingerKey,
			"userName":      userName,
			"name":          metalName,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/user/edit/remove-metal")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostRemoveMetalByUserId 通过用户ID移除勋章 未验证
// Deprecated: 改用 PostMedalAdminRevoke
func (s *FishPiSDK) PostRemoveMetalByUserId(userId, metalName string) (*types.SimpleResponse, error) {
	resp := new(types.SimpleResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"goldFingerKey": s.GetConfig().MetalGoldFingerKey,
			"userId":        userId,
			"name":          metalName,
		}).
		SetSuccessResult(resp).
		SetErrorResult(resp).
		Post("/user/edit/remove-metal-by-user-id")

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// PostUserQueryItems 查询用户背包
func (s *FishPiSDK) PostUserQueryItems(userName string) (*types.ApiResponse[map[types.ItemType]int64], error) {
	response := new(types.ApiResponse[map[types.ItemType]int64])

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"goldFingerKey": s.GetConfig().ItemGoldFingerKey,
			"userName":      userName,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/user/query/items")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostUserEditItems 调整用户背包
func (s *FishPiSDK) PostUserEditItems(userName string, item types.ItemType, sum int64) (*types.ApiResponse[map[types.ItemType]int64], error) {
	response := new(types.ApiResponse[map[types.ItemType]int64])

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]interface{}{
			"goldFingerKey": s.GetConfig().ItemGoldFingerKey,
			"userName":      userName,
			"item":          item,
			"sum":           sum,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/user/edit/items")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostUserEditPoints 调整用户积分
func (s *FishPiSDK) PostUserEditPoints(userName string, point int, memo string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]interface{}{
			"goldFingerKey": s.GetConfig().PointGoldFingerKey,
			"userName":      userName,
			"point":         point,
			"memo":          memo,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/user/edit/points")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostUserLiveness 获取用户活跃度
func (s *FishPiSDK) PostUserLiveness(userName string) (*types.PostUserLivenessResponse, error) {
	response := new(types.PostUserLivenessResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"goldFingerKey": s.GetConfig().LivenessGoldFingerKey,
			"userName":      userName,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/user/liveness")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostYesterdayLivenessReward 领取指定用户的昨日活跃奖励
func (s *FishPiSDK) PostYesterdayLivenessReward(userName string) (*types.PostYesterdayLivenessRewardApiResponse, error) {
	response := new(types.PostYesterdayLivenessRewardApiResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"goldFingerKey": s.GetConfig().LivenessGoldFingerKey,
			"userName":      userName,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/activity/yesterday-liveness-reward-api")

	if err != nil {
		return nil, err
	}

	return response, nil
}
