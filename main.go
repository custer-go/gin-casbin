package main

import (
	"gin-casbin/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(lib.Middlewares()...)

	r.GET("/depts", func(c *gin.Context) {
		c.JSON(200, gin.H{"result": "部门列表"})
	})
	r.POST("/depts", func(c *gin.Context) {
		c.JSON(200, gin.H{"reult": "批量修改部门列表"})
	})
	r.Run(":8080")
}
