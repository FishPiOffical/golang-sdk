//go:generate go-enum --marshal --names --values --ptr --mustparse
package types

// UserAppRole 用户应用角色
/*
ENUM(
hack="0" // 黑客
painter="1" // 画家
)
*/
type UserAppRole string

// CanFollow 用户是否可以被关注
/*
ENUM(
hide="hide"
no="no"
yes="yes"
)
*/
type CanFollow string
