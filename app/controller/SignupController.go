package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scut2022-bishe/app/middleware"
	"scut2022-bishe/app/model"
	"scut2022-bishe/app/service"
	"scut2022-bishe/app/service/competition"
)

func Signup(c *gin.Context) {
	var cs model.CompetitionStudent
	//获取当前用户id
	claim := c.MustGet("claims").(*middleware.CustomClaims)
	user, err := service.FindUserByEmail(claim.Email)
	if err != nil {
		return
	}
	cs.StudentId = user.Id
	cs.Status = "signup"

	// 获取比赛id和备注信息
	if err := c.ShouldBindJSON(&cs); err != nil {
		return
	}

	err = competition.AddSignUp(&cs)
	if err != nil {
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    "报名成功",
			"data":   nil,
		})
	}
}
