package router

import (
	"github.com/gin-gonic/gin"
	"scut2022-bishe/app/controller"
)

func InitRouter() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/register", controller.Register)

	r.Run()
}
