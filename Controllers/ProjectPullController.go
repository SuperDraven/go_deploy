package Controllers

import (
	"github.com/chenhg5/go-task"
	"github.com/gin-gonic/gin"
	models "go_deploy/Models"
	"go_deploy/Services"
	"runtime"
	"strings"
)

func ProjectPull(c *gin.Context)  {

	task.InitTaskReceiver(runtime.NumCPU())
	id := c.Param("id")
	project := models.ProjectShow(id)
	Services.PullShellGo(project.Directory, "git pull origin master")

	a := strings.Split(project.ContentShell, "\r")
	ch := make(chan string)
	go sendData(ch, a)
	go getData(ch)
}

func sendData(ch chan string,a []string) {
	for i:=0;i<len(a);i++ {
		ch <- a[i]
	}
	 close(ch)
}

func getData(ch chan string) {
	//var input string
	// time.Sleep(2e9)
	for {
		input, open := <-ch
		if !open {
			break
		}
		Services.PullShellTest(input)
	}
}