package model

type CompetitionStudent struct {
	Id            int       `gorm:"column:id" json:"id"`
	CompetitionId int       `gorm:"column:competition_id" json:"competition_id"`
	StudentId     int       `gorm:"column:student_id" json:"student_id""`
	Remark        string    `gorm:"remark" json:"remark"`
	Status        string    `gorm:"status" json:"status"`
	WorkLink      string    `gorm:"work_link" json:"work_link"`
	Score         LocalTime `gorm:"score" json:"score"`
}

func (c CompetitionStudent) TableName() string {
	return "competition_student"
}

// AddSignup 新增报名信息
func AddSignup(cs *CompetitionStudent) error {
	return DB.Create(cs).Error
}
