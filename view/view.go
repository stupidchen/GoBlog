package view

import (
	"net/http"
	"fmt"
	"strconv"
	"blog/global"
)


type JsonHandler interface {
	Get(r *http.Request) *ResponseData
	Post(r *http.Request, body *RequestData) *ResponseData
	Put(r *http.Request, body *RequestData) *ResponseData
	Delete(r *http.Request) *ResponseData
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
			if body == nil {
				fmt.Fprint(w, InitError("Cannot read request from client.").ToString())
				logger.Printf("Cannot read request from client.")
				return
			}
			m = FromString(body)
			if m == nil {
				fmt.Fprint(w, InitError("Cannot marshal the request body.").ToString())
				logger.Printf("Cannot marshal the request body %s.", *body)
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


