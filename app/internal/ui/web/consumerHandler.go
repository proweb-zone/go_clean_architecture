package web

import (
	"clean/architector/internal/domain/adapter"
	"clean/architector/internal/domain/repository"
	"clean/architector/internal/domain/usecase"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CreateConsumerHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("create consumer handler")

	topicName := chi.URLParam(r, "topic_name")
	action := r.URL.Query().Get("action") // if value = action, closing goroutine this name context

	if topicName == "" {
		http.Error(w, "topic_name not found", http.StatusBadRequest)
		return
	}

	var iConsumerRepo repository.IConsumerRepo = repository.InitConsumerRepo(topicName)
	var iconsumerUseCase adapter.IconsumerUseCase = usecase.InitConsumerUseCase(iConsumerRepo)
	iconsumerUseCase.Run(topicName, action)
	return
}
