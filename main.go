package main

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "goniushop/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func RegisterDataBase()  {

	host := beego.AppConfig.String("db_host")
	database := beego.AppConfig.String("db_database")
	username := beego.AppConfig.String("db_username")
	password := beego.AppConfig.String("db_password")

	port := beego.AppConfig.String("db_port")

	createDB := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci", database)
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", username, password, host, port)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(createDB)
	if err != nil {
		panic(err)
	}

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", username, password, host, port, database)

	orm.RegisterDataBase("default", "mysql", dataSource)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	RegisterDataBase()
	beego.Run()
}
