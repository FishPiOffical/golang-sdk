//go:generate go-enum --marshal --names --values --ptr --mustparse
package types

type Medal struct {
	OId string `json:"oId"` // 1766144126546 创建时只会返回这个字段

	MedalId          string    `json:"medal_id,omitempty"`          // 唯一ID
	MedalType        MedalType `json:"medal_type,omitempty"`        // 精良
	MedalName        string    `json:"medal_name,,omitempty"`       // 摸鱼派1岁啦
	MedalDescription string    `json:"medal_description,omitempty"` // 一往无前
	MedalAttr        string    `json:"medal_attr,omitempty"`        // url=https://file.fishpi.cn/2022/02/奖牌-d667b38a.jpg&backcolor=db5a6b&fontcolor=ffffff

	Data         string `json:"data,omitempty"`           //
	UserId       string `json:"user_id,omitempty"`        // 1656984017362
	Display      bool   `json:"display,omitempty"`        // true
	DisplayOrder int    `json:"display_order,omitempty"`  // 8
	ExpireTime   int64  `json:"expire_time,omitempty"`    // 0
	UserMedalOId string `json:"user_medal_oId,omitempty"` // 1766144156722
}

// MedalType 勋章类型
/*
ENUM(
normal=普通 // 黑
fine=精良 // 蓝
rare=稀有 // 紫
epic=史诗 // 橙+略加粗
legendary=传说 // 金+加粗
mythic=神话 // 更亮的金色
)
*/
type MedalType string

// MedalReorderDirection 勋章重新排序方向
/*
ENUM(
up // 上移
down // 下移
)
*/
type MedalReorderDirection string

type PostMedalUserListRequest struct {
	UserId   string `json:"userId,omitempty"`   // 用户ID
	UserName string `json:"userName,omitempty"` // 用户名
}

type MedalOwner struct {
	OId          string `json:"oId"`
	UserId       string `json:"user_id"`
	MedalId      string `json:"medal_id"`
	ExpireTime   int64  `json:"expire_time"`
	Display      bool   `json:"display"`
	DisplayOrder int    `json:"display_order"`
	Data         string `json:"data"`
	UserName     string `json:"userName"`
}

type PostMedalAdminOwnersData struct {
	Total int           `json:"total"`
	Items []*MedalOwner `json:"items"`
}

type GetMedalUrlRequest struct {
	MedalId   string
	MedalName string
}
