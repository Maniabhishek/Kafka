package messaging

import (
	"github.com/Maniabhishek/Kafka/kafkaclient"
	"github.com/Maniabhishek/Kafka/messageconfig"
)

type IConsumer interface {
	FetchMessage()
	Close()
}

type IProducer interface {
	WriteMessage(key string, payload []byte) error
	TopicName() string
	Close()
}

func newProducer(broker messageconfig.BrokerConfig, topic string) (IProducer, error) {
	return kafkaclient.NewKafkaProducer(broker, topic)
}

func newConsumer(broker messageconfig.BrokerConfig, groupId string, topics []string) (IConsumer, error) {
	return kafkaclient.NewKafkaConsumer(broker, groupId, topics)
}
