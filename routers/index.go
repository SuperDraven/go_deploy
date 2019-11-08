package routers

import (
	"github.com/gin-gonic/gin"
	"go_deploy/controllers"
)

func SetRouter(r *gin.Engine) {
	r.POST("/test", controllers.Shell)

}