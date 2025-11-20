package sdk

import (
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
