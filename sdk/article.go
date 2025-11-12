package sdk

import (
	"fishpi-golang-sdk/types"
	"fmt"
)

// PostArticle 发布文章
func (s *FishPiSDK) PostArticle(req *types.PostArticleRequest) (string, error) {
	var resp types.ApiResponse[map[string]string]
	_, err := s.client.R().
		SetBodyJsonMarshal(req).
		SetSuccessResult(&resp).
		Post("/article")

	if err != nil {
		return "", err
	}

	if resp.Code != 0 {
		return "", fmt.Errorf(resp.Msg)
	}

	return resp.Data["articleId"], nil
}

// UpdateArticle 更新文章
func (s *FishPiSDK) UpdateArticle(articleId string, req *types.UpdateArticleRequest) error {
	var resp types.SimpleResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(req).
		SetSuccessResult(&resp).
		Put("/article/" + articleId)

	if err != nil {
		return err
	}

	if resp.Code != 0 {
		return fmt.Errorf(resp.Msg)
	}

	return nil
}

// GetArticleList 获取文章列表
func (s *FishPiSDK) GetArticleList(listType types.ArticleListType, tag string, page, size int) (*types.ArticleList, error) {
	url := fmt.Sprintf("/api/articles/%s?p=%d&size=%d", listType, page, size)
	if tag != "" {
		url = fmt.Sprintf("/api/articles/tag/%s/%s?p=%d&size=%d", tag, listType, page, size)
	}

	var resp types.ApiResponse[*types.ArticleList]
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

// GetUserArticles 获取用户文章列表
func (s *FishPiSDK) GetUserArticles(userName string, page, size int) (*types.ArticleList, error) {
	url := fmt.Sprintf("/api/user/%s/articles?p=%d&size=%d", userName, page, size)

	var resp types.ApiResponse[*types.ArticleList]
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

// GetArticleDetail 获取文章详情
func (s *FishPiSDK) GetArticleDetail(articleId string, page int) (*types.ArticleDetail, error) {
	url := fmt.Sprintf("/api/article/%s?p=%d", articleId, page)

	var resp types.ApiResponse[map[string]*types.ArticleDetail]
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get(url)

	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf(resp.Msg)
	}

	return resp.Data["article"], nil
}

// VoteArticle 文章投票（点赞/点踩）
func (s *FishPiSDK) VoteArticle(articleId string, voteType string) (types.VoteType, error) {
	url := fmt.Sprintf("/vote/%s/article", voteType)

	var resp types.ApiResponse[map[string]types.VoteType]
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"dataId": articleId,
		}).
		SetSuccessResult(&resp).
		Post(url)

	if err != nil {
		return -1, err
	}

	if resp.Code != 0 {
		return -1, fmt.Errorf(resp.Msg)
	}

	return resp.Data["type"], nil
}

// ThankArticle 感谢文章
func (s *FishPiSDK) ThankArticle(articleId string) error {
	var resp types.SimpleResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"articleId": articleId,
		}).
		SetSuccessResult(&resp).
		Post("/article/thank")

	if err != nil {
		return err
	}

	if resp.Code != 0 {
		return fmt.Errorf(resp.Msg)
	}

	return nil
}

// WatchArticle 关注文章
func (s *FishPiSDK) WatchArticle(articleId string, watch bool) error {
	var resp types.SimpleResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]interface{}{
			"articleId": articleId,
			"watch":     watch,
		}).
		SetSuccessResult(&resp).
		Post("/article/watch")

	if err != nil {
		return err
	}

	if resp.Code != 0 {
		return fmt.Errorf(resp.Msg)
	}

	return nil
}

// 旧版Client的Article方法
func (c *Client) PostArticle(req *types.PostArticleRequest) (string, error) {
	var resp types.ApiResponse[map[string]string]
	_, err := c.client.R().
		SetBodyJsonMarshal(req).
		SetSuccessResult(&resp).
		Post("/article")

	if err != nil {
		return "", err
	}

	if resp.Code != 0 {
		return "", fmt.Errorf(resp.Msg)
	}

	return resp.Data["articleId"], nil
}

func (c *Client) GetArticleList(listType, tag string, page, size int) (*types.ArticleList, error) {
	url := fmt.Sprintf("/api/articles/%s?p=%d&size=%d", listType, page, size)
	if tag != "" {
		url = fmt.Sprintf("/api/articles/tag/%s/%s?p=%d&size=%d", tag, listType, page, size)
	}

	var resp types.ApiResponse[*types.ArticleList]
	_, err := c.client.R().
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
