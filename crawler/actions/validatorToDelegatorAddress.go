package actions

import (
	"encoding/json"
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/models"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

func SetValidatorDelegatorAddress (config conf.Config, log zap.Logger){
	var genesisValidator models.GenesisFile
	var otherValidator models.CreatValidator
	var mappingValidatorDelegator models.ValidatorAddressAndDelegatorAddress
	genesisAddress := config.GenesisAddress
	URL := config.Remote.Lcd +"/txs?message.action=create_validator"
	c:=&http.Client{
		Timeout:time.Second * config.Param.HTTPGetTimeOut,
	}
	resp,err :=c.Get(genesisAddress)
	if err != nil{
		log.Error("Get genesis File error.",zap.String("URL",genesisAddress))
	}else {
		defer resp.Body.Close()
	}
	jsonStr,_:= ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(jsonStr,&genesisValidator)
	obj := genesisValidator.Result.Genesis.AppState.Genutil.Gentxs
	for _,item :=range obj{
		mappingValidatorDelegator.ValidatorAddress = item.Value.Msg[0].Value.ValidatorAddress
		mappingValidatorDelegator.DelegatorAddress = item.Value.Msg[0].Value.DelegatorAddress
		Sign ,_:=mappingValidatorDelegator.Check(item.Value.Msg[0].Value.ValidatorAddress)
		if Sign ==0{

			mappingValidatorDelegator.Set(log)
		}
	}

	for {
		resp, err = c.Get(URL)
		if err != nil {
			log.Error("Get genesis File error.Retry in 10s", zap.String("URL", genesisAddress))
			time.Sleep(time.Second*10)
			continue
		}else{

		}
		jsonStr, _ = ioutil.ReadAll(resp.Body)
		_ = json.Unmarshal(jsonStr, &otherValidator)
		_ = resp.Body.Close()
		obj := otherValidator.Txs
		for _,item :=range obj{
			mappingValidatorDelegator.ValidatorAddress = item.Tx.Value.Msg[0].Value.ValidatorAddress
			mappingValidatorDelegator.DelegatorAddress = item.Tx.Value.Msg[0].Value.DelegatorAddress
			Sign,_:=mappingValidatorDelegator.Check(item.Tx.Value.Msg[0].Value.ValidatorAddress)
			if Sign ==0{
				mappingValidatorDelegator.Set(log)
			}
		}
	time.Sleep(config.Param.DelegatorInterval)
	}

}
