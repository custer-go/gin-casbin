package lib

import (
	"log"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("token") == "" {
			c.AbortWithStatusJSON(400, gin.H{"message": "token required"})
		} else {
			c.Set("user_name", c.Request.Header.Get("token"))
			c.Next()
		}
	}
}

// 把 p.csv 持久化到数据库中
func RBAC() gin.HandlerFunc {
	adap, err := gormadapter.NewAdapterByDB(Gorm) // 数据库中没有就会创建表 casbin_rules
	if err != nil {
		log.Fatal(err)
	}
	e, err := casbin.NewEnforcer("resources/model.conf", adap)
	if err != nil {
		log.Fatal(err)
	}
	err = e.LoadPolicy()
	if err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		user, _ := c.Get("user_name")
		access, err := e.Enforce(user, c.Request.RequestURI, c.Request.Method)
		if err != nil || !access {
			c.AbortWithStatusJSON(403, gin.H{"message": "forbidden"})
		} else {
			c.Next()
		}
	}
}

func Middlewares() (fs []gin.HandlerFunc) {
	fs = append(fs, CheckLogin(), RBAC())
	return
}
