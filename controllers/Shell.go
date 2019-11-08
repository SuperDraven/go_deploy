package controllers

import (
	"github.com/gin-gonic/gin"
	"go_deploy/Services"
)

func Shell(c *gin.Context)  {
	Services.ShellGo("git pull")
}