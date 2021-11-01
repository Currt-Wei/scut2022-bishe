package service

import (
	"fmt"
	"scut2022-bishe/app/middleware"
	"scut2022-bishe/app/model"
)

func AddUser(user *model.User) (err error) {

	var user1 *model.User
	user1 = model.FindUserByEmail()
	if user1.Email != "" {
		middleware.Logger().Error("注册失败, 邮箱已注册")
		return fmt.Errorf("注册失败, 邮箱已注册")
	}

	err = AddUser(user1)
	if err != nil {
		middleware.Logger().Errorf("注册失败, %s", err)
		return fmt.Errorf("注册失败, 无法新建用户: %s", err)
	}

	return nil
}

func FindUserByEmail(user *model.User) error {

	var user1 *model.User
	user1 = model.FindUserByEmail()
	if user1.Email == "" {
		middleware.Logger().Error("登陆失败, 用户邮箱未注册")
		return fmt.Errorf("登陆失败, 用户邮箱未注册")
	}

	if user.Password != user1.Password {
		middleware.Logger().Error("登陆失败, 密码错误")
		return fmt.Errorf("登陆失败, 密码错误")
	}

	return nil
}
