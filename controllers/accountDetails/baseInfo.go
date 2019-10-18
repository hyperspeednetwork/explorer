package accountDetails

import (
	"github.com/astaxie/beego"
	"github.com/shopspring/decimal"
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/models"
	"github.com/wongyinlong/hsnNet/models/accountDetail"
	"strings"
)

type BaseInfoController struct {
	beego.Controller
}
type baseInfoerrMsg struct {
	Data error `json:"data"`
	Msg  string `json:"msg"`
	Code string `json:"code"`
}
type baseInfoMsg struct {
	Data accountDetail.BaseInfo `json:"data"`
	Msg   string`json:"msg"`
	Code string `json:"code"`
}
/**/
// @Title
// @Description
// @Success code 0
// @Failure code 1
//@router /
func (bic *BaseInfoController) Get() {
	bic.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", bic.Ctx.Request.Header.Get("Origin"))
	address := bic.GetString("address")
	if address == ""||  strings.Index(address,conf.NewConfig().Public.Bech32PrefixAccAddr)!=0 || strings.Index(address,conf.NewConfig().Public.Bech32PrefixValAddr)==0{
		var msg baseInfoerrMsg
		msg.Data=nil
		msg.Msg="Delegator address is empty Or Error address!"
		msg.Code ="1"
		bic.Data["json"]=msg
	}else {
		//获取验证人账户信息和获取提款地址
		var baseInfo accountDetail.BaseInfo
		var account accountDetail.Account
		var withdrawAddress accountDetail.WithdrawAddress
		var price models.Infomation
		decimalPrice,_:=decimal.NewFromString(price.GetInfo().Price)
		var msg baseInfoMsg
		baseInfo.Address,_ =account.GetInfo(address)
		decimalTotalAmount := GetAllKindsAmount(address).Data.TotalAmount[0]
		baseInfo.Amount ,_= decimalTotalAmount.Float64()
		baseInfo.RewardAddress = withdrawAddress.GetWithDrawAddress(address)
		baseInfo.TotalPrice,_= decimalTotalAmount.Mul(decimalPrice).Float64()
		baseInfo.Price ,_= decimalPrice.Float64()
		msg.Data = baseInfo
		msg.Code = "0"
		msg.Msg = "OK"
		bic.Data["json"] = msg
	}
	bic.ServeJSON()

}
func getDeciamlRewardAmount(address string)decimal.Decimal{
	var delegateReward accountDetail.DelegateRewards
	amount := delegateReward.GetDelegateReward(address)
	decimalAmount ,_:= decimal.NewFromString(amount)
	return decimalAmount
}
