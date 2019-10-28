package validatorDetails

import (
	"encoding/json"
	"errors"
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/models"
	"github.com/wongyinlong/hsnNet/models/validatorsDetail"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func GetDelegatorNums(config conf.Config, log zap.Logger, ) {
	for {

		utcH := time.Now().UTC().Hour()
		utcM := time.Now().UTC().Minute()
		h, _ := time.ParseDuration(strconv.Itoa(23-utcH) + "h")
		m, _ := time.ParseDuration(strconv.Itoa(50-utcM) + "m")
		time.Sleep(h)
		time.Sleep(m)
		err := updateInsertDelegatorData(config, log)
		if err != nil {
			//validator list is empty.
			continue
		}
		time.Sleep(time.Hour * 1)
	}
}
func updateInsertDelegatorData(config conf.Config, log zap.Logger, ) error {

	var validators models.ValidatorInfo
	var delegations validatorsDetail.Delegators
	var validatorDelegationNums validatorsDetail.ValidatorDelegatorNums
	// get validator info
	vaList := validators.GetInfo() //vaList == validators List
	if len(*vaList) == 0 {
		time.Sleep(time.Second * config.Param.DelegationInterval)
		return errors.New("validator list is empty")
	}
	c := &http.Client{
		Timeout: time.Second * config.Param.HTTPGetTimeOut,
	}
	for _, item := range *vaList {
		address := item.ValidatorAddress
		url := config.Remote.Lcd + "/staking/validators/" + address + "/delegations"
		resp, err := c.Get(url)
		if err != nil {
			log.Error("Get delegations error", zap.String("error", err.Error()))
			time.Sleep(time.Second * config.Param.CanNotGetErrorInterval)
			continue
		}
		// unmarshal resp.body
		jsonStr, _ := ioutil.ReadAll(resp.Body)
		_ = resp.Body.Close()
		err = json.Unmarshal(jsonStr, &delegations)
		if err != nil {
			log.Error("Unmarshal delegations Data error!", zap.String("error", err.Error()))
			time.Sleep(time.Second * config.Param.CanNotGetErrorInterval)
			continue
		}
		validatorDelegationNums.ValidatorAddress = address
		validatorDelegationNums.DelegatorNums = len(delegations.Result)
		validatorDelegationNums.SetInfo(log)
	}
	return nil
}


