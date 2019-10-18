package accountDetail

import (
	"encoding/json"
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/logger"
	"github.com/wongyinlong/hsnNet/models"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

type Delegators struct {
	Height string `json:"height"`
	Result []struct {
		DelegatorAddress string `json:"delegator_address"`
		ValidatorAddress string `json:"validator_address"`
		Shares           string `json:"shares"`
		Balance          struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"balance"`
		Reward string `json:"reward"`
		Name   string `json:"name"`
	} `json:"result"`
}
type DelegatorValidatorReward struct {
	Height string `json:"height"`
	Result []struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"result"`
}

func (d *Delegators) GetInfo(address string) *Delegators {
config := conf.NewConfig()
	log := logger.NewLogger()
	var delegators Delegators
	url := config.Remote.Lcd + "/staking/delegators/" + address + "/delegations"
	c:=&http.Client{
		Timeout:time.Second * config.Param.HTTPGetTimeOut,
	}
	resp, err := c.Get(url)

	if err != nil {
		log.Error("Get account's delegations error!", zap.String("error",err.Error()))
	}else {
		defer resp.Body.Close()
	}
	jsonStr, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(jsonStr, &delegators)
	//strBytes := []byte(delegateStr)
	//json.Unmarshal(strBytes,&delegators)
	for index, item := range delegators.Result {
		delegators.Result[index].Name = getName(item.ValidatorAddress)
		delegators.Result[index].Reward = getReward(config.Public.ChainName, config.Remote.Lcd, item.ValidatorAddress, item.DelegatorAddress)
	}
	return &delegators
}

func getReward(tokenName, baseURl, validatorAddress, delegatorAddress string) (string) {
	/*get validator's reward*/
	//distribution/delegators/{delegatorAddr}/rewards/{validatorAddr} Query a delegation reward
	var reward DelegatorValidatorReward
	url := baseURl + "/distribution/delegators/" + delegatorAddress + "/rewards/" + validatorAddress
	c:=&http.Client{
		Timeout:time.Second * conf.NewConfig().Param.HTTPGetTimeOut,
	}
	resp, err := c.Get(url)

	if err != nil {
		return ""
	}else {
		defer resp.Body.Close()
	}
	jsonStr, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(jsonStr, &reward)
	//	仅仅只返回名称与配置文件中名称相符的数据
	for _, item := range reward.Result {
		if item.Denom == tokenName {
			return item.Amount
		}
	}
	return ""
}

func getName(validatorAddress string) string {
	var validatorInfo models.ValidatorInfo
	validator := validatorInfo.GetOne(validatorAddress)
	return validator.AKA
}

//txs []map[string]interface{}
