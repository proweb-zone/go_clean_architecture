package usecase

import (
	"clean/architector/internal/domain/repository"
)

func StartListenersUseCase(listenerRepo repository.IlistenerRepo){
listenerRepo.GetListenerList()

// получаем список слушателей из БД
// Запускаем все слушатели через горутины
}
