package controllers

import (
	"github.com/gin-gonic/gin"
	models "go_deploy/Models"
	"go_deploy/Services"
	"net/http"
)

func ProjectPull(c *gin.Context)  {
	id := c.Param("id")
	project := models.ProjectShow(id)
	Services.PullShellGo(project.Directory, "git pull")
	c.JSON(http.StatusOK, gin.H{
		"data":project,
	})
}