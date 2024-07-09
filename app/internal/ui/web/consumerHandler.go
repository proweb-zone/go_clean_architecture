package web

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
)

func CreateConsumerHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("init consumerHandler")

	var topicName string = chi.URLParam(r, "topic_name")

	if topicName == "" {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Запустить горутину с именем
	ctx, _ := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "имя_пользователя", topicName)

	var wg sync.WaitGroup
    wg.Add(1)

		// Создать канал для отправки команды завершения
    done := make(chan struct{})

	go createConsumer(ctx, done)

// Завершить горутину с именем через 1 секунду
go func() {
	time.Sleep(1 * time.Second)
	done <- struct{}{}
}()

}

func createConsumer(ctx context.Context, done chan struct {}){

 name := ctx.Value("имя_пользователя").(string)

 fmt.Println(name)
 for {
		 select {
	case <-done:
			fmt.Println("Горутина с именем завершена по команде")
			return
		 default:
				 // Делать что-то
		 }
 }

}

func writeToFile(msg string){
	//prepairMsg := []byte(msg)
// file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, 0600)

// if err != nil {
// 	fmt.Println("err create file", err)
// 	os.Exit(1)
// }

// defer file.Close()
// file.WriteString(msg)
}
