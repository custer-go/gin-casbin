package lib

import (
	"log"

	"github.com/casbin/casbin/v2"
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

func RBAC() gin.HandlerFunc {
	e, err := casbin.NewEnforcer("resources/model.conf", "resources/p.csv")
	if err != nil {
		log.Panic(err)
	}
	return func(c *gin.Context) {
		user, _ := c.Get("user_name")
		ok, err := e.Enforce(user, c.Request.RequestURI, c.Request.Method)
		if err != nil {
			log.Panic(err)
		}
		if !ok {
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
