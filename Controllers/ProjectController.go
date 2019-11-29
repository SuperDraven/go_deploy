package Controllers

import (
	"fmt"
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
	project.Uuid = uuid.NewV4()
	project.Created_at = time.Now().Unix()
	project.Directory = c.PostForm("directory")
	project.Secret = c.PostForm("secret")
	project.Updated_at = time.Now().Unix()
	project.Webhook = conf.LoadConf().SiteUrl + ":" +conf.LoadConf().SitePort
	project.ContentShell = c.PostForm("contentshell")
	Services.ServiceCreateProject(project)
	Services.GenerateShell(project.Uuid.String(), project.ContentShell)
}

type Users struct {
	Id int
	Name string
	Password string
}
func ShowProjectList(c *gin.Context)  {
	//准备sql语句
	projects := Services.ServiceProjectShowList()
	fmt.Println(projects)
	c.HTML(http.StatusOK, "ProjectList.html", gin.H{
		"data":projects,
	})
	//c.HTML(200, gin.H{})
}
func ShowProject(c *gin.Context)  {
	id := c.Param("id")
	project := Services.ServiceShowProject(id)
	c.HTML(http.StatusOK, "createProject.html", gin.H{
		"data":project,
	})
}
func EditProject(c *gin.Context)  {
	id := c.Param("id")
	project := new(models.Project)
	project.Name = c.PostForm("name")
	project.Created_at = time.Now().Unix()
	project.Directory = c.PostForm("directory")
	project.Secret = c.PostForm("secret")
	project.ContentShell = c.PostForm("contentshell")
	Services.ServicesEditProject(id, project)
	projects := Services.ServiceShowProject(id)
	Services.GenerateShell(projects.Uuid.String(), project.ContentShell)
}
func DeleteProject(c *gin.Context)  {
	id := c.Param("id")
	Services.ServicesDeleteProject(id)
}