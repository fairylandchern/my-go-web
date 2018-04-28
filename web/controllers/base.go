package contorllers

import (
	"sync"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	mutex sync.Mutex
}

type RespData struct {
	Status bool
	Data   interface{}
}
//Response 结构体
type Response struct {
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}

//Response 结构体
type ErrResponse struct {
	Errcode int         `json:"errcode"`
	Errmsg  interface{} `json:"errmsg"`
}

func (base *BaseController) setdata(status bool, data interface{}) {
	respdata := RespData{
		Status: status,
		Data:   data}
	base.Data["json"] = &respdata
	base.ServeJSON()
}
