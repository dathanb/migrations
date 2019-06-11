package config

type ServerConfig interface {
	Port() int
}

type SimpleServerConfig struct {
	ServerPort int `default:"8080" envconfig:"port"`
}

func (cfg *SimpleServerConfig) Port() int {
	return cfg.ServerPort
}
