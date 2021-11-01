package policy

import (
	"github.com/casbin/casbin/v2"
	"scut2022-bishe/app/middleware"
	"scut2022-bishe/app/model"
)

// Role 对应的是role_permission这张表啊
type Role struct {
	Id           int `json:"id"`            // 角色id
	PermissionId int `json:"permission_id"` // 权限id

	Enforcer *casbin.Enforcer
}

// LoadAllPolicy 加载所有的角色-权限策略
func (r *Role) LoadAllPolicy() {
	// 加载所有的角色
	roles, err := model.GetAllRoles()
	if err != nil {
		return
	}

	// 加载每个角色的权限，加入到casbin中
	for _, role := range roles {
		r.LoadPolicy(role.Id)
	}
}

// LoadPolicy 根据role_id加载角色的permission
func (r *Role) LoadPolicy(id int) {
	role, err := model.GetRoleById(id)
	if err != nil {
		return
	}

	// 先删除现有的角色
	_, err = r.Enforcer.DeleteRole(role.RoleName)

	if err != nil {
		return
	}

	// 加载每一个该角色的每一个权限
	for _, permission := range role.Permission {
		_, err := r.Enforcer.AddPermissionForUser(role.RoleName, permission.Url, permission.Method)
		if err != nil {
			middleware.Logger().Errorf("[policy]给角色添加权限失败，%s", err)
			return
		}
	}
}
