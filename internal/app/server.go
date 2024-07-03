package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	Host string
}

func NewServer(cfg *Config) *Server {
	return &Server{Host: cfg.Address}
}

func (c *Server) StartServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/kafka/topic/{topic_name}", addMsgInKafka)

	http.ListenAndServe(c.Host, r)
}

func addMsgInKafka(w http.ResponseWriter, r *http.Request) {
	// topicName := chi.URLParam(r, "topic_name")

	var data Data
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Println(data)

}

type Data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
