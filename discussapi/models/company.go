package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

type Company struct {
	CompanyId int64 `json:"company_id" orm:"column(company_id);pk;auto;unique"`
	CompanyName string `json:"company_name" orm:"column(company_name);unique;size(50)"`
	CompanyAddre string `json:"company_addr" orm:"column(company_addr);size(100)"`
}

func init() {
	orm.RegisterModel(new(Company))
}

func (this *Company)TableName() string {
	return  TableName("company")
}

func (this *Company) Company() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Company))
}

//查询
func (this *Company) query(company *Company)  bool{
	cond:=orm.NewCondition().And("company_id").Or("company_name")
	count,err:=this.Company().SetCond(cond).Count()
	if err!=nil{
		logs.Error("company:条件查询出错",err.Error())
		return false
	}
	if count!=0{
		return true
	}else {
		return false
	}
}


