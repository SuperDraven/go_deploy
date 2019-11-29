package Routers

import (
	"github.com/gin-gonic/gin"
	"go_deploy/Controllers"
)
//测试
func SetRouter(r *gin.Engine) {

	r.GET("/", Controllers.Site)
	r.POST("/CreateProject", Controllers.CreateProject)
	r.GET("/CreateProjectForm", Controllers.CreateProjectForm)
	r.GET("/ShowProjectList", Controllers.ShowProjectList)
	r.GET("/ProjectPull/:id", Controllers.ProjectPull)
	r.GET("/Project/show/:id", Controllers.ShowProject)
	r.POST("/Project/edit/:id", Controllers.EditProject)
	r.GET("/Project/delete/:id", Controllers.DeleteProject)

	r.GET("/test", Controllers.Testa)
}