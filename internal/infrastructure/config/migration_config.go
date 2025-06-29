package config

import (
	"fmt"
	"os"
	"reflect"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type MigrationConfig struct {
	AppEnv AppEnvType `env:"APP_ENV,required"`

	PostgresqlUsername string `env:"POSTGRESQL_USERNAME,required"`
	PostgresqlPassword string `env:"POSTGRESQL_PASSWORD,required"`
	PostgresqlHost     string `env:"POSTGRESQL_HOST,required"`
	PostgresqlDatabase string `env:"POSTGRESQL_DATABASE,required"`
	PostgresqlSchema   string `env:"POSTGRESQL_SCHEMA,required"`
	PostgresqlSSL      bool   `env:"POSTGRESQL_SSL" envDefault:"true"`
}

func LoadMigrationConfig() (*APIConfig, error) {
	dotenvPath := os.Getenv("DOTENV_PATH")
	if dotenvPath != "" {
		if err := godotenv.Load(dotenvPath); err != nil {
			return nil, fmt.Errorf("error loading .env file from %s: %w", dotenvPath, err)
		}
	} else {
		if err := godotenv.Load(); err != nil {
			return nil, fmt.Errorf("error loading .env file: %w", err)
		}
	}

	cfg := APIConfig{}
	if err := env.ParseWithOptions(&cfg, env.Options{
		FuncMap: map[reflect.Type]env.ParserFunc{
			reflect.TypeOf(AppEnvType("")): func(v string) (interface{}, error) {
				appEnv := AppEnvType(v)
				if !appEnv.IsValid() {
					return nil, fmt.Errorf("invalid app environment: %s", v)
				}

				return appEnv, nil
			},
		},
	}); err != nil {
		return nil, fmt.Errorf("error parsing environment variables: %w", err)
	}

	return &cfg, nil
}
