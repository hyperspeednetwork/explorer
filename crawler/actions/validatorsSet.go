package actions

import (
	"encoding/json"
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/db"
	"github.com/wongyinlong/hsnNet/models"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type validatorsSets struct {
	Result models.ValidatorsSet `json:"result"`
}

func GetValidatorsSet(config conf.Config, log zap.Logger) {
	/*
		获取当前高度， 开始同步区块验证人信息直至当前高度。
		再次获取高度，若高度相同sleep，若不同再次同步
	*/
	// Defer is useless in loops, does not apply defer or organizes part of the code into a function.
	var sets validatorsSets
	URL := config.Remote.Lcd + "/validatorsets/"

	for {
		c:=&http.Client{
			Timeout:time.Second * config.Param.HTTPGetTimeOut,
		}
		aimHeight, ValidatorSetHeight := GetHeight()
		//如果高度差大于110，仅同步最新的100个高度地址
		if aimHeight-ValidatorSetHeight >= 110{
			ValidatorSetHeight = aimHeight-101
		}
		if aimHeight > ValidatorSetHeight {
			for aimHeight > ValidatorSetHeight {
				ValidatorSetHeight = ValidatorSetHeight + 1
				resp, err := c.Get(URL + strconv.Itoa(ValidatorSetHeight))
				if err != nil {
					log.Error("get ValidatorsSet error", zap.String("error",err.Error()))
					ValidatorSetHeight = ValidatorSetHeight - 1
					time.Sleep(time.Second * config.Param.ValidatorsSetsInterval * 2)
				} else {
					jsonStr, _ := ioutil.ReadAll(resp.Body)
					_ = resp.Body.Close()
					_ = json.Unmarshal(jsonStr, &sets)
					intHeight, _ := strconv.Atoi(sets.Result.BlockHeight)
					sets.Result.Height = intHeight
					sets.Result.Time = time.Now()
					sets.Result.SetInfo(log)
				}
				time.Sleep(time.Millisecond * 2)
			}

		}
		time.Sleep(time.Second * config.Param.ValidatorsSetsInterval)
	}
}

func GetHeight() (int, int) {
	// get public's Height,ValidatorsSet's Height
	var session = db.NewDBConn() //db
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	var tempValidatorsSet models.ValidatorsSet
	_ = dbConn.C("validatorsSet").Find(nil).Sort("-height").One(&tempValidatorsSet)
	ValidatorsSetHeight := tempValidatorsSet.Height
	var tempPublic models.Infomation
	_ = dbConn.C("public").Find(nil).Sort("-height").One(&tempPublic)
	PublicHeight := tempPublic.Height
	return PublicHeight, ValidatorsSetHeight
}
