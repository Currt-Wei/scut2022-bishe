package controller

import (
	"github.com/gin-gonic/gin"
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

	util.GetDB().Create(&u)

}
