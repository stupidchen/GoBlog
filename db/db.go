package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:Kiminonawa?Taki@tcp(120.55.56.82:32779)/goblog?charset=utf8&&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}

func close() {
	Db.Close()
}
