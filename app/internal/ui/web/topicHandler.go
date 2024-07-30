package web

import (
    "strings"

	"clean/architector/internal/domain/adapter"
	"clean/architector/internal/domain/entitie"
	"clean/architector/internal/domain/repository"
	"clean/architector/internal/domain/usecase"
	"encoding/json"
	"net/http"
)

func AddMsgToTopicHandler(w http.ResponseWriter, r *http.Request) {
	urlPathList := strings.Split(r.URL.Path, "/")

	if len(urlPathList) < 3 {
	    http.Error(w, "topicName not found", http.StatusBadRequest)
    	return
	}

	var topicName string = urlPathList[3]

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
