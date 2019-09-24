package accountDetails

import (
	"github.com/astaxie/beego"
	"github.com/shopspring/decimal"
	"github.com/wongyinlong/hsnNet/models"
	"github.com/wongyinlong/hsnNet/models/accountDetail"
	"strconv"
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

// @Title
// @Description
// @Success code 0
// @Failure code 1
//@router /
func (bic *BaseInfoController) Get() {
	address := bic.GetString("address")
	if address==""{
		var msg baseInfoerrMsg
		msg.Data=nil
		msg.Msg="Delegator address can not be empty!"
		msg.Code ="1"
		bic.Data["json"]=msg
	}else {
		//获取验证人账户信息和获取提款地址
		var baseInfo accountDetail.BaseInfo
		var account accountDetail.Account
		var withdrawAddress accountDetail.WithdrawAddress
		var price models.Infomation
		floatPrice,_:=strconv.ParseFloat(price.GetInfo().Price,64)
		var msg baseInfoMsg
		var strAmount string
		baseInfo.Address,strAmount =account.GetInfo(address)
		decimalAmount,_:= decimal.NewFromString(strAmount)
		decimalRewardAmount := getDeciamlRewardAmount(address)
		decimalTotalAmount := decimalRewardAmount.Add(decimalAmount)
		baseInfo.Amount ,_= decimalTotalAmount.Float64()
		baseInfo.RewardAddress = withdrawAddress.GetWithDrawAddress(address)
		baseInfo.TotalPrice= baseInfo.Amount*floatPrice
		baseInfo.Price = floatPrice
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
