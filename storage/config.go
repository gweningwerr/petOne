package storage

type Config struct {
	DatabaseURI string `toml:"db_uri"`
}

func NewConfig() *Config {
	return &Config{}
}
