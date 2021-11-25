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

// GetStuCompetitions 获取学生参与的所有比赛
func GetStuCompetitions(c *gin.Context) {
	// 获取该用户的信息
	claim := c.MustGet("claims").(*middleware.CustomClaims)
	user, _ := service.FindUserByEmail(claim.Email)
	// 在中间表中查找到所有的比赛id
	competitionIds, err := competition.GetStuCompetitions(user.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"msg":    err.Error(),
			"data":   nil,
		})
		return
	}
	// 遍历比赛，逐个查出来
	var coms []*model.Competition
	for _, id := range competitionIds {
		com, err := competition.GetCompetitionById(int(id))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"msg":    err.Error(),
				"data":   nil,
			})
			return
		}
		// 添加到比赛集合中
		coms = append(coms, com)
	}
	// 将结果返回
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "查询成功",
		"data":   coms,
	})
}
