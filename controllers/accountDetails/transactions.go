package accountDetails

import (
	"github.com/astaxie/beego"
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/models"
	"strings"
)

type DelegatorTxController struct {
	beego.Controller
}
type txInfo struct {
	Height int     `json:"height"`
	Hash   string  `json:"hash"`
	Types  string  `json:"types"`
	Result bool    `json:"result"`
	Amount float64 `json:"amount"`
	Fee    float64 `json:"fee"`
	Nums   int     `json:"nums"`
	Time   string  `json:"time"`
}
type TxBlocks struct {
	Code  string   `json:"code"`
	Data  []txInfo `json:"data"`
	Total int      `json:"total"`
	Msg   string   `json:"msg"`
}

// @Title
// @Description
// @Success code 0
// @Failure code 1
//@router /
func (dtc *DelegatorTxController) Get() {
	dtc.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", dtc.Ctx.Request.Header.Get("Origin"))
	address := dtc.GetString("address")
	page, _ := dtc.GetInt("page", 0)
	size, _ := dtc.GetInt("size", 5)
	if address == "" || strings.Index(address, conf.NewConfig().Public.Bech32PrefixAccAddr) != 0 || strings.Index(address, conf.NewConfig().Public.Bech32PrefixValAddr) == 0 {
		var errMsg TxBlocks
		errMsg.Data = nil
		errMsg.Code = "1"
		errMsg.Msg = "Delegator address is empty Or Error address!"
		dtc.Data["json"] = errMsg
	} else {
		var txList models.Txs
		var txsSet = make([]txInfo, size)
		var respJson TxBlocks
		list, count := txList.GetDelegatorTxs(address, page, size)
		for i, item := range *list {
			txsSet[i].Height = item.Height
			txsSet[i].Hash = item.TxHash
			txsSet[i].Fee = item.Fee
			txsSet[i].Result = item.Result
			txsSet[i].Time = item.TxTime
			txsSet[i].Types = item.Type
			txsSet[i].Nums = item.Plus

			if item.Type == "reward" {
				txsSet[i].Amount = getRewardAmount(item.WithDrawRewardAmout)
			} else {
				txsSet[i].Amount = getAmount(item.Amount)
			}

		}
		respJson.Code = "0"
		respJson.Msg = "OK"
		respJson.Total = count
		respJson.Data = txsSet
		dtc.Data["json"] = respJson

	}

	dtc.ServeJSON()
}
func getAmount(amounts []float64) float64 {
	var totalAmount float64
	if len(amounts) <= 0 {
		return 0.0
	} else {
		for i := 0; i < len(amounts); i++ {
			totalAmount = totalAmount + amounts[i]
		}
	}
	return totalAmount
}

func getRewardAmount(amounts []float64) float64 {
	if len(amounts) == 1 {
		return amounts[0]
	}
	return 0.0
}
