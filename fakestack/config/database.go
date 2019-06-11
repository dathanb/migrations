package config

import "fmt"

type DbConfig interface {
	ConnectionString() string
	DriverName() string
}

type PostgresDbConfig struct {
	Username string `default:"postgres"`
	Password string `default:"password"`
	Port     int    `default:"5432"`
	Host     string `default:"localhost"`
	DbName   string `default:"fakestack"`
	SslMode  string `split_words:"true"`
}

func (config *PostgresDbConfig) DriverName() string {
	return "postgres"
}

func (config *PostgresDbConfig) ConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.Username, config.Password,
		config.Host, config.Port, config.DbName)
}

