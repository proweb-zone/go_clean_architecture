package usecase

import (
	"clean/architector/internal/domain/repository"
	"fmt"
)

type TopicUseCase struct {
	ITopicRepo repository.ItopicRepo
}

func InitTopicUseCase(iTopicRepo repository.ItopicRepo) *TopicUseCase {
	return &TopicUseCase{ITopicRepo: iTopicRepo}
}

func (t *TopicUseCase) SendMsgToTopic() {
	fmt.Println("call func SendMsgToTopic to topicUseCase.go")
	t.ITopicRepo.NewMsgToTopic()
}
