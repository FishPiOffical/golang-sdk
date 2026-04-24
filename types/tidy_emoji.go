//go:generate go-enum --names --values --ptr --mustparse
package types

// EmojiGroupType 表情包分组类型
/*
ENUM(
custom = 0 // 通用表情表分组
all = 1 // 默认全部表情表分组
)
*/
type EmojiGroupType int

type EmojiGroup struct {
	Name string         `json:"name"`
	OId  string         `json:"oId"`
	Sort int            `json:"sort"`
	Type EmojiGroupType `json:"type"`
}

type Emoji struct {
	EmojiId string `json:"emojiId"`
	Name    string `json:"name"`
	OId     string `json:"oId"`
	Sort    int    `json:"sort"`
	Url     string `json:"url"`
}
