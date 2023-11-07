package kafkaclient

import (
	"github.com/Maniabhishek/Kafka/messageconfig"
	"github.com/segmentio/kafka-go/sasl"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func GetAuth(broker messageconfig.BrokerConfig) sasl.Mechanism {
	if broker.EnableAuth {
		return plain.Mechanism{
			Username: broker.Username,
			Password: broker.Password,
		}
	}
	return nil
}
