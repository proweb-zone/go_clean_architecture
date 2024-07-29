package usecase

import (
	"clean/architector/internal/domain/entitie"
	"clean/architector/internal/domain/repository"
	"fmt"
)

func StartListenersUseCase() {
	var listenerRepo repository.IlistenerRepo = repository.InitListenerRepo()
	var listenerList []*entitie.ListenerEntitie = listenerRepo.GetListenerList()

	for _, item := range listenerList {
		fmt.Printf("%s, \n", item.Settings)
	}

	// получаем список слушателей из БД
	// Запускаем все слушатели через горутины
}

func AddListenerUseCase(newListener entitie.ListenerEntitie) (bool, error) {
	var addListenerRepo repository.IlistenerRepo = repository.InitListenerRepo()
	res, err := addListenerRepo.AddListenerDb(newListener)

	if err != nil {
		return false, fmt.Errorf("Ошибка записи в БД")
	}

	return res, nil
}
