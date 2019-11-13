package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Shell(c *gin.Context)  {
	//buf := make([]byte, 1024)
	//n, _ := c.Request.Body.Read(buf)
	//fmt.Println(string(buf[0:n]))
	//获取body体里的数据
	//Services.ShellGo("git pull")
	c.HTML(http.StatusOK, "site/index.tmpl", gin.H{
		"title":"aaaa",
	})
}