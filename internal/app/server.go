package app

import (
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
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	http.ListenAndServe(c.Host, r)
}
