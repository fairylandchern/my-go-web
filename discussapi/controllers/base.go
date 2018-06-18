package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

//Response 结构体
type Response struct {
	Status STATUS        `json:"status"`
	Msg string      `json:"msg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

