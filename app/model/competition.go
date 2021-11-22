package model

type Competition struct {
	Id               int       `gorm:"column:id" json:"id"`
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

// GetCompanyCompetitions 查找比赛管理员能看的比赛
func GetCompanyCompetitions(id int) (coms []*Competition, err error) {
	coms = make([]*Competition, 10)
	err = DB.Where("company_id", id).Find(&coms).Error
	return
}

// GetAllCompetitions 查找所有的比赛
func GetAllCompetitions() (coms []*Competition, err error) {
	coms = make([]*Competition, 10)
	err = DB.Find(&coms).Error
	return
}

// GetCompetitionById 查找一个比赛
func GetCompetitionById(id int) (*Competition, error) {
	var com Competition
	err := DB.Preload("Student").First(&com, id).Error
	return &com, err
}
