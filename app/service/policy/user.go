package permission

import "github.com/casbin/casbin/v2"

// User 关联的是user-role这张表
type User struct {
	Id   int `json:"id"`   // 用户的id
	Role int `json:"role"` // 用户的角色id

	Enforce *casbin.Enforcer // casbin校验的对象
}
