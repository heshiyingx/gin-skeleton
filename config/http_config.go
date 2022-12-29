package config

import (
	"fmt"
	"time"
)

const (
	DEFAULT_HOST          = "0.0.0.0"
	DEFAULT_PROT          = 8089
	DEFAULT_READ_TIMEOUT  = time.Second * 50
	DEFAULT_WRITE_TIMEOUT = time.Second * 50
)

var (
	httpConfig = HttpConfig{
		Host:         DEFAULT_HOST,
		Port:         DEFAULT_PROT,
		ReadTimeout:  DEFAULT_READ_TIMEOUT,
		WriteTimeout: DEFAULT_WRITE_TIMEOUT,
	}
)

// HttpConfig http服务配置
type HttpConfig struct {
	Host         string
	Port         int64
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func init() {

}

// GetAddr 获取地址
func (c *HttpConfig) GetAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
