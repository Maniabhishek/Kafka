package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Maniabhishek/Kafka/apis"
	"github.com/Maniabhishek/Kafka/internal/services"
	"github.com/Maniabhishek/Kafka/messaging"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	go func() {
		mfactory := messaging.NewKafkaClientFactory()

		mservices := services.NewMessageService(mfactory)

		mcontroller := apis.NewMessageController(mservices)

		e := echo.New()
		e.POST("/sendmessage", mcontroller.SendKafkaMessage)
		e.Logger.Fatal(e.Start(":1234"))
	}()

	go ConsumeMessage()

	stopC := make(chan os.Signal, 1)
	signal.Notify(stopC, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-stopC
}

func ProduceMessage() {
	kafkaClient := messaging.NewKafkaClientFactory()

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
	fmt.Println("running....")
	consumerClient := messaging.NewKafkaClientFactory()
	cc, err := consumerClient.NewConsumer("user-group", "user")

	if err != nil {
		fmt.Println("err", err)
	}

	cc.FetchMessage()
}
