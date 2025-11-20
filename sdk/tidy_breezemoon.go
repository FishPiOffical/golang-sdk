package sdk

import (
	"github.com/FishPiOffical/golang-sdk/types"
)

// GetBreezemoons 获取清风明月列表
func (s *FishPiSDK) GetBreezemoons(page, size int) (*types.GetBreezemoonsResponse, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 20
	}

	response := new(types.GetBreezemoonsResponse)

	_, err := s.client.R().
		SetQueryParamsAnyType(map[string]any{
			"p":    page,
			"size": size,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/api/breezemoons")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostBreezemoon 发送清风明月
func (s *FishPiSDK) PostBreezemoon(content string) (*types.ApiResponse[*types.BreezemoonInfo], error) {
	response := new(types.ApiResponse[*types.BreezemoonInfo])

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"breezemoonContent": content,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/breezemoon")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetUserBreezemoons 获取指定用户的明月清风列表
func (s *FishPiSDK) GetUserBreezemoons(userName string, page, size int) (*types.ApiResponse[*types.GetUserBreezemoonsData], error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 20
	}
	response := new(types.ApiResponse[*types.GetUserBreezemoonsData])

	_, err := s.client.R().
		SetQueryParamsAnyType(map[string]any{
			"p":    page,
			"size": size,
		}).
		SetPathParam("userName", userName).
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/api/user/{userName}/breezemoons")

	if err != nil {
		return nil, err
	}

	return response, nil
}
