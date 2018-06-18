package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"my-go-web/discussapi/models"
)

type AdministratorController struct {
	BaseController
}


//@router /login [post]
func (this *AdministratorController) Login()  {
	beego.Info("admin::begin to login",string(this.Ctx.Input.RequestBody))
	var user models.Administrator
	err:=json.Unmarshal(this.Ctx.Input.RequestBody,&user)
	if err!=nil{
		this.Data["json"]=Response{READERR,"USER::JSON parsed failed",nil}
		beego.Info("user::json parsed failed")
	}else {
		flag:=user.CheckHasAdmin()
		if !flag{
			this.Data["json"]=Response{READERR,"user::read failed",nil}
			beego.Info("user::read failed",err.Error())
		}else {
			this.Data["json"]=Response{SUCCESS,"",user}
			beego.Info("user::",user)
		}
	}
	this.ServeJSON()
}
