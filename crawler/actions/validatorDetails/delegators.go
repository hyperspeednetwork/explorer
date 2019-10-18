package validatorDetails

import (
	"encoding/json"
	"fmt"
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/models"
	"github.com/wongyinlong/hsnNet/models/validatorsDetail"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

/*
validator's delegations
*/
func GetDelegations(config conf.Config, log zap.Logger, ) {
	/* need validator's opAddress*/
	/*get opAddress form db*/
	sign := 0
	errorCount := 0
	var delegationObj validatorsDetail.DelegatorObj
	var delegations validatorsDetail.Delegators
	for {
		//用于标志delegations信息，删除无用的信息。
		if sign >100 {
			sign =0
		}

		var validators models.ValidatorInfo
		// get validator info
		vaList := validators.GetInfo() //vaList == validators List
		if len(*vaList) == 0 {
			time.Sleep(time.Second * config.Param.DelegationInterval)
			errorCount++
			if errorCount >= 5 {
				log.Error("Get Validators Error")
				time.Sleep(time.Second * config.Param.DelegationInterval * 2)
			}
			continue
		} else {
			if errorCount > 0 {
				errorCount--
			}
		}
		c:=&http.Client{
			Timeout:time.Second * config.Param.HTTPGetTimeOut,
		}
		for _, item := range *vaList {
			address := item.ValidatorAddress
			url := config.Remote.Lcd + "/staking/validators/" + address + "/delegations"
			resp, err := c.Get(url)

			if err != nil {
				log.Error("Get delegations error", zap.String("error",err.Error()))
				continue
			}
			// unmarshal resp.body
			jsonStr, _ := ioutil.ReadAll(resp.Body)
			_ = resp.Body.Close()
			err = json.Unmarshal(jsonStr, &delegations)

			if err != nil {
				log.Error("Unmarshal delegations Data error!",zap.String("error",err.Error()))
				continue
			}

			for _, item := range delegations.Result {
				delegationObj.Shares ,_=strconv.ParseFloat(item.Balance.Amount,64)
				delegationObj.DelegatorAddress = item.DelegatorAddress
				delegationObj.Address = item.ValidatorAddress
				delegationObj.Sign = sign
				delegationObj.Time = time.Now()
				delegationObj.SetInfo(log)
			}

		}
		fmt.Println(sign)
		delegationObj.DeleteInfo(sign)
		time.Sleep(time.Second * config.Param.DelegationInterval)
		sign ++
	}

}
