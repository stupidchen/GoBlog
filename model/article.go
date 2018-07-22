package model

import (
	"net/http"
	"blog/db"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Article struct {
	gorm.Model
	Title string `gorm:"not null"`
	Content string `gorm:"type:text"`
	Author string `gorm:"not null"`
}

type Articles struct {
	data []Article
}

type ArticleHandler struct {
}

func (h *ArticleHandler) Get (r *http.Request) *Model {
	s2 := getSubPath(r.URL.Path, 2)
	if s2 != nil {
		id, err := strconv.ParseInt(*s2, 10, 64)
		if err != nil {
			return InitError(err.Error())
		}
		var article Article
		db.Db.Find(&article, id)
		return &Model{
			Articles: Articles{
				data: []Article{article},
			},
			ok: true,
			modelType: "Articles",
		}

	} else {
		var a Articles
		db.Db.Find(&a.data)
		return &Model{
			Articles:  a,
			ok:        true,
			modelType: "Articles",
		}
	}
	return nil
}

func (h *ArticleHandler) Post (r *http.Request, body *Model) *Model {
	return nil
}

func (h *ArticleHandler) Put (r *http.Request, body *Model) *Model {
	return nil
}

func (h *ArticleHandler) Delete (r *http.Request) *Model {
	return nil
}

func init() {
	http.HandleFunc("/Article/", JsonWrapper(&ArticleHandler{}))
}