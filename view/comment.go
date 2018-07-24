package view

import "net/http"

type CommentHandler struct {
}

func (h *CommentHandler) Get (r *http.Request) *ResponseData {
	return nil
}

func (h *CommentHandler) Post (r *http.Request, body *RequestData) *ResponseData {
	return nil
}

func (h *CommentHandler) Put (r *http.Request, body *RequestData) *ResponseData {
	return nil
}

func (h *CommentHandler) Delete (r *http.Request) *ResponseData {
	return nil
}

func init() {
	http.HandleFunc("/Comment/", JsonWrapper(&CommentHandler{}))
}
