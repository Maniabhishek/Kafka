package kafkaclient

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"

	"github.com/Maniabhishek/Kafka/messageconfig"
	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	reader   *kafka.Reader
	isClosed bool
}

func NewKafkaConsumer(broker messageconfig.BrokerConfig, groupId string, topics []string) (*KafkaConsumer, error) {
	client := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     strings.Split(broker.Url, ","),
		MaxBytes:    26214400, //25mb
		GroupID:     groupId,
		GroupTopics: topics,
		Dialer: &kafka.Dialer{
			SASLMechanism: GetAuth(broker),
			TLS:           &tls.Config{},
		},
	})
	return &KafkaConsumer{
		reader: client,
	}, nil
}

// function to read message from kafka topics, need to explicity commit messages
func (k *KafkaConsumer) FetchMessage() {
	ctx := context.Background()
	for {
		if k.isClosed {
			break
		}
		message, err := k.reader.FetchMessage(ctx)
		if err != nil {
			fmt.Println("error while fetching message", err)
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", message.Topic, message.Partition, message.Offset, string(message.Key), string(message.Value))
	}
}

func (k *KafkaConsumer) Close() {
	k.isClosed = true
	k.reader.Close()
}
