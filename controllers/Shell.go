package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_deploy/Services"
)

func Shell(c *gin.Context)  {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	fmt.Println(string(buf[0:n]))
	//获取body体里的数据
	Services.ShellGo("git pull")
}