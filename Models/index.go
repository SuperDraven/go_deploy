package models

import (
	//"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_"github.com/mattn/go-sqlite3"
)


var DB *sqlx.DB

func init() {
	var errs error
	DB, errs = sqlx.Open("sqlite3", "./deploy.db")
	if errs != nil {
		fmt.Println(errs)
	}
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return
	}
	//defer DB.Close()
	fmt.Println("connnect success")
}
