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

//@router /getallusers [get]
func (this *UserController)GetAllUsers()  {
	beego.Info("begin user::getallusers",string(this.Ctx.Input.RequestBody))
	users,err:=models.GetAllUsers()
	if err!=nil{
		this.Data["json"]=Response{READERR,"",nil}
		beego.Info("查询失败")
	}else {
		this.Data["json"]=Response{SUCCESS,"",users}
	}
	this.ServeJSON()
}

//@router /delete [post]
func (this *UserController)Delete()  {
	beego.Info("issue::begin",string(this.Ctx.Input.RequestBody))
	var issue models.User
	err:=json.Unmarshal(this.Ctx.Input.RequestBody,&issue)
	if err!=nil{
		beego.Info("parsed error")
		this.Data["json"]=Response{INSERTERR,"解析帖子失败",nil}
	}
	err=issue.Delete()
	if err!=nil{
		beego.Info("删除失败")
		this.Data["json"]=Response{DELETEERR,"删除解析失败",nil}
	}else {
		this.Data["json"]=Response{SUCCESS,"删除成功",nil}
	}
	this.ServeJSON()
}

//@router /querydetail [post]
func (this *UserController)QueryDetail()  {
	beego.Info("issue::begin",string(this.Ctx.Input.RequestBody))
	var user models.User
	err:=json.Unmarshal(this.Ctx.Input.RequestBody,&user)
	if err!=nil{
		beego.Info("parsed error",err.Error())
		this.Data["json"]=Response{READERR,"解析帖子失败",nil}
	}
	err=user.GetUserDetail()
	if err!=nil{
		this.Data["json"]=Response{READERR,"",nil}
		beego.Info("查询失败")
	}else {
		this.Data["json"]=Response{SUCCESS,"",user}
	}
	this.ServeJSON()
}
