package duinflux

import "log"

var dataBuffer chan Metric

const (
	bufferLen = 10000
)

func init() {
	dataBuffer = make(chan Metric, bufferLen)
}
func push(m Metric) {
	if len(dataBuffer) == bufferLen {
		log.Println("InfluxDB data buffer overflow,data discard:", m.Name)
		return
	}
	dataBuffer <- m
}
