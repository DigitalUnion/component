/**
@ServicePath : chenwei
@Time : 2021/7/31 下午11:11
**/
package duinflux

import (
	"context"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/influxdata/influxdb-client-go/v2/domain"
	"log"
	"strings"
	"time"
)

var (
	client    influxdb2.Client
	writeAPI  api.WriteAPI
	pathsType = 0
	paths     []string
	batchSize = 50000
)

// Connect Connect to influxdb
func Connect(cfg InfluxConfig) {
	log.Println("Connect influxDB:", cfg.Url, cfg.Bucket, cfg.Token)
	client = influxdb2.NewClient(cfg.Url, cfg.Token)
	if client == nil {
		return
	}
	checkBucket(cfg)
	if cfg.BatchSize != 0 {
		batchSize = cfg.BatchSize
	}
	paths = cfg.PathList
	pathsType = cfg.PathListType
	
	writeOption := write.DefaultOptions()
	writeOption.SetPrecision(time.Microsecond)
	writeOption.SetMaxRetries(0)
	writeAPI = api.NewWriteAPI(cfg.Org, cfg.Bucket, client.HTTPService(), writeOption)
	go watchErrors()
	go watchMetrics()
	go uptimeJob()
}

// checkBucket : create bucket if not exists
func checkBucket(cfg InfluxConfig) {
	bucket, _ := client.BucketsAPI().FindBucketByName(context.Background(), cfg.Bucket)
	if bucket == nil {
		var grafanaAddr string
		var grafanaToken string
		if strings.Contains(cfg.Url, "172.17.147.39:") {
			grafanaAddr = "http://172.17.147.30:3000/"
			grafanaToken = "Bearer eyJrIjoiRzgzRUVYREpjcDEyMnZXMXVMOUVva1c0Mmc4WHVRV3kiLCJuIjoiY2xvdWQiLCJpZCI6MX0="
		} else if strings.Contains(cfg.Url, "172.17.147.41:") {
			grafanaAddr = "http://172.17.147.39:3000/"
			grafanaToken = "Bearer eyJrIjoiNHBYSE8yNWRKVTc5OHZ4ZDA1bk4ySWM4R1lEd01pekwiLCJuIjoiY2xvdWQiLCJpZCI6MX0="
		} else if strings.Contains(cfg.Url, "172.17.129.178:") {
			grafanaAddr = "http://172.17.129.178:3000/"
			grafanaToken = "Bearer eyJrIjoiNU5WbmNuaEdTMUdTWGtqYzVjTHRwZzNIMlFHeVY3aDEiLCJuIjoiVEVTVCIsImlkIjoxfQ=="
		} else {
			return
		}
		log.Printf("Bucket: [%s] not found,try to create it...\n", cfg.Bucket)
		org, err := client.OrganizationsAPI().FindOrganizationByName(context.Background(), cfg.Org)
		if err != nil {
			log.Printf("Bucket: [%s] create faild: org [%s] not found!", cfg.Bucket, cfg.Org)
			//panic(err.Error())
		}
		if org == nil {
			log.Printf("Bucket: [%s] create faild: org [%s] not found!", cfg.Bucket, cfg.Org)
			//panic(err)
		}
		_, err = client.BucketsAPI().CreateBucketWithName(context.Background(), org, cfg.Bucket, domain.RetentionRule{EverySeconds: 86400 * 30, Type: domain.RetentionRuleTypeExpire})
		if err != nil {
			log.Printf("Bucket: [%s] create faild:%s\n", cfg.Bucket, err.Error())
			return
		} else {
			log.Printf("Bucket: [%s] create success!\n", cfg.Bucket)
		}
		
		log.Printf("Try to create Grafana Dashboard: [%s] at: %s...\n", cfg.Bucket, grafanaAddr)
		createGrafanaDashboard(grafanaAddr, grafanaToken, cfg.Bucket)
	}
}
func createGrafanaDashboard(grafanaAddr, grafanaToken string, bucket string) {
	grafanaAddr += "api/dashboards/db"
	req := strings.ReplaceAll(grafanaTemplate, "$title", bucket)
	req = strings.ReplaceAll(req, "$bucket", bucket)
	res := post(grafanaAddr, grafanaToken, req)
	log.Printf("Create Grafana Dashboard: [%s] result:%s\n", bucket, res)
}
func Count(name string) {
	if client == nil {
		return
	}
	m := Metric{Name: name, Cost: 0}
	push(m)
}
func Cost(name string, costMicroSeconds int64) {
	if client == nil {
		return
	}
	m := Metric{Name: name, Cost: costMicroSeconds}
	push(m)
}
func Done(name string, startTime time.Time) time.Time {
	if client == nil {
		return time.Now()
	}
	now := time.Now()
	m := Metric{Name: name, Cost: now.Sub(startTime).Microseconds()}
	push(m)
	return now
}

// watchErrors watch for write errors
func watchErrors() {
	if writeAPI == nil {
		return
	}
	for data := range writeAPI.Errors() {
		log.Println(data.Error())
	}
}
func Close() {
	client.Close()
}
