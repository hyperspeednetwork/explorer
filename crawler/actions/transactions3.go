package actions

import (
	"github.com/bitly/go-simplejson"
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/models"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

// get txs list from the listed urls
//http://172.38.8.89:1317/txs?message.action=send
//http://172.38.8.89:1317/txs?message.action=delegate
//http://172.38.8.89:1317/txs?message.action=vote
//http://172.38.8.89:1317/txs?message.action=begin_unbonding
//http://172.38.8.89:1317/txs?message.action=withdraw_delegator_reward
//http://172.38.8.89:1317/txs?message.action=withdraw_validator_commission
//http://172.38.8.89:1317/txs?message.action=multisend
// .. Unfinished

func GetTxs(config conf.Config, log zap.Logger, chainName string) {

	var Lcd = config.Remote.Lcd
	var SendUrl = Lcd + "/txs?message.action=send"
	var DelegateUrl = Lcd + "/txs?message.action=delegate"
	var VoteUrl = Lcd + "/txs?message.action=vote"
	var UnDelegateUrl = Lcd + "/txs?message.action=begin_unbonding"
	var RewardUrl = Lcd + "/txs?message.action=withdraw_delegator_reward"
	var RewardCommissionUrl = Lcd + "/txs?message.action=withdraw_validator_commission"
	var MultiSendUrl = Lcd + "/txs?message.action=multisend"
	var ReDelegateUrl = Lcd + "/txs?message.action=begin_redelegate"
	//get the transaction judge whether it is stored in the database
	for {
		getTxs(SendUrl, log, chainName, "send")
		getTxs(DelegateUrl, log, chainName, "delegate")
		getTxs(RewardCommissionUrl, log, chainName, "commission")
		getTxs(RewardUrl, log, chainName, "reward")
		getTxs(VoteUrl, log, chainName, "vote")
		getTxs(UnDelegateUrl, log, chainName, "unbonding")
		getTxs(MultiSendUrl, log, chainName, "multisend")
		getTxs(ReDelegateUrl, log, chainName, "redelegate")
		time.Sleep(time.Second * config.Param.TransactionsInterval) //Avoid frequent request api
	}

	// Todo: get amount,
}
func getTxs(url string, log zap.Logger, chainName string, types string) {
	page := 1
	for {
		time.Sleep(time.Second*1)
		tempUrl := url
		URL := tempUrl + "&page=" + strconv.Itoa(page)
		c:=&http.Client{
			Timeout:time.Second * conf.NewConfig().Param.HTTPGetTimeOut,
		}
		resp, err := c.Get(URL)
		if err != nil {
			log.Error("get "+types+" Txs failed!", zap.String("url", url))
			time.Sleep(time.Second * conf.NewConfig().Param.TransactionsInterval)
			continue
		} else {

		}

		var txsInfo models.Txs
		jsonObj, _ := simplejson.NewFromReader(resp.Body)
		_ = resp.Body.Close()
		jsonTxs, _ := jsonObj.Get("txs").Array()
		txsError, _ := jsonObj.Get("error").String()
		if txsError != "" {
			break
		}
		lenTxs := len(jsonTxs)
		for i := 0; i < lenTxs; i++ {
			hash, _ := jsonObj.Get("txs").GetIndex(i).Get("txhash").String()
			flage := txsInfo.CheckHash(hash)
			if flage == 0 {
				height, _ := jsonObj.Get("txs").GetIndex(i).Get("height").String()
				status, _ := jsonObj.Get("txs").GetIndex(i).Get("logs").GetIndex(0).Get("success").Bool()
				txTime, _ := jsonObj.Get("txs").GetIndex(i).Get("timestamp").String()
				feeArray, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("fee").Get("amount").Array()
				var fee float64
				// get fee

				for index := 0; index < len(feeArray); index++ {
					demo, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("fee").Get("amount").GetIndex(index).Get("denom").String()

					if demo == chainName {
						strFee, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("fee").Get("amount").GetIndex(index).Get("amount").String()
						floatFee, _ := strconv.ParseFloat(strFee, 64)
						fee = fee + floatFee
					}
				}
				types := types
				logs, _ := jsonObj.Get("txs").GetIndex(i).Get("logs").Array()
				pluse := len(logs)
				//, amount , validator ,delegator,from ,to
				msgArray, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").Array()
				var realAmount, withDrawRewardAmout, withDrawCommissionAmout []float64
				var delegatorList, validatorList, fromAddress, toAddress, outputsAddress, inputsAddress, voterAddress, options []string
				for index := 0; index < len(msgArray); index++ {
					msgType, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("type").String()
					switch msgType {
					case "cosmos-sdk/MsgSend":
						// get amount,from address
						from, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("from_address").String()
						to, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("to_address").String()
						fromAddress = append(fromAddress, from)
						toAddress = append(toAddress, to)

						amount, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("amount").Array()
						for index := 0; index < len(amount); index++ {
							demo, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("amount").GetIndex(index).Get("denom").String()
							if demo == chainName {
								strAmount, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("amount").GetIndex(index).Get("amount").String()
								floatAmount, _ := strconv.ParseFloat(strAmount, 64)
								realAmount = append(realAmount, floatAmount)
							}

						}
					case "cosmos-sdk/MsgMultiSend":
						//Get input calculation amount,output address
						outputsArray, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("outputs").Array()
						for outputIndex := 0; outputIndex < len(outputsArray); outputIndex++ {
							output, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("outputs").GetIndex(outputIndex).Get("address").String()
							outputsAddress = append(outputsAddress, output)
						}

						inputsArray, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("inputs").Array()
						for index2 := 0; index2 < len(inputsArray); index2++ {
							input, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("inputs").GetIndex(index2).Get("address").String()
							inputsAddress = append(inputsAddress, input)

							coinArray, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("inputs").GetIndex(index2).Get("coins").Array()
							for innerIndex := 0; innerIndex < len(coinArray); innerIndex++ {
								denom, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("inputs").GetIndex(index2).Get("coins").GetIndex(innerIndex).Get("denom").String()
								if denom == chainName {
									strAmount, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("inputs").GetIndex(index2).Get("coins").GetIndex(innerIndex).Get("amount").String()
									floatAmount, _ := strconv.ParseFloat(strAmount, 64)
									realAmount = append(realAmount, floatAmount)
								}

							}
						}
					case "cosmos-sdk/MsgVote":
						// No amount attribute get voter,options
						voter, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("voter").String()
						option, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("option").String()
						voterAddress = append(voterAddress, voter)
						options = append(options, option)
					case "cosmos-sdk/MsgWithdrawValidatorCommission":
						delegator, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("delegator_address").String()
						validator, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("validator_address").String()
						eventsArrery, _ := jsonObj.Get("txs").GetIndex(i).Get("events").Array()
						for iEvent := 0; iEvent < len(eventsArrery); iEvent++ {
							eventsType, _ := jsonObj.Get("txs").GetIndex(i).Get("events").GetIndex(iEvent).Get("type").String()
							if eventsType == "withdraw_commission" {
								attributesArrery, _ := jsonObj.Get("txs").GetIndex(i).Get("events").GetIndex(iEvent).Get("attributes").Array()
								for iAttributes := 0; iAttributes < len(attributesArrery); iAttributes++ {
									value, _ := jsonObj.Get("txs").GetIndex(i).Get("events").GetIndex(iEvent).Get("attributes").GetIndex(iAttributes).Get("value").String()
									if len(value) > 3 && value[len(value)-3:len(value)] == chainName {
										floatAmout, _ := strconv.ParseFloat(value[0:len(value)-3], 64)
										withDrawCommissionAmout = append(withDrawCommissionAmout, floatAmout)
									}
								}
							}
						}
						if validator != ""{
							validatorList = append(validatorList, validator)
						}
						if delegator!=""{
							delegatorList = append(delegatorList, delegator)
						}
					case "cosmos-sdk/MsgWithdrawDelegationReward":
						delegator, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("delegator_address").String()
						validator, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("validator_address").String()
						eventsArrery, _ := jsonObj.Get("txs").GetIndex(i).Get("events").Array()
						for iEvent := 0; iEvent < len(eventsArrery); iEvent++ {
							eventsType, _ := jsonObj.Get("txs").GetIndex(i).Get("events").GetIndex(iEvent).Get("type").String()
							if eventsType == "withdraw_rewards" {
								attributesArrery, _ := jsonObj.Get("txs").GetIndex(i).Get("events").GetIndex(iEvent).Get("attributes").Array()
								for iAttributes := 0; iAttributes < len(attributesArrery); iAttributes++ {
									key, _ := jsonObj.Get("txs").GetIndex(i).Get("events").GetIndex(iEvent).Get("attributes").GetIndex(iAttributes).Get("key").String()
									amountValue, _ := jsonObj.Get("txs").GetIndex(i).Get("events").GetIndex(iEvent).Get("attributes").GetIndex(iAttributes).Get("value").String()

									if len(amountValue) > 3 && amountValue[len(amountValue)-3:len(amountValue)] == chainName && key == "amount" {
										floatAmout, _ := strconv.ParseFloat(amountValue[0:len(amountValue)-3], 64)
										withDrawRewardAmout = append(withDrawRewardAmout, floatAmout)
									}
								}
							}
						}
						if validator != ""{
							validatorList = append(validatorList, validator)
						}
						if delegator!=""{
							delegatorList = append(delegatorList, delegator)
						}

					case "cosmos-sdk/MsgDelegate":
						delegator, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("delegator_address").String()
						validator, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("validator_address").String()
						strAmount, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("amount").Get("amount").String()
						floatAmount, _ := strconv.ParseFloat(strAmount, 64)
						realAmount = append(realAmount, floatAmount)
						validatorList = append(validatorList, validator)
						delegatorList = append(delegatorList, delegator)
					case "cosmos-sdk/MsgBeginRedelegate":
						delegator, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("delegator_address").String()
						validator, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("validator_src_address").String()
						strAmount, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("amount").Get("amount").String()
						floatAmount, _ := strconv.ParseFloat(strAmount, 64)
						realAmount = append(realAmount, floatAmount)
						validatorList = append(validatorList, validator)
						delegatorList = append(delegatorList, delegator)
					case "cosmos-sdk/MsgUndelegate":
						// get amount ,delegatorList, validatorList
						delegator, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("delegator_address").String()
						validator, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("validator_address").String()
						strAmount, _ := jsonObj.Get("txs").GetIndex(i).Get("tx").Get("value").Get("msg").GetIndex(index).Get("value").Get("amount").Get("amount").String()
						floatAmount, _ := strconv.ParseFloat(strAmount, 64)
						realAmount = append(realAmount, floatAmount)
						validatorList = append(validatorList, validator)
						delegatorList = append(delegatorList, delegator)

					}

				}
				txsInfo.Height, _ = strconv.Atoi(height)
				txsInfo.TxHash = hash
				txsInfo.Result = status
				txsInfo.Amount = realAmount
				txsInfo.Plus = pluse
				txsInfo.Fee = fee
				txsInfo.Type = types
				txsInfo.Time = time.Now()
				txsInfo.TxTime, _ = time.Parse(time.RFC3339Nano, txTime) //string to time
				txsInfo.ValidatorAddress = validatorList
				txsInfo.DelegatorAddress = delegatorList
				txsInfo.WithDrawCommissionAmout = withDrawCommissionAmout
				txsInfo.WithDrawRewardAmout = withDrawRewardAmout
				txsInfo.FromAddress = fromAddress
				txsInfo.ToAddress = toAddress
				txsInfo.OutPutsAddress = outputsAddress
				txsInfo.InputsAddress = inputsAddress
				txsInfo.VoterAddress = voterAddress
				txsInfo.Options = options
				txsInfo.SetInfo(log)
				//fmt.Println(fromAddress)
				//fmt.Println(toAddress)
				//fmt.Println(outputsAddress)
				//fmt.Println(inputsAddress)
				//fmt.Println(voterAddress,options)
			}

		}
		page++
	}

}
