package lib

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var E *casbin.Enforcer

func init() {
	initDB()
	adapter, err := gormadapter.NewAdapterByDB(Gorm)
	if err != nil {
		log.Fatal()
	}
	e, err := casbin.NewEnforcer("resources/model.conf", adapter)
	if err != nil {
		log.Fatal()
	}
	err = e.LoadPolicy()
	if err != nil {
		log.Fatal()
	}
	E = e
	initPolicy()
}

// 从数据库中初始化策略数据
func initPolicy() {
	// E.AddPolicy("member", "/depts", "GET")
	// E.AddPolicy("admin", "/depts", "POST")
	// E.AddRoleForUser("zhangsan", "member")
	m := make([]*RoleRel, 0)
	GetRoles(0, &m, "")
	for _, r := range m {
		_, err := E.AddRoleForUser(r.PRole, r.Role)
		if err != nil {
			log.Fatal(err)
		}
	}
	// 初始化用户角色
	userRoles := GetUserRoles()
	fmt.Println(userRoles)
	for _, user := range userRoles {
		_, err := E.AddRoleForUser(user.UserName, user.RoleName)
		if err != nil {
			log.Fatal(err)
		}
	}
	// 初始化路由角色
	routerRoles := GetRouterRoles()
	for _, rr := range routerRoles {
		_, err := E.AddPolicy(rr.RoleName, rr.RouterUri, rr.RouterMethod)
		if err != nil {
			log.Fatal(err)
		}
	}
}
