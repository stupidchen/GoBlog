package view

import (
	"net/http"
	"fmt"
	"strings"
	"strconv"
	"blog/global"
)


type JsonHandler interface {
	Get(r *http.Request) *ResponseData
	Post(r *http.Request, body *RequestData) *ResponseData
	Put(r *http.Request, body *RequestData) *ResponseData
	Delete(r *http.Request) *ResponseData
}

func getRequestBody(r *http.Request) *string {
	l := r.ContentLength
	buf := make([]byte, l)
	rl, _ := r.Body.Read(buf)
	if rl == 0 {
		return nil
	}
	s := string(buf)
	return &s
}

func getSubPath(path string, index int) *string {
	t := strings.Split(path, "/")
	if len(t) <= index {
		return nil
	} else {
		return &t[index]
	}
}

func SecurityWrapper(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if loginCheck(r) != nil {
			h.ServeHTTP(w, r)
		} else {
			m := InitError("Unauthorized.")
			fmt.Fprint(w, m.ToString())
		}
	}
}

func JsonWrapper(h JsonHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := global.GLOBAL.Logger
		var m *RequestData
		var body *string
		if r.Method == "POST" || r.Method == "PUT" {
			body = getRequestBody(r)
			m = FromString(body)
			if m == nil {
				fmt.Fprint(w, InitError("Cannot read request from client."))
				logger.Panicln("Cannot read request from client.")
				return
			}
		}

		if m != nil {
			logger.Printf("Request %s:%s with body %s is accepted.", r.Method, r.URL, *body)
		} else {
			logger.Printf("Request %s:%s is accepted.", r.Method, r.URL)
		}
		var tm *ResponseData
		switch r.Method {
		case "GET": tm = h.Get(r)
		case "POST": tm = h.Post(r, m)
		case "PUT": tm = h.Put(r, m)
		case "DELETE": tm = h.Delete(r)
		default:
			tm = InitError("Unsupported HTTP method.")
		}
		logger.Printf("Request %s:%s is processed. Return status is %s, Message is %s.",
			r.Method, r.URL, strconv.FormatBool(tm.Ok), tm.Message)
		fmt.Fprint(w, tm.ToString())
	}
}


