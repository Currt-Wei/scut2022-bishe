package controller

import (
	"github.com/gin-gonic/gin"
	"scut2022-bishe/app/middleware"
	"scut2022-bishe/app/model"
	"scut2022-bishe/util"
)

func Register(c *gin.Context) {
	u := model.Student{
		StuNo:       "201830582036",
		StuName:     "weijiahuan",
		StuCollege:  "cs",
		StuEmail:    "2698230239",
		StuGrade:    "大四",
		StuPassword: "root",
	}

	result := util.GetDB().Create(&u)

	if result.Error != nil {
		middleware.Logger().Errorf("注册失败, %s", result.Error)
	}
}
