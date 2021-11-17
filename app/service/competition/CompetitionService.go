package competition

import "scut2022-bishe/app/model"

// CreateCompetition 创建比赛
func CreateCompetition(competition *model.Competition) error {
	return model.AddCompetition(competition)
}

// UpdateCompetition 更新比赛
func UpdateCompetition(competition *model.Competition) error {
	return model.UpdateCompetition(competition)
}

// GetCompanyCompetitions 查找比赛管理员能看的所有比赛
func GetCompanyCompetitions(id int) ([]*model.Competition, error) {
	return model.GetCompanyCompetitions(id)
}

// GetAllCompetitions 查找所有比赛
func GetAllCompetitions() ([]*model.Competition, error) {
	return model.GetAllCompetitions()
}

func GetCompetitionById(id int) ([]*model.Competition, error) {
	return model.GetCompetitionById(id)
}
