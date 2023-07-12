package dukafka

import (
	"encoding/json"
	"git.du.com/cloud/du_component/dukafka/kafka"
	"log"
	"strings"
)

type Consumer struct {
	config *Config
	client *kafka.Consumer
}

func NewCustomer(config Config) (*Consumer, error) {
	var kafkaConf = &kafka.ConfigMap{
		"api.version.request": "true",
		"bootstrap.servers":   config.Hosts,
	}
	
	if config.GroupId != "" {
		kafkaConf.SetKey("group.id", config.GroupId)
	}
	
	if config.OffsetReset != "" {
		kafkaConf.SetKey("auto.offset.reset", config.OffsetReset)
	}
	if config.SecurityProtocol == "sasl_ssl" {
		kafkaConf.SetKey("security.protocol", "sasl_ssl")
		kafkaConf.SetKey("ssl.ca.location", config.SslCaLocation)
		kafkaConf.SetKey("sasl.username", config.SaslUsername)
		kafkaConf.SetKey("sasl.password", config.SaslPassword)
		kafkaConf.SetKey("sasl.mechanism", "PLAIN")
		
		j, e := json.Marshal(kafkaConf)
		log.Println("cfg:", string(j), e)
	}
	
	var err error
	c := Consumer{}
	c.client, err = kafka.NewConsumer(kafkaConf)
	if err != nil {
		return nil, err
	}
	c.client.SubscribeTopics(strings.Split(config.Topic, ","), nil)
	return &c, err
}

func (p Consumer) StartCustomer(handle func(msg string)) {
	log.Println("start customer")
	for {
		msg, err := p.client.ReadMessage(-1)
		if err == nil {
			//log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
		if handle != nil {
			handle(string(msg.Value))
		}
	}
}

func (p Consumer) Close() {
	if p.client != nil {
		p.client.Close()
	}
}
