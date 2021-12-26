package rbac

import (
	"github.com/casbin/casbin/v2"
)

var (
	e   *casbin.Enforcer
	err error
)

func init() {
	e, err = casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
	}
}
