package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)



//关注表
type Follow struct {
	FollowId int64 `json:"follow_id" orm:"column(follow_id);pk;auto;unique"`
	FollowerId int64 `json:"follower_id" orm:"column(follower_id);"`  //粉丝id
	BeFollowerId int64 `json:"be_follower_id" orm:"column(be_follwer_id);"` //关注者id
}

func init() {
	orm.RegisterModel(new (Follow));
}

func (f *Follow)TableName() string {
	return TableName("follow");
}

func GetFollowQueryer() orm.QuerySeter  {
	return orm.NewOrm().QueryTable(new(Follow))
}

func (f *Follow) InsertFollow() error  {
	_,err:=orm.NewOrm().Insert(f)
	return  err;
}

func (f *Follow) DelFollow() bool {
	count,_:=orm.NewOrm().Delete(f,"FollowerId","BeFollowerId")
	beego.Info("count is:",count)
	if count>0{
		beego.Info("count is:",count)
		return true;
	}
	return false
}

//获取关注数
func (f *Follow) CountFollower() (int64, error) {
	count,err:=orm.NewOrm().QueryTable(new(Follow)).Filter("FollowerId",f.FollowerId).Count()
	return count,err
}

//获取粉丝数
func (f*Follow) CountToBeFollower() (int64,error) {
	count,err:=orm.NewOrm().QueryTable(new(Follow)).Filter("BeFollowerId",f.BeFollowerId).Count()
	return  count,err;
}

//检测是否已经关注过,基于orm框架的条件查询
func (f*Follow)CheckHasFollower() bool {
	cond:=orm.NewCondition()
	cond=cond.And("FollowerId",f.FollowerId).And("BeFollowerId",f.BeFollowerId)
	count,_:=orm.NewOrm().QueryTable(new(Follow)).SetCond(cond).Count()
	beego.Info(" follow count is:",count)
	if count==0{
		return  true
	}
	return false
}

//获取到所有关注者信息
func  (f *Follow)GetAllFollwer()([]*Follow,error)  {
	var follows []*Follow
	 _,err:=orm.NewOrm().QueryTable(new(Follow)).Filter("FollowerId",f.BeFollowerId).All(&follows,"FollowId","FollowerId","BeFollowerId")
	 if err!=nil{
	 	return nil,err
	 }
	 return follows,nil
}

//获取到所有粉丝信息
func  (f *Follow)GetAllFans()([]*Follow,error)  {
	var follows []*Follow
	_,err:=orm.NewOrm().QueryTable(new(Follow)).Filter("BeFollowerId",f.BeFollowerId).All(&follows,"FollowId","FollowerId","BeFollowerId")
	if err!=nil{
		return nil,err
	}
	return follows,nil
}


//分页查询如果有时间需要补充实现，下一步优化方向