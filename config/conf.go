package config

// Config ...
type Config struct {
	HttpServerConfig HttpConfig
}

func GetConfig() *Config {
	return &Config{httpConfig}
}
