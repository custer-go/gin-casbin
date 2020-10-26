package main

import (
	"log"

	"github.com/casbin/casbin/v2"
)

func main() {
	sub := "putongyonghu" // 想要访问资源的用户
	obj := "/depts"       // 将被访问的资源
	act := "GET"          // 用户对资源执行的操作
	e, err := casbin.NewEnforcer("resources/model.conf", "resources/p.csv")
	if err != nil {
		log.Panic(err)
	}

	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		log.Panic(err)
	}
	if ok {
		log.Println("运行通过")
	} else {
		log.Println("运行不通过")
	}
}
