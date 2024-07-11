package usecase

import (
	"clean/architector/internal/domain/repository"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type ConsumerUseCase struct {
	IrepoConsumer repository.IConsumerRepo
}

func InitConsumerUseCase(IrepoConsumer repository.IConsumerRepo) *ConsumerUseCase{
	fmt.Println("InitConsumerUseCase")
	return &ConsumerUseCase{IrepoConsumer: IrepoConsumer}
}

func (c *ConsumerUseCase) Run(topicName string, action string){
fmt.Println("consumerUseCase Run()")

// start goroutine by name
ctx, cancel := context.WithCancel(context.Background())
ctx = context.WithValue(ctx, "consumers", topicName)

go createConsumer(ctx)

if action == "close" {
	cancel()
}

}

 func createConsumer(ctx context.Context){
	name := ctx.Value("consumers").(string)
	fmt.Println("Создаем consumer "+name)

	for {
		select {
 case <-ctx.Done():
	fmt.Println("Удаляем горутину "+name)
		 return
		default:
			// Делать что-то

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "0.0.0.0:29092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	err = c.SubscribeTopics([]string{name, "^aRegex.*[Tt]opic"}, nil)

	if err != nil {
		panic(err)
	}

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true

	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			prepairMsg := string(*msg.TopicPartition.Topic)+" "+string(msg.Value)+"\n"
			writeToFile(prepairMsg)
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
	c.Close()
}
}

 }

 func writeToFile(msg string){
	//prepairMsg := []byte(msg)
file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, 0600)

if err != nil {
	fmt.Println("err create file", err)
	os.Exit(1)
}

defer file.Close()
file.WriteString(msg)
}
