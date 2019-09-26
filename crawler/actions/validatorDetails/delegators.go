package validatorDetails

import (
	"encoding/json"
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/models"
	"github.com/wongyinlong/hsnNet/models/validatorsDetail"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

/*
validator's delegations
*/
func GetDelegations(config conf.Config, log zap.Logger, ) {
	/* need validator's opAddress*/
	/*get opAddress form db*/
	errorCount := 0
	var delegations validatorsDetail.Delegators
	for {
		delegations.DeleteAllInfo()
		var validators models.ValidatorInfo
		// get validator info
		vaList := validators.GetInfo() //vaList == validators List
		if len(*vaList) == 0 {
			time.Sleep(time.Second * config.Param.DelegationInterval)
			errorCount++
			if errorCount >= 5 {
				log.Error("Get Validators Error")
				time.Sleep(time.Second * config.Param.DelegationInterval *2)
			}
			continue//
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
				log.Error("Get delegations error",  zap.String("error",err.Error()))
				continue
			}
			// unmarshal resp.body
			jsonStr, _ := ioutil.ReadAll(resp.Body)
			_ = resp.Body.Close()
			err = json.Unmarshal(jsonStr, &delegations)
			if err != nil {
				log.Error("Unmarshal delegations Data error!", zap.String("error",err.Error()))
				continue
			}
			// prevent duplicate storage
			//flage := delegations.CheckInfo() //no need
			//fmt.Println(flage)
			delegations.Address = address
			delegations.SetInfo(log) // store data
		}
		time.Sleep(time.Second*config.Param.DelegationInterval)
	}

}
