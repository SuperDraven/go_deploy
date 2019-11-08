package main

import (
	"github.com/gin-gonic/gin"
	"go_deploy/conf"
	"go_deploy/routers"
)

func main()  {
	r := gin.Default()
	//r.Use(Cors())
	routers.SetRouter(r)

	_ = r.Run(":" + conf.LoadConf().SitePort)
}
