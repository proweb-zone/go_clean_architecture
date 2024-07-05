package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"clean/architector/internal/domain/adapter"
	"clean/architector/internal/domain/entitie"
	"clean/architector/internal/domain/repository"
	"clean/architector/internal/domain/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Host string
}

func NewServer(cfg *Config) *Server {
	return &Server{Host: cfg.HTTPServer.Address}
}

func (c *Server) StartServer() {
	fmt.Println("start server")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/kafka/topic/{topic_name}", sendMsgToTopicHandler)

	http.ListenAndServe(c.Host, r)
}

func sendMsgToTopicHandler(w http.ResponseWriter, r *http.Request) {
	var topicName string = chi.URLParam(r, "topic_name")

	if topicName == "" {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var newMsgTopic entitie.MsgTopic
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newMsgTopic)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var iTopicrepo adapter.ItopicRepo = repository.InitTopicRepo(topicName, newMsgTopic.Msg)
	var topicUseCase *usecase.TopicUseCase = usecase.InitTopicUseCase(iTopicrepo)
	topicUseCase.SendMsgToTopic()

}
