package config

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/logger"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/migrations"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/postgres"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/probes"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/redis"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/tracing"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// ConfigPath is a path to config file
var configPath string

// InitFlags init flags
func init() {
	flag.StringVar(&configPath, "config", "", "User microservice config path")
}

type (
	Config struct {
		ServiceName string            `yaml:"serviceName" env:"SERVICE_NAME" env-default:"go-service"`
		GRPC        GRPC              `yaml:"grpc"`
		Logger      logger.Config     `yaml:"logger"`
		Postgres    postgres.Config   `yaml:"postgres"`
		Redis       redis.Config      `yaml:"redis"`
		Jaeger      tracing.Config    `yaml:"jaeger"`
		Probes      probes.Config     `yaml:"probes"`
		Migrations  migrations.Config `yaml:"migrations"`
		Timeouts    Timeouts          `yaml:"timeouts"`
	}

	GRPC struct {
		Port        string `yaml:"port" env:"GRPC_PORT" env-default:"5001"`
		Development bool   `yaml:"development" env:"GRPC_DEVELOPMENT" env-default:"false"`
	}

	Timeouts struct {
		PostgresInitMilliseconds int  `yaml:"postgresInitMilliseconds" env:"POSTGRES_INIT_MILLISECONDS" env-default:"1000"`
		PostgresInitRetryCount   uint `yaml:"postgresInitRetryCount" env:"POSTGRES_INIT_RETRY_COUNT" env-default:"10"`
	}
)

// InitConfig init config
func InitConfig() (*Config, error) {
	cfg := &Config{}

	if configPath == "" {
		dir, err := os.Getwd()
		if err != nil {
			return nil, errors.Wrap(err, "os.Getwd")
		}
		configPath = fmt.Sprintf("%s/config/config.yaml", dir)
	}

	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "cleanenv.ReadConfig")
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
