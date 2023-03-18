package main

import httpConfig "gitlab.myshuju.top/heshiying/gin-skeleton/config"

var config = &Config{}

// Config ...
type Config struct {
	HttpServerConfig httpConfig.HttpConfig `mapstructure:"http"`
}

func GetConfig() *Config {
	return config
}
