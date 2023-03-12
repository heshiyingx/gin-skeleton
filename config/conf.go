package config

var config = &Config{HttpServerConfig: DefaultHttpConfig}

// Config ...
type Config struct {
	HttpServerConfig HttpConfig `mapstructure:"http"`
}

func GetConfig() *Config {
	return config
}
