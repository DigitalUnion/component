package dunsq

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/nsqio/go-nsq"
	"log"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"unsafe"
)

var nsqdHandler *NsqdHandler

// Connect 初始化连接
func Connect(cfg NsqConfig) {
	if !cfg.isValid() {
		panic("nsq config error")
	}
	//连接nsqd集群
	connectNsqlookupd(cfg.Addrs)
}

// connectNsqlookupd 连接nsqlookupd集群
func connectNsqlookupd(addrs string) {
	//获取nsqd节点
	nsqlookupdNodes := strings.Split(addrs, ",")
	nsqNodes := getNsqdNodesInfo(nsqlookupdNodes)
	if len(nsqNodes) == 0 {
		panic("don't get nsqd node'info")
	}
	nsqdHandler = &NsqdHandler{
		producers: make(map[string]*nsq.Producer),
	}
	for _, nsqNode := range nsqNodes {
		config := nsq.NewConfig()
		producer, err := nsq.NewProducer(nsqNode, config)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		err = producer.Ping()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		nsqdHandler.producers[nsqNode] = producer
		nsqdHandler.nsqdNodes = append(nsqdHandler.nsqdNodes, nsqNode)
		nsqdHandler.size++
	}
	if nsqdHandler.isEmpty() {
		panic("connect nsq cluster fail")
	}
	nsqdHandler.nsqlookupdNodes = nsqlookupdNodes
}

// getNsqdNodesInfo 获取nsqd节点信息
func getNsqdNodesInfo(addrs []string) []string {
	var nodes []string
	for _, addr := range addrs {
		if addr == "" {
			continue
		}
		res, err := Http(http.MethodGet, buildGetNodesUrl(addr), nil)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		if res == nil {
			log.Printf("nsqlookupd %s don't connect\n", addr)
			continue
		}
		var nsqdNodesInfo NsqdNodesInfo
		err = jsoniter.Unmarshal(res, &nsqdNodesInfo)
		if err != nil {
			log.Println(err.Error())
		}
		for _, p := range nsqdNodesInfo.Producers {
			nsqAddr := p.buildNsqAddr()
			if nsqAddr == "" {
				continue
			}
			nodes = append(nodes, nsqAddr)
		}
		if len(nodes) != 0 {
			break
		}
	}
	return nodes
}

func PushBytes(topic string, data []byte) {
	nsqdHandler.SendBytes(topic, data)
}
func PushString(topic string, data string) {
	nsqdHandler.SendString(topic, data)
}

type NsqdHandler struct {
	producers       map[string]*nsq.Producer
	nsqdNodes       []string
	nsqlookupdNodes []string
	size            int
	index           int64
	l               sync.RWMutex
}

// AddProducer Add 添加连接
func (handler *NsqdHandler) AddProducer(producer *nsq.Producer) int {

	return 0
}

// Del 删除连接
func (handler *NsqdHandler) Del(i int) {

}

// GetProducer 获取连接
func (handler *NsqdHandler) GetProducer() *nsq.Producer {
	return handler.producers[handler.nsqdNodes[int(handler.index)%handler.size]]
}

// SendBytes 发送字节消息
func (handler *NsqdHandler) SendBytes(topic string, data []byte) {
	atomic.AddInt64(&handler.index, 1)
	err := handler.GetProducer().Publish(topic, data)
	if err != nil {
		return
	}
}

// SendString 发送字符串消息
func (handler *NsqdHandler) SendString(topic string, data string) {
	atomic.AddInt64(&handler.index, 1)
	err := handler.GetProducer().Publish(topic, string2Bytes(data))
	if err != nil {
		return
	}
}

//判断连接是否为空
func (handler *NsqdHandler) isEmpty() bool {
	handler.l.RLock()
	defer handler.l.RUnlock()
	if handler.size == 0 {
		return true
	}
	return false
}

type NsqConfig struct {
	Addrs  string   `json:"addrs" yaml:"addrs"`   //nsqlookupd集群 ip:port,ip:port
	Topics []string `json:"topics" yaml:"topics"` //topic 自动创建  会自动检测nsq集群有无此topic
}

func (cfg *NsqConfig) isValid() bool {
	if cfg.Addrs == "" || cfg.Topics == nil {
		return false
	}
	return true
}

func buildGetNodesUrl(addr string) string {
	return fmt.Sprintf("http://%s/nodes", addr)
}

func string2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
