package models

import "github.com/astaxie/beego/orm"

// type User struct {
// 	Id       int64     `json:"id" orm:"column(id);pk;auto;unique"`
// 	Phone    string    `json:"phone" orm:"column(phone);unique;size(11)"`
// 	Nickname string    `json:"nickname" orm:"column(nickname);unique;size(40);"`
// 	Password string    `json:"password" orm:"column(password);size(40)"`
// 	Created  time.Time `json:"-" orm:"column(create_at);auto_now_add;type(datetime)"`
// 	Updated  time.Time `json:"-" orm:"column(update_at);auto_now;type(datetime)"`
// }

type Follow struct {
	FollowId int64 `json:"follow_id" orm:"column(follow_id);pk;auto;unique"`
	FollowerId int64 `json:"follower_id" orm:"column(follower_id);"`
	BeFollwerId int64 `json:"be_follwer_id" orm:"column(be_follwer_id);"`
}

func init() {
	orm.RegisterModel(new (Follow));
}

func (f *Follow)TableName() string {
	return TableName("follow");
}
