package actions

import (
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/logger"
	"testing"
)

//func TestGetSendTxs(t *testing.T) {
//	log := logger.NewLogger() // LOG
//	url := "http://172.38.8.89:1317/txs?message.action=send"
//	GetSendTxs(url,log,"hsn")
//}
//
//func TestGetVoteTxs(t *testing.T) {
//	log := logger.NewLogger() // LOG
//	url := "http://172.38.8.89:1317/txs?message.action=withdraw_validator_commission"
//	GetVoteTxs(url,log,"hsn")
//}
func TestGetTxs(t *testing.T) {
	log := logger.NewLogger() // LOG
	config := conf.NewConfig()   // CONFIG
	GetTxs(config,log,"hsn")
}
//func TestGetRewardCommissionTxs(t *testing.T) {
//	log := logger.NewLogger() // LOG
//	url := "http://172.38.8.89:1317/txs?message.action=withdraw_validator_commission"
//	GetRewardCommissionTxs(url,log,"hsn")
//}
//func TestGetMultiSendTxs(t *testing.T) {
//	log := logger.NewLogger() // LOG
//	url := "http://172.38.8.89:1317/txs?message.action=multisend"
//	GetMultiSendTxs(url,log,"hsn")
//}