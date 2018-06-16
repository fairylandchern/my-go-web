package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type IssueComment struct {
	CommentId     int       `json:"comment_id" orm:"column(comment_id);pk;auto;unique"`
	IssueId       int       `json:"issue_id" orm:"column(issue_id)"`
	UserId        int       `json:"user_id"  orm:"column(user_id)"`
	ComentContent string    `json:"comment_content" orm:"column(coment_content);size(200)"`
	Created       time.Time `json:"created,omitempty" orm:"column(created);auto_now_add;type(datetime)"`
	Updated       time.Time `json:"updated,omitempty" orm:"column(updated);auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(IssueComment))
}

func (this *IssueComment) TableName() string {
	return TableName("issue_comment")
}

func IssueCommentQueryer() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(IssueComment))
}

func (this *IssueComment) Create() error {
	_, err := orm.NewOrm().Insert(this)
	return err
}

func (this *IssueComment) Update() error {
	_, err := orm.NewOrm().Update(this, "ComentContent")
	return err
}

func (this *IssueComment) Delete() error {
	_, err := orm.NewOrm().Delete(this)
	return err
}

func QueryCommentByIssueId(issueid int64) ([]*IssueComment, error) {
	var comments []*IssueComment
	_, err := IssueCommentQueryer().Filter("IssueId", issueid).All(&comments)
	if err != nil {
		return nil, err
	} else {
		return comments, nil
	}
}

func QueryCommentByUserId(userid int64) ([]*IssueComment, error) {
	var comments []*IssueComment
	_, err := IssueCommentQueryer().Filter("UserId", userid).All(&comments)
	if err != nil {
		return nil, err
	} else {
		return comments, nil
	}
}
