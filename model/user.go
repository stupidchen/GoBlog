package model

import (
	"github.com/jinzhu/gorm"
	"blog/global"
)

type User struct {
	Object `json:"object,omitempty"`
	gorm.Model
	Username string `json:"username",gorm:"not null;unique_index"`
	Email string `json:"email",gorm:"not null;unique_index"`
	Password string `json:"password",gorm:"not null"`
	Info string `json:"info",gorm:"type:text"`
}

func FindUserById(id uint) *User {
	db := global.GLOBAL.DB
	var a User
	if db.Find(&a, id).RecordNotFound() {
		return nil
	}
	return &a
}

func FindUser(user *User) *User {
	db := global.GLOBAL.DB
	var users []User
	if db.Where(user).Find(&users).RecordNotFound() || len(users) == 0 {
		return nil
	}
	return &users[0]
}

func AddUser(user *User) error {
	db := global.GLOBAL.DB
	return db.Create(user).Error
}