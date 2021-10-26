package model

type Student struct {
	Id          int    `gorm:"column:id" form:"id"'`
	StuNo       string `gorm:"column:stu_no" form:"stu_no"`
	StuName     string `gorm:"column:stu_name" form:"stu_name"`
	StuCollege  string `gorm:"column:stu_college" from:"stu_college"`
	StuGrade    string `gorm:"column:stu_grade" form:"stu_grade"`
	StuEmail    string `gorm:"column:stu_email" form:"stu_email"`
	StuPassword string `gorm:"column:stu_password" form:"stu_password"`
}

func (S Student) TableName() string {
	return "student"
}
