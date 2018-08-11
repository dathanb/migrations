package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/udacity/migration-demo/error"
	"fmt"
)

type DbConfig interface {
	ConnectionString() string
	DriverName() string
}

type PostgresDbConfig struct {
	Username string `default:"postgres"`
	Password string `default:"password"`
	Port     int    `default:"5432"`
	Host     string `default:"localhost"`
	DbName   string `default:"migration-demo"`
	SslMode  string `split_words:"true"`
}

func (config *PostgresDbConfig) DriverName() string {
	return "postgres"
}

func (config *PostgresDbConfig) ConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.Username, config.Password,
		config.Host, config.Port, config.DbName)
}

/*
Config encapsulates configuration from the environment
 */
type Config struct {
	Db     DbConfig
}

func LoadConfig() (*Config, error.Error) {
	var dbConfig PostgresDbConfig
	var err error.Error

	err = envconfig.Process("MIGRATION_DEMO_DB", &dbConfig)
	if err != nil {
		return nil, err
	}

	return &Config{
		&dbConfig,
	}, nil
}

