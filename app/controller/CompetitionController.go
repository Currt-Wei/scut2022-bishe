package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scut2022-bishe/app/middleware"
	"scut2022-bishe/app/model"
	"scut2022-bishe/app/service/competition"
	"scut2022-bishe/constant"
)

// CreateCompetition 发布比赛
func CreateCompetition(c *gin.Context) {
	// 验证字段
	var com model.Competition

	if err := c.ShouldBind(&com); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": constant.PostCompetitionFail,
			"msg":    err.Error(),
			"data":   nil,
		})
		return
	}

	// 添加上user的id作为company_id
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	user := model.GetUserByEmail(claims.Email)
	com.CompanyId = user.Id

	err := competition.CreateCompetition(&com)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": constant.PostCompetitionFail,
			"msg":    err.Error(),
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": constant.PostCompetitionSuccess,
		"msg":    "发布成功",
		"data":   nil,
	})
	return
}
