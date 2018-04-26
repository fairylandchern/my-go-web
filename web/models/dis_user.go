package models

import (
	"fmt"
	"log"
)

//User 存储用户信息
type User struct {
	UserId int
	Name   string
	Passwd string
}

type UserAction interface {
	AddUser() (bool, error)
	DelUserByUserid() (bool, error)
	UpdateUserByUserID() (bool, error)
	QueryByUserID() (*User, error)
	Query() (*[]User, error)
}

//AddUser 添加用户
func (user *User) AddUser() (bool, error) {
	stmt, err := db.Prepare("insert into user(user_name,user_pwd) values(?,?)")
	defer stmt.Close()
	if err != nil {
		log.Print("stmt获取失败", err.Error())
		return false, err
	}
	_, err = stmt.Exec(user.Name, user.Passwd)
	if err != nil {
		log.Print("插入数据失败", err)
		return false, err
	}
	return true, err
}

//DelUserByUserid 根据用户id删除用户
func (user *User) DelUserByUserid() (bool, error) {
	stmt, err := db.Prepare("delete from user where user_id=?")
	defer stmt.Close()
	if err != nil {
		log.Print("get the stmt failed", err.Error())
		return false, err
	}
	_, err = stmt.Exec(user.UserId)
	if err != nil {
		log.Println("删除数据失败", err.Error())
		return false, err
	}
	return true, err
}

//UpdateUserByUserID 通过用户id更新用户信息
func (user *User) UpdateUserByUserID() (bool, error) {
	stmt, err := db.Prepare("update user set user_name=?,user_pwd=? where  user_id=?")
	defer stmt.Close()
	if err != nil {
		log.Print("get the stmt failed", err.Error())
		return false, err
	}
	_, err = stmt.Exec(user.Name, user.Passwd, user.UserId)
	if err != nil {
		log.Println("更新数据失败", err.Error())
		return false, err
	}
	return true, err
}

//QueryByUserID 根据用户id查询用户信息
func (user *User) QueryByUserID() (*User, error) {
	stmt, err := db.Prepare("select user_name,user_pwd from user where user_id=?")
	defer stmt.Close()
	if err != nil {
		log.Print("query the database failed", err.Error())
		return nil, err
	}
	row := stmt.QueryRow(user.UserId)
	row.Scan(&user.Name, &user.Passwd)
	fmt.Println(user.Name, user.Passwd)
	return user, err
}

//Query 查询所有用户信息
func (user *User) Query() (*[]User, error) {
	rows, err := db.Query("select user_name,user_pwd,user_id")
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		rows.NextResultSet()
	}
	return nil, nil
}
