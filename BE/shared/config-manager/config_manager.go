// package config_manager

// import (
// 	"context"
// 	"os"
// 	"shared/common/utils"
// )

// type ConfigLoader interface {
// 	GetConfig(ctx context.Context) ([]byte, error)
// }

// type ConfigManager[T any] struct {
// 	client ConfigLoader
// }

// type ConfigManagerSettings struct {
// 	Namespace   string
// 	ServiceName string
// 	Environment string
// }

// func NewConfigManager[T any](settings ConfigManagerSettings) *ConfigManager[T] {
// 	return &ConfigManager[T]{
// 		client: NewAWSAppConfigManager(settings.Namespace, settings.ServiceName, settings.Environment),
// 	}
// }

// func GetDefaultSettings() ConfigManagerSettings {
// 	return ConfigManagerSettings{
// 		Namespace:   os.Getenv("app_project"),
// 		ServiceName: os.Getenv("app_name"),
// 		Environment: os.Getenv("app_env"),
// 	}
// }

// func (c *ConfigManager[T]) GetConfig(ctx context.Context) (*T, error) {
// 	bytes, err := c.client.GetConfig(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	config, err := utils.Decode[T](bytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &config, nil
// }

package config_manager

import (
	"context"
	"fmt"
	"os"
	"shared/common/utils"
)

type ConfigManager[T any] struct {
	path string
}

func NewConfigManager[T any](localPath string) *ConfigManager[T] {
	return &ConfigManager[T]{path: localPath}
}

func (c *ConfigManager[T]) GetConfig(_ context.Context) (*T, error) {
	data, err := os.ReadFile(c.path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	cfg, err := utils.Decode[T](data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode config: %w", err)
	}

	return &cfg, nil
}
