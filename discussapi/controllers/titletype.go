package controllers

import (
	"github.com/astaxie/beego/logs"
	"my-go-web/discussapi/models"
	"encoding/json"
)

type TitleTypeController struct {
	BaseController
}

//@router /createmaintype [post]
func (this *TitleTypeController) CreateMainType()  {
	var mtype models.MainTitleType
	json.Unmarshal(this.Ctx.Input.RequestBody,&mtype)
	logs.Error("the error message is:",mtype.MainTitleTypeName)
	exist:=mtype.QueryByName()
	if exist{
		this.Data["json"]=Response{DATAEXIST,"the data is exist",nil}
	}else {
		 err:=mtype.Insert()
		 if err!=nil{
		 	this.Data["json"]=Response{INSERTERR,"insert failured",nil}
		 }else {
		 	this.Data["json"]=Response{SUCCESS,"",nil}
		 }
	}
	this.ServeJSON()
}
//@router /getallmaintype [get]
func (this *TitleTypeController) GetAllMainType(){
	var mtypes []models.MainTitleType
	var err error
	mtypes,err=models.QueryAllMainType()
	if err!=nil {
		this.Data["json"]=Response{GETALLDATAERROR,"数据未读取到",nil}
	}else {
		this.Data["json"]=Response{SUCCESS,"",mtypes}
	}
	this.ServeJSON()
}