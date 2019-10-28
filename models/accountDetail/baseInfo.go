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

/*
struct and operates,store
*/

type Account struct {
	Height string `json:"height"`
	Result struct {
		Type  string `json:"type"`
		Value struct {
			Address string `json:"address"`
			Coins   []struct {
				Denom  string `json:"denom"`
				Amount string `json:"amount"`
			} `json:"coins"`
			PublicKey struct {
				Type  string `json:"type"`
				Value string `json:"value"`
			} `json:"public_key"`
			AccountNumber string `json:"account_number"`
			Sequence      string `json:"sequence"`
		} `json:"value"`
	} `json:"result"`
}

type WithdrawAddress struct {
	Height string `json:"height"`
	Result string `json:"result"`
}

type BaseInfo struct {
	Address       string `json:"address"`
	RewardAddress string `json:"reward_address"`
	Amount        float64 `json:"amount"`
	TotalPrice    float64 `json:"total_price"`
	Price float64 `json:"price"`
}

func (a *Account) GetInfo(address string)(string,string) {



	var account Account
	var Token string
	config := conf.NewConfig()
	url := config.Remote.Lcd + "/auth/accounts/" + address
	jsonStr :=getInfo(url)
	_ = json.Unmarshal(jsonStr,&account)
	amounts := account.Result.Value.Coins
	for _,amount :=range amounts{
		if amount.Denom == config.Public.ChainName{
			Token = amount.Amount
		}
	}

	return account.Result.Value.Address,Token
}

func (wa *WithdrawAddress) GetWithDrawAddress(address string) string{
	var withdrawAddress WithdrawAddress
	config := conf.NewConfig()
	url := config.Remote.Lcd + "/distribution/delegators/" + address + "/withdraw_address"
	jsonStr :=getInfo(url)
	json.Unmarshal(jsonStr,&withdrawAddress)
	return withdrawAddress.Result
}
func getInfo(url string,)[]byte{
	c:=&http.Client{
		Timeout:time.Second * conf.NewConfig().Param.HTTPGetTimeOut,
	}
	resp ,err :=c.Get(url)

	if err!=nil{
		log :=logger.NewLogger()
		log.Error("Get info error!",zap.String("error",err.Error()))
		return nil
	}else {
		defer resp.Body.Close()
	}
	jsonStr,_ :=ioutil.ReadAll(resp.Body)
	return jsonStr

}