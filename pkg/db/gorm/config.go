package gorm

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"sync"
)

// LogLevel ...
type LogLevel int

const (
	Silent LogLevel = iota + 1
	Error
	Warn
	Info
)
const DefaultClient = "default"

var (
	clientLock         sync.RWMutex
	clientMap          = make(map[string]*gorm.DB)
	ErrDBUninitialized = errors.New("database uninitialized or initial failed")
)

type Config struct {
	Database string   `env:"MYSQL_DATABASE" default:""`
	Host     string   `env:"MYSQL_HOST" default:"127.0.0.1"`
	Port     int      `env:"MYSQL_PORT" default:"3306"`
	User     string   `env:"MYSQL_USER" default:"root"`
	Password string   `env:"MYSQL_PASSWORD" default:""`
	Retry    int      `env:"MYSQL_RETRY" default:"3"`
	MaxOpen  int      `env:"MYSQL_MAX_OPEN" default:"100"`
	MaxIdle  int      `env:"MYSQL_MAX_IDLE" default:"10"`
	Lifetime int      `env:"MYSQL_CONN_LIVETIME" default:"60"` //单位：分钟
	Timeout  int      `env:"MYSQL_CONN_TIMEOUT" default:"0"`   //单位：秒
	Tracing  bool     `env:"MYSQL_TRACING" default:"false"`
	LogLevel LogLevel `env:"MYSQL_LOG_LEVEL" default:"3"` //日志级别，默认为 WARN
}

func (cfg *Config) DSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=PRC",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database,
	)
	if cfg.Timeout > 0 {
		dsn = fmt.Sprintf("%s&timeout=%ds", dsn, cfg.Timeout)
	}

	return dsn
}
