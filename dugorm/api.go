package dugorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Config struct {
	//数据库地址   ip:port
	Address  string `json:"address" yaml:"address"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Database string `json:"database" yaml:"database"`
	//指定字符集 默认utf8mb4
	Charset string `json:"charset" yaml:"charset"`
	//数据库的最大连接数 默认不限制
	MaxOpenConns int `json:"max_open_conns" yaml:"max_open_conns"`
	//最大空闲连接数 默认2
	MaxIdleConns int `json:"max_idle_conns" yaml:"max_idle_conns"`
	//是否跳过事务，默认开启事务
	SkipDefaultTransaction bool `json:"skip_default_transaction" yaml:"skip_default_transaction"`
	//创建一个 prepared statement 并将其缓存 默认关闭
	PrepareStmt bool `json:"prepare_stmt" yaml:"prepare_stmt"`
}

func NewGormDb(cfg Config) *gorm.DB {
	if cfg.Address == "" || cfg.Username == "" || cfg.Password == "" {
		panic("mysql info is error")
	}
	if cfg.Charset == "" {
		cfg.Charset = "utf8mb4"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.Address, cfg.Database, cfg.Charset)
	gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: cfg.SkipDefaultTransaction,
		PrepareStmt:            cfg.PrepareStmt,
		NowFunc: func() time.Time {
			tmp := time.Now().Local().Format("2006-01-02 15:04:05")
			now, _ := time.ParseInLocation("2006-01-02 15:04:05", tmp, time.Local)
			return now
		},
	})
	if err != nil {
		panic(err)
	}
	if cfg.MaxOpenConns != 0 || cfg.MaxIdleConns != 0 {
		db, err := gormDb.DB()
		if err != nil {
			panic(err)
		}
		db.SetMaxOpenConns(cfg.MaxOpenConns)
		db.SetMaxIdleConns(cfg.MaxIdleConns)
	}
	return gormDb
}
