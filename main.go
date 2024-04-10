package main

import (
	"fmt"
	"github.com/WeiDashan/shop-go/cmd"
	"github.com/WeiDashan/shop-go/middleware/rabbitmq"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func test1() {
	dsn := "shop:shop@tcp(49.233.51.52:3306)/shop?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println(db)
	fmt.Println(err)

	e := gin.Default()
	e.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello world!",
		})
	})
	e.Run(":8084")
}

// @title Go-Web开发记录
// @version 0.0.1
// @description golang电商平台后端
func main() {
	defer cmd.Clean()
	cmd.Start()
	rabbitmq.EmailConsumer()
}
