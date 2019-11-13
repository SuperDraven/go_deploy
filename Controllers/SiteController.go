package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Site(c *gin.Context)  {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":"site",
	})
}