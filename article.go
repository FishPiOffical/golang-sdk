package fishPiSdk

import "strconv"

// ArticlePostRequest 发布文章请求
type ArticlePostRequest struct {
	ArticleTitle           string `json:"articleTitle"`
	ArticleContent         string `json:"articleContent"`
	ArticleTags            string `json:"articleTags"`
	ArticleCommentable     bool   `json:"articleCommentable"`
	ArticleNotifyFollowers bool   `json:"articleNotifyFollowers"`
	ArticleType            int    `json:"articleType"`
	ArticleShowInList      int    `json:"articleShowInList"` // 1:显示 0:不显示
	ArticleRewardContent   string `json:"articleRewardContent,omitempty"`
	ArticleRewardPoint     int    `json:"articleRewardPoint,omitempty"`
	ArticleAnonymous       bool   `json:"articleAnonymous,omitempty"`
	ArticleQnAOfferPoint   int    `json:"articleQnAOfferPoint,omitempty"`
}

// ArticlePostResponse 发布文章响应
type ArticlePostResponse struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	ArticleId string `json:"articleId"`
}

// ArticleUpdateRequest 更新文章请求
type ArticleUpdateRequest = ArticlePostRequest

// ArticleUpdateResponse 更新文章响应
type ArticleUpdateResponse = ArticlePostResponse

// ArticleListData 文章列表数据
type ArticleListData struct {
	Articles   []ArticleInfo `json:"articles"`
	Pagination Pagination    `json:"pagination"`
}

// ArticleInfo 文章信息
type ArticleInfo struct {
	OId                       string `json:"oId"`
	ArticleTitle              string `json:"articleTitle"`
	ArticlePermalink          string `json:"articlePermalink"`
	ArticleTags               string `json:"articleTags"`
	ArticleAuthorName         string `json:"articleAuthorName"`
	ArticleAuthorThumbnailURL string `json:"articleAuthorThumbnailURL"`
	ArticleCreateTime         int64  `json:"articleCreateTime"`
	ArticleUpdateTime         int64  `json:"articleUpdateTime"`
	ArticleViewCount          int    `json:"articleViewCount"`
	ArticleCommentCount       int    `json:"articleCommentCount"`
	ArticleThankCount         int    `json:"articleThankCount"`
	ArticleGoodCnt            int    `json:"articleGoodCnt"`
	ArticleBadCnt             int    `json:"articleBadCnt"`
	ArticleCollectCnt         int    `json:"articleCollectCnt"`
	ArticleWatchCnt           int    `json:"articleWatchCnt"`
	ArticleContent            string `json:"articleContent"`
	ArticlePreviewContent     string `json:"articlePreviewContent"`
	ArticleStick              int64  `json:"articleStick"`
	ArticleAnonymous          int    `json:"articleAnonymous"`
	ArticlePerfect            int    `json:"articlePerfect"`
	ArticleType               int    `json:"articleType"`
	ArticleStatus             int    `json:"articleStatus"`
	OfferedPoint              int    `json:"offeredPoint"`
	RewardedPoint             int    `json:"rewardedPoint"`
}

// Pagination 分页信息
type Pagination struct {
	PageNum     int `json:"paginationPageNum"`
	PageSize    int `json:"paginationPageSize"`
	PageCount   int `json:"paginationPageCount"`
	RecordCount int `json:"paginationRecordCount"`
}

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
