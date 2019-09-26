package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/wongyinlong/hsnNet/conf"
	"io/ioutil"
	"net/http"
	"time"
)

type TxsMsgs interface {
}
type TxDetailControllers struct {
	beego.Controller
}
type ErrorTxInfoBlock struct {
	Code string `json:"code"`
	Data error  `json:"data"`
	Msg  string `json:"msg"`
}
type Msgs struct {
	Code string `json:"code"`
	Data TXD    `json:"data"`
	Msg  string `json:"msg"`
}

type TXD struct {
	Height string `json:"height"`
	Txhash string `json:"txhash"`
	Logs   []struct {
		Success bool `json:"success"`
	} `json:"logs"`
	GasWanted string `json:"gas_wanted"`
	GasUsed   string `json:"gas_used"`
	Events    []struct {
		Type       string `json:"type"`
		Attributes []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"attributes"`
	} `json:"events"`
	Tx struct {
		Type  string `json:"type"`
		Value struct {
			Msg []interface{} `json:"msg"`
			Fee struct {
				Amount []struct {
					Denom  string `json:"denom"`
					Amount string `json:"amount"`
				} `json:"amount"`
				Gas string `json:"gas"`
			} `json:"fee"`
			Memo string `json:"memo"`
		} `json:"value"`
	} `json:"tx"`
	Timestamp time.Time `json:"timestamp"`
}

// @Title 获取tx detail
// @Description 通过hash 查询 tx详情
// @Success code 0
// @Failure code 1
// @router /
func (td *TxDetailControllers) Get() {
	td.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", td.Ctx.Request.Header.Get("Origin"))
	config := conf.NewConfig()
	hash := td.GetString("hash")
	var txd TXD
	//tx := item.GetDetail(hash)
	if hash == "" {
		var txd ErrorTxInfoBlock
		txd.Code = "1"
		txd.Msg = "Hash address can not be empty!"
		txd.Data = nil
		td.Data["json"] = txd
	} else {
		var msg Msgs
		url := config.Remote.Lcd + "/txs/" + hash
		c:=&http.Client{
			Timeout:time.Second * config.Param.HTTPGetTimeOut,
		}
		resp, err := c.Get(url)
		if err != nil {
		} else {
			defer resp.Body.Close()
		}
		jsonStr, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(jsonStr, &txd)
		//txd.Tx.Value.Msg = append(txd.Tx.Value.Msg,txd.Events)
		msg.Data = txd
		msg.Code = "0"
		msg.Msg = "OK"
		td.Data["json"] = msg
	}
	td.ServeJSON()
}
