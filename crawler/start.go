package crawler

import (
	"errors"
	"github.com/wongyinlong/hsnNet/conf"

	"github.com/wongyinlong/hsnNet/crawler/actions"
	"github.com/wongyinlong/hsnNet/crawler/actions/validatorDetails"
	"github.com/wongyinlong/hsnNet/logger"
)

func OnStart(){
	config := conf.NewConfig() // CONFIG
	log := logger.NewLogger()  // LOG
	if config.Remote.Lcd != "" && config.Remote.Rpc!=""{
		// check url
		go actions.GetPublic(config, log)
		go actions.GetBlock(config, log)
		go actions.GetValidators(config, log)
		go actions.GetValidatorsSet(config, log)
		go actions.GetTxs(config, log, config.Public.ChainName)
		go validatorDetails.GetDelegations(config, log)
		go validatorDetails.GetDelegations2(config, log)
		go actions.SetValidatorDelegatorAddress(config, log)
		//go validatorBlocksTable()
	} else {
		log.Info("LCD OR RPC address is empty")
		panic(errors.New("LCD/ RPC address is empty,"))
	}

}


