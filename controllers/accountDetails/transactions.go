package accountDetails

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/wongyinlong/hsnNet/models"
	"time"
)

type DelegatorTxController struct {
	beego.Controller
}
type txInfo struct {
	Height int       `json:"height"`
	Hash   string    `json:"hash"`
	Types  string    `json:"types"`
	Result bool      `json:"result"`
	Amount float64   `json:"amount"`
	Fee    float64   `json:"fee"`
	Nums   int       `json:"nums"`
	Time   time.Time `json:"time"`
}
type TxBlocks struct {
	Code string   `json:"code"`
	Data []txInfo `json:"data"`
	Total int `json:"total"`
	Msg  string   `json:"msg"`
}

// @Title
// @Description
// @Success code 0
// @Failure code 1
//@router /
func (dtc *DelegatorTxController) Get() {
	address := dtc.GetString("address")
	page, _ := dtc.GetInt("page", 0)
	size, _ := dtc.GetInt("size", 5)
	if address == "" {
		var errMsg TxBlocks
		errMsg.Data = nil
		errMsg.Code = "1"
		errMsg.Msg = "address can not be empty!"
		dtc.Data["json"] = errMsg
	}else {
		var txList models.Txs
		var txsSet = make([]txInfo, size)
		var respJson TxBlocks
		list,count := txList.GetDelegatorTxs(address, page, size)
		for i, item := range *list {
			txsSet[i].Height = item.Height
			txsSet[i].Hash = item.TxHash
			txsSet[i].Fee = item.Fee
			txsSet[i].Result = item.Result
			txsSet[i].Time = item.TxTime
			txsSet[i].Types = item.Type
			txsSet[i].Nums = item.Plus
			// judegd todo
			if item.Type =="reward"{
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
	var totalAmout float64
	fmt.Println(amounts)
	if len(amounts) <=0 {
		return 0.0
	}else{
		for i:=0;i<len(amounts);i++{
			totalAmout = totalAmout+amounts[i]
		}
	}
	return totalAmout
}

func getRewardAmount(amounts []float64)float64{
	if len(amounts) == 1{
		return amounts[0]
	}
	return 0.0
}