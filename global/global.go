package global

import (
	"log"
	"os"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

type Global struct {
	Logger *log.Logger
	DB *gorm.DB

}

var GLOBAL *Global

func connectDatabase() *gorm.DB {
	db, err := gorm.Open("mysql", "root:Kiminonawa?Taki@tcp(120.55.56.82:32779)/goblog?charset=utf8&&parseTime=True&loc=Local")
	if err != nil {
		GLOBAL.Logger.Fatalf("%s happened while connecting to database. Exit.", err.Error())
	}
	return db
}

func init() {
	logFile, err := os.Create("/var/log/blog.log")
	var logger *log.Logger
	if err != nil {
		fmt.Println(err)
		fmt.Println("Cannot create the log file. Use stdout.")
		logger = log.New(os.Stdout, "", log.LstdFlags | log.Lshortfile)
	} else {
		logger = log.New(logFile, "", log.LstdFlags | log.Lshortfile)
	}

	GLOBAL = &Global{Logger:logger}
	db := connectDatabase()
	GLOBAL.DB = db
}