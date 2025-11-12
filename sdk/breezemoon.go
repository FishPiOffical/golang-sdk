package sdk

import (
	"fishpi-golang-sdk/types"
	"fmt"
)

// PostBreezemoon 发布清风明月
func (s *FishPiSDK) PostBreezemoon(req *types.PostBreezemoonRequest) error {
	var resp types.SimpleResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(req).
		SetSuccessResult(&resp).
		Post("/breezemoon")
	if err != nil {
		return err
	}
	if resp.Code != 0 {
		return fmt.Errorf(resp.Msg)
	}
	return nil
}

// UpdateBreezemoon 更新清风明月
func (s *FishPiSDK) UpdateBreezemoon(breezemoonId string, req *types.UpdateBreezemoonRequest) error {
	var resp types.SimpleResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(req).
		SetSuccessResult(&resp).
		Put("/breezemoon/" + breezemoonId)
	if err != nil {
		return err
	}
	if resp.Code != 0 {
		return fmt.Errorf(resp.Msg)
	}
	return nil
}

// GetBreezemoonList 获取清风明月列表
func (s *FishPiSDK) GetBreezemoonList(page, size int) (*types.BreezemoonList, error) {
	url := fmt.Sprintf("/api/breezemoons?p=%d&size=%d", page, size)
	var resp types.ApiResponse[*types.BreezemoonList]
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

// GetUserBreezemoons 获取用户清风明月列表
func (s *FishPiSDK) GetUserBreezemoons(userName string, page, size int) (*types.BreezemoonList, error) {
	url := fmt.Sprintf("/api/user/%s/breezemoons?p=%d&size=%d", userName, page, size)
	var resp types.ApiResponse[*types.BreezemoonList]
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

// RemoveBreezemoon 删除清风明月
func (s *FishPiSDK) RemoveBreezemoon(breezemoonId string) error {
	var resp types.SimpleResponse
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Delete("/breezemoon/" + breezemoonId)
	if err != nil {
		return err
	}
	if resp.Code != 0 {
		return fmt.Errorf(resp.Msg)
	}
	return nil
}
