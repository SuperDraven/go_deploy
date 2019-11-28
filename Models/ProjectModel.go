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
	ContentShell string
}

func ProjectInsert(name string, secret string, directory string, created_at int64, updated_at int64, webhook string, uuid uuid.UUID, contentShell string)  {
	_,error :=DB.Exec("insert into `project`(`name`,`secret`,`directory`,`created_at`,`updated_at`, `webhook`,`uuid`,`contentshell`)VALUES(?,?,?,?,?,?,?,?)",name, secret, directory, created_at, updated_at, webhook, uuid, contentShell)
	fmt.Println(error)
}
func ProjectUpdate(name string, secret string, directory string, contentShell string, uuid string)  {
	_,error :=DB.Exec("UPDATE `project` SET `name` = ?, `secret` = ?, `directory` = ?, `contentShell` = ? WHERE `uuid` = ?",name, secret, directory, contentShell, uuid)
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
func ProjectDelete(id string)  {
	_,error :=DB.Exec("DELETE FROM `project` WHERE uuid = ?", id)
	fmt.Println(error)
}
func CreateTabelProject()  {
	_,errors := DB.Query(`CREATE TABLE project(
   id 			INTEGER PRIMARY KEY   AUTOINCREMENT,
   name          VARCHAR(255),
   directory      VARCHAR(255),
   secret         VARCHAR(50),
   webhook		  VARCHAR(255),
   uuid			  VARCHAR(255),
   created_at     int(11),
   updated_at     VARCHAR(255),
   contentshell   Text(0)
)`);
fmt.Println(errors)
}