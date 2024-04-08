package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte("Welcome to shortner!"))
		if err != nil {
			return
		}
	})
	//nolint: gosec // временно отключаем проверку сервера без таймаутов
	log.Fatal(http.ListenAndServe(":8081", mux))
}
