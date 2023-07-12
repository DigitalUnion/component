package dunsq

import (
	"fmt"
	"strings"
)

type NsqdNodesInfo struct {
	Producers []Producer `json:"producers"`
}

type Producer struct {
	RemoteAddress    string   `json:"remote_address"`
	HostName         string   `json:"host_name"`
	BroadcastAddress string   `json:"broadcast_address"`
	TcpPort          int      `json:"tcp_port"`
	HttpPort         int      `json:"http_port"`
	Version          string   `json:"version"`
	Tombstones       []bool   `json:"tombstones"`
	Topics           []string `json:"topics"`
}

func (p *Producer) buildNsqAddr() string {
	index := strings.Index(p.RemoteAddress, ":")
	if index == -1 {
		return ""
	}
	host := p.RemoteAddress[:index]
	return fmt.Sprintf("%s:%d", host, p.TcpPort)
}
