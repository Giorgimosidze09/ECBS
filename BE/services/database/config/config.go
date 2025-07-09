package db_config

import "sync"

type DBConfig struct {
	NatsURL string `json:"NATS_URL"`
	DBUrl   string `json:"DATABASE_URL"`
}

var (
	cfg  *DBConfig
	once sync.Once
)

func Set(config *DBConfig) {
	once.Do(func() {
		cfg = config
	})
}

func Get() *DBConfig {
	if cfg == nil {
		panic("config not initialized. Call config.Set() first.")
	}
	return cfg
}
