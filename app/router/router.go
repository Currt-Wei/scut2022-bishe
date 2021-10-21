package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"scut2022-bishe/app/controller"
	"scut2022-bishe/util"
)

func InitRouter() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		util.Logger().WithFields(logrus.Fields{
			"name":"wjh",
		}).Infof("测试日志使用","Info")

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/register", controller.Register)

	r.Run()
}
