package validatorDetails

import (
	"github.com/astaxie/beego"
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/models/validatorsDetail"
	"strings"
)

type DelegationsController struct {
	beego.Controller
}
type MsgErr struct {
	Code string `json:"code"`
	Data error  `json:"data"`
	Msg  string `json:"msg"`
}
type Msgs struct {
	Code string        `json:"code"`
	Data DelegationMsg `json:"data"`
	Msg  string        `json:"msg"`
}
type DelegationMsg struct {
	TotalDelegations     int           `json:"total_delegation"`
	OneDayAgoDelegations int           `json:"one_day_ago_delegations"`
	Delegations          []Delegations `json:"delegations"`
}
type Delegations struct {
	Address          string  `json:"address"`
	Amount           float64 `json:"amount"`
	AmountPercentage float64 `json:"share"`
}

// @Title Get
// @Description get delegations
// @Success code 0
// @Failure code 1
// @router /
func (dc *DelegationsController) Get() {
	dc.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", dc.Ctx.Request.Header.Get("Origin"))
	address := dc.GetString("address")
	page, _ := dc.GetInt("page", 0)
	size, _ := dc.GetInt("size", 0)
	if size == 0 {
		size = 5
	}
	if address == "" || strings.Index(address, conf.NewConfig().Public.Bech32PrefixValAddr) != 0 {
		var errorMessage MsgErr
		errorMessage.Code = "1"
		errorMessage.Data = nil
		errorMessage.Msg = "Validator address is empty! Or error address!"
		dc.Data["json"] = errorMessage
		dc.ServeJSON()
	}
	var msg Msgs
	var respJson DelegationMsg
	var respJsonDelegations []Delegations
	var respJsonDelegation Delegations
	var delegations validatorsDetail.DelegatorObj
	var baseInfo validatorsDetail.ExtraValidatorInfo
	var validatorDelegationNums validatorsDetail.ValidatorDelegatorNums
	validatorBaseInfo := baseInfo.GetOne(address)

	items, totalDelegations := delegations.GetInfo(address, page, size)
	oneDayAgoDelegations := validatorDelegationNums.GetInfo(address)
	for _, item := range *items {
		respJsonDelegation.Amount = getShares(items, item.DelegatorAddress)
		respJsonDelegation.Address = item.DelegatorAddress
		respJsonDelegation.AmountPercentage = getPercentage(respJsonDelegation.Amount, validatorBaseInfo.TotalToken)
		respJsonDelegations = append(respJsonDelegations, respJsonDelegation)
	}
	respJson.TotalDelegations = totalDelegations
	respJson.OneDayAgoDelegations = totalDelegations - oneDayAgoDelegations
	respJson.Delegations = respJsonDelegations
	msg.Data = respJson
	msg.Code = "0"
	msg.Msg = "OK"
	dc.Data["json"] = msg
	dc.ServeJSON()
}

func getShares(items *[]validatorsDetail.DelegatorObj, address string) float64 {
	var amount float64
	for _, item := range *items {
		if item.DelegatorAddress == address {
			share := item.Shares
			amount = amount + share
		}
	}
	return amount
}
func getPercentage(amout float64, totalToken float64) float64 {

	return amout / totalToken

}
