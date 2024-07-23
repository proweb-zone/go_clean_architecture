package web

import (
	"clean/architector/internal/domain/entitie"
	"clean/architector/internal/domain/repository"
	"clean/architector/internal/domain/usecase"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

func StartListenersHandler() {
	fmt.Println("StartListenerHandler - запуск слушателей событий")
	var listenerRepo repository.IlistenerRepo = repository.InitListenerRepo()
	usecase.StartListenersUseCase(listenerRepo)
}

func CreateListener(w http.ResponseWriter, r *http.Request) {

	var newListener entitie.ListenerEntitie

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newListener)

	if err != nil {
		http.Error(w, "Body MsgTopic not exist", http.StatusBadRequest)
		return
	}

	if newListener.Name == "" {
		http.Error(w, "field Name is empty ", http.StatusBadRequest)
		return
	}

	empty := reflect.ValueOf(newListener.Settings).IsZero()
	if empty {
		http.Error(w, "field Settings is empty ", http.StatusBadRequest)
		return
	}

	if newListener.Settings.Host == "" {
		http.Error(w, "field Host is empty ", http.StatusBadRequest)
		return
	}

	res, err := usecase.AddListenerUseCase(newListener)

	if err != nil {
		http.Error(w, "Ошибка создания записи в БД - ", http.StatusBadRequest)
		return
	}

	fmt.Println(res)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "success")

}
