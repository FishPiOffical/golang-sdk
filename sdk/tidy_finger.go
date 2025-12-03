package sdk

import (
	"github.com/FishPiOffical/golang-sdk/types"
)

// PostGiveMetal 添加勋章
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
