package main

import (
	"net/http"
	_ "blog/view"
)

func initServer() {
	fs := http.FileServer(http.Dir("static/"))
	server := http.Server{
		Addr: "0.0.0.0:8080",
		Handler: nil,
	}

	http.Handle("/static/", http.StripPrefix("/static", fs))
	server.ListenAndServe()
}

func main() {
	initServer()
}
