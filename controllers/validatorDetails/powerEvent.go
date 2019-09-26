package validatorDetails

import (
	"github.com/astaxie/beego"
	"github.com/wongyinlong/hsnNet/models"
	"time"
)

type PowerEventController struct {
	beego.Controller
}

type powerEvents struct {
	Height int       `json:"height"`
	Hash   string    `json:"hash"`
	Amount  float64   `json:"amount"`
	Type   string    `json:"type"`
	Time   time.Time `json:"time"`
}

type powerEventMsg struct {
	Code  string        `json:"code"`
	Data  []powerEvents `json:"data"`
	Msg   string        `json:"msg"`
	Total int           `json:"total"`
}
type powerEventErrMsg struct {
	Code string `json:"code"`
	Data error  `json:"data"`
	Msg  string `json:"msg"`
}

// @Title Get
// @Description get txs (delegate and undelegate)
// @Success code 0
// @Failure code 1
// @router /
func (pwc *PowerEventController) Get() {
	pwc.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", pwc.Ctx.Request.Header.Get("Origin"))
	address := pwc.GetString("address", "")
	page, _ := pwc.GetInt("page", 0)
	size, _ := pwc.GetInt("size", 0)
	if address == "" || address[0:10] != "hsnvaloper" {
		var err powerEventErrMsg
		err.Data = nil
		err.Msg = "Address can not be empty! Or error address!"
		err.Code = "1"
		pwc.Data["json"] = err
	} else {
		var txs models.Txs
		// get data form mongodb
		// type redelegate,delegate,unbonding ,validator's address == address ,
		txList, total := txs.GetPowerEventInfo(address, page, size)
		var pe powerEvents
		var msg powerEventMsg
		msg.Total = total
		msg.Code = "0"
		msg.Msg = "OK"
		for _, item := range *txList {
			pe.Time = item.TxTime
			pe.Height = item.Height
			pe.Hash = item.TxHash
			pe.Type = item.Type
			pe.Amount = getTxValidatorAmount(address, item)
			msg.Data = append(msg.Data, pe)
		}
		pwc.Data["json"] = msg

	}
	pwc.ServeJSON()
}
func getTxValidatorAmount(address string, item models.Txs) float64 {
	var tempAmount float64
	if len(item.Amount) == 0 {
		tempAmount = 0.0
		return tempAmount
	}
	for i, addr := range item.ValidatorAddress {
		if addr == address {
			tempAmount = tempAmount + item.Amount[i]
		}
	}
	return tempAmount
}
