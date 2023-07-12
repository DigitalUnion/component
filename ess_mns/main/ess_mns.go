package main

import (
	"flag"
	"git.du.com/cloud/du_component/ess_mns/mns"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var httpAddr string
var topic string

func main() {
	flag.StringVar(&httpAddr, "httpAddr", ":2222", "listen http addr")
	flag.StringVar(&topic, "topic", "", "mns topic")
	flag.Parse()

	if topic != "" {
		mns.MNS.Topic = topic
	}
	clinet, err := mns.MnsInit()

	if err != nil {
		// todo 报警
		log.Println("mns init err:", err)
		return
	}

	// 接收缩容消息
	clinet.ReceiveMsg()
	mns.StartHttpServer(httpAddr)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	<-quit
	clinet.Close()
}
