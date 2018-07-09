package model

import (
	"net/http"
	"blog/db"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Archive struct {
	gorm.Model
	title string
	content string
	author string
	pubTime string
	editTime string
}

type Archives struct {
	data []Archive
}

type ArchiveHandler struct {
}

func (h *ArchiveHandler) Get (r *http.Request) *Model {
	s2 := getSubPath(r.URL.Path, 2)
	if s2 != nil {
		id, err := strconv.ParseInt(*s2, 10, 64)
		if err != nil {
			return InitError(err.Error())
		}
		var archive Archive
		db.Db.Find(&archive, id)
		return &Model{
			Archives: Archives{
				data: []Archive {archive},
			},
			ok: true,
			modelType: "Archives",
		}

	} else {
		var a Archives
		db.Db.Find(&a.data)
		return &Model{
			Archives: a,
			ok: true,
			modelType: "Archives",
		}
	}
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
	http.HandleFunc("/Archive/", JsonWrapper(&ArchiveHandler{}))
}