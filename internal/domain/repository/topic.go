package repository

import (
	"fmt"
	"os"

	"clean/architector/internal/domain/entitie"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func AddMsgToTopic(topicName string, newMsgTopic entitie.MsgTopic) {

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "0.0.0.0:29092",
		"client.id":         "localhost",
		"acks":              "all"})

	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}

	// Go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Produced event to topic %s: key = %-10s value = %s\n",
						*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
				}
			}
		}
	}()

	// p.Produce(&kafka.Message{
	// 	TopicPartition: kafka.TopicPartition{Topic: &topicName, Partition: kafka.PartitionAny},
	// 	Key:            []byte(newMsgTopic.Id),
	// 	Value:          []byte(newMsgTopic.Msg),
	// }, nil)

	// p.Flush(3 * 1000)
	// p.Close()

}
