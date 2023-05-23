package config

import (
	"github.com/caarlos0/env/v8"
)

type ConfigList struct {
	Port int `env:"PORT" envDefault:"80"`
}

func LoadConfig() (*ConfigList, error) {
	cfg := &ConfigList{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
