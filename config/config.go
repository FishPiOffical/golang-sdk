package config

type Config struct {
	BaseUrl   string `json:"base_url,omitempty" yaml:"baseUrl,omitempty"`
	UserAgent string `json:"user_agent,omitempty" yaml:"userAgent,omitempty"`

	ApiKey      string `json:"api_key,omitempty" yaml:"apiKey,omitempty"`
	Username    string `json:"username,omitempty" yaml:"username,omitempty"`
	Password    string `json:"password,omitempty" yaml:"password,omitempty"`
	PasswordMd5 string `json:"password_md5,omitempty" yaml:"passwordMd5,omitempty"`
	MfaCode     string `json:"mfa_code,omitempty" yaml:"mfaCode,omitempty"`
	Totp        string `json:"totp,omitempty" yaml:"totp,omitempty"`

	PointGoldFingerKey    string `json:"point_gold_finger_key,omitempty" yaml:"pointGoldFingerKey,omitempty"`
	LivenessGoldFingerKey string `json:"liveness_gold_finger_key,omitempty" yaml:"livenessGoldFingerKey,omitempty"`
	GameGoldFingerKey     string `json:"game_gold_finger_key,omitempty" yaml:"gameGoldFingerKey,omitempty"`
	QueryGoldFingerKey    string `json:"query_gold_finger_key,omitempty" yaml:"queryGoldFingerKey,omitempty"`
	MetalGoldFingerKey    string `json:"metal_gold_finger_key,omitempty" yaml:"metalGoldFingerKey,omitempty"`
	ItemGoldFingerKey     string `json:"item_gold_finger_key,omitempty" yaml:"itemGoldFingerKey,omitempty"`
	MedalReadFingerKey    string `json:"medal_read_finger_key,omitempty" yaml:"medalReadFingerKey,omitempty"`
	MedalWriteFingerKey   string `json:"medal_write_finger_key,omitempty" yaml:"medalWriteFingerKey,omitempty"`
}
