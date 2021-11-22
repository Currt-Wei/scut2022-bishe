package competition

import "scut2022-bishe/app/model"

// AddSignUp 新增报名信息
func AddSignUp(cs *model.CompetitionStudent) error {
	return model.AddSignup(cs)
}

// GetCompetitionStudent 获取报名比赛的所有学生
func GetCompetitionStudent(competitionId int) (cs []*model.CompetitionStudent) {
	return model.GetCompetitionStudent(competitionId)
}
