package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers/accountDetails:BaseInfoController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers/accountDetails:BaseInfoController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers/accountDetails:DeleatorsController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers/accountDetails:DeleatorsController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers/accountDetails:DelegatorTxController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers/accountDetails:DelegatorTxController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers/accountDetails:KindsRewardController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers/accountDetails:KindsRewardController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers/accountDetails:UnbondingsController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers/accountDetails:UnbondingsController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
