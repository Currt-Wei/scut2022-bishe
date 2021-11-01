package policy

import "github.com/casbin/casbin/v2"

// User 关联的是user-role这张表
type User struct {
	Id     int `json:"id"`      // 用户的id
	RoleId int `json:"role_id"` // 用户的角色id

	Enforcer *casbin.Enforcer // casbin校验的对象
}
