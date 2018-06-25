package main

import (
	"net/http"
	_ "blog/model"
)

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
		Handler: nil,
	}

	server.ListenAndServe()
}
