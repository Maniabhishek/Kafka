package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	ProduceMessage()
	ConsumeMessage()

	fmt.Println(":::::::::::::::::::end:::::::::::::::::::::")
}

func ProduceMessage() {
	kafkaClient := NewKafkaClientFactory()

	producer, err := kafkaClient.NewProducer("user")

	if err != nil {
		fmt.Println("err:::::::", err)
		return
	}

	payload := []byte("abhishek")

	err = producer.WriteMessage("username", payload)
	if err != nil {
		fmt.Println("main line18", err)
	}
}

func ConsumeMessage() {
	consumerClient := NewKafkaClientFactory()
	cc, err := consumerClient.NewConsumer("user-group", "user")

	if err != nil {
		fmt.Println("err", err)
	}

	cc.FetchMessage()
}
