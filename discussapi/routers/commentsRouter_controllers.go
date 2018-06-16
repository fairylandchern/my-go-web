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

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:FollowController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:FollowController"],
		beego.ControllerComments{
			Method: "AddFollower",
			Router: `/addfollower`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:FollowController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:FollowController"],
		beego.ControllerComments{
			Method: "CountBeFollower",
			Router: `/countbefollwer`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:FollowController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:FollowController"],
		beego.ControllerComments{
			Method: "CountFollower",
			Router: `/countfollwer`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:FollowController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:FollowController"],
		beego.ControllerComments{
			Method: "DelBeFollower",
			Router: `/delbefollower`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:FollowController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:FollowController"],
		beego.ControllerComments{
			Method: "Getfans",
			Router: `/getfans`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:FollowController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:FollowController"],
		beego.ControllerComments{
			Method: "GetFollowers",
			Router: `/getfollowers`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:IssueCommentController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:IssueCommentController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:IssueCommentController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:IssueCommentController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:IssueCommentController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:IssueCommentController"],
		beego.ControllerComments{
			Method: "QueryByIssue",
			Router: `/querybyissue/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:IssueCommentController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:IssueCommentController"],
		beego.ControllerComments{
			Method: "QueryByUser",
			Router: `/querybyuser/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:IssueCommentController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:IssueCommentController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:IssueController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:IssueController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/create`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:IssueController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:IssueController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:IssueController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:IssueController"],
		beego.ControllerComments{
			Method: "QueryBasic",
			Router: `/querybasic`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:IssueController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:IssueController"],
		beego.ControllerComments{
			Method: "QueryDetail",
			Router: `/querydetail`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:IssueController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:IssueController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
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
			Method: "Delete",
			Router: `/delete`,
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

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:TitleTypeController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:TitleTypeController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
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
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["my-go-web/discussapi/controllers:UserController"] = append(beego.GlobalControllerRouter["my-go-web/discussapi/controllers:UserController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
