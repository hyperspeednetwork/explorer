package controllers

import (
	"github.com/astaxie/beego"
	"github.com/wongyinlong/hsnNet/models"
)

// Operations about txs
type TxsController struct {
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
	Time   string `json:"time"`
}
type TxBlocks struct {
	Total int `json:"total"`
	Code string   `json:"code"`
	Data []txInfo `json:"data"`
	Msg  string   `json:"msg"`
}

// @Title 获取tx列表
// @Description 默认获取after head的5个tx
// @Success code 0
// @Failure code 1
// @router /
func (txs *TxsController) Get() {
	txs.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", txs.Ctx.Request.Header.Get("Origin"))
	var txList models.Txs
	var respJson TxBlocks
	head, _ := txs.GetInt("head")
	page, _ := txs.GetInt("page")
	size, _ := txs.GetInt("size")
	if size == 0 {
		size = 5
	}
	var txsSet = make([]txInfo, size)
	list ,total := txList.GetInfo(head, page, size)
	for i, item := range list {
		txsSet[i].Height = item.Height
		txsSet[i].Hash = item.TxHash
		txsSet[i].Fee = item.Fee
		txsSet[i].Result = item.Result
		txsSet[i].Time = item.TxTime
		txsSet[i].Types = item.Type
		txsSet[i].Nums = item.Plus
		if item.Type =="reward"{
			txsSet[i].Amount = getRewardAmount(item.WithDrawRewardAmout)
		} else {
			txsSet[i].Amount = getAmount(item.Amount)
		}

	}
	respJson.Code = "0"
	respJson.Msg = "OK"
	respJson.Data = txsSet
	respJson.Total =total
	txs.Data["json"] = respJson
	txs.ServeJSON()
}
func getAmount(amounts []float64) float64 {
	var totalAmout float64
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
	//else {
	//	var aim float64
	//	for _,item := range amounts{
	//		aim = aim+item
	//	}
	//	return aim
	//}

	// 因为不需要该数值，所以返回0
	return 0.0 //
}