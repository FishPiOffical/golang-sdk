//go:generate go-enum --marshal --names --values --ptr --mustparse
package types

// ClientType 客户端类型
/*
ENUM(
Web
PC
Mobile
Windows
macOS
Linux
iOS
Android
IDEA
Chrome
Edge
VSCode
Python
Golang
Rust
Harmony
CLI
Bird
IceNet
ElvesOnline
Other
)
*/
type ClientType string
