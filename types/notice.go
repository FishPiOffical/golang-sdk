package types

// UserMessage 用户频道消息
type UserMessage struct {
	Type    string      `json:"type"`
	Command string      `json:"command"`
	Data    interface{} `json:"data"`
}

// BreezemoonData 清风明月通知数据
type BreezemoonData struct {
	OId                          string `json:"oId"`
	BreezemoonContent            string `json:"breezemoonContent"`
	BreezemoonAuthorName         string `json:"breezemoonAuthorName"`
	BreezemoonAuthorThumbnailURL string `json:"breezemoonAuthorThumbnailURL"`
	BreezemoonCreateTime         int64  `json:"breezemoonCreateTime"`
	BreezemoonUpdated            int64  `json:"breezemoonUpdated"`
	BreezemoonCity               string `json:"breezemoonCity"`
}
