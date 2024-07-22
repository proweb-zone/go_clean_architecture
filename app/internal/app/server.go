package app

import (
	"fmt"
	"net/http"

	"clean/architector/internal/ui/web"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/lib/pq"
)

type Server struct {
	IhttServer IHTTPServer
	IDb IDb
	IKafka IKafka
}

func NewServer(config *Config) *Server {
	var iHttpServer IHTTPServer = config.GetHTTPServer()
	var iDb IDb = config.GetConfigDb()
	return &Server{IhttServer: iHttpServer, IDb: iDb}
}

func (s *Server) StartServer() {
	fmt.Println("start server")
	// test.CreateListenerTable() // создаем БД
	web.StartListenersHandler() // запускаем слушатели

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/kafka/topic/{topic_name}", web.AddMsgToTopicHandler)
	r.Post("/kafka/create/listener/", web.CreateListener)
	// r.Post("/kafka/consumer/{topic_name}", web.CreateConsumerHandler)
	http.ListenAndServe(s.IhttServer.GetAddressHttpServer(), r)
}
