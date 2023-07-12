package duhbase

import (
	"fmt"
	"github.com/silenceper/pool"
	"time"
)

type HbaseConfig struct {
	//hbase addr
	Address  string `json:"address" yaml:"address" toml:"address"`
	User     string `json:"user" yaml:"user" toml:"user"`
	Password string `json:"password" yaml:"password" toml:"password"`
	//连接池中拥有的最小连接数 默认200
	InitialCap int `json:"initial_cap" yaml:"initial_cap" toml:"initial_cap"`
	//最大并发存活连接数 默认 600
	MaxCap int `json:"max_cap" yaml:"max_cap" toml:"max_cap"`
	//最大空闲连接 默认 200
	MaxIdle int `json:"max_idle" yaml:"max_idle" toml:"max_idle"`
	//连接最大空闲时间，超过该事件则将失效 默认60秒
	IdleTimeout int64 `json:"idle_timeout" yaml:"idle_timeout" toml:"idle_timeout"`
}

func (hc *HbaseConfig) DefaultConfig() {
	if hc.InitialCap == 0 {
		hc.InitialCap = 200
	}
	if hc.MaxCap == 0 {
		hc.MaxCap = 600
	}
	if hc.MaxIdle == 0 {
		hc.MaxIdle = 200
	}
	if hc.IdleTimeout == 0 {
		hc.IdleTimeout = 60
	}
}

func NewHbasePool(hbaseConfig *HbaseConfig) pool.Pool {
	cfg = hbaseConfig
	cfg.DefaultConfig()
	var err error
	poolConfig := &pool.Config{
		InitialCap:  cfg.InitialCap,
		MaxCap:      cfg.MaxCap,
		MaxIdle:     cfg.MaxIdle,
		Factory:     newClient,
		Close:       closeClient,
		IdleTimeout: time.Duration(cfg.IdleTimeout) * time.Second,
	}
	hbaseClientPool, err := pool.NewChannelPool(poolConfig)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Printf("successful link creation :%d \n", hbaseClientPool.Len())
	return hbaseClientPool
}

func ReleaseHbasePool(p pool.Pool) {
	p.Release()
}
