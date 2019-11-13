package controllers

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	conf "go_deploy/Conf"
	models "go_deploy/Models"
	"go_deploy/Services"
	"net/http"
	"time"
)

func CreateProjectForm(c *gin.Context)  {
	c.HTML(http.StatusOK, "createProject.html", gin.H{
		"title": "hello Go",
	})
}

func CreateProject(c *gin.Context)  {
	project := new(models.Project)
	project.Name = c.PostForm("name")
	project.Id = uuid.NewV4()
	project.Created_at = time.Now().Unix()
	project.Directory = c.PostForm("directory")
	project.Secret = c.PostForm("secret")
	project.Updated_at = time.Now().Unix()
	project.Webhook = conf.LoadConf().SiteUrl + ":" +conf.LoadConf().SitePort
	Services.ServiceCreateProject(project)
}

type Users struct {
	Id int
	Name string
	Password string
}
func ShowProjectList(c *gin.Context)  {
	//准备sql语句
	projects := Services.ServiceProjectShowList()
	c.HTML(http.StatusOK, "ProjectList.html", gin.H{
		"data":projects,
	})
	//c.HTML(200, gin.H{})
}