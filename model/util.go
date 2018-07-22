package model

import (
	"net/http"
	"fmt"
	"strings"
	"log"
	"os"
	"strconv"
)

type Global struct {
	logger *log.Logger
}

var GLOBAL *Global

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
		logger := GLOBAL.logger
		var m *Model
		if r.Method == "POST" || r.Method == "PUT" {
			m = FromString(getRequestBody(r))
			fmt.Fprint(w, "")
			logger.Panicln("Cannot read request from client")
			return
		}

		if m != nil {
			logger.Printf("Request %s:%s with body %s is accepted.", r.Method, r.URL, *m)
		} else {
			logger.Printf("Request %s:%s is accepted.", r.Method, r.URL)
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
		logger.Printf("Request %s:%s is processed. Return status is %s, message is %s",
			r.Method, r.URL, strconv.FormatBool(tm.ok), tm.message)
		fmt.Fprint(w, tm.ToString())
	}
}

func init() {
	logFile, err := os.Create("/var/log/blog.log")
	var logger *log.Logger
	if err != nil {
		fmt.Println(err)
		fmt.Println("Cannot create the log file. Use stdout.")
		logger = log.New(os.Stdout, "", log.LstdFlags | log.Lshortfile)
	} else {
		logger = log.New(logFile, "", log.LstdFlags | log.Lshortfile)
	}

	GLOBAL = &Global{logger:logger}
}
