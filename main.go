package main

import (
	"kallme/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/dist", "./dist")

	// 测试路由
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 启动服务器
	router.Run(":" + config.Config.Port)
}
