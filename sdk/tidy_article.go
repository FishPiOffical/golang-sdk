package sdk

import (
	"fmt"

	"github.com/FishPiOffical/golang-sdk/types"
)

// PostArticle 发布文章
func (s *FishPiSDK) PostArticle(req *types.PostArticleRequest) (*types.PostArticleResponse, error) {
	response := new(types.PostArticleResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(req).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/article")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PutArticle 更新文章
func (s *FishPiSDK) PutArticle(articleId string, req *types.PostArticleRequest) (*types.PostArticleResponse, error) {

	response := new(types.PostArticleResponse)
	_, err := s.client.R().
		SetBodyJsonMarshal(req).
		SetPathParam("articleId", articleId).
		SetSuccessResult(response).
		SetErrorResult(response).
		Put("/article/{articleId}")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetArticles 获取文章列表
func (s *FishPiSDK) GetArticles(req *types.GetArticlesRequest) (*types.ApiResponse[*types.ArticleList], error) {
	response := new(types.ApiResponse[*types.ArticleList])

	url := fmt.Sprintf("/api/articles%s", req.ToPath())

	_, err := s.client.R().
		SetBodyJsonMarshal(req).
		SetQueryParamsAnyType(map[string]any{
			"p":    req.Page,
			"size": req.Size,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Get(url)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetArticleDetail 获取文章详情
func (s *FishPiSDK) GetArticleDetail(articleId string) (*types.ApiResponse[*types.GetArticleData], error) {
	response := new(types.ApiResponse[*types.GetArticleData])

	_, err := s.client.R().
		SetPathParam("articleId", articleId).
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/api/article/{articleId}")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetUserArticles 获取用户文章列表
func (s *FishPiSDK) GetUserArticles(userName string, page, size int) (*types.ApiResponse[*types.ArticleList], error) {
	response := new(types.ApiResponse[*types.ArticleList])

	_, err := s.client.R().
		SetQueryParamsAnyType(map[string]any{
			"p":    page,
			"size": size,
		}).
		SetPathParam("userName", userName).
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/api/user/{userName}/articles")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostVoteUpArticle 给文章点赞
func (s *FishPiSDK) PostVoteUpArticle(articleId string) (*types.PostVoteUpArticleResponse, error) {
	response := new(types.PostVoteUpArticleResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"dataId": articleId,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/vote/up/article")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostArticleThank 感谢文章
func (s *FishPiSDK) PostArticleThank(articleId string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetQueryParam("articleId", articleId).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/article/thank")

	if err != nil {
		return nil, err
	}

	return response, nil
}
