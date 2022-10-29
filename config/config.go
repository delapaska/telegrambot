package config

import (
	"github.com/caarlos0/env"
	"log"
	"time"
)

const (
	CtxTimeout = time.Second * 30
)

// Переменные окружения

type CommonEnvConfigs struct {
	DevEnv   bool   `env:"DEV_ENV" envDefault:"false"`
	LogLevel string `json:"LOG_LEVEL" env:"LOG_LEVEL" envDefault:"info"`
}

func GetCommonEnvConfigs() CommonEnvConfigs {
	envConfig := CommonEnvConfigs{}
	if err := env.Parse(&envConfig); err != nil {
		log.Panicf("Error parse env config: %s", err)
		return envConfig
	}

	return envConfig
}
