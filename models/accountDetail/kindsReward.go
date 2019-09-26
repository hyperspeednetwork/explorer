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

type DelegateRewards struct {
	Result struct {
		Total []struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"total"`
	} `json:"result"`
}

func (dr *DelegateRewards) GetDelegateReward(address string) string {
	config := conf.NewConfig()
	log := logger.NewLogger()
	var delegateRewards DelegateRewards
	url := config.Remote.Lcd + "/distribution/delegators/" + address + "/rewards"
	c:=&http.Client{
		Timeout:time.Second * config.Param.HTTPGetTimeOut,
	}
	resp, err := c.Get(url)
	if err != nil {
		log.Error("get delegator rewards error,",zap.String("error",err.Error()))
	} else {
		defer resp.Body.Close()
	}
	jsonBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(jsonBytes, &delegateRewards)
	for _, item := range delegateRewards.Result.Total {
		if item.Denom == config.Public.ChainName {
			return item.Amount
		}

	}
	return "0.0"
}
