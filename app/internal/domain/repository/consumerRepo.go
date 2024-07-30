package repository

import "fmt"

type IConsumerRepo interface {
	CreateConsumer()
}

type ConsumerRepo struct {
	ConsumerName string
}

func InitConsumerRepo(consumerName string) *ConsumerRepo{
	return &ConsumerRepo{ConsumerName: consumerName}
}

func(c *ConsumerRepo) CreateConsumer(){
fmt.Println("создаем consumer")
}
