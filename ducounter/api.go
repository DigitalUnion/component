package ducounter

import "sync/atomic"

// Init mysqlConfig为mysql的dsn，messageTypeTableMap为消息类型和对应表明的映射关系
func Init(mysqlConfig string, messageTypeTableNameMap map[string]string) {
	connectDb(mysqlConfig)
	initVariable(messageTypeTableNameMap)
	go startJob()
}

func CounterAdd(messageType string) {
	_, ok := counterMap[messageType]
	if !ok {
		return
	}
	atomic.AddInt64(&counterMap[messageType].Counter, 1)
}
