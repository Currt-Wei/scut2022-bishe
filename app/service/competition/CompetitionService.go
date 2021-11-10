package competition

import "scut2022-bishe/app/model"

// CreateCompetition 创建比赛
func CreateCompetition(competition *model.Competition) error {
	return model.AddCompetition(competition)
}
