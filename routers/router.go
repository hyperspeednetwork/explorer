// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/wongyinlong/hsnNet/controllers"
	ad "github.com/wongyinlong/hsnNet/controllers/accountDetails"
	vd "github.com/wongyinlong/hsnNet/controllers/validatorDetails"
)

func init() {

	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/public",
			beego.NSInclude(
				&controllers.PublicController{},
			),
		),
		beego.NSNamespace("/drawing",
			beego.NSInclude(
				&controllers.DrawingDataController{},
			),
		),
		beego.NSNamespace("/blocks",
			beego.NSInclude(
				&controllers.BlockController{},
			),
		),
		beego.NSNamespace("/txs",
			beego.NSInclude(
				&controllers.TxsController{},
			),
		),
		beego.NSNamespace("/tx",
			beego.NSInclude(
				&controllers.TxDetailControllers{},
			),
		),
		beego.NSNamespace("/blockTxs",
			beego.NSInclude(
				&controllers.BlockTxController{},
			),
		),
		beego.NSNamespace("/validators",
			beego.NSInclude(
				&controllers.ValidatorsController{},
			),
		),
		beego.NSNamespace("/validatorBase",
			beego.NSInclude(
				&vd.VaBaseInfoController{},
			),
		),
		beego.NSNamespace("/validatorDelegations",
			beego.NSInclude(
				&vd.DelegationsController{},
			),
		),
		beego.NSNamespace("/validatorPowerEvent",
			beego.NSInclude(
				&vd.PowerEventController{},
			),
		),
		beego.NSNamespace("/validatorProposedBlock",
			beego.NSInclude(
				&vd.ProposedBlocksController{},
			),
		),
		beego.NSNamespace("/accountInfo",
			beego.NSInclude(
				&ad.BaseInfoController{},
			),
		),
		beego.NSNamespace("/delegators",
			beego.NSInclude(
				&ad.DeleatorsController{},
			),
		),
		beego.NSNamespace("/unbonding",
			beego.NSInclude(
				&ad.UnbondingsController{},
			),
		),
		beego.NSNamespace("/delegatorTxs",
			beego.NSInclude(
				&ad.DelegatorTxController{},
			),
		),
		beego.NSNamespace("/delegatorAllKindsReward",
			beego.NSInclude(
				&ad.KindsRewardController{},
			),
		),
	)

	beego.AddNamespace(ns)
}
