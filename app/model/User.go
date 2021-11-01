package model

import "scut2022-bishe/app/middleware"

type User struct {
	Id         int    `gorm:"column:id" form:"id"`
	Name       string `gorm:"column:name" form:"name"`
	Email      string `gorm:"column:email" form:"email"`
	Password   string `gorm:"column:password" form:"password"`
	StuNo      string `gorm:"column:stu_no" form:"stu_no"`
	StuCollege string `gorm:"column:stu_college" form:"stu_college"`
	StuGrade   string `gorm:"column:stu_grade" form:"stu_grade"`
}

func (S User) TableName() string {
	return "users"
}

// FindUserByEmail 根据email查找用户
func FindUserByEmail() *User {
	var user User
	err := DB.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		middleware.Logger().Errorf("[user]根据email查找用户失败，%s", err)
		return nil
	}
	return &user
}

// AddUser 添加用户
func AddUser(user User) error {
	return DB.Create(&user).Error
}
