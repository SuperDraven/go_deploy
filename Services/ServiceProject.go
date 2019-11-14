package Services

import (
	"fmt"
	models "go_deploy/Models"
)


func ServiceCreateProject(project *models.Project)  {
	models.ProjectInsert(project.Name, project.Secret, project.Directory, project.Created_at, project.Updated_at, project.Webhook, project.Uuid)
}
func ServiceProjectShowList() (projects []models.Project)  {
	project := models.ProjectShowList()
	fmt.Println(project)
	return project
}