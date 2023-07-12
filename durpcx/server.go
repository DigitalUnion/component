package durpcx

import (
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/rcrowley/go-metrics"
	"github.com/rpcxio/libkv/store"
	consulServerPlugin "github.com/rpcxio/rpcx-consul/serverplugin"
	redisServerPlugin "github.com/rpcxio/rpcx-redis/serverplugin"
	"github.com/smallnest/rpcx/server"
	"log"
	"net"
	"os"
	"runtime"
	"strings"
	"time"
)

var localIp string
var osType string

func init() {
	var err error
	osType = runtime.GOOS
	if osType == "windows" {
		localIp, err = getLocalIpWindows()
	} else {
		localIp, err = getLocalIp()
	}
	if err != nil {
		panic("get local ip error")
	}
}

type RpcServer struct {
	cfg          Config
	Rcvr         interface{}
	srv          *server.Server
	consulPlugin *consulServerPlugin.ConsulRegisterPlugin
	redisPlugin  *redisServerPlugin.RedisRegisterPlugin
}

// NewDNSServer : 创建并初始化一个基于DNS的server，兼容 k8s
func NewDNSServer(cfg Config, rcvr interface{}) (*RpcServer, error) {
	j, _ := jsoniter.Marshal(&cfg)
	log.Println("RPC config:", string(j))
	s := &RpcServer{cfg: cfg}
	var options []server.OptionFn
	if cfg.ReadTimeout > 0 {
		options = append(options, server.WithReadTimeout(time.Duration(cfg.ReadTimeout)*time.Millisecond))
	}
	if cfg.WriteTimeout > 0 {
		options = append(options, server.WithWriteTimeout(time.Duration(cfg.WriteTimeout)*time.Millisecond))
	}
	s.srv = server.NewServer(options...)
	addr := fmt.Sprintf("%s:%d", localIp, cfg.RpcPort)
	if cfg.Group != "" {
		if cfg.Metadata == "" {
			cfg.Metadata = "group=" + cfg.Group
		} else {
			cfg.Metadata += "&group=" + cfg.Group
		}
	}
	err := s.srv.RegisterName(cfg.ServicePath, rcvr, cfg.Metadata)
	if err != nil {
		return nil, err
	}
	go s.srv.Serve("tcp", addr)
	return s, nil
}

// NewServer : 创建并初始化一个server
func NewServer(cfg Config, rcvr interface{}) (*RpcServer, error) {
	j, _ := jsoniter.Marshal(&cfg)
	log.Println("RPC config:", string(j))
	s := &RpcServer{cfg: cfg}
	var options []server.OptionFn
	if cfg.ReadTimeout > 0 {
		options = append(options, server.WithReadTimeout(time.Duration(cfg.ReadTimeout)*time.Millisecond))
	}
	if cfg.WriteTimeout > 0 {
		options = append(options, server.WithWriteTimeout(time.Duration(cfg.WriteTimeout)*time.Millisecond))
	}
	s.srv = server.NewServer(options...)
	addr := fmt.Sprintf("%s:%d", localIp, cfg.RpcPort)
	if len(s.cfg.ConsulServer) != 0 {
		s.addConsulRegistryPlugin(addr)
	}
	if len(s.cfg.RedisServer) != 0 {
		s.addRedisRegistryPlugin(addr)
	}
	if cfg.Group != "" {
		if cfg.Metadata == "" {
			cfg.Metadata = "group=" + cfg.Group
		} else {
			cfg.Metadata += "&group=" + cfg.Group
		}
	}
	s.srv.RegisterOnShutdown(Unregister)
	err := s.srv.RegisterName(cfg.ServicePath, rcvr, cfg.Metadata)
	if err != nil {
		return nil, err
	}
	go s.srv.Serve("tcp", addr)
	go s.startStateChecker()
	return s, nil
}

func (p *RpcServer) addConsulRegistryPlugin(addr string) {
	p.consulPlugin = &consulServerPlugin.ConsulRegisterPlugin{
		ServiceAddress: "tcp@" + addr,
		ConsulServers:  []string{p.cfg.ConsulServer},
		BasePath:       "cloud",
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: 10 * time.Second,
	}
	err := p.consulPlugin.Start()
	if err != nil {
		log.Fatal(err)
	}
	p.srv.Plugins.Add(p.consulPlugin)
}
func (p *RpcServer) addRedisRegistryPlugin(addr string) {
	p.redisPlugin = &redisServerPlugin.RedisRegisterPlugin{
		ServiceAddress: "tcp@" + addr,
		RedisServers:   []string{p.cfg.RedisServer},
		Options:        &store.Config{Password: p.cfg.RedisPassword, Bucket: p.cfg.RedisDb},
		BasePath:       "cloud",
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: 10 * time.Second,
	}
	err := p.redisPlugin.Start()
	if err != nil {
		log.Fatal(err)
	}
	p.srv.Plugins.Add(p.redisPlugin)
}

func Unregister(s *server.Server) {
	log.Println("Unregister Server")
	if s == nil {
		return
	}
	s.UnregisterAll()
	time.Sleep(80 * time.Second)
}

// Stop : 停止服务端，graceful shutdown 时须调用次方法，以显式通知注册中心该服务已下线，减少客户端重试次数
func (p *RpcServer) Stop() {
	log.Println("Stop RpcServer")
	if p.cfg.OnUnregister != nil {
		p.cfg.OnUnregister()
	}
	Unregister(p.srv)
	if p.srv != nil {
		log.Println("Close Server")
		err := p.srv.Close()
		if err != nil {
			log.Println("stop rpcx server error:", err.Error())
		}
	}
}

func (p *RpcServer) startStateChecker() {
	if p.cfg.CheckStateUrl == "" {
		return
	}
	log.Println("StartHealthyChecker,url:", p.cfg.CheckStateUrl)
	for {
		if isMovingOut(p.cfg.CheckStateUrl) {
			log.Println("MovingOut!")
			p.Stop()
			break
		}
		time.Sleep(10 * time.Second)
	}
}
func isMovingOut(url string) bool {
	bs, err := Http("GET", url, nil)
	res := string(bs)
	res = strings.ReplaceAll(res, "\"", "")
	if err != nil {
		log.Println(err.Error())
		return true
	}
	return string(res) == "MovingOut"
}
func getLocalIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("Can not find the client ip address!")
}

func getLocalIpWindows() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	addrs, err := net.LookupIP(hostname)
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if strings.Contains(addr.String(), ":") {
			continue // 跳过ipv6地址
		}
		return addr.String(), nil
	}
	return "", errors.New("no ip found")
}
