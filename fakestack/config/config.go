package config

import (
	"github.com/kelseyhightower/envconfig"
)

/*
Config encapsulates configuration from the environment
 */
type Config struct {
	Db     DbConfig
	Server ServerConfig
}

func LoadConfig() (*Config, error) {
	var dbConfig PostgresDbConfig
	var serverConfig SimpleServerConfig
	var err error

	err = envconfig.Process("FAKESTACK_DB", &dbConfig)
	if err != nil {
		return nil, err
	}

	err = envconfig.Process("FAKESTACK_SERVER", &serverConfig)

	return &Config{
		&dbConfig,
		&serverConfig,
	}, nil
}

