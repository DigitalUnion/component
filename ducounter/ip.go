package ducounter

import "net"

func GetIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "getLocalIpErr"
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "getLocalIpErr"
}
