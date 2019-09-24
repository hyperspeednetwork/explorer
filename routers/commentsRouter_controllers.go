package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:BlockController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:BlockController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:DrawingDataController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:DrawingDataController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:test/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:PublicController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:PublicController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:TxDetailControllers"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:TxDetailControllers"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:TxsController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:TxsController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:ValidatorsController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:ValidatorsController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
