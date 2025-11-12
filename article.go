package fishPiSdk

import (
	"fishpi-golang-sdk/types"
	"strconv"
)

// ArticlePostRequest 发布文章请求（别名）
type ArticlePostRequest = types.PostArticleRequest

// ArticlePostResponse 发布文章响应
type ArticlePostResponse struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	ArticleId string `json:"articleId"`
}

// ArticleUpdateRequest 更新文章请求（别名）
type ArticleUpdateRequest = types.UpdateArticleRequest

// ArticleUpdateResponse 更新文章响应
type ArticleUpdateResponse = ArticlePostResponse

// ArticleListData 文章列表数据（别名）
type ArticleListData = types.ArticleList

// ArticleInfo 文章信息（别名）
type ArticleInfo = types.ArticleInfo

// PostArticle 发布文章
func (c *Client) PostArticle(req *ArticlePostRequest) (*ArticlePostResponse, error) {
	res, err := c.client.R().
		SetBodyJsonMarshal(req).
		Post("/article")
	if err != nil {
		return nil, err
	}

	var response ArticlePostResponse
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// PutArticle 更新文章
func (c *Client) PutArticle(articleId string, req *ArticleUpdateRequest) (*ArticleUpdateResponse, error) {
	res, err := c.client.R().
		SetBodyJsonMarshal(req).
		Put("/article/" + articleId)
	if err != nil {
		return nil, err
	}

	var response ArticleUpdateResponse
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetArticleList 获取文章列表
// listType: "", "/hot", "/good", "/reply", "/perfect"
// tag: 可选，指定标签
func (c *Client) GetArticleList(listType, tag string, page, size int) (*ApiResponse[ArticleListData], error) {
	path := "/api/articles/recent" + listType
	if tag != "" {
		path = "/api/articles/tag/" + tag + listType
	}

	res, err := c.client.R().
		SetQueryParams(map[string]string{
			"p":    strconv.Itoa(page),
			"size": strconv.Itoa(size),
		}).
		Get(path)
	if err != nil {
		return nil, err
	}

	var response ApiResponse[ArticleListData]
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// GetUserArticles 获取用户文章列表
func (c *Client) GetUserArticles(userName string, page, size int) (*ApiResponse[ArticleListData], error) {
	res, err := c.client.R().
		SetQueryParams(map[string]string{
			"p":    strconv.Itoa(page),
			"size": strconv.Itoa(size),
		}).
		Get("/api/user/" + userName + "/articles")
	if err != nil {
		return nil, err
	}

	var response ApiResponse[ArticleListData]
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
