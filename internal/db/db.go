package db

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags),
	logger.Config{
		LogLevel: logger.Info,
	},
)

var DB *gorm.DB

func init() {
	res, err := gorm.Open(mysql.Open("root:NYSY2nysy@mysql@tcp(127.0.0.1:3306)/day10?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Error connecting to database:%s", err.Error())
	}
	DB = res
}
