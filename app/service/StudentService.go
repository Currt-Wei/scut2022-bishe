package service

import (
	"scut2022-bishe/app/middleware"
	"scut2022-bishe/app/model"
	"scut2022-bishe/util"
)

func AddStudent(stu *model.Student){
	result := util.GetDB().Create(&stu)

	if result.Error != nil {
		middleware.Logger().Errorf("注册失败, %s", result.Error)
	}
}