package api_config

import "sync"

type ApiConfig struct {
	NatsURL string `json:"NATS_URL"`
}

var (
	cfg  *ApiConfig
	once sync.Once
)

func Set(config *ApiConfig) {
	once.Do(func() {
		cfg = config
	})
}

func Get() *ApiConfig {
	if cfg == nil {
		panic("config not initialized. Call config.Set() first.")
	}
	return cfg
}
