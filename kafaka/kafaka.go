package main

import (
	"github.com/Shopify/sarama"
)

const (
	brokerAddress = "localhost:9092"
	topic         = "kafka-topic"
)

func produceMessage() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

}
