package fishPiSdk

import (
	"fishpi-golang-sdk/types"
	"strconv"
)

// BreezemoonContent 清风明月内容（别名）
type BreezemoonContent = types.BreezemoonInfo

// BreezemoonListData 清风明月列表数据（别名）
type BreezemoonListData = types.BreezemoonList

// BreezemoonPostRequest 发布清风明月请求（别名）
type BreezemoonPostRequest = types.PostBreezemoonRequest

// GetBreezemoonList 获取清风明月列表
func (c *Client) GetBreezemoonList(page, size int) (*types.ApiResponse[types.BreezemoonList], error) {
	res, err := c.client.R().
		SetQueryParams(map[string]string{
			"p":    strconv.Itoa(page),
			"size": strconv.Itoa(size),
		}).
		Get("/api/breezemoons")
	if err != nil {
		return nil, err
	}

	var response types.ApiResponse[types.BreezemoonList]
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetUserBreezemoonList 获取用户清风明月列表
func (c *Client) GetUserBreezemoonList(userName string, page, size int) (*types.ApiResponse[types.BreezemoonList], error) {
	res, err := c.client.R().
		SetQueryParams(map[string]string{
			"p":    strconv.Itoa(page),
			"size": strconv.Itoa(size),
		}).
		Get("/api/user/" + userName + "/breezemoons")
	if err != nil {
		return nil, err
	}

	var response ApiResponse[BreezemoonListData]
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// PostBreezemoon 发布清风明月
func (c *Client) PostBreezemoon(content string) (*SimpleResponse, error) {
	req := &BreezemoonPostRequest{
		BreezemoonContent: content,
	}

	res, err := c.client.R().
		SetBodyJsonMarshal(req).
		Post("/breezemoon")
	if err != nil {
		return nil, err
	}

	var response SimpleResponse
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
