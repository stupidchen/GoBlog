package view

import (
	"net/http"
	"fmt"
)

type UserHandler struct {
}

func (h *UserHandler) Get(r *http.Request) *ResponseData {
	username := getSubPath(r.URL.Path, 2)
	if username == nil {
		return InitError("User name is missing!")
	}
	return InitHint(fmt.Sprintf("Getting User %s", *username))
}

func (h *UserHandler) Post(r *http.Request, body *RequestData) *ResponseData {
	return nil
}

func (h *UserHandler) Put(r *http.Request, body *RequestData) *ResponseData {
	return nil
}

func (h *UserHandler) Delete(r *http.Request) *ResponseData {
	return nil
}

func init() {
	http.HandleFunc("/user/", SecurityWrapper(JsonWrapper(&UserHandler{})))
}