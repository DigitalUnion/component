package dumysql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func Connect(cfg Config) (*sqlx.DB, error) {
	if cfg.Charset == "" {
		cfg.Charset = "utf8mb4"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.Address, cfg.Database, cfg.Charset)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if cfg.MaxOpenConns <= 0 {
		cfg.MaxOpenConns = 10
	}
	if cfg.MaxIdleConns <= 0 {
		cfg.MaxIdleConns = 10
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return db, nil
}
