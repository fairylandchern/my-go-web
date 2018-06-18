package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

//const M=1024*1024*1024
//var m=M
type Issue struct {
	IssueId      int       `json:"issue_id" orm:"column(issue_id);pk;auto;uinque" `
	UserId       int       `json:"user_id"  orm:"column(user_id)"`
	IssueContent string    `json:"issue_content" orm:"column(issue_content);type(text);size(1073741824)"`
	IssueTitle   string    `json:"issue_title" orm:"column(issue_title);size(100)"`
	Created      time.Time `json:"created,omitempty" orm:"column(created);auto_now_add;type(datetime)"`
	Updated      time.Time `json:"updated,omitempty" orm:"column(updated);auto_now;type(datetime)"`
	TypeId       int       `json:"type_id,omitempty" orm:"column(type_id)" `
}

func init() {
	orm.RegisterModel(new(Issue))
}
func (this *Issue) TableName() string {
	return TableName("issue")
}

func IssueQueryer() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Issue))
}

func (this *Issue) Create() (error, bool) {
	i, err := orm.NewOrm().Insert(this)
	// i, err := orm.NewOrm().InsertOrUpdate(this)
	if err != nil {
		return err, false
	}
	if i == 0 {
		return err, false
	}
	return nil, true
}

func (this *Issue) Update() error {
	_, err := orm.NewOrm().Update(this, "IssueContent", "IssueTitle", "Updated", "TypeId")
	return err
}

func (this *Issue) Delete() error {
	_, err := orm.NewOrm().Delete(this)
	return err
}

func QueryBasicIssue() ([]*Issue, error) {
	var issues []*Issue
	_, err := IssueQueryer().OrderBy("-Updated").All(&issues, "IssueId", "IssueTitle", "Created", "TypeId", "UserId")
	if err != nil {
		return nil, err
	}
	return issues, nil
}

func (this *Issue) QueryDetailIssue() error {
	err := IssueQueryer().Filter("IssueId", this.IssueId).One(this)
	return err
}
