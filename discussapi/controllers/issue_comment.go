package controllers

import (
	"encoding/json"
	"my-go-web/discussapi/models"
	"strconv"

	"github.com/astaxie/beego"
)

type IssueCommentController struct {
	BaseController
}

//@router /create [post]
func (this *IssueCommentController) Create() {
	beego.Info("IssueComment::begin", string(this.Ctx.Input.RequestBody))
	var issue models.IssueComment
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &issue)
	beego.Info("the comment is", issue.ComentContent)
	if err != nil {
		beego.Info("parsed error", err.Error())
		this.Data["json"] = Response{INSERTERR, "解析帖子评论失败", nil}
	} else if err = issue.Create(); err != nil {
		this.Data["json"] = Response{INSERTERR, "插入数据出错", nil}
		beego.Info("err is:", err.Error())
	} else {
		this.Data["json"] = Response{SUCCESS, "插入成功", nil}
	}
	this.ServeJSON()
}

//@router /update [post]
func (this *IssueCommentController) Update() {
	beego.Info("issue::begin", string(this.Ctx.Input.RequestBody))
	var issue models.IssueComment
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &issue)
	if err != nil {
		beego.Info("parsed error")
		this.Data["json"] = Response{UPDATEERR, "更新解析失败", nil}
	}
	err = issue.Update()
	if err != nil {
		beego.Info("更新失败")
		this.Data["json"] = Response{UPDATEERR, "更新解析失败", nil}
	} else {
		this.Data["json"] = Response{SUCCESS, "更新成功", issue}
	}
	this.ServeJSON()
}

//@router /delete [post]
func (this *IssueCommentController) Delete() {
	beego.Info("issue::begin", string(this.Ctx.Input.RequestBody))
	var issue models.IssueComment
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &issue)
	if err != nil {
		beego.Info("parsed error", err.Error())
		this.Data["json"] = Response{INSERTERR, "解析帖子失败", nil}
	}
	err = issue.Delete()
	if err != nil {
		beego.Info("删除失败")
		this.Data["json"] = Response{DELETEERR, "删除解析失败", nil}
	} else {
		this.Data["json"] = Response{SUCCESS, "删除成功", nil}
	}
	this.ServeJSON()
}

//@router /querybyissue/:id [get]
func (this *IssueCommentController) QueryByIssue() {
	beego.Info("issue::begin", string(this.Ctx.Input.RequestBody))
	str_issueid := this.Ctx.Input.Param(":id")
	issueid, err := strconv.ParseInt(str_issueid, 10, 64)
	if err != nil {
		beego.Info("转换失败", err.Error())
		this.Data["json"] = Response{READERR, "", nil}
		return
	}
	issues, err := models.QueryCommentByIssueId(issueid)
	if err != nil {
		this.Data["json"] = Response{READERR, "", nil}
		beego.Info("查询失败")
	} else {
		this.Data["json"] = Response{SUCCESS, "", issues}
	}
	defer this.ServeJSON()
}

//@router /querybyuser/:id [get]
func (this *IssueCommentController) QueryByUser() {
	beego.Info("issue::begin", string(this.Ctx.Input.RequestBody))
	str_issueid := this.Ctx.Input.Param(":id")
	userid, err := strconv.ParseInt(str_issueid, 10, 64)
	if err != nil {
		beego.Info("转换失败", err.Error())
		this.Data["json"] = Response{READERR, "", nil}
		return
	}
	issues, err := models.QueryCommentByUserId(userid)
	if err != nil {
		this.Data["json"] = Response{READERR, "", nil}
		beego.Info("查询失败")
	} else {
		this.Data["json"] = Response{SUCCESS, "", issues}
	}
	defer this.ServeJSON()
}

//@router /getallcomments [get]
func (this *IssueCommentController)GetAllComments()  {
	beego.Info("begin user::getallusers",string(this.Ctx.Input.RequestBody))
	comments,err:=models.QueryComments()
	if err!=nil{
		this.Data["json"]=Response{READERR,"",nil}
		beego.Info("查询失败")
	}else {
		this.Data["json"]=Response{SUCCESS,"",comments}
	}
	this.ServeJSON()
}