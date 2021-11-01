package policy

import "github.com/casbin/casbin/v2"

type Permission struct {
	Enforcer *casbin.Enforcer
}
