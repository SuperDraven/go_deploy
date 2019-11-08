package routers

import (
	"github.com/gin-gonic/gin"
	"go_deploy/controllers"
)

func SetRouter(r *gin.Engine) {
	r.GET("/test", controllers.Shell)

}