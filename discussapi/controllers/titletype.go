package controllers

import "log"

type TitleTypeController struct {
	BaseController
}

//@descriptor get the main title type
//@router /createmaintype [post]
func (this *TitleTypeController) CreateMainType()  {
	log.Fatal(string(this.Ctx.Input.RequestBody))
	this.Data["json"]="hello";
	this.ServeJSON()
}