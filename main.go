package main

import (
	"net/http"
	_ "app/model"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: nil,
	}

	server.ListenAndServe()
}
