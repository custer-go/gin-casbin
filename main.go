package main

import (
	"github.com/casbin/casbin/v2"
	"log"
)

func main() {
	sub:= "lisi" // 想要访问资源的用户。
	obj:= "/depts" // 将被访问的资源。
	act:= "POST" // 用户对资源执行的操作。
	e,_:= casbin.NewEnforcer("resources/model_t.conf","resources/p_t.csv")

	ok,err:= e.Enforce(sub,"domain2", obj, act)
	if err==nil && ok {
		log.Println("运行通过")
	}
}
