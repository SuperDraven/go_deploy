package Controllers

import (
	"github.com/gin-gonic/gin"
	models "go_deploy/Models"
	"net/http"
)

func Site(c *gin.Context)  {
	models.CreateTabelProject()
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":"site",
	})
}