//go:generate go-enum --marshal --names --values --ptr --mustparse
package types

// UserAppRole 用户应用角色
/*
ENUM(
hack=0 // 黑客
painter=1 // 画家
)
*/
type UserAppRole int
