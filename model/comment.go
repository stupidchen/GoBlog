package model

import (
	"net/http"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Author string
	Content string
}

type Comments struct {
	data []Comment
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
	http.HandleFunc("/Comment/", JsonWrapper(&CommentHandler{}))
}
