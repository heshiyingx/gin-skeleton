package config

import httpConfig "gitlab.myshuju.top/base/ginskeleton/config"

var config = &Config{}

// Config ...
type Config struct {
	HttpServerConfig httpConfig.HttpConfig `mapstructure:"http"`
}

func GetConfig() *Config {
	return config
}
