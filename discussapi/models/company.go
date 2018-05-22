package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
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
func (this *Company) Query(company *Company)  bool{
	cond:=orm.NewCondition().And("company_id",company.CompanyId).Or("company_name",company.CompanyName)
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

//插入新数据
func (this *Company)Create() error  {
	o:=orm.NewOrm()
	_,err:=o.Insert(this)
	if err!=nil{
		beego.Error("插入数据时出错",err)
		return err
	}
	return nil
}

//更新数据
func (this *Company) Update(fileds ...string) error {
	o:=orm.NewOrm()
	_,err:=o.Update(this,fileds...)
	if err!=nil{
		beego.Error("company更新出错：",err)
		return err
	}
	return nil
}

//删除company
func (this *Company) Delete()error  {
	_,err:=orm.NewOrm().Delete(this)
	if err!=nil{
		beego.Error("删除company出错:",err)
		return err
	}
	return nil
}

//查询company
func (this *Company) Read(fields ...string) error {
	err:=orm.NewOrm().Read(this,fields...)
	if err!=nil{
		beego.Error("更新company时出错:",err)
		return err
	}
	return nil
}

