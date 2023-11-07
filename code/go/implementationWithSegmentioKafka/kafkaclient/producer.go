package kafkaclient

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"github.com/Maniabhishek/Kafka/messageconfig"
	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(broker messageconfig.BrokerConfig, topic string) (*KafkaProducer, error) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:    strings.Split(broker.Url, ","),
		Topic:      topic,
		BatchBytes: 26214400,
		Dialer: &kafka.Dialer{
			SASLMechanism: GetAuth(broker),
			TLS:           &tls.Config{},
		},
	})

	return &KafkaProducer{
		writer: writer,
	}, nil
}

func (k *KafkaProducer) Close() {
	k.writer.Close()
}

func (k *KafkaProducer) TopicName() string {
	return k.writer.Topic
}

func (k *KafkaProducer) WriteMessage(key string, payload []byte) error {
	attempts := 0
	var err error
	for attempts < 5 {
		err := k.writer.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(key),
			Value: []byte(payload),
		})

		if err != nil {
			fmt.Println("error while writing", err)
			time.Sleep(5 * time.Second)
			attempts += 1
			continue
		} else {
			break
		}
	}
	return err
}
