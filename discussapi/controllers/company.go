package controllers

import (
	"my-go-web/discussapi/models"
	"encoding/json"
)

//operation about company
type	CompanyController struct {
	BaseController
}

//@Title Create company
//@Description 添加公司信息
//@Param body body models.company "the company content"
//@Success 200 {string} models.company.companyid
//@Failure 403 body is empty
//@router /create [post]
func (this *CompanyController) Post()  {
	var company models.Company
	json.Unmarshal(this.Ctx.Input.RequestBody,&company)
	if err:=company.Create();err!=nil{
		this.Ctx.ResponseWriter.WriteHeader(403)
		this.Data["json"]=ErrResponse{403,err.Error()}
		this.ServeJSON()
		return
	}
	this.Data["json"]=Response{0,"nothing error",company}
	this.ServeJSON()
}




