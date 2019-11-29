package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	models "go_deploy/Models"
	"go_deploy/Services"
	"strings"
)

func ProjectPull(c *gin.Context)  {
	id := c.Param("id")
	project := models.ProjectShow(id)
	Services.PullShellGo(project.Directory, "git pull")

	 a := strings.Split(project.ContentShell, "\r")
	fmt.Println(a)
	for i:=0;i<len(a);i++ {
		//Services.PullShellTest(a[i])
	}
	//c.JSON(http.StatusOK, gin.H{
	//	"data":project,
	//})
}