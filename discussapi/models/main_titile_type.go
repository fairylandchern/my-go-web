package models

import (
	"github.com/astaxie/beego/orm"
)
//主栏目标题
type MainTitleType struct {
	MainTitleTypeId int `json:"main_title_type_id" orm:"column(main_title_type_id);pk;auto;unique"`
	MainTitleTypeName string `json:"main_title_type_name" orm:"column(main_title_type_name);unique;size(30)"`
}

func init() {
	orm.RegisterModel(new(MainTitleType))
}

func (m *MainTitleType)TableName() string  {
	return TableName("main_title_type")
}

func MtypeQuerySetter() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(MainTitleType))
}

func NewMainTitleType()*MainTitleType  {
	return new(MainTitleType)
}

func QueryAllMainType()([]MainTitleType,error)  {
	var mains []MainTitleType
	_,err:=MtypeQuerySetter().All(mains)
	return  mains,err
}

func (m *MainTitleType)QueryByName()bool {
	i:=MtypeQuerySetter().Filter("MainTitleTypeName",m.MainTitleTypeName).Exist()
	return i
}

func (m *MainTitleType)Insert() error  {
	_,err:=orm.NewOrm().Insert(m)
	return err
}

func (m *MainTitleType) Update() error {
	_,err:=orm.NewOrm().Update(m,"MainTitleTypeName")
	return err
}

func (m *MainTitleType)Delete()  error {
	_,err:=orm.NewOrm().Delete(m)
	return err
}