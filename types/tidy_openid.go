package types

type GetUserInfoByIdData struct {
	UserAvatarURL string `json:"userAvatarURL"`
	UserNickname  string `json:"userNickname"`
	UserName      string `json:"userName"`
}
