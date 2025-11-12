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
		return "", fmt.Errorf("failed to post article: %w", err)
	}

	if resp.Code != 0 {
		return "", fmt.Errorf("post article failed: %s", resp.Msg)
	}

	articleId, ok := resp.Data["articleId"]
	if !ok {
		return "", fmt.Errorf("articleId not found in response")
	}

	return articleId, nil
}

// UpdateArticle 更新文章
func (s *FishPiSDK) UpdateArticle(articleId string, req *types.UpdateArticleRequest) error {
	if articleId == "" {
		return fmt.Errorf("articleId is required")
	}

	var resp types.SimpleResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(req).
		SetSuccessResult(&resp).
		Put("/article/" + articleId)

	if err != nil {
		return fmt.Errorf("failed to update article: %w", err)
	}

	if resp.Code != 0 {
		return fmt.Errorf("update article failed: %s", resp.Msg)
	}

	return nil
}

// GetArticleList 获取文章列表
func (s *FishPiSDK) GetArticleList(listType types.ArticleListType, tag string, page, size int) (*types.ArticleList, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20
	}

	url := fmt.Sprintf("/api/articles/%s?p=%d&size=%d", listType, page, size)
	if tag != "" {
		url = fmt.Sprintf("/api/articles/tag/%s/%s?p=%d&size=%d", tag, listType, page, size)
	}

	var resp types.ApiResponse[*types.ArticleList]
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get(url)

	if err != nil {
		return nil, fmt.Errorf("failed to get article list: %w", err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("get article list failed: %s", resp.Msg)
	}

	if resp.Data == nil {
		return nil, fmt.Errorf("article list data is nil")
	}

	return resp.Data, nil
}

// GetUserArticles 获取用户文章列表
func (s *FishPiSDK) GetUserArticles(userName string, page, size int) (*types.ArticleList, error) {
	if userName == "" {
		return nil, fmt.Errorf("userName is required")
	}
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20
	}

	url := fmt.Sprintf("/api/user/%s/articles?p=%d&size=%d", userName, page, size)

	var resp types.ApiResponse[*types.ArticleList]
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get(url)

	if err != nil {
		return nil, fmt.Errorf("failed to get user articles: %w", err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("get user articles failed: %s", resp.Msg)
	}

	if resp.Data == nil {
		return nil, fmt.Errorf("user articles data is nil")
	}

	return resp.Data, nil
}

// GetArticleDetail 获取文章详情
func (s *FishPiSDK) GetArticleDetail(articleId string, page int) (*types.ArticleDetail, error) {
	if articleId == "" {
		return nil, fmt.Errorf("articleId is required")
	}
	if page < 1 {
		page = 1
	}

	url := fmt.Sprintf("/api/article/%s?p=%d", articleId, page)

	var resp types.ApiResponse[map[string]*types.ArticleDetail]
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Get(url)

	if err != nil {
		return nil, fmt.Errorf("failed to get article detail: %w", err)
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("get article detail failed: %s", resp.Msg)
	}

	if resp.Data == nil {
		return nil, fmt.Errorf("article detail data is nil")
	}

	article, ok := resp.Data["article"]
	if !ok || article == nil {
		return nil, fmt.Errorf("article not found in response")
	}

	return article, nil
}

// VoteArticle 文章投票（点赞/点踩）
func (s *FishPiSDK) VoteArticle(articleId string, voteType string) (types.VoteType, error) {
	if articleId == "" {
		return types.VoteTypeUnvoted, fmt.Errorf("articleId is required")
	}
	if voteType != "up" && voteType != "down" {
		return types.VoteTypeUnvoted, fmt.Errorf("voteType must be 'up' or 'down'")
	}

	url := fmt.Sprintf("/vote/%s/article", voteType)

	var resp types.ApiResponse[map[string]types.VoteType]
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"dataId": articleId,
		}).
		SetSuccessResult(&resp).
		Post(url)

	if err != nil {
		return types.VoteTypeUnvoted, fmt.Errorf("failed to vote article: %w", err)
	}

	if resp.Code != 0 {
		return types.VoteTypeUnvoted, fmt.Errorf("vote article failed: %s", resp.Msg)
	}

	if resp.Data == nil {
		return types.VoteTypeUnvoted, fmt.Errorf("vote response data is nil")
	}

	voteResult, ok := resp.Data["type"]
	if !ok {
		return types.VoteTypeUnvoted, fmt.Errorf("vote type not found in response")
	}

	return voteResult, nil
}

// ThankArticle 感谢文章
func (s *FishPiSDK) ThankArticle(articleId string) error {
	if articleId == "" {
		return fmt.Errorf("articleId is required")
	}

	var resp types.SimpleResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"articleId": articleId,
		}).
		SetSuccessResult(&resp).
		Post("/article/thank")

	if err != nil {
		return fmt.Errorf("failed to thank article: %w", err)
	}

	if resp.Code != 0 {
		return fmt.Errorf("thank article failed: %s", resp.Msg)
	}

	return nil
}

// WatchArticle 关注文章
func (s *FishPiSDK) WatchArticle(articleId string, watch bool) error {
	if articleId == "" {
		return fmt.Errorf("articleId is required")
	}

	var resp types.SimpleResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]interface{}{
			"articleId": articleId,
			"watch":     watch,
		}).
		SetSuccessResult(&resp).
		Post("/article/watch")

	if err != nil {
		return fmt.Errorf("failed to watch article: %w", err)
	}

	if resp.Code != 0 {
		return fmt.Errorf("watch article failed: %s", resp.Msg)
	}

	return nil
}

// FollowArticle 收藏/取消收藏文章
func (s *FishPiSDK) FollowArticle(followingId string) error {
	if followingId == "" {
		return fmt.Errorf("followingId is required")
	}

	var resp types.SimpleResponse
	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"followingId": followingId,
		}).
		SetSuccessResult(&resp).
		Post("/follow/article")

	if err != nil {
		return fmt.Errorf("failed to follow article: %w", err)
	}

	if resp.Code != 0 {
		return fmt.Errorf("follow article failed: %s", resp.Msg)
	}

	return nil
}

// RewardArticle 打赏文章
func (s *FishPiSDK) RewardArticle(articleId string) error {
	if articleId == "" {
		return fmt.Errorf("articleId is required")
	}

	var resp types.SimpleResponse
	_, err := s.client.R().
		SetSuccessResult(&resp).
		Post("/article/reward?articleId=" + articleId)

	if err != nil {
		return fmt.Errorf("failed to reward article: %w", err)
	}

	if resp.Code != 0 {
		return fmt.Errorf("reward article failed: %s", resp.Msg)
	}

	return nil
}
