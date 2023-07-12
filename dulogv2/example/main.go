package main

import (
	"git.du.com/cloud/du_component/dulogv2"
)

var (
	bizLog *dulogv2.Logger
	//runLog *dulogv2.Logger
)

func main() {
	//eg 1:远程
	bizLog = dulogv2.NewLogger(dulogv2.Config{
		Addrs: []string{"172.17.129.178:30514"},
		//Addrs: []string{
		//	"172.17.0.166:30514",
		//	"172.17.0.167:30514",
		//},
		Appid:   "test2", //appid中不能含有 '-'
		Module:  "bbb",   //module中不能含有 '-'
		IsCheck: false,
		IsEs:    true, //支持索引到es
	})
	for i := 0; i < 2; i++ {
		bizLog.PlainTextfn(`{"a":"a%d","b":"%d"}`, i, i)
	}
	//for i := 0; i < 10; i++ {
	//	bizLog.Infoln("bb", "cc", "dd")
	//}
	//for i := 0; i < 10; i++ {
	//	bizLog.Debugln("bb", "cc", "dd")
	//}
	//for i := 0; i < 10; i++ {
	//	bizLog.Errorln("bb", "cc", "dd")
	//}
	//
	////eg 2:本地
	//runLog = dulogv2.NewLogger(dulogv2.Config{
	//	Appid:       "test", //appid中不能含有 '-'
	//	Module:      "run",  //module中不能含有 '-'
	//	IsLoc:       true,
	//	LocDir:      "./logs",
	//	LocCompress: true,
	//	LocMaxAge:   7,
	//	LocMaxSize:  1024, //1024MB
	//})
	//
	//for i := 0; i < 10; i++ {
	//	runLog.PlainTextln("aa", "bb", "cc")
	//}
	//for i := 0; i < 10; i++ {
	//	runLog.Infoln("aa", "bb", "cc")
	//}
	//for i := 0; i < 10; i++ {
	//	runLog.Debugln("aa", "bb", "cc")
	//}
	//
	////eg 3:中间件
	//r := gin.New()
	//r.Use(gin.RecoveryWithWriter(runLog))
	//r.Use(dulogv2.GinDulogMiddleware(runLog))
	//r.GET("/ping", func(c *gin.Context) {
	//	c.String(200, "pong")
	//	//panic("paniccccshdasghcvghdsvcghdsvdggh")
	//})
	//r.Run(":8080")
}
