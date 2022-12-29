package config

type Config struct {
	HttpServerConfig HttpConfig
}

func GetConfig() *Config {
	return &Config{httpConfig}
}
