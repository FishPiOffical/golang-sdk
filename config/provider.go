package config

// Provider 配置提供者接口
type Provider interface {
	// Get 获取配置
	Get() *Config
	// Update 更新配置
	Update(config *Config) error
}
