package casbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"runtime"
	"scut2022-bishe/app/middleware"
	"scut2022-bishe/app/service/policy"
)

type CasbinObject struct {
	UserAPI       policy.User
	RoleAPI       policy.Role
	PermissionAPI policy.Permission

	Enforcer *casbin.Enforcer
}

var CasbinObj CasbinObject

// Init 初始化CasbinObj这个变量
func init() {
	osType := runtime.GOOS // 运行时操作系统
	var path string        // 配置文件的位置
	if osType == "windows" {
		path = "config\\rbac_models.conf"
	} else if osType == "linux" || osType == "darwin" {
		path = "config/rbac_models.conf"
	}

	enforcer, err := casbin.NewEnforcer(path)
	if err != nil {
		middleware.Logger().Errorf("[policy]加载casbin策略出错，%s", err)
		return
	}

	// 初始化每个的Enforcer都为同一个
	userApi := policy.User{Enforcer: enforcer}
	roleApi := policy.Role{Enforcer: enforcer}
	permissionApi := policy.Permission{Enforcer: enforcer}

	// 初始化这个变量为同一个
	CasbinObj = CasbinObject{
		UserAPI:       userApi,
		RoleAPI:       roleApi,
		PermissionAPI: permissionApi,
		Enforcer:      enforcer,
	}

	return
}

// InitCasbinPolicyData 从数据库中加载casbin策略
func InitCasbinPolicyData() {
	// 加载所有的角色-权限关系
	CasbinObj.RoleAPI.LoadAllPolicy()
	// 加载所有的用户-角色关系
	CasbinObj.UserAPI.LoadAllPolicy()
	CheckAllPolicy()
}

func CheckAllPolicy() {
	list := CasbinObj.Enforcer.GetPolicy()
	for _, vlist := range list {
		for _, v := range vlist {
			fmt.Printf("value: %s, ", v)
		}
		fmt.Println()
	}
	ok, _ := CasbinObj.Enforcer.Enforce("zhangsan@qq.com", "/api/v1/setting/permission", "GET")
	fmt.Println(ok)
}
