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
	Config *Config
	Token *map[string]uint
}

var GLOBAL *Global

func connectDatabase(config DatabaseConfig) *gorm.DB {
	connStr := fmt.Sprintf(DefaultDBConnectionFormat, config.Username, config.Password, config.Host, config.Port)
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		GLOBAL.Logger.Fatalf("%s happened while connecting to database. Exit.", err.Error())
	}
	return db
}

func init() {
	config := initConfig()
	if config == nil {
		panic("Cannot init configuration by both file and environment variables. Exit.")
	}

	logFile, err := os.Create(config.Sys.LogFile)
	var logger *log.Logger
	if err != nil {
		fmt.Println(err)
		fmt.Println("Cannot create the log file. Use stdout.")
		logger = log.New(os.Stdout, "", log.LstdFlags | log.Lshortfile)
	} else {
		logger = log.New(logFile, "", log.LstdFlags | log.Lshortfile)
	}

	db := connectDatabase(config.Db)
	token := make(map[string]uint)
	GLOBAL = &Global{
		Config: config,
		Logger: logger,
		DB: db,
		Token: &token,
	}
}