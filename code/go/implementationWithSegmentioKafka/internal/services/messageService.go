package services

import (
	"github.com/Maniabhishek/Kafka/messaging"
)

type mservice struct {
	mclient messaging.MessagingClientFactory
}

type MInterface interface {
	SendTokafka(key string, payload []byte) error
}

func NewMessageService(mclient messaging.MessagingClientFactory) MInterface {
	return &mservice{
		mclient: mclient,
	}
}

func (m *mservice) SendTokafka(key string, payload []byte) error {
	producer, err := m.mclient.NewProducer("user")

	if err != nil {
		return err
	}

	err = producer.WriteMessage(key, payload)
	if err != nil {
		return err
	}

	return nil
}
