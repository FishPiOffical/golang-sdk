package fishPiSdk

import "strconv"

// BreezemoonContent 清风明月内容
type BreezemoonContent struct {
	OId                          string `json:"oId"`
	BreezemoonAuthorName         string `json:"breezemoonAuthorName"`
	BreezemoonAuthorThumbnailURL string `json:"breezemoonAuthorThumbnailURL"`
	BreezemoonContent            string `json:"breezemoonContent"`
	BreezemoonCreateTime         int64  `json:"breezemoonCreateTime"`
	BreezemoonCity               string `json:"breezemoonCity"`
	BreezemoonAuthorURL          string `json:"breezemoonAuthorURL"`
}

// BreezemoonListData 清风明月列表数据
type BreezemoonListData struct {
	Breezemoons []BreezemoonContent `json:"breezemoons"`
	Pagination  Pagination          `json:"pagination"`
}

// BreezemoonPostRequest 发布清风明月请求
type BreezemoonPostRequest struct {
	BreezemoonContent string `json:"breezemoonContent"`
}

// GetBreezemoonList 获取清风明月列表
func (c *Client) GetBreezemoonList(page, size int) (*ApiResponse[BreezemoonListData], error) {
	res, err := c.client.R().
		SetQueryParams(map[string]string{
			"p":    strconv.Itoa(page),
			"size": strconv.Itoa(size),
		}).
		Get("/api/breezemoons")
	if err != nil {
		return nil, err
	}

	var response ApiResponse[BreezemoonListData]
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetUserBreezemoonList 获取用户清风明月列表
func (c *Client) GetUserBreezemoonList(userName string, page, size int) (*ApiResponse[BreezemoonListData], error) {
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
