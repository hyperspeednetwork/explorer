package actions

import (
	"encoding/json"
	"fmt"
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
		// Delete validatorList
		/*Temporary use, late change code may be dedicated to maintaining the list of validatros*/
		// 获取验证人列表集合 unbonding bonded unbonded
		// http://172.38.8.89:1317/staking/validators?status=unbonding&page=1
		// http://172.38.8.89:1317/staking/validators?status=bonded&page=1
		// http://172.38.8.89:1317/staking/validators?status=unbonded&page=1
		//var validatorList ValidatorList
		var validators models.Validators
		var validatorInfos []models.ValidatorInfo
		ValidatorsSet := getValidatorsSets(config.Public.ValidatorsSetLimit)
		bondedUrl := config.Remote.Lcd + "/staking/validators?status=bonded"
		unbondedUrl := config.Remote.Lcd + "/staking/validators?status=unbonded"
		unbondingdUrl := config.Remote.Lcd + "/staking/validators?status=unbonding"
		c:=&http.Client{
			Timeout:time.Second * config.Param.HTTPGetTimeOut,
		}
		resp, err := c.Get(bondedUrl)
		if err != nil {
			log.Error("Can not get Validators", zap.String("url", bondedUrl))
			continue
		}
		jsonStrBondedValidators, _ := ioutil.ReadAll(resp.Body)

		resp, err = c.Get(unbondedUrl)
		if err != nil {
			log.Error("Can not get Validators", zap.String("url", unbondedUrl))
			continue
		}
		jsonStrUnBondedValidators, _ := ioutil.ReadAll(resp.Body)

		resp, err = c.Get(unbondingdUrl)

		if err != nil {
			log.Error("Can not get Validators", zap.String("url", unbondingdUrl))
			continue
		}else {
			resp.Body.Close()
		}


		//TODO:解析错误，导致出错。
		jsonStrUnBondingValidators, _ := ioutil.ReadAll(resp.Body)

		// get validators basic information
		err = json.Unmarshal(jsonStrBondedValidators, &validators)
		if err != nil {
			log.Error("Bonded validator list is empty!")
		}
		for _, item := range validators.Result {
			info := dealWithValidatorList(item, config.Public.CoinToVoitingPower, ValidatorsSet,log)
			validatorInfos = append(validatorInfos, info)
		}
		err = json.Unmarshal(jsonStrUnBondedValidators, &validators)
		if err != nil {
			log.Error("UnBonded validator list is empty!")
		}
		for _, item := range validators.Result {
			info := dealWithValidatorList(item, config.Public.CoinToVoitingPower, ValidatorsSet,log)
			validatorInfos = append(validatorInfos, info)
		}

		//TODO:解析错误，导致出错。 if result is empty array, Can not Unmarshal Validators.
		err = json.Unmarshal(jsonStrUnBondingValidators, &validators)
		if err != nil {
			log.Error("Unbonding validator list is empty!", )
		}
		for _, item := range validators.Result {
			info := dealWithValidatorList(item, config.Public.CoinToVoitingPower, ValidatorsSet,log)
			validatorInfos = append(validatorInfos, info)
		}
		deleteValidatorInfo() // delete validators info
		storeValidatorsInfo(&validatorInfos,log)
		time.Sleep(time.Second * config.Param.ValidatorsSetsInterval)
	}
}
func deleteValidatorInfo() {
	var Info models.ValidatorInfo
	Info.DeleteAllInfo()

}
func storeValidatorsInfo(vi *[]models.ValidatorInfo,log zap.Logger) {
	for _, item := range *vi {
		item.SetInfo(log)
	}

}

func getAllPledgenTokens() decimal.Decimal {
	/* GET PLEDGETN TOKESN FROM DB*/
	var Info models.Infomation
	session := db.NewDBConn()
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	dbConn.C("public").Find(nil).Sort("-height").One(&Info)
	tokens := strconv.Itoa(Info.TotalHsn)
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
func dealWithValidatorList(item models.Result, CoinToVoitingPower float32, VS *[]models.ValidatorsSet,log zap.Logger) models.ValidatorInfo {
	time.Sleep(time.Second*1) // need to fix panic.被除数可能为0
	go validatorDetails.MakeBaseInfo(item, VS,log)
	go SetValidatorHashAddress(item.OperatorAddress,item.ConsensusPubkey,log)
	var validatorInfo models.ValidatorInfo
	validatorInfo.AKA = item.Description.Moniker // get nick name
	validatorInfo.Avater = ""                    // avater address
	validatorInfo.ValidatorAddress = item.OperatorAddress
	validatorInfo.Jailed = item.Jailed
	validatorInfo.Commission = item.Commission.CommissionRates.Rate

	selfDelegation, _ := decimal.NewFromString(item.MinSelfDelegation)
	othersDelegation, _ := decimal.NewFromString(item.DelegatorShares)
	floatAmount := selfDelegation.Add(othersDelegation)
	floatCoinToVoitingPower := decimal.NewFromFloat32(CoinToVoitingPower)
	tempAmount := floatAmount.Div(floatCoinToVoitingPower)

	validatorInfo.VotingPower.Amount, _ = tempAmount.Float64()
// may be has some problem
	tempPledgenTokens :=getAllPledgenTokens()
	if tempPledgenTokens.LessThan(decimal.NewFromFloat(1)){
		fmt.Println("tpis",tempPledgenTokens)
		tempPledgenTokens = decimal.NewFromFloat(1.0)
	}
	tempPercent := tempAmount.Div(tempPledgenTokens)

	validatorInfo.VotingPower.Percent, _ = tempPercent.Float64()
	validatorInfo.Uptime = getUptime(VS, item.ConsensusPubkey)
	validatorInfo.Time = time.Now()
	return validatorInfo
}


//func OnlyTest(){
//	errorCount := 0
//	for   {
//	url := "http://172.38.8.89:1317/staking/validators?status=unbonding"
//	resp, _:=http.Get(url)
//	resp.Body.Close()
//	var validators models.Validators
//	jsonStrUnBondingValidators,_:= ioutil.ReadAll(resp.Body)
//	err := json.Unmarshal(jsonStrUnBondingValidators, &validators)
//	if err != nil {
//		fmt.Println("Can not Unmarshal Validators", )
//		errorCount ++
//		fmt.Println(errorCount)
//	}else {
//		fmt.Println("OK")
//	}
//	}
//}