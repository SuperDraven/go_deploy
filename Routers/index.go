package Routers

import (
	"github.com/gin-gonic/gin"
	"go_deploy/controllers"
)
//测试
func SetRouter(r *gin.Engine) {

	r.GET("/", controllers.Site)
	r.POST("/CreateProject", controllers.CreateProject)
	r.GET("/CreateProjectForm", controllers.CreateProjectForm)
	r.GET("/ShowProjectList", controllers.ShowProjectList)
	r.GET("/ProjectPull/:id", controllers.ProjectPull)
}