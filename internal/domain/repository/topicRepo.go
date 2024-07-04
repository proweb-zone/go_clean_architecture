package repository

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type TopicRepo struct {
	TopicName string
	MsgTopic  string
}

func InitTopicRepo(topicId string, msgTopic string) *TopicRepo {
	return &TopicRepo{TopicName: topicId, MsgTopic: msgTopic}
}

func (t *TopicRepo) NewMsgToTopic() {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "0.0.0.0:29092"})

		if err != nil {
			panic(err)
		}

		defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &t.TopicName, Partition: kafka.PartitionAny},
		Value:          []byte(t.MsgTopic),
	}, nil)

	// Wait for message deliveries before shutting down
	p.Flush(3 * 1000)

}
