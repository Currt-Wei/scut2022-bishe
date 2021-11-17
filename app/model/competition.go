package model

import (
	"gorm.io/gorm"
)

type Competition struct {
	gorm.Model
	Title            string    `gorm:"column:title" json:"title"`
	Description      string    `gorm:"column:description" json:"description""`
	Reward           string    `gorm:"reward" json:"reward"`
	EntryRequirement string    `gorm:"entry_requirement" json:"entry_requirement"`
	WorkRequirement  string    `gorm:"work_requirement" json:"work_requirement"`
	SignupDeadline   LocalTime `gorm:"signup_deadline" json:"signup_deadline"`
	SubmitDeadline   LocalTime `gorm:"submit_deadline" json:"submit_deadline"`

	CompanyId int `json:"company_id"`
}

func (c Competition) TableName() string {
	return "competition"
}

// AddCompetition 创建比赛
func AddCompetition(competition *Competition) error {
	return DB.Create(competition).Error
}

// UpdateCompetition 更新比赛
func UpdateCompetition(competition *Competition) error {
	return DB.Updates(competition).Error
}
