package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Context struct {
	Port string
	Host string
}

func NewContext() *Context {

	os.Setenv("PORT", "3000")
	os.Setenv("HOST", "localhost")

	return &Context{
		Port: os.Getenv("PORT"),
		Host: os.Getenv("HOST"),
	}
}

func (c *Context) StartWebServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	var host string = fmt.Sprintf(c.Host+":"+c.Port)
	fmt.Println(host)
	http.ListenAndServe(host, r)
}
