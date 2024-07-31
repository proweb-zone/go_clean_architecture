package web

import (
	"clean/architector/internal/domain/adapter"
	"clean/architector/internal/domain/entitie"
	"clean/architector/internal/domain/repository"
	"clean/architector/internal/domain/usecase"
	"encoding/json"
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

	var iConsumerRepo repository.IConsumerRepo = repository.InitConsumerRepo()
	var iconsumerUseCase adapter.IconsumerUseCase = usecase.InitConsumerUseCase(iConsumerRepo)
	iconsumerUseCase.Run(topicName, action)
	return
}


func CreateConsumer(w http.ResponseWriter, r *http.Request) {

	var newConsumer entitie.ConsumerEntitie

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newConsumer)

	if err != nil {
		http.Error(w, "incorrect Object", http.StatusBadRequest)
		return
	}

	if newConsumer.Name == "" {
		http.Error(w, "field Name is empty ", http.StatusBadRequest)
		return
	}

	if newConsumer.Host == "" {
		http.Error(w, "field Host is empty ", http.StatusBadRequest)
		return
	}

	if newConsumer.Port == "" {
		http.Error(w, "field Port is empty ", http.StatusBadRequest)
		return
	}

	res, err := usecase.AddConsumerUseCase(newConsumer)

	if err != nil {
		fmt.Print(err)
		http.Error(w, "Ошибка создания записи в БД - ", http.StatusBadRequest)
		return
	}

	fmt.Println(res)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "success")

}