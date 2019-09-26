package actions
//
//import (
//	"encoding/json"
//	"github.com/wongyinlong/hsnNet/conf"
//	"github.com/wongyinlong/hsnNet/models"
//	"go.uber.org/zap"
//	"io/ioutil"
//	"net/http"
//	"strconv"
//	"time"
//)
//
//// get txs list from the listed urls
////http://172.38.8.89:1317/txs?message.action=send
////http://172.38.8.89:1317/txs?message.action=delegate
////http://172.38.8.89:1317/txs?message.action=vote
////http://172.38.8.89:1317/txs?message.action=begin_unbonding
////http://172.38.8.89:1317/txs?message.action=withdraw_delegator_reward
////http://172.38.8.89:1317/txs?message.action=withdraw_validator_commission
////http://172.38.8.89:1317/txs?message.action=multisend
//// .. Unfinished
//
//func GetTxs(config conf.Config, log zap.Logger, chainName string) {
//
//	var Lcd = config.Remote.Lcd
//	var SendUrl = Lcd + "/txs?message.action=send"
//	var DelegateUrl = Lcd + "/txs?message.action=delegate"
//	var VoteUrl = Lcd + "/txs?message.action=vote"
//	var UnDelegateUrl = Lcd + "/txs?message.action=begin_unbonding"
//	var RewardUrl = Lcd + "/txs?message.action=withdraw_delegator_reward"
//	var RewardCommissionUrl = Lcd + "/txs?message.action=withdraw_validator_commission"
//	var MultiSendUrl = Lcd + "/txs?message.action=multisend"
//	// get the transaction judge whether it is stored in the database
//	for {
//		getTxs(SendUrl, log, chainName, "send")
//		getTxs(DelegateUrl, log, chainName, "delegate")
//		getTxs(RewardCommissionUrl, log, chainName, "commission")
//		getTxs(RewardUrl, log, chainName, "reward")
//		getTxs(VoteUrl, log, chainName, "vote")
//		getTxs(UnDelegateUrl, log, chainName, "unbonding")
//		getTxs(MultiSendUrl, log, chainName, "multisend")
//		time.Sleep(time.Second * config.Param.TransactionsInterval) //Avoid frequent request api
//	}
//
//	// Todo: get amount,
//}
//func getTxs(url string, log zap.Logger, chainName string, types string) {
//	page := 1
//	var txsObject Txs
//	for {
//		tempUrl :=url
//		URL := tempUrl + "&page=" + strconv.Itoa(page)
//		resp, err := http.Get(URL)
//		defer resp.Body.Close()
//		if err != nil {
//			log.Error("get"+types+"Txs failed!", zap.String("url", url))
//			return
//		}
//		jsonStr, _ := ioutil.ReadAll(resp.Body)
//		err = json.Unmarshal(jsonStr, &txsObject)
//		if err != nil {
//			log.Error("Unmarshal jsonStr failed!")
//			return
//		}
//		if txsObject.Error!=""{
//			break
//		}
//		var txs models.Txs
//		for _, item := range txsObject.Txs {
//			txhash := item.Txhash
//			flage := txs.CheckHash(txhash)
//			if flage == 0 {
//				txs.Height, _ = strconv.Atoi(item.Height)
//				txs.TxHash = item.Txhash
//				txs.Result = item.Logs[0].Success
//				txs.TxTime = item.Timestamp
//				txs.Time = time.Now()
//				txs.Plus = len(item.Logs) // len(msg,or logs)
//				txs.Type = types
//				for _, amount := range item.Tx.Value.Fee.Amount {
//					if amount.Denom == chainName {
//						txs.Fee = amount.Amount
//					} else {
//						txs.Fee = ""
//					}
//				}
//				// get amout,types
//				if len(item.Logs) > 1 {
//					txs.Amount = ""
//				} else {
//					if types == "send" || types == "delegate" || types == "unbonding" {
//						for _, obj := range item.Events {
//							if obj.Type == "unbond" || obj.Type == "delegate" || obj.Type == "transfer" {
//								for _, innerObj := range obj.Attributes {
//									if innerObj.Key == "amount" {
//										if obj.Type == "delegate" || obj.Type == "unbond" {
//
//											txs.Amount = innerObj.Value + "hsn"
//										} else {
//											txs.Amount = innerObj.Value
//										}
//									}
//								}
//							}
//						}
//					}
//
//					//Todo type multisend
//					//if types == "multisend" {
//					//	//TOdo
//					//	//var index []int
//					//	//index = append(index,0)
//					//	//msg := reflect.ValueOf(item.Tx.Value.Msg).Index(0)
//					//	i :=0
//					//	for _,val := range item.Tx.Value.Msg{
//					//		values := reflect.ValueOf(val).MapIndex(reflect.ValueOf("value"))
//					//		fmt.Println(values)
//					//		inputsInter := reflect.ValueOf(values).Interface()
//					//		//getType := reflect.TypeOf(inputsInter)
//					//		fmt.Println(inputsInter)
//					//		i++
//					//	}
//
//						//length :=  reflect.ValueOf(item.Tx.Value.Msg).Len()
//						//for i:=0;i<length;i++{
//						//	msg := reflect.ValueOf(item.Tx.Value.Msg).Index(i)
//						//
//						//	fmt.Println(OBJS)
//						//
//						//	fmt.Println(msg.MapIndex(reflect.ValueOf("type")))
//						//
//						//}
//						//value :=reflect.ValueOf(msg).FieldByName("value")
//						//outputs :=reflect.ValueOf(value).FieldByName("outputs")
//
//						//for _,inneritem :=range msg
//						//value :=reflect.ValueOf(msg).MapIndex(reflect.ValueOf("outputs"))
//						//outputs :=reflect.ValueOf(value).MapIndex(reflect.ValueOf("outputs"))
//						//coins :=reflect.ValueOf(outputs).MapIndex(reflect.ValueOf("coins"))
//						//for _,inner :=range msg{
//						//
//						//}
//
//					//}
//					if types == "reward" {
//						for _, obj := range item.Events {
//							if obj.Type == "withdraw_rewards" {
//								for _, innerObj := range obj.Attributes {
//									if innerObj.Key == "amount" {
//										txs.Amount = innerObj.Value
//									}
//								}
//							}
//						}
//					}
//
//				}
//
//
//				// get  validators and delegators
//
//				txs.SetInfo()
//			}
//		}
//	page++}
//
//}
//
//
//type Txs struct {
//	TotalCount string `json:"total_count"`
//	Count      string `json:"count"`
//	PageNumber string `json:"page_number"`
//	PageTotal  string `json:"page_total"`
//	Limit      string `json:"limit"`
//	Error string `json:"error"`
//	Txs        []struct {
//		Height string `json:"height"`
//		Txhash string `json:"txhash"`
//		RawLog string `json:"raw_log"`
//		Logs   []struct {
//			MsgIndex int    `json:"msg_index"`
//			Success  bool   `json:"success"`
//			Log      string `json:"log"`
//		} `json:"logs"`
//		GasWanted string `json:"gas_wanted"`
//		GasUsed   string `json:"gas_used"`
//		Events    []struct {
//			Type       string `json:"type"`
//			Attributes []struct {
//				Key   string `json:"key"`
//				Value string `json:"value"`
//			} `json:"attributes"`
//		} `json:"events"`
//		Tx struct {
//			Type  string `json:"type"`
//			Value struct {
//				Msg []interface{} `json:"msg"`
//				Fee struct {
//					Amount []struct {
//						Denom  string `json:"denom"`
//						Amount string `json:"amount"`
//					} `json:"amount"`
//					Gas string `json:"gas"`
//				} `json:"fee"`
//				Signatures []struct {
//					PubKey struct {
//						Type  string `json:"type"`
//						Value string `json:"value"`
//					} `json:"pub_key"`
//					Signature string `json:"signature"`
//				} `json:"signatures"`
//				Memo string `json:"memo"`
//			} `json:"value"`
//		} `json:"tx"`
//		Timestamp time.Time `json:"timestamp"`
//	} `json:"txs"`
//}
