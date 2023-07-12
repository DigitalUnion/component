package ducounter

import (
	"gorm.io/gorm"
	"sync"
	"time"
)

type CounterCell struct {
	Counter     int64
	Lock        sync.RWMutex
	MessageType string
	TableName   string
}

var tk = time.NewTicker(1 * time.Second)
var counterMap = make(map[string]*CounterCell)
var counterMapTmp = make(map[string]*CounterCell)
var ip string

func startJob() {
	for range tk.C {
		upload()
	}
}

func upload() {
	for k, v := range counterMap {
		v.Lock.Lock()
		counterMapTmp[k].Counter = v.Counter
		v.Counter = 0
		v.Lock.Unlock()
	}
	//上传至数据库
	for _, v := range counterMapTmp {
		uploadToDb(v.TableName, v.Counter)
	}
}

func uploadToDb(dbName string, count int64) {
	ts := time.Now()
	tx := db.Table(dbName)
	tx.Where("time=?", ts)
	tx.Where("name=?", ip)
	var total int64
	tx.Count(&total)
	if total == 0 {
		tmp := Info{
			Time:  ts,
			Name:  ip,
			Count: count,
		}
		tx.Create(tmp)
		return
	}
	tx.Limit(1).Update("count", gorm.Expr("count+?", count))
}
