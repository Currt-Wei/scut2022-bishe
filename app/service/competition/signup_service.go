package competition

import "scut2022-bishe/app/model"

// AddSignUp 新增报名信息
func AddSignUp(cs *model.CompetitionStudent) error {
	return model.AddSignup(cs)
}
