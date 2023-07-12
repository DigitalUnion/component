#云端基础组件封装

## 目录
- [du_mongo](#du_mongo)
- [du_kafka](#du_kafka)




## du_mongo

- Installation

`go get git.du.com/cloud/du_component/dukafka`
- Quickstart
``` go
package main

import (
	"git.du.com/cloud/du_component/dumongo"
	"log"
)

func main() {
	config := dumongo.Config{
		Hosts: "172.17.129.199:27017",
		Username: "dna2_lab_rw",
		Password: "DUw@s2xo3pm4ds#%WWS",
		DbName: "dna2_lab",
		AuthSource: "dna2_lab",
	}
	// new client
	client, err := dumongo.NewMongo(config)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	// close
	err = dumongo.CloseMongo(client)
	if err != nil {
		log.Fatalf("%+v", err)
	}
}
```

## du_kafka

- Installation

`go get git.du.com/cloud/du_component/dukafka`
- Quickstart Producer
``` go
package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"git.du.com/cloud/du_component/dukafka"
	"os"
	"os/signal"
)

var client *kafka.Producer

func initKafka() {
	var err error
	client, err = dukafka.NewProducer(dukafka.ProducerConfig{
		Hosts: "xxx",
		Topic: "confluent-test",
	}, func(partition kafka.TopicPartition) { // 处理发送失败的消息
		fmt.Printf("error: %s\n", partition.Error)
		fmt.Printf("Delivered message to %v\n", partition)
	})
	fmt.Println(err)
}

func main() {
	initKafka()
	dukafka.SendChannelMsg(client, []byte("test message"))
	
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	dukafka.CloseProducer(client)
}
```
- Quickstart Consumer
``` go
package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"new_test/v1/func/kafka/dukafka"
	"os"
	"os/signal"
)

var consumerClient *kafka.Consumer

func intConsumer() {
	var err error
	consumerClient, err = dukafka.NewCustomer(dukafka.ConsumerConfig{
		Hosts:   "xxx",
		Topic:   []string{"confluent-test"},
		GroupId: "confluent-kafak-test",
	})
	fmt.Println("init consumer:", err)
}

func main() {

	intConsumer()
	go dukafka.StartCustomer(consumerClient, func(msg string, err error) {
		fmt.Println(msg, err)
	})

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	dukafka.CloseCustomer(consumerClient)
}
```