package usecase

import (
	"clean/architector/internal/domain/adapter"
	"fmt"
)

type TopicUseCase struct {
	ITopicRepo adapter.ItopicRepo
}

func NewTopicUseCase(iTopicRepo adapter.ItopicRepo) *TopicUseCase {
	return &TopicUseCase{ITopicRepo: iTopicRepo}
}

func (t *TopicUseCase) SendMsgToTopic() {
	fmt.Println("call func SendMsgToTopic to topicUseCase.go")
}
