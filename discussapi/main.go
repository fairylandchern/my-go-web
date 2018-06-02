package main

import (
	_ "my-go-web/discussapi/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	"my-go-web/discussapi/controllers"
	"my-go-web/discussapi/models"
	cros "github.com/astaxie/beego/plugins/cors"
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
	//服务器端添加对跨域的支持
	beego.InsertFilter("*", beego.BeforeRouter, cros.Allow(&cros.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.Run()
}
