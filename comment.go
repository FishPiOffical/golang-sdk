package fishPiSdk

import "fishpi-golang-sdk/types"

// CommentPostRequest 发布评论请求（别名）
type CommentPostRequest = types.PostCommentRequest

// CommentPostResponse 发布评论响应
type CommentPostResponse struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	CommentId string `json:"commentId"`
}

// CommentUpdateRequest 更新评论请求（别名）
type CommentUpdateRequest = types.UpdateCommentRequest

// CommentUpdateResponse 更新评论响应
type CommentUpdateResponse struct {
	Code           int    `json:"code"`
	Msg            string `json:"msg"`
	CommentContent string `json:"commentContent"`
}

// CommentVoteRequest 评论投票请求
type CommentVoteRequest struct {
	DataId string `json:"dataId"`
}

// CommentVoteResponse 评论投票响应
type CommentVoteResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Type int    `json:"type"` // -1:未投票 0:点赞 1:点踩
}

// PostComment 发布评论
func (c *Client) PostComment(req *CommentPostRequest) (*CommentPostResponse, error) {
	res, err := c.client.R().
		SetBodyJsonMarshal(req).
		Post("/comment")
	if err != nil {
		return nil, err
	}

	var response CommentPostResponse
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// PutComment 更新评论
func (c *Client) PutComment(commentId string, req *CommentUpdateRequest) (*CommentUpdateResponse, error) {
	res, err := c.client.R().
		SetBodyJsonMarshal(req).
		Put("/comment/" + commentId)
	if err != nil {
		return nil, err
	}

	var response CommentUpdateResponse
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// VoteUpComment 评论点赞
func (c *Client) VoteUpComment(commentId string) (*CommentVoteResponse, error) {
	req := &CommentVoteRequest{
		DataId: commentId,
	}

	res, err := c.client.R().
		SetBodyJsonMarshal(req).
		Post("/vote/up/comment")
	if err != nil {
		return nil, err
	}

	var response CommentVoteResponse
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// VoteDownComment 评论点踩
func (c *Client) VoteDownComment(commentId string) (*CommentVoteResponse, error) {
	req := &CommentVoteRequest{
		DataId: commentId,
	}

	res, err := c.client.R().
		SetBodyJsonMarshal(req).
		Post("/vote/down/comment")
	if err != nil {
		return nil, err
	}

	var response CommentVoteResponse
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}

// ThankComment 感谢评论
func (c *Client) ThankComment(commentId string) (*SimpleResponse, error) {
	req := &CommentVoteRequest{
		DataId: commentId,
	}

	res, err := c.client.R().
		SetBodyJsonMarshal(req).
		Post("/comment/thank")
	if err != nil {
		return nil, err
	}

	var response SimpleResponse
	if err = res.Unmarshal(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
