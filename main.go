package main

import (
	"github.com/gin-gonic/gin"
	conf "go_deploy/Conf"
	"go_deploy/Routers"
	"net/http"
)

func main()  {
	r := gin.Default()
	//r.Use(Cors())
	r.LoadHTMLGlob("Resources/views/*")
	r.Static("/assets", "./assets")
	r.StaticFS("/more_static", http.Dir("assets"))
	Routers.SetRouter(r)

	_ = r.Run(":" + conf.LoadConf().SitePort)
}
