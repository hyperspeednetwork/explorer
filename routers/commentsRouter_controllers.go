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

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:BlockTxController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:BlockTxController"],
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

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:ValidatorsController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers:ValidatorsController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
