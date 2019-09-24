package actions

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/models"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)
var tokenPrice string
var fiveMinAgo time.Time

func GetPublic(config conf.Config, log zap.Logger) {
	/*
	Dashboard info ,

			Price            float32   `json:"price"`
			Height           int       `json:"height"`
			PledgeHsn        float32   `json:"pledge_hsn"`
			TotalHsn         float32   `json:"total_hsn"`
			Inflation        float32   `json:"inflation"`
			TotalValidators  int       `json:"total_validators"`
			OnlineValidators int       `json:"online_validators"`
			BlockTime     float64   `json:"block_time"`

	*/
	// todo:
	info := models.NewInfomation()

	for {

		price := getPriceFormDragonex(config, log)

		height, pledgen, total, err := pledgenAndTotalHsn(config, log)
		if err != nil {
			time.Sleep(time.Second * 4)
			log.Error("get height,pledgen,total Error!,retry in 4s.")
			continue
		}
		inflation, err := getInflation(config, log)
		if err != nil {
			time.Sleep(time.Second * 4)
			log.Error("get height,pledgen,total Error!,retry in 4s.")
			continue
		}
		onlineV, totalV, err := getValidators(config, log)
		if err != nil {
			time.Sleep(time.Second * 4)
			log.Error("get height,pledgen,total Error!,retry in 4s.")
			continue
		}
		blockTime, err := getBLockTime(config, log, height)
		if err != nil {
			time.Sleep(time.Second * 4)
			log.Error("get height,pledgen,total Error!,retry in 4s.")
			continue
		}
		err = info.SetInfo(log,price, height, pledgen, total, inflation, totalV, onlineV, blockTime)
		if err != nil {
			log.Error("insert public data error")
			log.Sync()
			time.Sleep(time.Second * 4)
			continue
		}

		defer func() {
			if err := recover(); err != nil {
				log.Error("panic,retry in 10s.")
				time.Sleep(time.Second*10)
			}
		}()
		//fmt.Println(price, height, pledgen, total, inflation, totalV, onlineV, blockTime)
		time.Sleep(time.Second * 4)
	}


}

func getBLockTime(config conf.Config, log zap.Logger, height int) (float64, error) {
	var block models.BlockInfo
	lastHeightUrl := config.Remote.Lcd + "/blocks/" + strconv.Itoa(height)
	aheadHeightUrl := config.Remote.Lcd + "/blocks/" + strconv.Itoa(height-1)
	c:=&http.Client{
		Timeout:time.Second * config.Param.HTTPGetTimeOut,
	}
	resp, err := c.Get(lastHeightUrl)
	if err != nil {
		log.Error("Cannot get block info! ")
		log.Sync()
		return 0.0, errors.New("get block error")
	}
	jsonStr, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(jsonStr, &block)
	if err != nil {
		log.Error("Cannot parse block info! ")
		log.Sync()
		return 0.0, errors.New("parse block error")
	}
	lastHeightTime := block.Block.Header.Time

	resp, err = c.Get(aheadHeightUrl)

	if err != nil {
		log.Error("Cannot get block info! ")
		log.Sync()
		return 0.0, errors.New("get block error")
	}else {
		defer resp.Body.Close()
	}
	jsonStr, _ = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(jsonStr, &block)
	if err != nil {
		log.Error("Cannot parse block info! ")
		log.Sync()
		return 0.0, errors.New("parse block error")
	}
	aheadHeightTime := block.Block.Header.Time
	t1, _ := time.Parse(time.RFC3339Nano, lastHeightTime)
	t2, _ := time.Parse(time.RFC3339Nano, aheadHeightTime)
	blockTime := t1.Sub(t2).Seconds()
	return blockTime, nil
}


func getPriceFormDragonex (config conf.Config, log zap.Logger) string{

	if tokenPrice == "" {
		fmt.Println("get new price,and Time,type1")
		tokenPrice = getPrice()
		fiveMinAgo = time.Now()
	}else {
		fmt.Println("judged time")
		now := time.Now()
		m, _ := time.ParseDuration("-1m")
		fiveMinAgoFromNow := now.Add(m * 5)
		if fiveMinAgo.Before(fiveMinAgoFromNow){
			fmt.Println("get new price,Time and update ,type2")

			tokenPrice = getPrice()
			fiveMinAgo = time.Now()
		}else {
			fmt.Println("in five min ,do nothing")
		}
	}
	fmt.Println(tokenPrice,fiveMinAgo)
	return tokenPrice
}


func getPrice() (string) {
	/*
		5分钟从网站取一次价格
	*/
	var price Price
	c:=&http.Client{
		Timeout:time.Second * conf.NewConfig().Param.HTTPGetTimeOut,
	}
	resp, err := c.Get("https://openapi.dragonex.im/api/v1/market/real/?symbol_id=302")
	if err != nil {
	}else {
		defer  resp.Body.Close()
	}
	jsonStr, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(jsonStr, &price)
	if err != nil {
	}
	return price.Data[0].ClosePrice


}

//func getPrice(config conf.Config, log zap.Logger) (string, error) {
//	/*
//	5分钟从网站取一次价格
//	*/
//	//now := time.Now()
//	//if now <fiveMinsAgo{
//	//	return Price
//	//}else {
//	//
//	//}
//
//	var price Price
//	resp, err := http.Get("https://openapi.dragonex.im/api/v1/market/real/?symbol_id=302")
//	defer resp.Body.Close()
//	if err != nil {
//		log.Error("Cannot get hsn price info! ")
//		return "", errors.New("get hsn price error")
//	}
//	jsonStr, _ := ioutil.ReadAll(resp.Body)
//	err = json.Unmarshal(jsonStr, &price)
//	if err != nil {
//		log.Error("Cannot parse hsn price info! ")
//		return "", errors.New("parse hsn price error")
//	}
//	//err := resp.Body.Close()
//
//	return price.Data[0].ClosePrice, nil
//
//	// return hsn price  https://openapi.dragonex.im/api/v1/market/real/?symbol_id=302
//
//}
func pledgenAndTotalHsn(config conf.Config, log zap.Logger) (int, int, int, error) {
	//return pledge and total http://localhost:1317/staking/pool
	// Cannot specify height
	var pledgenAndTotalHsn PledgenAndTotalHsn
	url := config.Remote.Lcd + "/staking/pool"
	c:=&http.Client{
		Timeout:time.Second * config.Param.HTTPGetTimeOut,
	}
	resp, err := c.Get(url)

	if err != nil {
		log.Error("Cannot get pledge info! ")
		log.Sync()
		return 0, 0, 0, errors.New("get pledge error")
	}else {
		defer resp.Body.Close()
	}
	jsonStr, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(jsonStr, &pledgenAndTotalHsn)
	if err != nil {
		log.Error("Cannot parse pledge info! ")
		log.Sync()
		return 0, 0, 0, errors.New("parse pledge error")
	}
	bonded, _ := strconv.Atoi(pledgenAndTotalHsn.Result.BondedTokens)
	unbonded, _ := strconv.Atoi(pledgenAndTotalHsn.Result.NotBondedTokens)
	total := bonded + unbonded
	height, _ := strconv.Atoi(pledgenAndTotalHsn.Height)
	return height, bonded, total, nil
}

func getInflation(config conf.Config, log zap.Logger) (string, error) {
	// return inflation http://localhost:1317/minting/inflation
	var inflation Inflation
	url := config.Remote.Lcd + "/minting/inflation"
	c:=&http.Client{
		Timeout:time.Second * config.Param.HTTPGetTimeOut,
	}
	resp, err := c.Get(url)

	if err != nil {
		log.Error("Cannot get inflation info! ")
		log.Sync()
		return "", errors.New("get inflation error")
	}else {
		defer resp.Body.Close()
	}
	jsonStr, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(jsonStr, &inflation)
	if err != nil {
		log.Error("Cannot parse inflation info! ")
		log.Sync()
		return "", errors.New("parse inflation error")
	}
	result := inflation.Result
	return result, nil

}
func getValidators(config conf.Config, log zap.Logger) (int, int, error) {
	// bonded, 	unbonding  http://172.38.8.89:1317/staking/validators?status=unbonding&page=1
	//http://172.38.8.89:1317/staking/validators?status=bonded&page=1
	var validators models.Validators
	bondedUrl := config.Remote.Lcd + "/staking/validators?status=bonded&page=1"
	unbondedUrl := config.Remote.Lcd + "/staking/validators?status=unbonded&page=1"
	unbondingdUrl := config.Remote.Lcd + "/staking/validators?status=unbonding&page=1"
	var Jailed int
	c:=&http.Client{
		Timeout:time.Second * config.Param.HTTPGetTimeOut,
	}
	resp, err := c.Get(bondedUrl)
	if err != nil {
		log.Error("Cannot get validator's info! ")
		log.Sync()
		return 0, 0, errors.New("get validators error")
	}
	jsonStr, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(jsonStr, &validators)
	if err != nil {
		log.Error("Cannot parse validators information correctly!")
		log.Sync()
		return 0, 0, errors.New("parse validators information error")
	}
	for _,item := range validators.Result{
		if item.Jailed {
			Jailed = Jailed+1
		}
	}
	bonded := len(validators.Result)

	resp, err = c.Get(unbondingdUrl)
	if err != nil {
		log.Error("Cannot get validator's info! ")
		log.Sync()
		return 0, 0, errors.New("get validators error")
	}
	jsonStr, _ = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(jsonStr, &validators)
	if err != nil {
		log.Error("Cannot parse validators information correctly!")
		log.Sync()
		return 0, 0, errors.New("parse validators information error")
	}
	for _,item := range validators.Result{
		if item.Jailed {
			Jailed = Jailed+1
		}
	}
	unbonding := len(validators.Result)

	resp, err = c.Get(unbondedUrl)

	if err != nil {
		log.Error("Cannot get validator's info! ")
		log.Sync()
		return 0, 0, errors.New("get validators error")
	}else {
		defer resp.Body.Close()
	}
	jsonStr, _ = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(jsonStr, &validators)
	if err != nil {
		log.Error("Cannot parse validators information correctly!")
		log.Sync()
		return 0, 0, errors.New("parse validators information error")
	}
	for _,item := range validators.Result{
		if item.Jailed {
			Jailed = Jailed+1
		}
	}
	unbonded := len(validators.Result)
	total := bonded + unbonded + unbonding
	alive := total -Jailed
	return alive, total, nil
}

type Inflation struct {
	Height string `json:"height"`
	Result string `json:"result"`
}

type PledgenAndTotalHsn struct {
	Height string `json:"height"`
	Result struct {
		NotBondedTokens string `json:"not_bonded_tokens"`
		BondedTokens    string `json:"bonded_tokens"`
	} `json:"result"`
}

type Price struct {
	Ok   bool `json:"ok"`
	Code int  `json:"code"`
	Data []struct {
		ClosePrice      string `json:"close_price"`
		CurrentVolume   string `json:"current_volume"`
		MaxPrice        string `json:"max_price"`
		MinPrice        string `json:"min_price"`
		OpenPrice       string `json:"open_price"`
		PriceBase       string `json:"price_base"`
		PriceChange     string `json:"price_change"`
		PriceChangeRate string `json:"price_change_rate"`
		Timestamp       int    `json:"timestamp"`
		TotalAmount     string `json:"total_amount"`
		TotalVolume     string `json:"total_volume"`
		UsdtAmount      string `json:"usdt_amount"`
		SymbolID        int    `json:"symbol_id"`
	} `json:"data"`
}

