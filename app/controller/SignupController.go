package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path"
	"scut2022-bishe/app/middleware"
	"scut2022-bishe/app/model"
	"scut2022-bishe/app/service"
	"scut2022-bishe/app/service/competition"
	"strconv"
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

// PostWork 上传作品
func PostWork(c *gin.Context) {
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"msg":    err.Error(),
			"data":   nil,
		})
		return
	}

	// 当前的比赛
	competitionId, _ := strconv.Atoi(c.Query("competition_id"))
	// 用户的学号
	claim := c.MustGet("claims").(*middleware.CustomClaims)
	user, _ := service.FindUserByEmail(claim.Email)
	// 创建文件
	basePath := fmt.Sprintf("upload/%d", competitionId)
	if _, err := os.Stat(basePath); err != nil {
		err := os.MkdirAll(basePath, 0777)
		if err != nil {
			log.Println("Error creating directory")
			log.Println(err)
			return
		}
	}
	// 构造存储文件的路径名
	dstFilename := fmt.Sprintf("%s/%s%s", basePath, user.StuNo, path.Ext(file.Filename))
	if err := c.SaveUploadedFile(file, dstFilename); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"msg":    "存储文件失败",
			"data":   err.Error(),
		})
		return
	}
	// 将结果存到数据库中
	cs, err := competition.GetSignUpInfo(competitionId, user.Id)
	// 更新作品链接
	cs.WorkLink = dstFilename
	err = competition.UpdateSignup(cs)
	if err != nil {
		log.Println("更新失败")
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "提交作品成功",
		"data":   nil,
	})
	return
}
