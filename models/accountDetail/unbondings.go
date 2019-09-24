package accountDetail

import (
	"encoding/json"
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/logger"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

type Unbonding struct {
	Height string `json:"height"`
	Result []struct {
		DelegatorAddress string `json:"delegator_address"`
		ValidatorAddress string `json:"validator_address"`
		Entries          []struct {
			CreationHeight string    `json:"creation_height"`
			CompletionTime time.Time `json:"completion_time"`
			InitialBalance string    `json:"initial_balance"`
			Balance        string    `json:"balance"`
		} `json:"entries"`
		Name string `json:"name"`
	} `json:"result"`
}

func (u *Unbonding) GetInfo(address string) *Unbonding {

	config := conf.NewConfig()
	log := logger.NewLogger()
	var unbonding Unbonding
	url := config.Remote.Lcd + "/staking/delegators/" + address + "/unbonding_delegations"
	c:=&http.Client{
		Timeout:time.Second * config.Param.HTTPGetTimeOut,
	}
	resp, err := c.Get(url)

	if err != nil {
		log.Error("Get account's delegations error!", zap.String("URL", url))
	} else {
		defer resp.Body.Close()
	}
	jsonStr, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(jsonStr, &unbonding)
	//testUnbondingBytes := []byte(unbondingStr)
	//json.Unmarshal(testUnbondingBytes, &unbonding)
	for index, item := range unbonding.Result {
		unbonding.Result[index].Name = getName(item.ValidatorAddress)
	}
	return &unbonding
}
