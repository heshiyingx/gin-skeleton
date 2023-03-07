package config

var config = &Config{HttpServerConfig: httpConfig}

// Config ...
type Config struct {
	HttpServerConfig HttpConfig `mapstructure:"http"`
}

func GetConfig() *Config {
	return config
}
