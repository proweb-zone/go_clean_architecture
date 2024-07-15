package app

import (
	"fmt"
	"net/http"

	"clean/architector/internal/ui/web"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

type Server struct {
	Host string
	Cfg *Config
}

func NewServer(cfg *Config) *Server {
	return &Server{Host: cfg.HTTPServer.Address}
}

func (c *Server) StartServer() {
	fmt.Println("start server")

	web.StartListenerHandler()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/kafka/topic/{topic_name}", web.AddMsgToTopicHandler)
	// r.Post("/kafka/consumer/{topic_name}", web.CreateConsumerHandler)
	http.ListenAndServe(c.Host, r)
}
