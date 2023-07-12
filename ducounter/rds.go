package ducounter

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"
)

var db *gorm.DB

func connectDb(dsn string) {
	db = connect(dsn)
}

func connect(connectStr string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(connectStr), &gorm.Config{SkipDefaultTransaction: true, Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		log.Println("connect db error:", err.Error())
	}
	return db
}

func initVariable(messageTypeTableNameMap map[string]string) {
	ip = GetIp()
	for k, v := range messageTypeTableNameMap {
		counterMap[k] = &CounterCell{
			Counter:     0,
			Lock:        sync.RWMutex{},
			MessageType: k,
			TableName:   v,
		}
		counterMapTmp[k] = &CounterCell{
			Counter:     0,
			Lock:        sync.RWMutex{},
			MessageType: k,
			TableName:   v,
		}
	}
}
