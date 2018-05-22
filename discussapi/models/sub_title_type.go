package models

import (
	"github.com/astaxie/beego/orm"
)
//子栏目标题
type SubTitleType struct {
	SubTitleTypeId int `json:"sub_title_type_id" orm:"column(sub_title_type_id);pk;auto;unique"`
	SubTitleName string `json:"sub_title_name" orm:"column(sub_title_name);unique;size(30)"`
	MainTitleTypeId int `json:"main_title_type_id" orm:"column(main_title_type_id);unique"`
}

func init() {
	orm.RegisterModel(new(SubTitleType))
}

func (s *SubTitleType) TableName() string {
	return TableName("sub_title_type")
}

func NewSubTitleType() *SubTitleType{
	return new(SubTitleType)
}

func (s *SubTitleType)QuerySetter() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(SubTitleType))
}

func (s *SubTitleType)QueryAll()([]*SubTitleType,error){
	var subtitles []*SubTitleType
	_,err:=s.QuerySetter().All(subtitles)
	return subtitles,err
}

func (s *SubTitleType) QueryByMainTypeId()([]*SubTitleType,error)  {
	var subtitles []*SubTitleType
	_,err:=s.QuerySetter().Filter("MainTitleTypeId",s.MainTitleTypeId).All(subtitles)
	return subtitles,err
}

func (s *SubTitleType)Insert()error  {
	_,err:=orm.NewOrm().Insert(s)
	return err
}

func (s *SubTitleType)Update()(err error) {
	_,err=orm.NewOrm().Update(s,"SubTitleName")
	return
}

func (s *SubTitleType) Delete() error {
	_,err:=orm.NewOrm().Delete(s)
	return err;
}
