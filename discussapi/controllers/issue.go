package controllers

import (
	"github.com/astaxie/beego"
	"my-go-web/discussapi/models"
	"encoding/json"
)

type IssueController struct{
	BaseController
}

//@router /create [post]
func (this *IssueController)Create()  {
	beego.Info("issue::begin",string(this.Ctx.Input.RequestBody))
	var issue models.Issue
	err:=json.Unmarshal(this.Ctx.Input.RequestBody,&issue)
	if err!=nil{
		beego.Info("parsed error")
		this.Data["json"]=Response{INSERTERR,"解析帖子失败",nil}
	}
	err,flag:=issue.Create()
	if err!=nil{
		this.Data["json"]=Response{INSERTERR,"插入数据出错",nil}
	}else if !flag{
		this.Data["json"]=Response{INSERTERR,"插入数据失败",nil}
	}else {
		this.Data["json"]=Response{SUCCESS,"插入成功",nil}
	}
	this.ServeJSON()
}

//@router /update [post]
func (this *IssueController) Update()  {
	beego.Info("issue::begin",string(this.Ctx.Input.RequestBody))
	var issue models.Issue
	err:=json.Unmarshal(this.Ctx.Input.RequestBody,&issue)
	if err!=nil{
		beego.Info("parsed error")
		this.Data["json"]=Response{UPDATEERR,"更新解析失败",nil}
	}
	err=issue.Update()
	if err!=nil{
		beego.Info("更新失败")
		this.Data["json"]=Response{UPDATEERR,"更新解析失败",nil}
	}else {
		this.Data["json"]=Response{SUCCESS,"更新成功",issue}
	}
	this.ServeJSON()
}

//@router /delete [post]
func (this *IssueController)Delete()  {
	beego.Info("issue::begin",string(this.Ctx.Input.RequestBody))
	var issue models.Issue
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

//@router /querybasic [get]
func (this *IssueController) QueryBasic()  {
	beego.Info("issue::begin",string(this.Ctx.Input.RequestBody))
	issues,err:=models.QueryBasicIssue()
	if err!=nil{
		this.Data["json"]=Response{READERR,"",nil}
		beego.Info("查询失败")
	}else {
		this.Data["json"]=Response{SUCCESS,"",issues}
	}
	this.ServeJSON()
}

//@router /querydetail [post]
func (this *IssueController)QueryDetail()  {
	beego.Info("issue::begin",string(this.Ctx.Input.RequestBody))
	var issue models.Issue
	err:=json.Unmarshal(this.Ctx.Input.RequestBody,&issue)
	if err!=nil{
		beego.Info("parsed error")
		this.Data["json"]=Response{READERR,"解析帖子失败",nil}
	}
	err=issue.QueryDetailIssue()
	if err!=nil{
		this.Data["json"]=Response{READERR,"",nil}
		beego.Info("查询失败")
	}else {
		this.Data["json"]=Response{SUCCESS,"",issue}
	}
	this.ServeJSON()
}