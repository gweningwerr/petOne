package api

import "github.com/gweningwarr/petOne/storage"

type Config struct {
	// Port
	BindAddr string `toml:"bind_addr"`
	// logger
	LoggerLevel string `toml:"logger_level"`
	// storage
	Storage *storage.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8989",
		LoggerLevel: "debug",
		Storage:     storage.NewConfig(),
	}
}
