package sdk

import (
	"fmt"

	"github.com/FishPiOffical/golang-sdk/types"
)

// Finger 金手指API（需要金手指密钥）
type Finger struct {
	sdk           *FishPiSDK
	goldFingerKey string
}

// NewFinger 创建金手指实例
func (s *FishPiSDK) NewFinger(goldFingerKey string) *Finger {
	return &Finger{
		sdk:           s,
		goldFingerKey: goldFingerKey,
	}
}

// AddMofishScore 上传摸鱼大闯关关卡数据
func (f *Finger) AddMofishScore(userName, stage string, time int64) error {
	var resp types.SimpleResponse
	_, err := f.sdk.client.R().
		SetBodyJsonMarshal(map[string]interface{}{
			"goldFingerKey": f.goldFingerKey,
			"userName":      userName,
			"stage":         stage,
			"time":          time,
		}).
		SetSuccessResult(&resp).
		Post("/api/games/mofish/score")

	if err != nil {
		return err
	}

	if resp.Code != 0 {
		return fmt.Errorf(resp.Msg)
	}

	return nil
}

// QueryLatestLoginIP 查询用户最近登录的IP地址
func (f *Finger) QueryLatestLoginIP(userName string) (*types.UserIP, error) {
	var resp types.ApiResponse[*types.UserIP]
	_, err := f.sdk.client.R().
		SetBodyJsonMarshal(map[string]string{
			"goldFingerKey": f.goldFingerKey,
			"userName":      userName,
		}).
		SetSuccessResult(&resp).
		Post("/user/query/latest-login-iP")

	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Msg)
	}

	return resp.Data, nil
}

// RemoveMetalByUserId 通过用户ID移除勋章
func (f *Finger) RemoveMetalByUserId(userId, metalName string) error {
	if userId == "" {
		return fmt.Errorf("userId is required")
	}
	if metalName == "" {
		return fmt.Errorf("metalName is required")
	}

	var resp types.SimpleResponse
	_, err := f.sdk.client.R().
		SetBodyJsonMarshal(map[string]string{
			"goldFingerKey": f.goldFingerKey,
			"userId":        userId,
			"name":          metalName,
		}).
		SetSuccessResult(&resp).
		Post("/user/edit/remove-metal-by-user-id")

	if err != nil {
		return fmt.Errorf("failed to remove metal by user id: %w", err)
	}

	if resp.Code != 0 {
		return fmt.Errorf("remove metal by user id failed: %s", resp.Msg)
	}

	return nil
}

// QueryUserBag 查询用户背包
func (f *Finger) QueryUserBag(userName string) (*types.UserBag, error) {
	if userName == "" {
		return nil, fmt.Errorf("userName is required")
	}

	var resp types.ApiResponse[*types.UserBag]
	_, err := f.sdk.client.R().
		SetBodyJsonMarshal(map[string]string{
			"goldFingerKey": f.goldFingerKey,
			"userName":      userName,
		}).
		SetSuccessResult(&resp).
		Post("/user/query/items")

	if err != nil {
		return nil, fmt.Errorf("failed to query user bag: %w", err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("query user bag failed: %s", resp.Msg)
	}

	return resp.Data, nil
}

// EditUserBag 调整用户背包
func (f *Finger) EditUserBag(userName, item string, sum int) error {
	if userName == "" {
		return fmt.Errorf("userName is required")
	}
	if item == "" {
		return fmt.Errorf("item is required")
	}

	var resp types.SimpleResponse
	_, err := f.sdk.client.R().
		SetBodyJsonMarshal(map[string]interface{}{
			"goldFingerKey": f.goldFingerKey,
			"userName":      userName,
			"item":          item,
			"sum":           sum,
		}).
		SetSuccessResult(&resp).
		Post("/user/edit/items")

	if err != nil {
		return fmt.Errorf("failed to edit user bag: %w", err)
	}

	if resp.Code != 0 {
		return fmt.Errorf("edit user bag failed: %s", resp.Msg)
	}

	return nil
}

// EditUserPoints 调整用户积分
func (f *Finger) EditUserPoints(userName string, point int, memo string) error {
	if userName == "" {
		return fmt.Errorf("userName is required")
	}

	var resp types.SimpleResponse
	_, err := f.sdk.client.R().
		SetBodyJsonMarshal(map[string]interface{}{
			"goldFingerKey": f.goldFingerKey,
			"userName":      userName,
			"point":         point,
			"memo":          memo,
		}).
		SetSuccessResult(&resp).
		Post("/user/edit/points")

	if err != nil {
		return fmt.Errorf("failed to edit user points: %w", err)
	}

	if resp.Code != 0 {
		return fmt.Errorf("edit user points failed: %s", resp.Msg)
	}

	return nil
}

// GetUserLiveness 查询用户当前活跃度
func (f *Finger) GetUserLiveness(userName string) (int, error) {
	if userName == "" {
		return 0, fmt.Errorf("userName is required")
	}

	var resp types.ApiResponse[map[string]int]
	_, err := f.sdk.client.R().
		SetBodyJsonMarshal(map[string]string{
			"goldFingerKey": f.goldFingerKey,
			"userName":      userName,
		}).
		SetSuccessResult(&resp).
		Post("/user/liveness")

	if err != nil {
		return 0, fmt.Errorf("failed to get user liveness: %w", err)
	}

	if resp.Code != 0 {
		return 0, fmt.Errorf("get user liveness failed: %s", resp.Msg)
	}

	if resp.Data != nil {
		if liveness, ok := resp.Data["liveness"]; ok {
			return liveness, nil
		}
	}

	return 0, nil
}
