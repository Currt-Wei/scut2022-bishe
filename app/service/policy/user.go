package policy

import (
	"github.com/casbin/casbin/v2"
	"scut2022-bishe/app/middleware/log"
	"scut2022-bishe/app/model"
)

// User 关联的是user-role这张表
type User struct {
	Id     int `json:"id"`      // 用户的id
	RoleId int `json:"role_id"` // 用户的角色id

	Enforcer *casbin.Enforcer // casbin校验的对象
}

// LoadAllPolicy 加载所有的用户-角色关系
func (u *User) LoadAllPolicy() {
	// 加载所有的用户
	users, err := model.GetAllUsers()
	if err != nil {
		return
	}

	// 加载每个用户的角色，加入到casbin中
	for _, user := range users {
		u.LoadPolicy(user.Id)
	}
}

// LoadPolicy 根据user_id加载出该用户的角色
func (u *User) LoadPolicy(id int) {
	user, err := model.GetUserById(id)
	if err != nil {
		return
	}

	// 先删除现有的角色
	_, err = u.Enforcer.DeleteRolesForUser(user.Email)

	if err != nil {
		return
	}

	// 加载每一个该角色的每一个权限
	for _, role := range user.Role {
		_, err := u.Enforcer.AddRoleForUser(user.Email, role.RoleName)
		if err != nil {
			log.Logger().Errorf("[policy]给用户添加角色失败，%s", err)
			return
		}
	}
}
