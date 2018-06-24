package model

import (
	"strings"
	"net/http"
	"fmt"
)

type user struct {
	id string
	username string
	email string
	password string
	info string
}

func getSubPath(path string, index int) *string {
	t := strings.Split(path, "/")
	if len(t) < index {
		return nil
	} else {
		return &t[index]
	}
}


type UserHandler struct {
}

func (h *UserHandler) Get(r *http.Request) *Model {
	username := getSubPath(r.URL.Path, 2)
	if username == nil {
		return InitError("User name is missing!")
	}
	return InitHint(fmt.Sprintf("Getting user %s", *username))
}

func (h *UserHandler) Post(r *http.Request, body *Model) *Model {
	return nil
}

func (h *UserHandler) Put(r *http.Request, body *Model) *Model {
	return nil
}

func (h *UserHandler) Delete(r *http.Request) *Model {
	return nil
}

func init() {
	http.HandleFunc("/user/", JsonWrapper(&UserHandler{}))
}