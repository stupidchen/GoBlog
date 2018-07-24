package model

import (
	"blog/global"
)

type Object interface {

}

func init() {
	db := global.GLOBAL.DB

	db.AutoMigrate(&Article{}, &Comment{}, &User{})
}