package router

import (
	"blog/global"
	"blog/internal/router/route"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func StartServer() {
	gin.SetMode(global.Conf.App.Mode)
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiGroup := router.Group("/api")
	route.RegisterUserRouter(apiGroup)

	log.Fatal(router.Run(fmt.Sprintf(":%s", global.Conf.App.Port)))
}
