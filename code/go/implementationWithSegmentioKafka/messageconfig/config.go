package messageconfig

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type BrokerConfig struct {
	Url                string
	Username           string
	Password           string
	EnableAuth         bool
	InsecureSkipVerify bool
}

type EventMessage interface {
	GetTopic() string
	GetPayload() []byte
	CommitMessage() error
	GetKey() []byte
}

type emessage struct {
	context context.Context
	payload *kafka.Message
	reader  *kafka.Reader
}

func NewMessageConsumer(ctx context.Context, payload *kafka.Message, reader *kafka.Reader) EventMessage {
	return &emessage{
		context: ctx,
		payload: payload,
		reader:  reader,
	}
}

func (e *emessage) GetTopic() string {
	return e.payload.Topic
}

func (e *emessage) GetPayload() []byte {
	return e.payload.Value
}

func (e *emessage) CommitMessage() error {
	return e.reader.CommitMessages(e.context, *e.payload)
}

func (e *emessage) GetKey() []byte {
	return e.payload.Key
}
