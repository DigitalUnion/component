package main

import (
	"fmt"
	"git.du.com/cloud/du_component/dukafka"
	"log"
	"time"
)

func main() {

	cfg := dukafka.Config{
		ApiVersionRequest: true,
		TimeoutMs:         1000,
		Acks:              "1",
		Hosts:             "alikafka-pre-cn-tl32n3ipw007-1-vpc.alikafka.aliyuncs.com:9092,alikafka-pre-cn-tl32n3ipw007-2-vpc.alikafka.aliyuncs.com:9092,alikafka-pre-cn-tl32n3ipw007-3-vpc.alikafka.aliyuncs.com:9092",
		Topic:             "songhz_test",
	}
	producer, err := dukafka.NewProducer(cfg, nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for i := 0; i < 1000; i++ {
		time.Sleep(1 * time.Second)
		log.Println("send msg")
		producer.SendMsg([]byte("bbb"))
	}
	fmt.Println("end")
}
