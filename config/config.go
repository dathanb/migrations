package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/udacity/migration-demo/error"
)

/*
Config encapsulates configuration from the environment
 */
type Config struct {
	Db     DbConfig
	Server ServerConfig
}

func LoadConfig() (*Config, error.Error) {
	var dbConfig PostgresDbConfig
	var serverConfig SimpleServerConfig
	var err error.Error

	err = envconfig.Process("MIGRATION_DEMO_DB", &dbConfig)
	if err != nil {
		return nil, err
	}

	err = envconfig.Process("MIGRATION_DEMO_SERVER", &serverConfig)

	return &Config{
		&dbConfig,
		&serverConfig,
	}, nil
}

