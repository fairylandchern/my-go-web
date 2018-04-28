package main

import (
	"github.com/astaxie/beego"
	controllers "my-go-web/web/controllers"
	_ "my-go-web/web/routers"
)


func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
