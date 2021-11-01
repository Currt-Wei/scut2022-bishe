package permission

import "github.com/casbin/casbin/v2"

// Role 对应的是role_permission这张表啊
type Role struct {
	Id         int `json:"id"`         // 角色id
	Permission int `json:"permission"` // 权限id

	Enforce *casbin.Enforcer
}
