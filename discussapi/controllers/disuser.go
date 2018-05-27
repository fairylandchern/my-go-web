package controllers

import (
	"my-go-web/discussapi/models"
	"github.com/astaxie/beego"
	"encoding/json"
)

type UserController struct {
	BaseController
}

//@router /register [post]
func (this *UserController)Register()  {
	beego.Info("begin to register",string(this.Ctx.Input.RequestBody))
	var user models.User
	error:=json.Unmarshal(this.Ctx.Input.RequestBody,&user)
	if error!=nil{
		beego.Error("json parsed user failed")
		this.Data["json"]=Response{INSERTERR,"parsed failed",nil}
		beego.Info("user::json parsed failed")
	}else {
		error=user.Insert()
		if error!=nil{
			this.Data["json"]=Response{INSERTERR,"registered failed",nil}
			beego.Error("user::insert failed")
		}else {
			this.Data["json"]=Response{SUCCESS,"registered success",nil}
		}
	}
	this.ServeJSON()
}

//@router /login [post]
func (this *UserController) Login()  {
	beego.Info("user::begin to login",string(this.Ctx.Input.RequestBody))
	var user models.User
	err:=json.Unmarshal(this.Ctx.Input.RequestBody,&user)
	if err!=nil{
		this.Data["json"]=Response{READERR,"USER::JSON parsed failed",nil}
		beego.Info("user::json parsed failed")
	}else {
		err=(&user).Read("Nickname","Password")
		if err!=nil{
			this.Data["json"]=Response{READERR,"user::read failed",nil}
			beego.Info("user::read failed",err.Error())
		}else {
			this.Data["json"]=Response{SUCCESS,"",user}
			beego.Info("user::",user)
		}
	}
	this.ServeJSON()
}

//@router /update [post]
func (this *UserController)Update()  {
	beego.Info("begin user::update",string(this.Ctx.Input.RequestBody))
	var user models.User
	err:=json.Unmarshal(this.Ctx.Input.RequestBody,&user)
	if err!=nil{
		beego.Error("user::json parsed failed",err.Error())
		this.Data["json"]=Response{UPDATEERR,"",nil}
	}else {
		err=user.Update()
		if err!=nil{
			beego.Error("user::update error")
			this.Data["json"]=Response{UPDATEERR,"",nil}
		}else {
			this.Data["json"]=Response{SUCCESS,"",user}
		}
	}
	this.ServeJSON()
}

//***后续补充吧