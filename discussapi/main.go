package main

import (
	_ "my-go-web/discussapi/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	"my-go-web/discussapi/controllers"
	"my-go-web/discussapi/models"
)

func init()  {
	models.Init()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	orm.DefaultTimeLoc = time.UTC
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
