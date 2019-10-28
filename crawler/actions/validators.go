package actions

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/crawler/actions/validatorDetails"
	"github.com/wongyinlong/hsnNet/db"
	"github.com/wongyinlong/hsnNet/models"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func GetValidators(config conf.Config, log zap.Logger, ) {
	for {
		// 获取验证人列表集合 unbonding bonded unbonded
		// http://172.38.8.89:1317/staking/validators?status=unbonding&page=1
		// http://172.38.8.89:1317/staking/validators?status=bonded&page=1
		// http://172.38.8.89:1317/staking/validators?status=unbonded&page=1
		//var validatorList ValidatorList

		var validators models.Validators

		var validatorInfos []models.ValidatorInfo
		//var validatorInfo models.ValidatorInfo
		ValidatorsSet := getValidatorsSets(config.Public.ValidatorsSetLimit)

		bondedUrl := config.Remote.Lcd + "/staking/validators?status=bonded"
		unbondedUrl := config.Remote.Lcd + "/staking/validators?status=unbonded"
		unbondingdUrl := config.Remote.Lcd + "/staking/validators?status=unbonding"
		c := &http.Client{
			Timeout: time.Second * config.Param.HTTPGetTimeOut,
		}
		resp, err := c.Get(bondedUrl)
		if err != nil {
			log.Info("Can not get Validators", zap.String("err", err.Error()))
			time.Sleep(time.Second * config.Param.CanNotGetErrorInterval)
			continue
		}
		jsonStrBondedValidators, _ := ioutil.ReadAll(resp.Body)
		resp, err = c.Get(unbondedUrl)
		if err != nil {
			log.Info("Can not get Validators", zap.String("err", err.Error()))
			time.Sleep(time.Second * config.Param.CanNotGetErrorInterval)
			continue
		}
		jsonStrUnBondedValidators, _ := ioutil.ReadAll(resp.Body)
		resp, err = c.Get(unbondingdUrl)
		if err != nil {
			log.Info("Can not get Validators", zap.String("err", err.Error()))
			time.Sleep(time.Second * config.Param.CanNotGetErrorInterval)
			continue
		}
		jsonStrUnBondingValidators, _ := ioutil.ReadAll(resp.Body)
		// get validators basic information
		resp.Body.Close()
		err = json.Unmarshal(jsonStrBondedValidators, &validators)
		if err != nil {
			log.Info("Bonded validator list is empty!")
		} else {
			for _, item := range validators.Result {
				info := dealWithValidatorList(item, config.Public.CoinToVoitingPower, ValidatorsSet, log)
				validatorInfos = append(validatorInfos, info)
			}
		}

		err = json.Unmarshal(jsonStrUnBondedValidators, &validators)

		if err != nil {
			log.Info("UnBonded validator list is empty!")
		} else {
			for _, item := range validators.Result {
				//test
				info := dealWithValidatorList(item, config.Public.CoinToVoitingPower, ValidatorsSet, log)
				validatorInfos = append(validatorInfos, info)

			}
		}
		err = json.Unmarshal(jsonStrUnBondingValidators, &validators)

		if err != nil {
			log.Info("Unbonding validator list is empty!")
		} else {
			for _, item := range validators.Result {
				info := dealWithValidatorList(item, config.Public.CoinToVoitingPower, ValidatorsSet, log)
				validatorInfos = append(validatorInfos, info)
			}
		}
		//validatorInfo.DeleteAllInfo()
		storeValidatorsInfo(&validatorInfos, log)
		time.Sleep(time.Second * config.Param.ValidatorsSetsInterval)
	}
}

func storeValidatorsInfo(vi *[]models.ValidatorInfo, log zap.Logger) {
	for _, item := range *vi {
		// mongo upsert
		item.SetInfo(log)
	}

}

func getAllPledgenTokens() decimal.Decimal {
	/* GET PLEDGEN TOKENS FROM DB*/
	var Info models.Infomation
	session := db.NewDBConn()
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	dbConn.C("public").Find(nil).Sort("-height").One(&Info)
	tokens := strconv.Itoa(Info.PledgeHsn)
	decimalTotalHsn, _ := decimal.NewFromString(tokens)
	return decimalTotalHsn
}
func getUptime(vs *[]models.ValidatorsSet, pbKey string) int {
	count := 0 //记录一百个块中该验证着参与的次数（通过公钥）

	for _, Sets := range *vs {
		for _, item := range Sets.Validators {
			if item.PubKey == pbKey {
				count++
			}
		}
	}
	return count
}

func getValidatorsSets(limit int) *[]models.ValidatorsSet {
	var vSets models.ValidatorsSet
	vs := vSets.GetInfo(limit)
	return vs
}
func dealWithValidatorList(item models.Result, CoinToVoitingPower float32, VS *[]models.ValidatorsSet, log zap.Logger) models.ValidatorInfo {
	//time.Sleep(time.Second * 1) // need to fix panic.被除数可能为0
	go validatorDetails.MakeBaseInfo(item, VS, log)
	go SetValidatorHashAddress(item.OperatorAddress, item.ConsensusPubkey, log)
	var validatorInfo models.ValidatorInfo
	validatorInfo.AKA = item.Description.Moniker // get nick name
	validatorInfo.Status = item.Status
	validatorInfo.Avater = ""                    // avater address
	validatorInfo.ValidatorAddress = item.OperatorAddress
	validatorInfo.Jailed = item.Jailed
	validatorInfo.Commission = item.Commission.CommissionRates.Rate
	othersDelegation, _ := decimal.NewFromString(item.Tokens)
	floatAmount := othersDelegation
	floatCoinToVoitingPower := decimal.NewFromFloat32(CoinToVoitingPower)
	tempAmount := floatAmount.Div(floatCoinToVoitingPower)
	validatorInfo.VotingPower.Amount, _ = tempAmount.Float64()
	// may be has some problem
	tempPledgenTokens := getAllPledgenTokens()
	if tempPledgenTokens.LessThan(decimal.NewFromFloat(1)) {
		tempPledgenTokens = decimal.NewFromFloat(1.0)
	}
	tempPercent := tempAmount.Div(tempPledgenTokens)
	validatorInfo.VotingPower.Percent, _ = tempPercent.Float64()
	validatorInfo.Uptime = getUptime(VS, item.ConsensusPubkey)
	validatorInfo.Time = time.Now()
	return validatorInfo
}
