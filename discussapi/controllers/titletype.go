package controllers

import (
	"log"
	"github.com/astaxie/beego/logs"
)

type TitleTypeController struct {
	BaseController
}


//@descriptor get the main title type
//@router /createmaintype [post]
func (this *TitleTypeController) CreateMainType()  {
	log.Println(this.Ctx.Input.RequestBody)	//果然不能用log.fatal，否则会报错，并且程序停止运行,以字节的形式打印输出（数字，所以没有看出来）
	logs.Error(string(this.Ctx.Input.RequestBody))
	this.Data["json"]=Response{200,"no error","data is success"};
	this.ServeJSON()
}