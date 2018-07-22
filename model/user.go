package model

import (
	"net/http"
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null"`
	Email string `gorm:"not null"`
	Password string `gorm:"not null"`
	Info string `gorm:"type:text"`
}

type Users struct {
	data []User
}

type UserHandler struct {
}

func (h *UserHandler) Get(r *http.Request) *Model {
	username := getSubPath(r.URL.Path, 2)
	if username == nil {
		return InitError("User name is missing!")
	}
	return InitHint(fmt.Sprintf("Getting User %s", *username))
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
	http.HandleFunc("/User/", JsonWrapper(&UserHandler{}))
}
