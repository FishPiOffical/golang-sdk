//go:generate go-enum --marshal --names --values --ptr --mustparse
package types

type Medal struct {
	OId string `json:"oId"` // 1766144126546 创建时只会返回这个字段

	MedalId          string    `json:"medal_id,omitempty"`          // 唯一ID
	MedalType        MedalType `json:"medal_type,omitempty"`        // 精良
	MedalName        string    `json:"medal_name,omitempty"`        // 摸鱼派1岁啦
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

type Badge struct {
	Ver       string  `json:"ver"`       // 0.1
	Scale     float64 `json:"scale"`     // 0.79
	Txt       string  `json:"txt"`       // 摸鱼派1岁啦
	Url       string  `json:"url"`       // https://file.fishpi.cn/2022/02/奖牌-d667b38a.jpg
	Backcolor string  `json:"backcolor"` // 多个6位16进制色值（用英文逗号分隔），或单独auto，不能与颜色混用 ffffff,000000,ffa500,ff0000 或 auto
	Way       string  `json:"way"`       // 背景渐变方向 BadgeWay 或 0deg~359deg
	Fontcolor string  `json:"fontcolor"` // 多个6位16进制色值（用英文逗号分隔），或单独auto，不能与颜色混用 ffffff,000000 或 auto
	FontWay   string  `json:"fontway"`   // 同 Way bottom 120deg
	Shadow    float64 `json:"shadow"`    // 背景阴影浓度 0.8
	Anime     float64 `json:"anime"`     // 动画时间（秒）建议范围0.1~10 5
	Size      int     `json:"size"`      // 徽章尺寸 建议范围16~512 32
	Border    int     `json:"border"`    // 边距和阴影扩散范围 整数（建议范围(0- Size )/4） 3
	BarLen    int     `json:"barlen"`    // 徽章的文字条的长度 默认由文字长度决定 100
	Fontsize  int     `json:"fontsize"`  // 建议范围12~48 默认15 30
	BarRadius int     `json:"barradius"` // 文字条圆角大小 整数（建议范围0~ Size /2）, 默认为 Size 的一半 15
}

// BadgeWay 徽章渐变方向
/*
ENUM(
top // 从上到下
bottom // 从下到上
left // 从左到右
right // 从右到左
top-left // 从左上到右下
top-right // 从右上到左下
bottom-left // 从左下到右上
bottom-right // 从右下到左上
)
*/
type BadgeWay string
