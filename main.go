package main

import (
	"net/http"
	_ "blog/view"
)

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
		Handler: nil,
	}

	server.ListenAndServe()
}
