package model

type CompetitionStudent struct {
	Id            int    `gorm:"column:id" json:"id"`
	CompetitionId int    `gorm:"column:competition_id" json:"competition_id"`
	StudentId     int    `gorm:"column:student_id" json:"student_id""`
	Remark        string `gorm:"remark" json:"remark"`
	Status        string `gorm:"status" json:"status"`
	WorkLink      string `gorm:"work_link" json:"work_link"`
	Score         int    `gorm:"score" json:"score"`
}

type ComStuInfo struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	StuNo      string `json:"stu_no"`
	StuCollege string `json:"stu_college"`
	StuGrade   string `json:"stu_grade"`
	Remark     string `json:"remark"`
	WorkLink   string `json:"work_link"`
	Score      int    `json:"score"`
}

func (c CompetitionStudent) TableName() string {
	return "competition_student"
}

// AddSignup 新增报名信息
func AddSignup(cs *CompetitionStudent) error {
	return DB.Create(cs).Error
}

// GetCompetitionStudent 获取报名该比赛的所有学生的id
func GetCompetitionStudent(competitionId int) (cs []*CompetitionStudent) {
	cs = make([]*CompetitionStudent, 10)
	DB.Where("competition_id", competitionId).Find(&cs)
	return
}

// GetStuCompetitionsId 获取学生参与的所有比赛Id
func GetStuCompetitionsId(stuId int) ([]uint8, error) {
	competitionIds := make([]uint8, 10)
	err := DB.Table("competition_student").Where("student_id", stuId).Select("competition_id").Scan(&competitionIds).Error
	return competitionIds, err
}

// GetSignUpInfo 获取报名详情
func GetSignUpInfo(competitionId int, stuId int) (*CompetitionStudent, error) {
	var cs CompetitionStudent
	err := DB.Where("competition_id", competitionId).Where("student_id", stuId).Find(&cs).Error
	return &cs, err
}

// UpdateSignup 更新报名信息
func UpdateSignup(cs *CompetitionStudent) error {
	return DB.Updates(cs).Error
}
