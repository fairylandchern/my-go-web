package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["json"] = Response{
		Status: 404,
		Msg:  "Not Found",
	}
	c.ServeJSON()
}
func (c *ErrorController) Error401() {
	c.Data["json"] = Response{
		Status: 401,
		Msg:  "Permission denied",
	}
	c.ServeJSON()
}
func (c *ErrorController) Error403() {
	c.Data["json"] = Response{
		Status: 403,
		Msg:  "Forbidden",
	}
	c.ServeJSON()
}
