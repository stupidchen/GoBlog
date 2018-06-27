package model

import (
	"net/http"
	"fmt"
	"strings"
)

type JsonHandler interface {
	Get(r *http.Request) *Model
	Post(r *http.Request, body *Model) *Model
	Put(r *http.Request, body *Model) *Model
	Delete(r *http.Request) *Model
}

func getRequestBody(r *http.Request) *string {
	l := r.ContentLength
	buf := make([]byte, l)
	_, err := r.Body.Read(buf)
	if err != nil {
		return nil
	}
	s := string(buf)
	return &s
}

func getSubPath(path string, index int) *string {
	t := strings.Split(path, "/")
	if len(t) < index {
		return nil
	} else {
		return &t[index]
	}
}

func JsonWrapper(h JsonHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var m *Model
		if r.Method == "POST" || r.Method == "PUT" {
			m = FromString(getRequestBody(r))
			fmt.Fprint(w, "Cannot read request from client")
		}

		var tm *Model
		switch r.Method {
		case "GET": tm = h.Get(r)
		case "POST": tm = h.Post(r, m)
		case "PUT": tm = h.Put(r, m)
		case "DELETE": tm = h.Delete(r)
		default:
			tm = InitError("Unsupported HTTP method")
		}
		fmt.Fprint(w, tm.ToString())
	}
}
