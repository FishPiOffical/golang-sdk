package sdk

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

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
	mu     sync.RWMutex
}

// NewMemoryConfigProvider 创建内存配置提供者
func NewMemoryConfigProvider(config *Config) *MemoryConfigProvider {
	if config == nil {
		config = &Config{}
	}
	return &MemoryConfigProvider{
		config: config,
	}
}

// Get 获取配置
func (m *MemoryConfigProvider) Get() *Config {
	m.mu.RLock()
	defer m.mu.RUnlock()

	// 返回配置的副本，避免并发修改
	configCopy := *m.config
	return &configCopy
}

// Update 更新配置
func (m *MemoryConfigProvider) Update(config *Config) error {
	if config == nil {
		return fmt.Errorf("config cannot be nil")
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.config = config
	return nil
}

// FileConfigProvider 文件配置提供者
type FileConfigProvider struct {
	filePath string
	config   *Config
	mu       sync.RWMutex
}

// NewFileConfigProvider 创建文件配置提供者
func NewFileConfigProvider(filePath string) *FileConfigProvider {
	provider := &FileConfigProvider{
		filePath: filePath,
		config:   &Config{},
	}

	// 尝试加载配置文件
	if err := provider.load(); err != nil {
		// 如果文件不存在或加载失败，使用默认配置
		provider.config = &Config{}
	}

	return provider
}

// Get 获取配置
func (p *FileConfigProvider) Get() *Config {
	p.mu.RLock()
	defer p.mu.RUnlock()

	// 返回配置的副本
	configCopy := *p.config
	return &configCopy
}

// Update 更新配置并保存到文件
func (p *FileConfigProvider) Update(config *Config) error {
	if config == nil {
		return fmt.Errorf("config cannot be nil")
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	p.config = config
	return p.save()
}

// load 从文件加载配置
func (p *FileConfigProvider) load() error {
	data, err := os.ReadFile(p.filePath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	if err := json.Unmarshal(data, p.config); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return nil
}

// save 保存配置到文件
func (p *FileConfigProvider) save() error {
	data, err := json.MarshalIndent(p.config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(p.filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
