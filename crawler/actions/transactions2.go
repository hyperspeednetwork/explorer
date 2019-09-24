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
//func GetTxs2(config conf.Config, log zap.Logger, chainName string) {
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
//		GetSendTxs(SendUrl, log, chainName)
//		GetDelegateTxs(DelegateUrl, log, chainName)
//		GetRewardCommissionTxs(RewardCommissionUrl, log, chainName)
//		GetRewardTxs(RewardUrl, log, chainName)
//		GetVoteTxs(VoteUrl, log, chainName)
//		GetUndelegateTxs(UnDelegateUrl, log, chainName)
//		GetMultiSendTxs(MultiSendUrl, log, chainName)
//		// TODO  get MUTIL Send Txs
//		time.Sleep(time.Second * config.Param.TransactionsInterval) //Avoid frequent request api
//	}
//	// TODO : GET TXS
//	// Todo: get amount,
//}
//
//func GetMultiSendTxs2(url string, log zap.Logger, chainName string) {
//	var multiSendTxs MultiSendTxs
//	resp, err := http.Get(url)
//	defer resp.Body.Close()
//	if err != nil {
//		log.Error("get SendTxs failed!", zap.String("url", url))
//		return
//	}
//	jsonStr, _ := ioutil.ReadAll(resp.Body)
//	err = json.Unmarshal(jsonStr, &multiSendTxs)
//	if err != nil {
//		log.Error("Unmarshal jsonStr failed!")
//		return
//	}
//	// get the transaction, judge whether it is stored in the database
//	var txs models.Txs
//	for _, item := range multiSendTxs.Txs {
//		txhash := item.Txhash
//		flage := txs.CheckHash(txhash)
//		if flage == 0 {
//			txs.Height, _ = strconv.Atoi(item.Height)
//			txs.TxHash = item.Txhash
//			txs.Result = item.Logs[0].Success
//			txs.Gas.Wanted = item.GasWanted
//			txs.Gas.Used = item.GasUsed
//			txs.Memo = item.Tx.Value.Memo
//			txs.TxTime = item.Timestamp
//			txs.Time = time.Now()
//			txs.Plus = 2 // MultiSend has inputs and outputs
//			txs.Type = "multisend"
//			txs.MultiSendMsgs = item.Tx.Value.Msg
//
//			for _, amount := range item.Tx.Value.Fee.Amount {
//				if amount.Denom == chainName {
//					txs.Fee = amount.Amount
//				} else {
//					txs.Fee = ""
//				}
//			}
//			txs.SetInfo()
//		}
//	}
//
//}
//
//func GetRewardCommissionTxs2(url string, log zap.Logger, chainName string) {
//	var rewardCommission RewardCommission
//	resp, err := http.Get(url)
//	defer resp.Body.Close()
//	if err != nil {
//		log.Error("get SendTxs failed!", zap.String("url", url))
//		return
//	}
//	jsonStr, _ := ioutil.ReadAll(resp.Body)
//	err = json.Unmarshal(jsonStr, &rewardCommission)
//	if err != nil {
//		log.Error("Unmarshal jsonStr failed!")
//		return
//	}
//	// get the transaction, judge whether it is stored in the database
//	var txs models.Txs
//	for _, item := range rewardCommission.Txs {
//		txhash := item.Txhash
//		flage := txs.CheckHash(txhash)
//		if flage == 0 {
//			txs.Height, _ = strconv.Atoi(item.Height)
//			txs.TxHash = item.Txhash
//			txs.Result = item.Logs[0].Success
//			txs.Gas.Wanted = item.GasWanted
//			txs.Gas.Used = item.GasUsed
//			txs.Memo = item.Tx.Value.Memo
//			txs.TxTime = item.Timestamp
//			txs.Time = time.Now()
//
//			txs.Type = "rewardCommission"
//			txs.RewardCommisionMsgs = item.Tx.Value.Msg
//			for _, amount := range item.Tx.Value.Fee.Amount {
//				if amount.Denom == chainName {
//					txs.Fee = amount.Amount
//				} else {
//					txs.Fee = ""
//				}
//			}
//
//			// get reward amount
//			// get commission amount
//			var amountValue []string
//			var amountValue2 []string
//			for _, event := range item.Events {
//				if event.Type == "withdraw_rewards" {
//					for _, attribute := range event.Attributes {
//						if attribute.Key == "amount" {
//							amountValue = append(amountValue, attribute.Value)
//						}
//					}
//				}
//				if event.Type == "withdraw_commission" {
//					for _, attribute := range event.Attributes {
//						if attribute.Key == "amount" {
//							amountValue2 = append(amountValue2, attribute.Value)
//						}
//					}
//				}
//			}
//			countReward := 0
//			countCommission := 0
//			for index, _ := range txs.RewardCommisionMsgs {
//				// set reward amount
//				if txs.RewardCommisionMsgs[index].Type == "cosmos-sdk/MsgWithdrawDelegationReward" {
//					txs.RewardCommisionMsgs[index].Value.Amount = amountValue[countReward]
//					countReward++
//				}
//				// set commission amount
//				if txs.RewardCommisionMsgs[index].Type == "cosmos-sdk/MsgWithdrawValidatorCommission" {
//					txs.RewardCommisionMsgs[index].Value.Amount = amountValue2[countCommission]
//					countCommission++
//				}
//			}
//			//todo :remove the reward with out amount.
//			var tempMsgs []models.GetRewardCommissionMsg
//			for index, _ := range txs.RewardCommisionMsgs {
//				// Remove the reward with out amount.
//				if txs.RewardCommisionMsgs[index].Value.Amount == "" && txs.RewardCommisionMsgs[index].Type == "cosmos-sdk/MsgWithdrawDelegationReward" {
//				} else {
//					tempMsgs = append(tempMsgs, txs.RewardCommisionMsgs[index])
//				}
//			}
//			txs.RewardCommisionMsgs = tempMsgs
//			txs.Plus = len(txs.RewardCommisionMsgs)
//			txs.SetInfo()
//		}
//
//	}
//}
//
//func GetSendTxs2(url string, log zap.Logger, chainName string) {
//	var sendTxs SendTxs
//	resp, err := http.Get(url)
//	defer resp.Body.Close()
//	if err != nil {
//		log.Error("get SendTxs failed!", zap.String("url", url))
//		return
//	}
//	jsonStr, _ := ioutil.ReadAll(resp.Body)
//	err = json.Unmarshal(jsonStr, &sendTxs)
//	if err != nil {
//		log.Error("Unmarshal jsonStr failed!")
//		return
//	}
//	// get the transaction, judge whether it is stored in the database
//	var txs models.Txs
//	for _, item := range sendTxs.Txs {
//		txhash := item.Txhash
//		flage := txs.CheckHash(txhash)
//		if flage == 0 {
//			txs.Height, _ = strconv.Atoi(item.Height)
//			txs.TxHash = item.Txhash
//			txs.Result = item.Logs[0].Success
//			txs.Gas.Wanted = item.GasWanted
//			txs.Gas.Used = item.GasUsed
//			txs.Memo = item.Tx.Value.Memo
//			txs.TxTime = item.Timestamp
//			txs.Time = time.Now()
//			txs.Plus = len(item.Tx.Value.Msg)
//			txs.Type = "send"
//			txs.SendMsgs = item.Tx.Value.Msg
//			for _, amount := range item.Tx.Value.Fee.Amount {
//				if amount.Denom == chainName {
//					txs.Fee = amount.Amount
//				} else {
//					txs.Fee = ""
//				}
//			}
//			if txs.Plus == 1 {
//				txs.Amount = item.Tx.Value.Msg[0].Value.Amount[0].Amount
//			}
//			txs.SetInfo()
//		}
//	}
//
//}
//
//func GetDelegateTxs2(url string, log zap.Logger, chainName string) {
//	var delegateTxs DelegateTxs
//	resp, err := http.Get(url)
//	defer resp.Body.Close()
//	if err != nil {
//		log.Error("get DelegateTxs failed", zap.String("url", url))
//		return
//	}
//	jsonStr, _ := ioutil.ReadAll(resp.Body)
//	err = json.Unmarshal(jsonStr, &delegateTxs)
//	if err != nil {
//		log.Error("Unmarshal jsonStr failed!")
//		return
//	}
//	// get the transaction, judge whether it is stored in the database
//	var txs models.Txs
//	for _, item := range delegateTxs.Txs {
//		txhash := item.Txhash
//		flage := txs.CheckHash(txhash)
//		if flage == 0 {
//			txs.Height, _ = strconv.Atoi(item.Height)
//			txs.TxHash = item.Txhash
//			txs.Result = item.Logs[0].Success
//			txs.Gas.Wanted = item.GasWanted
//			txs.Gas.Used = item.GasUsed
//			txs.Memo = item.Tx.Value.Memo
//			txs.TxTime = item.Timestamp
//			txs.Time = time.Now()
//			txs.Plus = len(item.Tx.Value.Msg)
//			txs.Type = "delegate"
//			txs.DelegateMsgs = item.Tx.Value.Msg
//			for _, amount := range item.Tx.Value.Fee.Amount {
//				if amount.Denom == chainName {
//					txs.Fee = amount.Amount
//				} else {
//					txs.Fee = ""
//				}
//			}
//			if txs.Plus == 1 {
//				txs.Amount = item.Tx.Value.Msg[0].Value.Amount.Amount
//			}
//			txs.SetInfo()
//		}
//	}
//
//}
//
//func GetVoteTxs2(url string, log zap.Logger, chainName string) {
//	var voteTxs VoteTxs
//	resp, err := http.Get(url)
//	defer resp.Body.Close()
//	if err != nil {
//		log.Error("get VoteTxs failed", zap.String("url", url))
//		return
//	}
//	jsonStr, _ := ioutil.ReadAll(resp.Body)
//	err = json.Unmarshal(jsonStr, &voteTxs)
//	if err != nil {
//		log.Error("Unmarshal jsonStr failed!")
//		return
//	}
//
//	// get the transaction, judge whether it is stored in the database
//	var txs models.Txs
//	for _, item := range voteTxs.Txs {
//		txhash := item.Txhash
//		flage := txs.CheckHash(txhash)
//		if flage == 0 {
//			txs.Height, _ = strconv.Atoi(item.Height)
//			txs.TxHash = item.Txhash
//			txs.Result = item.Logs[0].Success
//			txs.Gas.Wanted = item.GasWanted
//			txs.Gas.Used = item.GasUsed
//			txs.Memo = item.Tx.Value.Memo
//			txs.TxTime = item.Timestamp
//			txs.Time = time.Now()
//			txs.Plus = len(item.Tx.Value.Msg)
//			txs.Type = "vote"
//			txs.VoteMsgs = item.Tx.Value.Msg
//			for _, amount := range item.Tx.Value.Fee.Amount {
//				if amount.Denom == chainName {
//					txs.Fee = amount.Amount
//				} else {
//					txs.Fee = ""
//				}
//			}
//			txs.SetInfo()
//		}
//	}
//
//}
//
//func GetUndelegateTxs2(url string, log zap.Logger, chainName string) {
//	var undelegateTxs UnDelegateTxs
//	resp, err := http.Get(url)
//	defer resp.Body.Close()
//	if err != nil {
//		log.Error("get UndelegateTxs failed", zap.String("url", url))
//		return
//	}
//	jsonStr, _ := ioutil.ReadAll(resp.Body)
//	err = json.Unmarshal(jsonStr, &undelegateTxs)
//	if err != nil {
//		log.Error("Unmarshal jsonStr failed!")
//		return
//	}
//	// get the transaction, judge whether it is stored in the database
//	var txs models.Txs
//	for _, item := range undelegateTxs.Txs {
//		txhash := item.Txhash
//		flage := txs.CheckHash(txhash)
//		if flage == 0 {
//			txs.Height, _ = strconv.Atoi(item.Height)
//			txs.TxHash = item.Txhash
//			txs.Result = item.Logs[0].Success
//			txs.Gas.Wanted = item.GasWanted
//			txs.Gas.Used = item.GasUsed
//			txs.Memo = item.Tx.Value.Memo
//			txs.TxTime = item.Timestamp
//			txs.Time = time.Now()
//			txs.Plus = len(item.Tx.Value.Msg)
//			txs.Type = "UnDelegate"
//			txs.UndelegateMsgs = item.Tx.Value.Msg
//			for _, amount := range item.Tx.Value.Fee.Amount {
//				if amount.Denom == chainName {
//					txs.Fee = amount.Amount
//				} else {
//					txs.Fee = ""
//				}
//			}
//			if txs.Plus == 1 {
//				txs.Amount = item.Tx.Value.Msg[0].Value.Amount.Amount
//			}
//			txs.SetInfo()
//		}
//	}
//}
//
//func GetRewardTxs2(url string, log zap.Logger, chainName string) {
//	var rewardTxs RewardTxs
//	resp, err := http.Get(url)
//	defer resp.Body.Close()
//	if err != nil {
//		log.Error("get RewardTxs failed", zap.String("url", url))
//		return
//	}
//	jsonStr, _ := ioutil.ReadAll(resp.Body)
//	err = json.Unmarshal(jsonStr, &rewardTxs)
//	if err != nil {
//		log.Error("Unmarshal jsonStr failed!")
//		return
//	}
//	// get the transaction, judge whether it is stored in the database
//	var txs models.Txs
//	for _, item := range rewardTxs.Txs {
//		txhash := item.Txhash
//		flage := txs.CheckHash(txhash)
//		if flage == 0 {
//			txs.Height, _ = strconv.Atoi(item.Height)
//			txs.TxHash = item.Txhash
//			txs.Result = item.Logs[0].Success
//			txs.Gas.Wanted = item.GasWanted
//			txs.Gas.Used = item.GasUsed
//			txs.Memo = item.Tx.Value.Memo
//			txs.TxTime = item.Timestamp
//			txs.Time = time.Now()
//			txs.Plus = len(item.Tx.Value.Msg)
//			txs.Type = "reward"
//			txs.GetRewardMsgs = item.Tx.Value.Msg
//
//			// get amount
//			var amoutVlaue []string
//			for _, event := range item.Events {
//				if event.Type == "withdraw_rewards" {
//					for _, attribute := range event.Attributes {
//						if attribute.Key == "amount" {
//							amoutVlaue = append(amoutVlaue, attribute.Value)
//						}
//					}
//				}
//			}
//			// set amount
//			for index, value := range amoutVlaue {
//				txs.GetRewardMsgs[index].Value.Amount = value
//			}
//			for _, amount := range item.Tx.Value.Fee.Amount {
//				if amount.Denom == chainName {
//					txs.Fee = amount.Amount
//				} else {
//					txs.Fee = ""
//				}
//			}
//			if txs.Plus == 1 {
//				txs.Amount = txs.GetRewardMsgs[0].Value.Amount
//			}
//			txs.SetInfo()
//		}
//	}
//}
//
////type VoteTxs struct {
////	Txs []struct {
////		Height string `json:"height"`
////		Txhash string `json:"txhash"`
////		RawLog string `json:"raw_log"`
////		Logs   []struct {
////			MsgIndex int    `json:"msg_index"`
////			Success  bool   `json:"success"`
////			Log      string `json:"log"`
////		} `json:"logs"`
////		GasWanted string `json:"gas_wanted"`
////		GasUsed   string `json:"gas_used"`
////		Events    []struct {
////			Type       string `json:"type"`
////			Attributes []struct {
////				Key   string `json:"key"`
////				Value string `json:"value"`
////			} `json:"attributes"`
////		} `json:"events"`
////		Tx struct {
////			Type  string `json:"type"`
////			Value struct {
////				Msg []models.VoteMsg `json:"msg"`
////				Fee struct {
////					Amount []struct {
////						Denom  string `json:"denom"`
////						Amount string `json:"amount"`
////					} `json:"amount"`
////					Gas string `json:"gas"`
////				} `json:"fee"`
////				Signatures []struct {
////					PubKey struct {
////						Type  string `json:"type"`
////						Value string `json:"value"`
////					} `json:"pub_key"`
////					Signature string `json:"signature"`
////				} `json:"signatures"`
////				Memo string `json:"memo"`
////			} `json:"value"`
////		} `json:"tx"`
////		Timestamp time.Time `json:"timestamp"`
////	} `json:"txs"`
////}
////
////type SendTxs struct {
////	Txs []struct {
////		Height string `json:"height"`
////		Txhash string `json:"txhash"`
////		RawLog string `json:"raw_log"`
////		Logs   []struct {
////			MsgIndex int    `json:"msg_index"`
////			Success  bool   `json:"success"`
////			Log      string `json:"log"`
////		} `json:"logs"`
////		GasWanted string `json:"gas_wanted"`
////		GasUsed   string `json:"gas_used"`
////		Events    []struct {
////			Type       string `json:"type"`
////			Attributes []struct {
////				Key   string `json:"key"`
////				Value string `json:"value"`
////			} `json:"attributes"`
////		} `json:"events"`
////		Tx struct {
////			Type  string `json:"type"`
////			Value struct {
////				Msg []models.Msg `json:"msg"`
////				Fee struct {
////					Amount []struct {
////						Denom  string `json:"denom"`
////						Amount string `json:"amount"`
////					} `json:"amount"`
////				} `json:"fee"`
////				Signatures []struct {
////					PubKey struct {
////						Type  string `json:"type"`
////						Value string `json:"value"`
////					} `json:"pub_key"`
////					Signature string `json:"signature"`
////				} `json:"signatures"`
////				Memo string `json:"memo"`
////			} `json:"value"`
////		} `json:"tx"`
////		Timestamp time.Time `json:"timestamp"`
////	} `json:"txs"`
////}
////
////type DelegateTxs struct {
////	Txs []struct {
////		Height string `json:"height"`
////		Txhash string `json:"txhash"`
////		RawLog string `json:"raw_log"`
////		Logs   []struct {
////			MsgIndex int    `json:"msg_index"`
////			Success  bool   `json:"success"`
////			Log      string `json:"log"`
////		} `json:"logs"`
////		GasWanted string `json:"gas_wanted"`
////		GasUsed   string `json:"gas_used"`
////		Events    []struct {
////			Type       string `json:"type"`
////			Attributes []struct {
////				Key   string `json:"key"`
////				Value string `json:"value"`
////			} `json:"attributes"`
////		} `json:"events"`
////		Tx struct {
////			Type  string `json:"type"`
////			Value struct {
////				Msg []models.DelegateMsg `json:"msg"`
////				Fee struct {
////					Amount []struct {
////						Denom  string `json:"denom"`
////						Amount string `json:"amount"`
////					} `json:"amount"`
////					Gas string `json:"gas"`
////				} `json:"fee"`
////				Signatures []struct {
////					PubKey struct {
////						Type  string `json:"type"`
////						Value string `json:"value"`
////					} `json:"pub_key"`
////					Signature string `json:"signature"`
////				} `json:"signatures"`
////				Memo string `json:"memo"`
////			} `json:"value"`
////		} `json:"tx"`
////		Timestamp time.Time `json:"timestamp"`
////	} `json:"txs"`
////}
////
////type UnDelegateTxs struct {
////	Txs []struct {
////		Height string `json:"height"`
////		Txhash string `json:"txhash"`
////		Data   string `json:"data"`
////		RawLog string `json:"raw_log"`
////		Logs   []struct {
////			MsgIndex int    `json:"msg_index"`
////			Success  bool   `json:"success"`
////			Log      string `json:"log"`
////		} `json:"logs"`
////		GasWanted string `json:"gas_wanted"`
////		GasUsed   string `json:"gas_used"`
////		Events    []struct {
////			Type       string `json:"type"`
////			Attributes []struct {
////				Key   string `json:"key"`
////				Value string `json:"value"`
////			} `json:"attributes"`
////		} `json:"events"`
////		Tx struct {
////			Type  string `json:"type"`
////			Value struct {
////				Msg []models.UndelegateMsg `json:"msg"`
////				Fee struct {
////					Amount []struct {
////						Denom  string `json:"denom"`
////						Amount string `json:"amount"`
////					} `json:"amount"`
////					Gas string `json:"gas"`
////				} `json:"fee"`
////				Signatures []struct {
////					PubKey struct {
////						Type  string `json:"type"`
////						Value string `json:"value"`
////					} `json:"pub_key"`
////					Signature string `json:"signature"`
////				} `json:"signatures"`
////				Memo string `json:"memo"`
////			} `json:"value"`
////		} `json:"tx"`
////		Timestamp time.Time `json:"timestamp"`
////	} `json:"txs"`
////}
////
////type RewardTxs struct {
////	Txs []struct {
////		Height string `json:"height"`
////		Txhash string `json:"txhash"`
////		RawLog string `json:"raw_log"`
////		Logs   []struct {
////			MsgIndex int    `json:"msg_index"`
////			Success  bool   `json:"success"`
////			Log      string `json:"log"`
////		} `json:"logs"`
////		GasWanted string `json:"gas_wanted"`
////		GasUsed   string `json:"gas_used"`
////		Events    []struct {
////			Type       string `json:"type"`
////			Attributes []struct {
////				Key   string `json:"key"`
////				Value string `json:"value"`
////			} `json:"attributes"`
////		} `json:"events"`
////		Tx struct {
////			Type  string `json:"type"`
////			Value struct {
////				Msg []models.GetRewardMsg `json:"msg"`
////				Fee struct {
////					Amount []struct {
////						Denom  string `json:"denom"`
////						Amount string `json:"amount"`
////					} `json:"amount"`
////					Gas string `json:"gas"`
////				} `json:"fee"`
////				Signatures []struct {
////					PubKey struct {
////						Type  string `json:"type"`
////						Value string `json:"value"`
////					} `json:"pub_key"`
////					Signature string `json:"signature"`
////				} `json:"signatures"`
////				Memo string `json:"memo"`
////			} `json:"value"`
////		} `json:"tx"`
////		Timestamp time.Time `json:"timestamp"`
////	} `json:"txs"`
////}
////
////type RewardCommission struct {
////	Txs [] struct {
////		Height string `json:"height"`
////		Txhash string `json:"txhash"`
////		RawLog string `json:"raw_log"`
////		Logs   []struct {
////			MsgIndex int    `json:"msg_index"`
////			Success  bool   `json:"success"`
////			Log      string `json:"log"`
////		} `json:"logs"`
////		GasWanted string `json:"gas_wanted"`
////		GasUsed   string `json:"gas_used"`
////		Events    []struct {
////			Type       string `json:"type"`
////			Attributes []struct {
////				Key   string `json:"key"`
////				Value string `json:"value"`
////			} `json:"attributes"`
////		} `json:"events"`
////		Tx struct {
////			Type  string `json:"type"`
////			Value struct {
////				Msg []models.GetRewardCommissionMsg `json:"msg"`
////				Fee struct {
////					Amount []struct {
////						Denom  string `json:"denom"`
////						Amount string `json:"amount"`
////					} `json:"amount"`
////					Gas string `json:"gas"`
////				} `json:"fee"`
////				Signatures []struct {
////					PubKey struct {
////						Type  string `json:"type"`
////						Value string `json:"value"`
////					} `json:"pub_key"`
////					Signature string `json:"signature"`
////				} `json:"signatures"`
////				Memo string `json:"memo"`
////			} `json:"value"`
////		} `json:"tx"`
////		Timestamp time.Time `json:"timestamp"`
////	} `json:"txs"`
////}
////type MultiSendTxs struct {
////	TotalCount string `json:"total_count"`
////	Count      string `json:"count"`
////	PageNumber string `json:"page_number"`
////	PageTotal  string `json:"page_total"`
////	Limit      string `json:"limit"`
////	Txs        []struct {
////		Height string `json:"height"`
////		Txhash string `json:"txhash"`
////		RawLog string `json:"raw_log"`
////		Logs   []struct {
////			MsgIndex int    `json:"msg_index"`
////			Success  bool   `json:"success"`
////			Log      string `json:"log"`
////		} `json:"logs"`
////		GasWanted string `json:"gas_wanted"`
////		GasUsed   string `json:"gas_used"`
////		Events    []struct {
////			Type       string `json:"type"`
////			Attributes []struct {
////				Key   string `json:"key"`
////				Value string `json:"value"`
////			} `json:"attributes"`
////		} `json:"events"`
////		Tx struct {
////			Type  string `json:"type"`
////			Value struct {
////				Msg []models.MultiSendMsg `json:"msg"`
////				Fee struct {
////					Amount []struct {
////						Denom  string `json:"denom"`
////						Amount string `json:"amount"`
////					} `json:"amount"`
////					Gas string `json:"gas"`
////				} `json:"fee"`
////				Signatures []struct {
////					PubKey struct {
////						Type  string `json:"type"`
////						Value string `json:"value"`
////					} `json:"pub_key"`
////					Signature string `json:"signature"`
////				} `json:"signatures"`
////				Memo string `json:"memo"`
////			} `json:"value"`
////		} `json:"tx"`
////		Timestamp time.Time `json:"timestamp"`
////	} `json:"txs"`
////}
