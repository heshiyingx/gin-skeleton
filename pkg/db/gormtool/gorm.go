package gormtool

import (
	"fmt"
	"gitlab.myshuju.top/base/ginskeleton/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// Client 获取默认数据库连接
func Client() (*gorm.DB, error) {
	return ClientByName(DefaultClient)
}

func ClientByName(name string) (*gorm.DB, error) {
	clientLock.RLock()
	defer clientLock.RUnlock()
	if db, ok := clientMap[name]; ok {
		return db, nil
	}
	return nil, ErrDBUninitialized
}
func InitClient(c Config, opts ...gorm.Option) error {
	if _, err := ClientByName(DefaultClient); err == nil {
		return nil
	}
	return InitClientByName(DefaultClient, c)
}
func InitClientByName(name string, c Config, opts ...gorm.Option) error {
	if _, err := ClientByName(name); err == nil {
		return nil
	}
	clientLock.Lock()
	defer clientLock.Unlock()
	db, err := createMySQLConnection(&c, opts...)
	if err != nil {
		return err
	}
	clientMap[name] = db

	return err
}
func createMySQLConnection(cfg *Config, opts ...gorm.Option) (*gorm.DB, error) {

	utils.FillDefault(cfg)
	//level := gormlog.LogLevel(cfg.LogLevel)
	//if ormCfg.Logger == nil {
	//	if cfg.LogLevel > 0 {
	//
	//		ormCfg.Logger = DefaultLogger.LogMode(level)
	//	} else {
	//		ormCfg.Logger = DefaultLogger.LogMode(gormlog.Warn)
	//	}
	//}

	var (
		err   error
		times = 10
		db    *gorm.DB
	)

	if cfg.Retry > 0 {
		times = cfg.Retry
	}

	for {
		times--
		if times < 0 {
			return nil, fmt.Errorf("connect to mysql %s failed", cfg.Host)
		}
		if db, err = gorm.Open(mysql.Open(cfg.DSN()), opts...); err != nil {
			time.Sleep(time.Second * 2)
			continue
		}
		break
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpen)
	if cfg.MaxIdle > 0 {
		sqlDB.SetMaxIdleConns(cfg.MaxIdle)
	}
	if cfg.Lifetime < 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(cfg.Lifetime) * time.Minute)
	}
	//if cfg.Tracing {
	//	db.Use(NewTracing())
	//}
	return db, nil
}
