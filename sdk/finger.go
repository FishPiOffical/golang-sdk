package sdk

import (
	"fishpi-golang-sdk/types"
	"fmt"
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

// AddMetal 添加勋章
func (f *Finger) AddMetal(userName string, metal *types.MetalBase) error {
	var resp types.SimpleResponse
	_, err := f.sdk.client.R().
		SetBodyJsonMarshal(map[string]interface{}{
			"goldFingerKey":   f.goldFingerKey,
			"userName":        userName,
			"name":            metal.Name,
			"attr":            metal.Attr,
			"description":     metal.Description,
			"data":            metal.Data,
			"backgroundImage": metal.BackgroundImage,
		}).
		SetSuccessResult(&resp).
		Post("/user/edit/give-metal")

	if err != nil {
		return err
	}

	if resp.Code != 0 {
		return fmt.Errorf(resp.Msg)
	}

	return nil
}

// RemoveMetal 移除勋章
func (f *Finger) RemoveMetal(userName, metalName string) error {
	var resp types.SimpleResponse
	_, err := f.sdk.client.R().
		SetBodyJsonMarshal(map[string]string{
			"goldFingerKey": f.goldFingerKey,
			"userName":      userName,
			"name":          metalName,
		}).
		SetSuccessResult(&resp).
		Post("/user/edit/remove-metal")

	if err != nil {
		return err
	}

	if resp.Code != 0 {
		return fmt.Errorf(resp.Msg)
	}

	return nil
}
