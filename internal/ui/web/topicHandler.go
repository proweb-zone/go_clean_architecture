package web

import (
	"clean/architector/internal/domain/adapter"
	"clean/architector/internal/domain/entitie"
	"clean/architector/internal/domain/repository"
	"clean/architector/internal/domain/usecase"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func AddMsgToTopicHandler(w http.ResponseWriter, r *http.Request) {
	var topicName string = chi.URLParam(r, "topic_name")

	if topicName == "" {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var newMsgTopic entitie.MsgTopic
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newMsgTopic)
	if err != nil {
		http.Error(w, "Body MsgTopic not exist", http.StatusBadRequest)
		return
	}

	var iTopicrepo repository.ItopicRepo = repository.InitTopicRepo(topicName, newMsgTopic.Msg)
	var iTopicUseCase adapter.ItopicUseCase = usecase.InitTopicUseCase(iTopicrepo)
	iTopicUseCase.SendMsgToTopic()
}
