package dukafka

import (
	"git.du.com/cloud/du_component/dukafka/kafka"
	"log"
)

type Producer struct {
	config *Config
	client *kafka.Producer
}

func NewProducer(config Config, handle func(partition kafka.TopicPartition)) (*Producer, error) {
	if config.TimeoutMs == 0 {
		config.TimeoutMs = 15 * 1000
	}
	
	var kafkaConf = &kafka.ConfigMap{
		"bootstrap.servers":   config.Hosts,
		"api.version.request": config.ApiVersionRequest,
	}
	
	if config.Acks != "" {
		kafkaConf.SetKey("acks", config.Acks)
	}
	
	var err error
	p := Producer{config: &config}
	p.client, err = kafka.NewProducer(kafkaConf)
	if err != nil {
		return nil, err
	}
	
	// Delivery report handler for produced messages
	go func() {
		for e := range p.client.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					if handle != nil {
						handle(ev.TopicPartition)
					} else {
						log.Printf("Delivered message to %v\n", ev.TopicPartition)
					}
				}
			}
		}
	}()
	return &p, nil
}

func (p Producer) SendMsg(msg []byte) error {
	if p.client == nil {
		return nil
	}
	// Produce messages to topic (asynchronously)
	err := p.client.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &p.config.Topic,
			Partition: kafka.PartitionAny,
		},
		Value: msg,
	}, nil)
	
	// Wait for message deliveries before shutting down
	p.client.Flush(p.config.TimeoutMs)
	return err
}

func (p Producer) Close() {
	if p.client != nil {
		p.client.Close()
	}
}
