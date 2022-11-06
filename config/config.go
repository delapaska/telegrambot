package config

import (
	"log"
	"time"

	"github.com/caarlos0/env"
)

const (
	CtxTimeout = time.Second * 30
)

// Переменные окружения

type CommonEnvConfigs struct {
	DevEnv   bool   `env:"DEV_ENV" envDefault:"false"`
	LogLevel string `json:"LOG_LEVEL" env:"LOG_LEVEL" envDefault:"info"`

	PostgreSQL struct {
		Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
		Port     int    `env:"POSTGRES_PORT" envDefault:"5432"`
		User     string `env:"POSTGRES_USER" envDefault:"postgres"`
		Password string `env:"POSTGRES_PASSWORD" envDefault:"postgres"`
		DBName   string `env:"POSTGRES_DB" envDefault:"postgres"`
	}

	Binance struct {
		URL string `env:"BINANCE_URL" envDefault:"NULL"`
		//URL
		//TOKEN
	}
	TgBot struct {
		TOKEN string `env:"TOKEN" envDefault:"NUll"`
	}
}

func GetCommonEnvConfigs() CommonEnvConfigs {
	envConfig := CommonEnvConfigs{}
	if err := env.Parse(&envConfig); err != nil {
		log.Panicf("Error parse env config: %s", err)
		return envConfig
	}

	return envConfig
}
