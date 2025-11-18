package config

import (
	"fmt"
	"sync"
)

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
