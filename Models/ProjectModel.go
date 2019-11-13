package models

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)

type Project struct {
	Id  uuid.UUID
	Name string `json:"name"`
	Secret string `json:"secret"`
	Directory string `json:"directory"`
	Webhook string
	Created_at int64 `json:"created_at"`
	Updated_at int64 `json:"updated_at"`
}

func ProjectInsert(id uuid.UUID,name string, secret string, directory string, created_at int64, updated_at int64, webhook string)  {
	_,error :=DB.Exec("insert into `project`(`id`, `name`,`secret`,`directory`,`created_at`,`updated_at`, `webhook`)VALUES(?,?,?,?,?,?,?)", id, name, secret, directory, created_at, updated_at, webhook)
	fmt.Println(error)
}
func ProjectShowList() (projects []Project) {
	project := make([]Project, 0)
	err := DB.Select(&project,"select * from `project`")
	fmt.Println(err)
	return project
}
func ProjectShow(id string) (projects Project)  {
	project := Project{}
	errors := DB.Get(&project,"select * from `project` where id=?", id)
	fmt.Println(project)
	fmt.Println(errors)
	return project
}