package config

import (
	"fmt"
)

const (
	DEFAULT_HOST          = "0.0.0.0"
	DEFAULT_PROT          = 8089
	DEFAULT_READ_TIMEOUT  = 50
	DEFAULT_WRITE_TIMEOUT = 50
)

var (
	DefaultHttpConfig = HttpConfig{
		Host:            DEFAULT_HOST,
		Port:            DEFAULT_PROT,
		ReadTimeoutSec:  DEFAULT_READ_TIMEOUT,
		WriteTimeoutSec: DEFAULT_WRITE_TIMEOUT,
	}
)

// HttpConfig http服务配置
type HttpConfig struct {
	Host            string `mapstructure:"host"`
	Port            int64  `mapstructure:"port"`
	ReadTimeoutSec  int64  `mapstructure:"read_timeout_sec"`
	WriteTimeoutSec int64  `mapstructure:"write_timeout_sec"`
}

func init() {

}

// GetAddr 获取地址.
func (c *HttpConfig) GetAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
