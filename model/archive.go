package model

import "net/http"

type archive struct {
	id string
	title string
	content string
	author string
	pubTime string
	editTime string
}

type ArchiveHandler struct {
}

func (h *ArchiveHandler) Get (r *http.Request) *Model {
	return nil
}

func (h *ArchiveHandler) Post (r *http.Request, body *Model) *Model {
	return nil
}

func (h *ArchiveHandler) Put (r *http.Request, body *Model) *Model {
	return nil
}

func (h *ArchiveHandler) Delete (r *http.Request) *Model {
	return nil
}

func init() {
	http.HandleFunc("/archive/", JsonWrapper(&ArchiveHandler{}))
}