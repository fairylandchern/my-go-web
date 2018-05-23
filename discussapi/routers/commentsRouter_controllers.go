package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:CompanyController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:CompanyController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:DefaultController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:DefaultController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"any"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/regist`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:TitleTypeController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:TitleTypeController"],
		beego.ControllerComments{
			Method: "CreateMainType",
			Router: `/createmaintype`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:TitleTypeController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:TitleTypeController"],
		beego.ControllerComments{
			Method: "GetAllMainType",
			Router: `/getallmaintype`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:UserController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Auth",
			Router: `/auth`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:UserController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:UserController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Registered",
			Router: `/reg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
