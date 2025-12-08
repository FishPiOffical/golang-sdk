package sdk

import (
	"fmt"
	"strings"

	"github.com/FishPiOffical/golang-sdk/types"
)

// ArticleChannelWebSocket 用户通知WebSocket客户端
type ArticleChannelWebSocket struct {
	*Client[types.ArticleChannelMsg]
}

// NewArticleChannelWebSocket 创建文章通知WebSocket连接
func (s *FishPiSDK) NewArticleChannelWebSocket(articleId string, articleType types.ArticleType, opts ...ClientOption[types.ArticleChannelMsg]) *ArticleChannelWebSocket {
	baseUrl := strings.Replace(s.GetConfig().BaseUrl, "http", "ws", 1)
	endpoint := fmt.Sprintf("%s/article-channel?apiKey=%s&articleId=%s&articleType=%s", baseUrl, s.GetAPIKey(), articleId, articleType)

	// 将 endpoint 作为选项传入
	allOpts := append([]ClientOption[types.ArticleChannelMsg]{
		WithEndpoint[types.ArticleChannelMsg](endpoint),
	}, opts...)

	client := NewClient(allOpts...)

	return &ArticleChannelWebSocket{
		Client: client,
	}
}
