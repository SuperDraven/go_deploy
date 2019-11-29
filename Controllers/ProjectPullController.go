package Controllers

import (
	"github.com/chenhg5/go-task"
	"github.com/gin-gonic/gin"
	models "go_deploy/Models"
	"go_deploy/Services"
	"runtime"
)

func ProjectPull(c *gin.Context)  {

	task.InitTaskReceiver(runtime.NumCPU())
	id := c.Param("id")
	project := models.ProjectShow(id)
	Services.PullShellGo(project.Directory, "git pull origin master")
	ch := make(chan string)
	go sendData(ch, project.Uuid.String())
	go getData(ch)
}

func sendData(ch chan string, uuid string) {
		ch <- uuid
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