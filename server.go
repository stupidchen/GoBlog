package main

import (
	"net/http"
	_ "./model"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: nil,
	}

	server.ListenAndServe()
}