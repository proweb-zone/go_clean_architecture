package usecase

import (
	"clean/architector/internal/domain/entitie"
	"clean/architector/internal/domain/repository"
	"fmt"
)

func StartListenersUseCase(listenerRepo repository.IlistenerRepo){
listenerRepo.GetListenerList()

// получаем список слушателей из БД
// Запускаем все слушатели через горутины
}

func AddListenerUseCase(newListener entitie.ListenerEntitie){
	fmt.Println("add listener use case")
	var addListenerRepo repository.IlistenerRepo = repository.InitListenerRepo()
	addListenerRepo.AddListenerDb(newListener)
}
