package sdk

// ConfigProvider 配置提供者接口
type ConfigProvider interface {
	// Get 获取配置
	Get() *Config
	// Update 更新配置
	Update(config *Config) error
}

// MemoryConfigProvider 内存配置提供者
type MemoryConfigProvider struct {
	config *Config
}

// NewMemoryConfigProvider 创建内存配置提供者
func NewMemoryConfigProvider(config *Config) *MemoryConfigProvider {
	return &MemoryConfigProvider{
		config: config,
	}
}

// Get 获取配置
func (m *MemoryConfigProvider) Get() *Config {
	return m.config
}

// Update 更新配置
func (m *MemoryConfigProvider) Update(config *Config) error {
	m.config = config
	return nil
}
