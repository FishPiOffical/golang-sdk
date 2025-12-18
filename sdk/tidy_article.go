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

// GetArticleComments 获取帖子的评论列表
func (s *FishPiSDK) GetArticleComments(articleId string, page, size int) (*types.ApiResponse[*types.GetArticleCommentsData], error) {
	response := new(types.ApiResponse[*types.GetArticleCommentsData])

	_, err := s.client.R().
		SetQueryParamsAnyType(map[string]any{
			"p":    page,
			"size": size,
		}).
		SetPathParam("articleId", articleId).
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/api/comment/{articleId}")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostComment 评论/回复
func (s *FishPiSDK) PostComment(req *types.PostCommentRequest) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(req).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/comment")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PutComment 更新评论
func (s *FishPiSDK) PutComment(commentId string, req *types.PutCommentRequest) (*types.PutCommentResponse, error) {
	response := new(types.PutCommentResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(req).
		SetPathParam("commentId", commentId).
		SetSuccessResult(response).
		SetErrorResult(response).
		Put("/comment/{commentId}")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostVoteUpComment 评论点赞
func (s *FishPiSDK) PostVoteUpComment(commentId string) (*types.PostVoteUpCommentResponse, error) {
	response := new(types.PostVoteUpCommentResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"dataId": commentId,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/vote/up/comment")
	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostCommentThank 感谢评论
func (s *FishPiSDK) PostCommentThank(commentId string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]string{
			"commentId": commentId,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/comment/thank")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostCommentRemove 删除评论
func (s *FishPiSDK) PostCommentRemove(commentId string) (*types.PostCommentRemoveResponse, error) {
	response := new(types.PostCommentRemoveResponse)

	_, err := s.client.R().
		SetPathParam("commentId", commentId).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/comment/{commentId}/remove")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostArticleHeat 获取帖子当前正在阅读的人数
func (s *FishPiSDK) PostArticleHeat(articleId string) (*types.PostArticleHeatResponse, error) {
	response := new(types.PostArticleHeatResponse)

	_, err := s.client.R().
		SetPathParam("articleId", articleId).
		SetSuccessResult(response).
		SetErrorResult(response).
		Get("/api/article/heat/{articleId}")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostFollowArticle 收藏文章
func (s *FishPiSDK) PostFollowArticle(articleId string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"followingId": articleId,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/follow/article")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostUnfollowArticle 取消收藏文章
func (s *FishPiSDK) PostUnfollowArticle(articleId string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"followingId": articleId,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/unfollow/article")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostFollowArticleWatch 关注文章
func (s *FishPiSDK) PostFollowArticleWatch(articleId string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"followingId": articleId,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/follow/article-watch")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostUnfollowArticleWatch 取消关注文章
func (s *FishPiSDK) PostUnfollowArticleWatch(articleId string) (*types.SimpleResponse, error) {
	response := new(types.SimpleResponse)

	_, err := s.client.R().
		SetBodyJsonMarshal(map[string]any{
			"followingId": articleId,
		}).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/unfollow/article-watch")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostArticleReward 打赏文章
func (s *FishPiSDK) PostArticleReward(articleId string) (*types.PostArticleRewardResponse, error) {
	response := new(types.PostArticleRewardResponse)

	_, err := s.client.R().
		SetQueryParam("articleId", articleId).
		SetSuccessResult(response).
		SetErrorResult(response).
		Post("/article/reward")

	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetArticleMd 获取帖子的Markdown原文
func (s *FishPiSDK) GetArticleMd(articleId string) (*string, error) {
	resp, err := s.client.R().
		SetPathParam("articleId", articleId).
		Get("/api/article/md/{articleId}")

	if err != nil {
		return nil, err
	}

	result := resp.String()

	return &result, nil
}
