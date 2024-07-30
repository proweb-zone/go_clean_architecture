package usecase

import (
	"clean/architector/internal/domain/entitie"
	"clean/architector/internal/domain/repository"
	"context"
	"fmt"
	"runtime"
	"time"
)

func StartListenersUseCase() {
	var listenerRepo repository.IlistenerRepo = repository.InitListenerRepo()
	var listenerList []*entitie.ListenerEntitie = listenerRepo.GetListenerList()

	for _, item := range listenerList {
		fmt.Printf("%s, %s, %s, %s, \n", item.Name, item.Host, item.Port, item.Status)

		ctx, _ := context.WithCancel(context.Background())
		go startListenerService(ctx, item)
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

// функция для создания внешних слушателей
func startListenerService(ctx context.Context, item *entitie.ListenerEntitie) {
	fmt.Println("start test Goroutine " + item.Name)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина завершена по сигналу отмены " + item.Name)
			runtime.Goexit()
		default:
			writeToFile(time.Now().Format(time.RFC3339) + " " + item.Name + "\n")
			time.Sleep(time.Duration(item.Deelay) * time.Second)
			fmt.Println("сигнал от горутины " + item.Name)
		}
	}

}
