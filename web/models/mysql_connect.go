package models

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql" //导出mysql数据库操作
)

type connectInfo struct {
	Host     string `json:host`
	Port     int    `json:port`
	Dbtype   string `json:dbtype`
	Dbname   string `json:dbname`
	Username string `json:username`
	Passwd   string `json:passwd`
	Protocol string `json:protocol`
}

var db *sql.DB

func init() {
	db, _ = getMysqlConnect()
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
}

//GetMysqlConnect 获取mysql数据库连接
func getMysqlConnect() (*sql.DB, error) {
	// db, err := sql.Open("mysql", "user:user123456@tcp(127.0.0.1:3306)/discuss")
	conf, err := concatConnectString()
	if err != nil {
		log.Print("获取连接字符串失败")
		return nil, err
	}
	port := strconv.Itoa(conf.Port)
	connectString := conf.Username + ":" + conf.Passwd + "@" + conf.Protocol + "(" + conf.Host + ":" + port + ")" + "/" + conf.Dbname
	db, err := sql.Open(conf.Dbtype, connectString)
	if err != nil {
		log.Print("获取连接失败:", err.Error())
		return nil, err
	}
	return db, err
}

//CloseConnect close the connect with the database
func closeConnect(db *sql.DB) error {
	err := db.Close()
	return err
}

//concatConnectString 从配置文件中获得数据库的连接配置信息
func concatConnectString() (*connectInfo, error) {
	file, err := os.Open("F:/go-test/src/my-go-web/web/conf/mysql.json")
	if err != nil {
		log.Print("打开文件失败", err.Error())
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := connectInfo{}
	err = decoder.Decode(&conf)
	if err != nil {
		log.Print("转义文件内容失败", err.Error())
		return nil, err
	}
	return &conf, err
}
