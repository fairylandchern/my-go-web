package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"my-go-web/discussapi/models"
)

type FollowController struct {
	BaseController
}

//@router /addfollower [post]
func (this *FollowController) AddFollower()  {
	beego.Info("follower::begin",string(this.Ctx.Input.RequestBody))
	var follow models.Follow;
	err:=json.Unmarshal(this.Ctx.Input.RequestBody,&follow)
	beego.Info(follow.FollowerId,follow.BeFollowerId)
	if err!=nil{
		beego.Info("parsed error")
		this.Data["json"]=Response{INSERTERR,"添加关注解析失败",nil}
	}else if models.CheckUserExist(follow.FollowerId) &&  models.CheckUserExist(follow.BeFollowerId)&&follow.CheckHasFollower(){
		err=follow.InsertFollow()
		if err!=nil{
			beego.Info("添加关注失败")
			this.Data["json"]=Response{INSERTERR,"添加关注失败",nil}
		}else {
			this.Data["json"]=Response{SUCCESS,"关注成功",nil}
		}
	}else {
		beego.Info("用户不存在或已关注")
		this.Data["json"]=Response{INSERTERR,"用户不存在",nil}
	}
	this.ServeJSON()
}

//@router /delbefollower [post]
func (this *FollowController) DelBeFollower()  {
	beego.Info("follow::del",string(this.Ctx.Input.RequestBody))
	var follow models.Follow
	err:=json.Unmarshal(this.Ctx.Input.RequestBody,&follow)
	beego.Info(follow.FollowerId,follow.BeFollowerId)
	if err!=nil{
		beego.Info("parsed error")
		this.Data["json"]=Response{DELETEERR,"取消关注解析失败",nil}
	}else {
		flag:=follow.DelFollow()
		if !flag{
			beego.Info("取消关注失败")
			this.Data["json"]=Response{DELETEERR,"取消关注失败",nil}
		}else {
			this.Data["json"]=Response{SUCCESS,"取消关注成功",nil}
		}
	}
	this.ServeJSON()
}

//@router /countfollwer [post]
func (this *FollowController)CountFollower()  {
	beego.Info("计算粉丝数",string(this.Ctx.Input.RequestBody))
	var follow models.Follow
	err:=json.Unmarshal(this.Ctx.Input.RequestBody,&follow)
	if err!=nil{
		beego.Info("parsed error")
		this.Data["json"]=Response{DELETEERR,"解析失败",nil}
	}else if models.CheckUserExist(follow.BeFollowerId) {
		count,err:=follow.CountToBeFollower()
		if err!=nil{
			beego.Info("粉丝数失败")
			this.Data["json"]=Response{DELETEERR,"获取粉丝数失败",nil}
		}else {
			this.Data["json"]=Response{SUCCESS,"获取粉丝数成功成功",count}
		}
	}else {
		beego.Info("用户不存在")
		this.Data["json"]=Response{INSERTERR,"用户不存在",nil}
	}
	this.ServeJSON()
}

//@router /getfans [post]
func (this *FollowController)Getfans()  {
	beego.Info("获取粉丝",string(this.Ctx.Input.RequestBody))
	var follow models.Follow
	err:=json.Unmarshal(this.Ctx.Input.RequestBody,&follow)
	if err!=nil{
		beego.Info("parsed error")
		this.Data["json"]=Response{DELETEERR,"解析失败",nil}
	}else if models.CheckUserExist(follow.BeFollowerId) {
		fans,err:=follow.GetAllFans()
		if err!=nil{
			beego.Info("粉丝获取失败")
			this.Data["json"]=Response{DELETEERR,"获取粉丝失败",nil}
		}else {
			this.Data["json"]=Response{SUCCESS,"获取粉丝成功",fans}
		}
	}else {
		beego.Info("用户不存在")
		this.Data["json"]=Response{INSERTERR,"用户不存在",nil}
	}
	this.ServeJSON()
}

//@router /countbefollwer [post]
func (this *FollowController)CountBeFollower()  {
	beego.Info("计算关注数",string(this.Ctx.Input.RequestBody))
	var follow models.Follow
	err:=json.Unmarshal(this.Ctx.Input.RequestBody,&follow)
	beego.Info(follow)
	if err!=nil{
		beego.Info("parsed error")
		this.Data["json"]=Response{DELETEERR,"解析失败",nil}
	}else if models.CheckUserExist(follow.FollowerId) {
		count,err:=follow.CountFollower()
		if err!=nil{
			beego.Info("关注数失败")
			this.Data["json"]=Response{DELETEERR,"获取关注数失败",nil}
		}else {
			this.Data["json"]=Response{SUCCESS,"获取关注数成功成功",count}
		}
	}else {
		beego.Info("用户不存在")
		this.Data["json"]=Response{INSERTERR,"用户不存在",nil}
	}
	this.ServeJSON()
}

//@router /getfollowers [post]
func (this *FollowController)GetFollowers()  {
	beego.Info("获取关注",string(this.Ctx.Input.RequestBody))
	var follow models.Follow
	err:=json.Unmarshal(this.Ctx.Input.RequestBody,&follow)
	beego.Info(follow)
	if err!=nil{
		beego.Info("parsed error")
		this.Data["json"]=Response{DELETEERR,"解析失败",nil}
	}else if models.CheckUserExist(follow.FollowerId) {
		follwers,err:=follow.GetAllFollwer()
		if err!=nil{
			beego.Info("获取关注失败")
			this.Data["json"]=Response{DELETEERR,"获取关注失败",nil}
		}else {
			this.Data["json"]=Response{SUCCESS,"获取关注成功",follwers}
		}
	}else {
		beego.Info("用户不存在")
		this.Data["json"]=Response{INSERTERR,"用户不存在",nil}
	}
	this.ServeJSON()
}