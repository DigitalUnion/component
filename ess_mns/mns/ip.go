package mns

import (
	"net"
	"sync"
)

const (
	GET_IP_FAILED = "GET_IP_FAILED"
)

var (
	once     sync.Once
	LOCAL_IP string
)

func GetIp() string {
	once.Do(func() {
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			return
		}
		for _, address := range addrs {
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					LOCAL_IP = ipnet.IP.String()
					return
				}
			}
		}
		LOCAL_IP = GET_IP_FAILED
	})
	return LOCAL_IP
}
