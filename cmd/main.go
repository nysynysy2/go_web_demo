package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Result struct {
	code    int
	message string
	data    interface{}
}

func RSuccess(data interface{}) Result {
	return Result{
		code:    1,
		message: "success",
		data:    data,
	}
}

func RError(data interface{}) Result {
	return Result{
		code:    0,
		message: "error",
		data:    data,
	}
}

type Dept struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
}

var db *gorm.DB

func init() {
	res, err := gorm.Open(mysql.Open("root:NYSY2nysy%40@mysql!tcp(127.0.0.1:3306)/day10?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database:%e", err)
	}
	db = res
}

func main() {
	r := gin.Default()

	dr := r.Group("/dept")
	dr.GET("", func(ctx *gin.Context) {
		depts := []Dept{}
		res := db.Table("dept").Find(&depts)
		if res.Error != nil {
			ctx.JSON(500, RError(res.Error))
			return
		}
		ctx.JSON(200, RSuccess(depts))
	})

	r.Run(":8080")
}
