package router

import (
	"github.com/gin-gonic/gin"
	"scut2022-bishe/app/controller"
	"scut2022-bishe/app/middleware"
	"scut2022-bishe/app/middleware/log"
)

func InitRouter() {
	r := gin.New()
	// 使用自定义的日志中间件
	r.Use(log.LoggerToFile())
	// 默认跨域
	r.Use(middleware.Cors())

	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	// 使用自定义的jwt认证
	r.Use(middleware.JWTAuth())
	// 权限验证
	r.Use(middleware.Authorize())

	g1 := r.Group("/test")
	g1.Use(middleware.JWTAuth())
	{
		g1.POST("/testtoken", controller.TestToken)
	}

	err := r.Run()
	if err != nil {
		log.Logger().Errorf("路由初始化失败, %s", err)
		return
	}
}
