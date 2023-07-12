package durpcx

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/rpcxio/libkv/store"
	consulClient "github.com/rpcxio/rpcx-consul/client"
	redisClient "github.com/rpcxio/rpcx-redis/client"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"log"
	"time"
)

type RpcClient struct {
	xClient client.XClient
	timeout time.Duration
}

// NewDNSClient : 创建并初始化一个client，包含了必要的心跳检测及其他配置以增强后端服务下线感知
func NewDNSClient(cfg Config) *RpcClient {
	j, _ := jsoniter.Marshal(&cfg)
	log.Println("RPC config:", string(j))
	var e error
	var d client.ServiceDiscovery
	d, e = client.NewDNSDiscovery(cfg.K8sService, "tcp", cfg.RpcPort, time.Minute)
	if e != nil {
		log.Println(e)
	}
	if cfg.TimeoutMillseconds == 0 {
		cfg.TimeoutMillseconds = 1000
	}
	c := RpcClient{timeout: time.Duration(cfg.TimeoutMillseconds) * time.Millisecond}
	option := client.DefaultOption
	option.Heartbeat = true
	option.HeartbeatInterval = 1 * time.Second
	option.MaxWaitForHeartbeat = 2 * time.Second
	option.IdleTimeout = 3 * time.Second
	if cfg.Group != "" {
		option.Group = cfg.Group
	}
	c.xClient = client.NewXClient(cfg.ServicePath, client.Failover, client.RandomSelect, d, option)
	return &c
}

// NewClient : 创建并初始化一个client，包含了必要的心跳检测及其他配置以增强后端服务下线感知
func NewClient(cfg Config) *RpcClient {
	j, _ := jsoniter.Marshal(&cfg)
	log.Println("RPC config:", string(j))
	var e error
	var d client.ServiceDiscovery
	if cfg.RedisServer != "" {
		d, e = redisClient.NewRedisDiscovery("cloud", cfg.ServicePath, []string{cfg.RedisServer}, &store.Config{Password: cfg.RedisPassword, Bucket: cfg.RedisDb})
	} else {
		d, e = consulClient.NewConsulDiscovery("cloud", cfg.ServicePath, []string{cfg.ConsulServer}, nil)
	}
	if e != nil {
		log.Println(e)
	}
	if cfg.TimeoutMillseconds == 0 {
		cfg.TimeoutMillseconds = 1000
	}
	c := RpcClient{timeout: time.Duration(cfg.TimeoutMillseconds) * time.Millisecond}
	option := client.DefaultOption
	option.Heartbeat = true
	option.HeartbeatInterval = 1 * time.Second
	option.MaxWaitForHeartbeat = 2 * time.Second
	option.IdleTimeout = 3 * time.Second
	if cfg.Group != "" {
		option.Group = cfg.Group
	}
	if cfg.SerializeType != 0 {
		option.SerializeType = protocol.SerializeType(cfg.SerializeType)
	}
	c.xClient = client.NewXClient(cfg.ServicePath, client.Failover, client.RandomSelect, d, option)
	return &c
}

// Call : 客户端调用统一方法，包含了超时及错误处理
func (p RpcClient) Call(method string, req Req, res Res) {
	ctx, cancelFn := context.WithTimeout(context.Background(), p.timeout)
	err := p.xClient.Call(ctx, method, req, res)
	if err != nil {
		res.SetError(err)
	}
	cancelFn()
}

// Close : 停止服客户端，graceful shutdown 时须调用次方法
func (p RpcClient) Close() {
	if p.xClient != nil {
		p.xClient.Close()
	}
}
