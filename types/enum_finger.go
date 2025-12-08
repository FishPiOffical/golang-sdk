//go:generate go-enum --marshal --names --values --ptr --mustparse
package types

// ItemType 物品类型
/*
ENUM(
checkin1day
patchCheckinCard
patchStart
sysCheckinRemain
)
*/
type ItemType string
