package duredis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"strings"
	"time"
)

// NewRedisSingleClient 创建单节点模式客户端对象
func NewRedisSingleClient(cfg RedisConfig) (*redis.Client, error) {
	return newRedisSingleApi(cfg)
}

// NewRedisFailoverClient 创建哨兵模式客户端
func NewRedisFailoverClient(cfg RedisFailoverConfig) (*redis.Client, error) {
	return newRedisFailoverApi(cfg.MasterName, cfg.Hosts, cfg.Password, cfg.Db, cfg.PoolSize)
}

// NewRedisClusterClient 创建集群模式客户端
func NewRedisClusterClient(cfg RedisConfig) (*redis.ClusterClient, error) {
	return newRedisClusterApi(cfg)
}

func CloseRedisSingle(r *redis.Client) {
	if r != nil {
		r.Close()
	}
}
func CloseRedisFailover(r *redis.Client) {
	if r != nil {
		r.Close()
	}
}

func CloseRedisCluster(r *redis.ClusterClient) {
	if r != nil {
		r.Close()
	}
}

// 单节点模式
func newRedisSingleApi(cfg RedisConfig) (*redis.Client, error) {
	cfg.DefaultConfig()
	r := redis.NewClient(&redis.Options{
		Addr:         cfg.Host,
		Password:     cfg.Password,
		MaxRetries:   cfg.MaxRetries,
		DialTimeout:  time.Second * time.Duration(cfg.DialTimeout),
		ReadTimeout:  time.Second * time.Duration(cfg.ReadTimeout),
		WriteTimeout: time.Second * time.Duration(cfg.WriteTimeout),
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
		MaxConnAge:   time.Minute * time.Duration(cfg.MaxConnAge),
		PoolTimeout:  time.Second * time.Duration(cfg.PoolTimeout),
		IdleTimeout:  time.Second * time.Duration(cfg.IdleTimeout),
		DB:           cfg.Db,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	_, err := r.Ping(ctx).Result()
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	return r, err
}

// 哨兵模式
func newRedisFailoverApi(masterName string, addrs []string, pw string, db, poolSize int) (*redis.Client, error) {
	if poolSize == 0 {
		poolSize = 100
	}
	r := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:       masterName,
		SentinelAddrs:    addrs,
		SentinelPassword: pw,
		MaxRetries:       3,
		DialTimeout:      time.Second * 30,
		ReadTimeout:      time.Second * 5,
		WriteTimeout:     time.Second * 5,
		PoolSize:         poolSize,
		MinIdleConns:     10,
		MaxConnAge:       time.Minute * 1,
		PoolTimeout:      time.Second * 30,
		IdleTimeout:      time.Second * 30,
		DB:               db,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.Ping(ctx).Result()
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	return r, err
}

// 集群模式
func newRedisClusterApi(cfg RedisConfig) (*redis.ClusterClient, error) {
	cfg.DefaultConfig()
	r := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        strings.Split(cfg.Host, ","),
		Password:     cfg.Password,
		MaxRetries:   cfg.MaxRetries,
		DialTimeout:  time.Second * time.Duration(cfg.DialTimeout),
		ReadTimeout:  time.Second * time.Duration(cfg.ReadTimeout),
		WriteTimeout: time.Second * time.Duration(cfg.WriteTimeout),
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
		MaxConnAge:   time.Minute * time.Duration(cfg.MaxConnAge),
		PoolTimeout:  time.Second * time.Duration(cfg.PoolTimeout),
		IdleTimeout:  time.Second * time.Duration(cfg.IdleTimeout),
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := r.Ping(ctx).Result()
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	return r, err
}
