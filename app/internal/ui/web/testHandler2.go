package web

import (
	"fmt"
	"net/http"
)

func TestHandler2(w http.ResponseWriter, r *http.Request){
	fmt.Println("сигнал отмены принято action close")
	// TODO
	//cancel()
}
