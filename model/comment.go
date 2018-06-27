package model

import "net/http"

type comment struct {
	id string
	author string
	pubTime string
	content string
}

type CommentHandler struct {
}

func (h *CommentHandler) Get (r *http.Request) *Model {
	return nil
}

func (h *CommentHandler) Post (r *http.Request, body *Model) *Model {
	return nil
}

func (h *CommentHandler) Put (r *http.Request, body *Model) *Model {
	return nil
}

func (h *CommentHandler) Delete (r *http.Request) *Model {
	return nil
}

func init() {
	http.HandleFunc("/comment/", JsonWrapper(&CommentHandler{}))
}
