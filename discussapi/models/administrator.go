package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Administrator struct {
	Id       int64     `json:"id" orm:"column(id);pk;auto;unique"`
	Phone    string    `json:"phone" orm:"column(phone);unique;size(11)"`
	Nickname string    `json:"nickname" orm:"column(nickname);unique;size(40);"`
	Password string    `json:"password" orm:"column(password);size(40)"`
	Created  time.Time `json:"-" orm:"column(create_at);auto_now_add;type(datetime)"`
	Updated  time.Time `json:"-" orm:"column(update_at);auto_now;type(datetime)"`
}

func (u *Administrator) TableName() string {
	return TableName("admin")
}
func init() {
	orm.RegisterModel(new(Administrator))
}

func (u *Administrator)Addadminastrator() error{
	_,err:=orm.NewOrm().Insert(u)
	return err;
}

func (u *Administrator) UpdateAdmin(fields ...string) (*Administrator,error) {
	_,err:=orm.NewOrm().Update(u,fields...)
	return u,err;
}

func (u *Administrator)CheckHasAdmin() bool {
	cond:=orm.NewCondition()
	cond=cond.And("Id",u.Id).And("Password",u.Password)
	exist:=orm.NewOrm().QueryTable(new(Administrator)).SetCond(cond).Exist()
	orm.NewOrm().QueryTable(new(Administrator)).SetCond(cond).One(u,"Id","Phone","Nickname","Password")
	return  exist
}

func (u *Administrator)DelAdmin() bool{
	count,_:=orm.NewOrm().Delete(u,"Id")
	if count>0{
		return  true
	}
	return false;
}
