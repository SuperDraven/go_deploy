package main

import (
	"github.com/gin-gonic/gin"
	conf "go_deploy/Conf"

	"go_deploy/Routers"
)

func main()  {
	r := gin.Default()
	//r.Use(Cors())
	r.LoadHTMLGlob("Resources/views/*")

	Routers.SetRouter(r)

	_ = r.Run(":" + conf.LoadConf().SitePort)
}
