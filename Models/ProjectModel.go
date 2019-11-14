package models

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)

type Project struct {
	Id  int64
	Name string `json:"name"`
	Secret string `json:"secret"`
	Directory string `json:"directory"`
	Webhook string
	Created_at int64 `json:"created_at"`
	Updated_at int64 `json:"updated_at"`
	Uuid uuid.UUID
}

func ProjectInsert(name string, secret string, directory string, created_at int64, updated_at int64, webhook string, uuid uuid.UUID)  {
	_,error :=DB.Exec("insert into `project`(`name`,`secret`,`directory`,`created_at`,`updated_at`, `webhook`,`uuid`)VALUES(?,?,?,?,?,?,?)",name, secret, directory, created_at, updated_at, webhook, uuid)
	fmt.Println(error)
}
func ProjectShowList() (projects []Project) {
	project := make([]Project, 0)
	err := DB.Select(&project,"select * from `project`")
	fmt.Println(err)
	fmt.Println(project)
	return project
}
func ProjectShow(id string) (projects Project)  {
	project := Project{}
	errors := DB.Get(&project,"select * from `project` where uuid=?", id)
	fmt.Println(project)
	fmt.Println(errors)
	return project
}
func CreateTabelProject()  {
	_,errors := DB.Prepare(`CREATE TABLE project(
   id 			INTEGER PRIMARY KEY   AUTOINCREMENT,
   name          VARCHAR(255)    ,
   directory      VARCHAR(255)     ,
   secret         VARCHAR(50),
   webhook		  VARCHAR(255),
   uuid			  VARCHAR(255),
   created_at     int(11),
   updated_at     VARCHAR(255)
)`);
fmt.Println(errors)
}