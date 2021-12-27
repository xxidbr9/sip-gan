package config

import (
	goConfig "github.com/golobby/config/v3"
	"github.com/golobby/config/v3/pkg/feeder"
)

type BaseConfig struct {
	EnvMode string `env:"ENV_MODE"`
}

func NewBaseConfig() (*BaseConfig, error) {
	envConfig := &BaseConfig{}

	envFeeder := feeder.DotEnv{Path: ".env"}
	enhanceConfig := goConfig.New()

	enhanceConfig.AddFeeder(envFeeder)
	enhanceConfig.AddStruct(envConfig)

	if err := enhanceConfig.Feed(); err != nil {
		return &BaseConfig{}, err
	}
	
	return envConfig, nil
}
