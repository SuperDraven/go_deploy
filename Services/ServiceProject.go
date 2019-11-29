package Services

import (
	"fmt"
	models "go_deploy/Models"
)


func ServiceCreateProject(project *models.Project)  {
	models.ProjectInsert(project.Name, project.Secret, project.Directory, project.Created_at, project.Updated_at, project.Webhook, project.Uuid, project.ContentShell)
}
func ServiceProjectShowList() (projects []models.Project)  {
	project := models.ProjectShowList()
	fmt.Println(project)
	return project
}

func ServiceShowProject(id string) (projects models.Project) {
	project := models.ProjectShow(id)

	return project
}
func ServicesEditProject(id string, project *models.Project)  {
	//models.ProjectUpdate()
	models.ProjectUpdate(project.Name, project.Secret, project.Directory, project.ContentShell, id)
}

func ServicesDeleteProject(id string)  {
	models.ProjectDelete(id)
}