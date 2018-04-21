package models

import (
	"log"
	"strconv"
)

type User struct {
	Name   string
	Passwd string
}

//对用户表的增删改查操作
func AddUser() (bool, error) {
	db, err := GetMysqlConnect()
	if err != nil {
		log.Print("获取连接失败")
		return false, err
	}
	stmt, err := db.Prepare("insert into user(user_name,user_pwd) values(?,?)")
	if err != nil {
		log.Print("stmt获取失败", err.Error())
		return false, err
	}
	for i := 0; i < 3; i++ {
		_, err := stmt.Exec("user"+strconv.Itoa(i), "pwd"+strconv.Itoa(i))
		if err != nil {
			log.Print("插入数据失败", err)
			return false, err
		}
	}
	return true, err
}
