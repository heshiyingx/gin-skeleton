package config

var config = &Config{}

// Config ...
type Config struct {
	HttpServerConfig HttpConfig `mapstructure:"http"`
}

func GetConfig() *Config {
	return config
}
