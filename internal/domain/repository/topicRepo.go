package repository

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type TopicRepo struct {
	TopicName string
	MsgTopic  string
}

func NewMsgTopic(topicId string, msgTopic string) *TopicRepo {
	return &TopicRepo{TopicName: topicId, MsgTopic: msgTopic}
}

func (t *TopicRepo) AddMsgToTopic() {

	fmt.Println(t)

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

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &t.MsgTopic, Partition: kafka.PartitionAny},
		Key:            []byte(t.TopicName),
		Value:          []byte(t.MsgTopic),
	}, nil)

	p.Flush(3 * 1000)
	p.Close()

}
