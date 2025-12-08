//go:generate go-enum --marshal --names --values --ptr --mustparse
package types

// ItemType 物品类型
/*
ENUM(
checkin1day // 单日免签卡
checkin2days // 两天免签卡
patchCheckinCard // 补签卡
metalTicket // 摸鱼派一周年纪念勋章领取券
nameCard // 改名卡
patchStart // 补签日期
sysCheckinRemain // 免签卡生效中，剩余X天
)
*/
type ItemType string
