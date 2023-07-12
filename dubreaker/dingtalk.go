package dubreaker

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

func send(tp, name, message string) {
	m := `{
		"msgtype": "markdown",
		"markdown": {
			"title": "状态切换",
			"text": "## %s 熔断
%s\n
告警 IP：%s\n
[%s]\n
%s\n
"
		}
	}`
	date := time.Now().Format("2006-01-02 15:04:05")
	ip, _ := getLocalIp()
	msg := fmt.Sprintf(m, name, tp, ip, date, message)

	resp, err := http.Post("https://oapi.dingtalk.com/robot/send?access_token=19d25e26fc3060641305dd92d350f7c278760ddf73d981064091b094d6722915",
		"application/json",
		strings.NewReader(msg),
	)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("send DingTalk alert error:", err.Error())
		return
	}
	log.Println("send DingTalk alert:", string(body))
}

func getLocalIp() (string, error) {
	var IpAddr string
	addrSlice, err := net.InterfaceAddrs()
	if nil != err {
		IpAddr = "localhost"
		return "nil", err
	}
	for _, addr := range addrSlice {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				IpAddr = ipnet.IP.String()
				return IpAddr, nil
			}
		}
	}
	return IpAddr, nil
}
