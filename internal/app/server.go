package app

import (
	"fmt"
	"net/http"

	"clean/architector/internal/ui/web"

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
	r.Post("/kafka/topic/{topic_name}", web.AddMsgToTopicHandler)
	http.ListenAndServe(c.Host, r)
}
