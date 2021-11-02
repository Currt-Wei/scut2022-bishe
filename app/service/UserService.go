package service

import (
	"fmt"
	"scut2022-bishe/app/middleware/log"
	"scut2022-bishe/app/model"
)

func AddUser(user *model.User) (err error) {

	var user1 *model.User
	user1 = model.GetUserByEmail(user.Email)
	if user1 != nil {
		log.Logger().Error("注册失败, 邮箱已注册")
		return fmt.Errorf("注册失败, 邮箱已注册")
	}

	err = model.AddUser(*user)
	if err != nil {
		log.Logger().Errorf("注册失败, %s", err)
		return fmt.Errorf("注册失败, 无法新建用户: %s", err)
	}

	return nil
}

func FindUserByEmail(user *model.User) error {

	var user1 *model.User
	user1 = model.GetUserByEmail(user.Email)
	if user1.Email == "" {
		log.Logger().Error("登陆失败, 用户邮箱未注册")
		return fmt.Errorf("登陆失败, 用户邮箱未注册")
	}

	if user.Password != user1.Password {
		log.Logger().Error("登陆失败, 密码错误")
		return fmt.Errorf("登陆失败, 密码错误")
	}

	return nil
}
