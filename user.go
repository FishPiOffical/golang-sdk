package fishPiSdk

// UserInfoResponse 用户信息响应
type UserInfoResponse struct {
	Code int            `json:"code"`
	Msg  string         `json:"msg"`
	Data ApiUserGetData `json:"data"`
}

// UserLivenessResponse 用户活跃度响应
type UserLivenessResponse struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	Liveness int    `json:"liveness"`
}

// UserCheckedInResponse 用户签到状态响应
type UserCheckedInResponse struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	CheckedIn bool   `json:"checkedIn"`
}

// IsCollectedLivenessResponse 是否已领取昨日活跃奖励响应
type IsCollectedLivenessResponse struct {
	Code                               int    `json:"code"`
	Msg                                string `json:"msg"`
	IsCollectedYesterdayLivenessReward bool   `json:"isCollectedYesterdayLivenessReward"`
}

// YesterdayLivenessRewardResponse 领取昨日活跃奖励响应
type YesterdayLivenessRewardResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Sum  int    `json:"sum"`
}

// UserEmotionsResponse 用户常用表情响应
type UserEmotionsResponse struct {
	Code int                `json:"code"`
	Msg  string             `json:"msg"`
	Data map[string]float64 `json:"data"`
}

// GetUserInfo 获取当前用户信息
func (c *Client) GetUserInfo() (*UserInfoResponse, error) {
	res, err := c.client.R().Get("/api/user")
	if err != nil {
		return nil, err
	}

	var response UserInfoResponse
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetUserLiveness 获取当前用户活跃度
func (c *Client) GetUserLiveness() (*UserLivenessResponse, error) {
	res, err := c.client.R().Get("/user/liveness")
	if err != nil {
		return nil, err
	}

	var response UserLivenessResponse
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetUserCheckedIn 检查是否已签到
func (c *Client) GetUserCheckedIn() (*UserCheckedInResponse, error) {
	res, err := c.client.R().Get("/user/checkedIn")
	if err != nil {
		return nil, err
	}

	var response UserCheckedInResponse
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetIsCollectedLiveness 检查是否已领取昨日活跃奖励
func (c *Client) GetIsCollectedLiveness() (*IsCollectedLivenessResponse, error) {
	res, err := c.client.R().Get("/api/activity/is-collected-liveness")
	if err != nil {
		return nil, err
	}

	var response IsCollectedLivenessResponse
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetYesterdayLivenessReward 领取昨日活跃度奖励
func (c *Client) GetYesterdayLivenessReward() (*YesterdayLivenessRewardResponse, error) {
	res, err := c.client.R().Get("/activity/yesterday-liveness-reward-api")
	if err != nil {
		return nil, err
	}

	var response YesterdayLivenessRewardResponse
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// PostPointTransfer 转账
func (c *Client) PostPointTransfer(req *TransferRequest) (*SimpleResponse, error) {
	res, err := c.client.R().
		SetBodyJsonMarshal(req).
		Post("/point/transfer")
	if err != nil {
		return nil, err
	}

	var response SimpleResponse
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetUserEmotions 获取用户常用表情
func (c *Client) GetUserEmotions() (*UserEmotionsResponse, error) {
	res, err := c.client.R().Get("/users/emotions")
	if err != nil {
		return nil, err
	}

	var response UserEmotionsResponse
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
