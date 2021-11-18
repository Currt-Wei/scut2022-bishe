package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scut2022-bishe/app/middleware"
	"scut2022-bishe/app/model"
	"scut2022-bishe/app/service"
	"scut2022-bishe/app/service/competition"
	"scut2022-bishe/constant"
	"strconv"
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
	user, _ := model.GetUserByEmail(claims.Email)
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
		"data":   com,
	})
	return
}

// UpdateCompetition 更新比赛
func UpdateCompetition(c *gin.Context) {
	// 验证字段
	var com model.Competition

	if err := c.ShouldBind(&com); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": constant.UpdateCompetitionFail,
			"msg":    err.Error(),
			"data":   nil,
		})
		return
	}

	// 添加上user的id作为company_id
	id, _ := strconv.Atoi(c.Param("competition_id"))
	com.Id = id

	err := competition.UpdateCompetition(&com)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": constant.UpdateCompetitionFail,
			"msg":    err.Error(),
			"data":   com,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": constant.UpdateCompetitionSuccess,
		"msg":    "更新比赛成功",
		"data":   com,
	})
	return
}

// GetCompanyCompetition 查询比赛
func GetCompanyCompetition(c *gin.Context) {
	// 获取competition_id
	idStr := c.Query("competition_id")
	id, _ := strconv.Atoi(idStr)

	// 根据id查找比赛
	com, err := competition.GetCompetitionById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    "查询一个比赛出错",
			"data":   nil,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    "查询一个比赛成功",
			"data":   com,
		})
		return
	}

}

// GetCompetitionList 获取比赛列表
func GetCompetitionList(c *gin.Context) {
	// 能进到这个界面的是管理员
	// 获取当前的登录用户
	claim := c.MustGet("claims").(*middleware.CustomClaims)
	user, err := service.FindUserByEmail(claim.Email)

	// 当前登录用户的角色
	role := user.Role[0].RoleName
	// 接收结果的数据集
	var coms = make([]*model.Competition, 10)
	// 超级管理员
	if role == "admin" {
		coms, err = competition.GetAllCompetitions()
	} else if role == "competition_manager" {
		coms, err = competition.GetCompanyCompetitions(user.Id)
	}

	// 数据集返回
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    "查询失败",
			"data":   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    "查询成功",
			"data":   coms,
		})
	}
}

func GetCompetitionListByUser(c *gin.Context) {
	coms, err := competition.GetAllCompetitions()
	// 数据集返回
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    "查询失败",
			"data":   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    "查询成功",
			"data":   coms,
		})
	}
}
