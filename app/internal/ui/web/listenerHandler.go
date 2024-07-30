package web

import (
	"clean/architector/internal/domain/entitie"
	"clean/architector/internal/domain/usecase"
	"encoding/json"
	"fmt"
	"net/http"
)

func StartListenersHandler() {
	fmt.Println("StartListenerHandler - запуск слушателей событий")
	usecase.StartListenersUseCase()
}

func CreateListener(w http.ResponseWriter, r *http.Request) {

	var newListener entitie.ListenerEntitie

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newListener)

	if err != nil {
		http.Error(w, "incorrect Object", http.StatusBadRequest)
		return
	}

	if newListener.Name == "" {
		http.Error(w, "field Name is empty ", http.StatusBadRequest)
		return
	}

	if newListener.Host == "" {
		http.Error(w, "field Host is empty ", http.StatusBadRequest)
		return
	}

	if newListener.Port == "" {
		http.Error(w, "field Port is empty ", http.StatusBadRequest)
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
