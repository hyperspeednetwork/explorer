package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers/validatorDetails:DelegationsController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers/validatorDetails:DelegationsController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers/validatorDetails:PowerEventController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers/validatorDetails:PowerEventController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers/validatorDetails:ProposedBlocksController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers/validatorDetails:ProposedBlocksController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers/validatorDetails:VaBaseInfoController"] = append(beego.GlobalControllerRouter["github.com/wongyinlong/hsnNet/controllers/validatorDetails:VaBaseInfoController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
