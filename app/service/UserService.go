package service

import (
	"fmt"
	"scut2022-bishe/app/middleware"
	"scut2022-bishe/app/model"
	"scut2022-bishe/util"
)

func AddUser(user *model.User) (err error){
	db := util.GetDB()

	var user1 *model.User
	db.Where("email = ?", user.Email).First(&user1)
	if user1.Email!="" {
		middleware.Logger().Error("注册失败, 邮箱已注册")
		return fmt.Errorf("注册失败, 邮箱已注册")
	}

	result := db.Create(&user)
	if result.Error != nil {
		middleware.Logger().Errorf("注册失败, %s", result.Error)
		return fmt.Errorf("注册失败, 无法新建用户: %s", result.Error)
	}

	return nil
}

func FindUserByEmail(user *model.User) (error){
	db := util.GetDB()

	var user1 *model.User
	db.Where("email = ?", user.Email).First(&user1)
	if user1.Email=="" {
		middleware.Logger().Error("登陆失败, 用户邮箱未注册")
		return fmt.Errorf("登陆失败, 用户邮箱未注册")
	}

	if user.Password!=user1.Password {
		middleware.Logger().Error("登陆失败, 密码错误")
		return fmt.Errorf("登陆失败, 密码错误")
	}

	return nil
}