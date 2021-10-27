package model

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
