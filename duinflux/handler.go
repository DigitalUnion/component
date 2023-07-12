package duinflux

import (
	"errors"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"net"
	"time"
)

var localIp string
var tk = time.NewTicker(1 * time.Second)
var tkTrigger = false
var metricMap = make(map[string]*Metric)

func init() {
	go timeTicker()
	localIp, _ = getLocalIp()
}

func watchMetrics() {
	handleCount := 0
	for m := range dataBuffer {
		add(m)

		handleCount++
		if handleCount >= batchSize {
			flush()
			handleCount = 0
		} else if tkTrigger {
			flush()
			tkTrigger = false
		}
	}
}

func timeTicker() {
	for range tk.C {
		tkTrigger = true
	}
}

func add(m Metric) {
	_, ok := metricMap[m.Name]
	if !ok {
		metricMap[m.Name] = &Metric{
			Count: 1,
			Cost:  m.Cost,
		}
	} else {
		metricMap[m.Name].Count += 1
		metricMap[m.Name].Cost += m.Cost
	}
}
func flush() {
	if len(metricMap) == 0 {
		return
	}
	if writeAPI == nil {
		return
	}
	now := time.Now()
	for k, v := range metricMap {
		if k == "" {
			continue
		}
		if v.Count == 0 {
			continue
		}

		fields := make(map[string]interface{})
		fields["qps"] = v.Count
		if v.Cost > 0 {
			fields["cost"] = v.Cost / int64(v.Count)
		}
		writeAPI.WritePoint(write.NewPoint(k, map[string]string{"ip": localIp}, fields, now))

		v.Count = 0
		v.Cost = 0
	}
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
