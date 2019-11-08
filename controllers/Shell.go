package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_deploy/Services"
)

func Shell(c *gin.Context)  {
	for k, v :=range c.Request.PostForm {
		fmt.Printf("k:%v\n", k)
		fmt.Printf("v:%v\n", v)
	}
	Services.ShellGo("git pull")
}