package config

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

// FileYamlProvider 文件配置提供者
type FileYamlProvider struct {
	filePath string
	config   *Config
	mu       sync.RWMutex
}

// NewFileYamlProvider 创建文件配置提供者
func NewFileYamlProvider(filePath string) *FileYamlProvider {
	provider := &FileYamlProvider{
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
func (p *FileYamlProvider) Get() *Config {
	p.mu.RLock()
	defer p.mu.RUnlock()

	// 返回配置的副本
	configCopy := *p.config
	return &configCopy
}

// Update 更新配置并保存到文件
func (p *FileYamlProvider) Update(config *Config) error {
	if config == nil {
		return fmt.Errorf("config cannot be nil")
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	p.config = config
	return p.save()
}

// load 从文件加载配置
func (p *FileYamlProvider) load() error {
	data, err := os.ReadFile(p.filePath)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	if err = yaml.Unmarshal(data, p.config); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return nil
}

// save 保存配置到文件
func (p *FileYamlProvider) save() error {
	data, err := yaml.Marshal(p.config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err = os.WriteFile(p.filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
