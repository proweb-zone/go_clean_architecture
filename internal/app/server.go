package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"clean/architector/internal/domain/entitie"
	"clean/architector/internal/domain/repository"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
	r.Post("/kafka/topic/{topic_name}", addMsgInKafka)

	http.ListenAndServe(c.Host, r)
}

func addMsgInKafka(w http.ResponseWriter, r *http.Request) {
	topicName := chi.URLParam(r, "topic_name")

	var newMsgTopic entitie.MsgTopic
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newMsgTopic)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	repository.AddMsgToTopic(topicName, newMsgTopic)

}
