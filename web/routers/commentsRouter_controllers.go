package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["my-go-web/web/controllers:DefaultController"] = append(beego.GlobalControllerRouter["my-go-web/web/controllers:DefaultController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"any"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/web/controllers:UserController"] = append(beego.GlobalControllerRouter["my-go-web/web/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/web/controllers:UserController"] = append(beego.GlobalControllerRouter["my-go-web/web/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/reg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
