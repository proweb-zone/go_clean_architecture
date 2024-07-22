package web

import (
	"clean/architector/internal/domain/entitie"
	"clean/architector/internal/domain/repository"
	"clean/architector/internal/domain/usecase"
	"encoding/json"
	"fmt"
	"net/http"
)

func StartListenersHandler(){
fmt.Println("StartListenerHandler - запуск слушателей событий")
var listenerRepo repository.IlistenerRepo = repository.InitListenerRepo()
usecase.StartListenersUseCase(listenerRepo)
}

func CreateListener(w http.ResponseWriter, r *http.Request){
fmt.Println("CreateListener - Добавляем слушателя")

var newListener entitie.ListenerEntitie
decoder := json.NewDecoder(r.Body)

err := decoder.Decode(&newListener)
	if err != nil {
		http.Error(w, "Body MsgTopic not exist", http.StatusBadRequest)
		return
	}

	usecase.AddListenerUseCase(newListener)
}
