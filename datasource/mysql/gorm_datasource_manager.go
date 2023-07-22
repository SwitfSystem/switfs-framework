package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type TDataSourceConfig struct {
	DbName       string
	Username     string
	Password     string
	Host         string
	Port         string
	Driver       string
	IdlePoolSize int
	MaxPoolSize  int
	MaxLifeTime  int64
	SqlDebug     int8
}

func New(cfg TDataSourceConfig) *gorm.DB {
	return openConn(cfg)
}

func openConn(cfg TDataSourceConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true&loc=Local&charset=utf8mb4&collation=utf8mb4_unicode_ci&readTimeout=10s&writeTimeout=10s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	if cfg.SqlDebug == 1 {
		db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Errorf("error : %v ", err.Error())
		return nil
	}
	sqlDB.SetMaxIdleConns(cfg.IdlePoolSize)
	sqlDB.SetMaxOpenConns(cfg.MaxPoolSize)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.MaxLifeTime) * time.Second)
	// 设置字符编码
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")
	return db
}
