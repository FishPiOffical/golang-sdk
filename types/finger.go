package types

// FingerAPI 金手指相关类型

// UserIP 用户IP信息
type UserIP struct {
	IP string `json:"ip"`
}

// UserBag 用户背包
type UserBag struct {
	Items []string `json:"items"`
}

// MetalBase 勋章基础信息
type MetalBase struct {
	Name            string   `json:"name"`
	Attr            []string `json:"attr"`
	Description     string   `json:"description"`
	Data            string   `json:"data"`
	BackgroundImage string   `json:"backgroundImage"`
}

// MofishScoreRequest 摸鱼大闯关分数请求
type MofishScoreRequest struct {
	GoldFingerKey string `json:"goldFingerKey"`
	UserName      string `json:"userName"`
	Stage         string `json:"stage"`
	Time          int64  `json:"time"`
}
