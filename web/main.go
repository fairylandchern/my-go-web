package main

import (
	"database/sql"
	"my-go-web/web/models"
)

var db *sql.DB

func getdb() *sql.DB {
	db, _ := models.GetMysqlConnect()
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	return db
}

func main() {
	// if beego.BConfig.RunMode == "dev" {
	// 	beego.BConfig.WebConfig.DirectoryIndex = true
	// 	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	// }
	// beego.Run()
	db = getdb()
	var user = new(models.User)
	user.UserId = 2
	user.QueryByUserID(db)
}
