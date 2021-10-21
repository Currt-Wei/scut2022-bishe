package model

type Student struct {
	Id          int    `gorm:"column:id"`
	StuNo       string `gorm:"column:stu_no"`
	StuName     string `gorm:"column:stu_name"`
	StuCollege  string `gorm:"column:stu_college"`
	StuGrade    string `gorm:"column:stu_grade"`
	StuEmail    string `gorm:"column:stu_email"`
	StuPassword string `gorm:"column:stu_password"`
}

func (S Student) TableName() string {
	return "student"
}
