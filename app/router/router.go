package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	// 使用自定义的jwt认证
	r.Use(middleware.JWTAuth())
	// 权限验证
	r.Use(middleware.Authorize())

	apiV1 := r.Group("/api/v1")
	{
		// 比赛模块
		// 管理员接口
		apiV1.POST("/setting/competition", controller.CreateCompetition)                // 创建比赛
		apiV1.PUT("/setting/competition/:competition_id", controller.UpdateCompetition) // 更新比赛
		apiV1.GET("/setting/competition", controller.GetCompanyCompetition)
		apiV1.GET("/setting/competition/get-list", controller.GetCompetitionList)
		// 普通用户接口
		apiV1.GET("/user/competition/get-list", controller.GetCompetitionListByUser)
		apiV1.GET("/user/competition", controller.GetCompanyCompetition)

		// 权限管理
		apiV1.GET("/setting/permission", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "请求权限管理成功",
			})
		})
		// 普通用户的
		apiV1.GET("/index", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "请求普通用户接口成功",
			})
		})

	}

	err := r.Run()
	if err != nil {
		log.Logger().Errorf("路由初始化失败, %s", err)
		return
	}
}
