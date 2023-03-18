package gormtool

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
	ErrDBNotSet        = errors.New("database is empty")
)

type Config struct {
	Database string   `mapstructure:"MYSQL_DATABASE" default:""`
	Host     string   `mapstructure:"MYSQL_HOST" default:"127.0.0.1"`
	Port     int      `mapstructure:"MYSQL_PORT" default:"3306"`
	User     string   `mapstructure:"MYSQL_USER" default:"root"`
	Password string   `mapstructure:"MYSQL_PASSWORD" default:""`
	Retry    int      `mapstructure:"MYSQL_RETRY" default:"3"`
	MaxOpen  int      `mapstructure:"MYSQL_MAX_OPEN" default:"100"`
	MaxIdle  int      `mapstructure:"MYSQL_MAX_IDLE" default:"10"`
	Lifetime int      `mapstructure:"MYSQL_CONN_LIVETIME" default:"60"` //单位：分钟
	Timeout  int      `mapstructure:"MYSQL_CONN_TIMEOUT" default:"0"`   //单位：秒
	Tracing  bool     `mapstructure:"MYSQL_TRACING" default:"false"`
	LogLevel LogLevel `mapstructure:"MYSQL_LOG_LEVEL" default:"3"` //日志级别，默认为 WARN
}

func (cfg *Config) DSN() string {
	if err := cfg.Check(); err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=PRC",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database,
	)
	if cfg.Timeout > 0 {
		dsn = fmt.Sprintf("%s&timeout=%ds", dsn, cfg.Timeout)
	}

	return dsn
}
func (cfg *Config) Check() error {
	if cfg.Database == "" {
		return ErrDBNotSet
	}
	return nil
}
