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
