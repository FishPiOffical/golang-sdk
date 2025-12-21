package sdk

import (
	"net/url"

	"github.com/FishPiOffical/golang-sdk/types"
)

// PostMedalMyList 读取当前登录用户的所有勋章
func (s *FishPiSDK) PostMedalMyList() (*types.ApiResponse[[]*types.Medal], error) {
	response := new(types.ApiResponse[[]*types.Medal])
	_, err := s.client.R().
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/api/medal/my/list")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostMedalMyReorder 调整单个勋章顺序（上移/下移）
func (s *FishPiSDK) PostMedalMyReorder(medalId string, direction types.MedalReorderDirection) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"medalId":   medalId,
			"direction": direction,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/api/medal/my/reorder")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostMedalMyDisplay 调整当前登录用户隐藏/显示指定勋章
func (s *FishPiSDK) PostMedalMyDisplay(medalId string, display bool) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"medalId": medalId,
			"display": display,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/api/medal/my/display")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostMedalUserList 读取指定用户当前展示中的勋章列表（用于主页展示）
func (s *FishPiSDK) PostMedalUserList(data *types.PostMedalUserListRequest) (*types.ApiResponse[[]*types.Medal], error) {
	response := new(types.ApiResponse[[]*types.Medal])
	_, err := s.client.R().
		SetBodyJsonMarshal(data).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/api/medal/user/list")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// ------------------------------------------------------------ 管理侧 ------------------------------------------------------------

// PostMedalAdminList 分页读取全部勋章列表
func (s *FishPiSDK) PostMedalAdminList(page, pageSize int) (*types.ApiResponse[[]*types.Medal], error) {
	response := new(types.ApiResponse[[]*types.Medal])
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"goldFingerKey": s.GetConfig().MedalReadFingerKey,
			"page":          page,
			"pageSize":      pageSize,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/api/medal/admin/list")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostMedalAdminSearch 按关键词搜索勋章列表
// 在medal_id、medal_name、medal_description中搜索
func (s *FishPiSDK) PostMedalAdminSearch(keyword string) (*types.ApiResponse[[]*types.Medal], error) {
	response := new(types.ApiResponse[[]*types.Medal])
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"goldFingerKey": s.GetConfig().MedalReadFingerKey,
			"keyword":       keyword,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/api/medal/admin/search")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostMedalAdminDetail 读取指定勋章详细信息
func (s *FishPiSDK) PostMedalAdminDetail(medalId string) (*types.ApiResponse[*types.Medal], error) {
	response := new(types.ApiResponse[*types.Medal])
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"goldFingerKey": s.GetConfig().MedalReadFingerKey,
			"medalId":       medalId,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/api/medal/admin/detail")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostMedalAdminDelete 删除指定勋章
func (s *FishPiSDK) PostMedalAdminDelete(medalId string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"goldFingerKey": s.GetConfig().MedalWriteFingerKey,
			"medalId":       medalId,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/api/medal/admin/delete")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostMedalAdminEdit 编辑指定勋章
func (s *FishPiSDK) PostMedalAdminEdit(medalId string, name string, medalType types.MedalType, description string, attr string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"goldFingerKey": s.GetConfig().MedalWriteFingerKey,
			"medalId":       medalId,
			"name":          name,
			"type":          medalType,
			"description":   description,
			"attr":          attr,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/api/medal/admin/edit")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostMedalAdminCreate 新建勋章
func (s *FishPiSDK) PostMedalAdminCreate(name string, medalType types.MedalType, description string, attr string) (*types.ApiResponse[*types.Medal], error) {
	response := new(types.ApiResponse[*types.Medal])
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"goldFingerKey": s.GetConfig().MedalWriteFingerKey,
			"name":          name,
			"type":          medalType,
			"description":   description,
			"attr":          attr,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/api/medal/admin/create")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostMedalAdminGrant 给指定用户发指定勋章
func (s *FishPiSDK) PostMedalAdminGrant(userId string, medalId string, expireTime int64, data string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"goldFingerKey": s.GetConfig().MedalWriteFingerKey,
			"userId":        userId,
			"medalId":       medalId,
			"expireTime":    expireTime,
			"data":          data,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/api/medal/admin/grant")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostMedalAdminRevoke 给指定用户移除指定勋章
func (s *FishPiSDK) PostMedalAdminRevoke(userId string, medalId string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"goldFingerKey": s.GetConfig().MedalWriteFingerKey,
			"userId":        userId,
			"medalId":       medalId,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/api/medal/admin/revoke")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostMedalAdminOwners 读取指定勋章拥有的用户和拥有总数（分页）
func (s *FishPiSDK) PostMedalAdminOwners(medalId string, page, pageSize int) (*types.ApiResponse[*types.PostMedalAdminOwnersData], error) {
	response := new(types.ApiResponse[*types.PostMedalAdminOwnersData])
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"goldFingerKey": s.GetConfig().MedalReadFingerKey,
			"medalId":       medalId,
			"page":          page,
			"pageSize":      pageSize,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/api/medal/admin/owners")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetMedalUrl 获取勋章图片完整URL
func (s *FishPiSDK) GetMedalUrl(data *types.GetMedalUrlRequest) string {
	query := url.Values{}

	if data.MedalId != "" {
		query.Set("id", data.MedalId)
	}
	if data.MedalName != "" {
		query.Set("name", data.MedalName)
	}
	addr := url.URL{
		Path:     "/gen",
		RawQuery: query.Encode(),
	}

	return s.GetConfig().BaseUrl + addr.String()
}
