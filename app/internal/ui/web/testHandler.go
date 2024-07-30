package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"
)

type Job struct {
	Name string
	Ctx context.Context
	Cancel context.CancelFunc
}

var jobs map[string]Job = make(map[string]Job)

func TestHandler(w http.ResponseWriter, r *http.Request){
fmt.Println("start test handler")

names := [...]string{
	"Меркурий",
	"Венера",
	"Земля",
	"Марс",
}

// Создать канал
action := r.URL.Query().Get("action")
topic := r.URL.Query().Get("topic")

if action == "close" {
	fmt.Println(jobs)
	jobs[topic].Cancel()
	return
}

	for _, name := range names {
		ctx, cancel := context.WithCancel(context.Background())

		var job Job
		job.Name = name
		job.Ctx = ctx
		job.Cancel = cancel
		jobs[name] = job

		go startTestGoroutines(ctx, name)
		}
}


func startTestGoroutines(ctx context.Context, name string){
	fmt.Println("start test Goroutine "+name)

	for {
		select {
		case <-ctx.Done():
				fmt.Println("Горутина завершена по сигналу отмены "+name)
				runtime.Goexit()
		default:
			writeToFile(time.Now().Format(time.RFC3339)+" "+name+"\n")
			time.Sleep(3 * time.Second)
		}
}
}

func writeToFile(msg string){
	//prepairMsg := []byte(msg)
file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0600)

if err != nil {
	fmt.Println("err create file", err)
	os.Exit(1)
}

defer file.Close()
file.WriteString(msg)
}
