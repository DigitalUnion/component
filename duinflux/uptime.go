package duinflux

import (
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/robfig/cron"
	"time"
)

var upseconds int64 = 0

func uptimeJob() {
	c := cron.New()
	c.AddFunc("0/10 * * * * *", flushUptime)
	c.Start()
}
func flushUptime() {
	upseconds += 10
	writeAPI.WritePoint(write.NewPoint(
		"uptime",
		map[string]string{"ip": localIp},
		map[string]interface{}{"seconds": upseconds},
		time.Now(),
	))

}
