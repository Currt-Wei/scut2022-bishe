package model

import (
	"scut2022-bishe/app/middleware/log"
)

type Role struct {
	Id       int    `json:"id"`
	RoleName string `json:"role_name"`

	Permission []Permission `gorm:"many2many:role_permission;"`
}

func (r Role) TableName() string {
	return "role"
}

// GetRoleById 根据role_id查找角色
func GetRoleById(id int) (*Role, error) {
	var role Role
	err := DB.Preload("Permission").Find(&role, id).Error
	if err != nil {
		log.Logger().Errorf("[policy]根据role_id查找角色失败，%s", err)
		return nil, err
	}
	return &role, nil
}

// GetAllRoles 加载所有的角色
func GetAllRoles() ([]Role, error) {
	roles := make([]Role, 10)
	err := DB.Preload("Permission").Find(&roles).Error
	if err != nil {
		log.Logger().Errorf("[policy]查询所有角色失败，%s", err)
		return nil, err
	}
	return roles, nil
}
