package main

import (
	"my-go-web/web/models"
)

func main() {
	// if beego.BConfig.RunMode == "dev" {
	// 	beego.BConfig.WebConfig.DirectoryIndex = true
	// 	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	// }
	// beego.Run()
	var user = new(models.User)
	user.UserId = 2
	user.QueryByUserID()
}
