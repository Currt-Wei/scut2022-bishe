package router

import (
	"github.com/gin-gonic/gin"
	"scut2022-bishe/app/controller"
	"scut2022-bishe/app/middleware"
)

func InitRouter() {
	r := gin.New()
	// 使用自定义的日志中间件
	r.Use(middleware.LoggerToFile())

	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/register", controller.Register)

	err := r.Run()
	if err != nil {
		middleware.Logger().Errorf("路由初始化失败, %s", err)
		return
	}
}