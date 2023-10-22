package main

import (
	"fmt"
	"kallme/api"
	"kallme/config"
	"kallme/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	dao.InitMongoDB()

	router.Static("/dist", "./dist")

	// 测试路由
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.POST("/login", func(c *gin.Context) { api.Login(c) })
	router.POST("/register", func(c *gin.Context) { api.Register(c) })

	// 启动服务器
	router.Run(":" + config.Config.Port)
}

func init() {
	fmt.Println("Server is running: http://localhost:" + config.Config.Port)
}
