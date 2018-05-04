//// @APIVersion 1.0.0
//// @Title beego Test API
//// @Description beego has a very cool tools to autogenerate documents for your API
//// @Contact astaxie@gmail.com
//// @TermsOfServiceUrl http://beego.me/
//// @License Apache 2.0
//// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"my-go-web/discussapi/controllers"

	"github.com/astaxie/beego"
)

// 使用注释路由
func init() {

	beego.Router("/", &controllers.DefaultController{}, "*:GetAll")
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
//
//func init() {
//	//ns := beego.NewNamespace("/v1",
//	//	beego.NSNamespace("/object",
//	//		beego.NSInclude(
//	//			&controllers.ObjectController{},
//	//		),
//	//	),
//	//	beego.NSNamespace("/user",
//	//		beego.NSInclude(
//	//			&controllers.UserController{},
//	//		),
//	//	),
//	//)
//
//	beego.AddNamespace(ns)
//}