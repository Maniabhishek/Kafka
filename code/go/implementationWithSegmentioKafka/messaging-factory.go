package main

import (
	"os"
	"strings"

	"github.com/Maniabhishek/Kafka/messageconfig"
)

type MessagingClientFactory interface {
	NewProducer(topic string) (IProducer, error)
	NewConsumer(groupId string, topics ...string) (IConsumer, error)
}

type clientFactory struct {
}

func NewKafkaClientFactory() MessagingClientFactory {
	return &clientFactory{}
}

func (c *clientFactory) NewProducer(topic string) (IProducer, error) {
	broker := messageconfig.BrokerConfig{
		Url:                os.Getenv("BROKERS"),
		Username:           os.Getenv("USERNAME"),
		Password:           os.Getenv("PASSWORD"),
		EnableAuth:         strings.ToLower(os.Getenv("ENABLE_AUTH")) == "true",
		InsecureSkipVerify: true,
	}

	return newProducer(broker, topic)
}

func (c *clientFactory) NewConsumer(groupdId string, topics ...string) (IConsumer, error) {
	broker := messageconfig.BrokerConfig{
		Url:                os.Getenv("BROKERS"),
		Username:           os.Getenv("USERNAME"),
		Password:           os.Getenv("PASSWORD"),
		EnableAuth:         strings.ToLower(os.Getenv("ENABLE_AUTH")) == "true",
		InsecureSkipVerify: true,
	}

	return newConsumer(broker, groupdId, topics)
}
