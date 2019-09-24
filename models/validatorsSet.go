package models

import (
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/db"
	"go.uber.org/zap"
	"time"
)

type ValidatorsSet struct {
	BlockHeight string `json:"block_height"`
	Height int `json:"height"`
	Time time.Time `json:"time"`
	Validators  []struct {
		Address     string `json:"address"`
		PubKey      string `json:"pub_key"`
		VotingPower string `json:"voting_power"`
	} `json:"validators"`
}

func (vs ValidatorsSet) SetInfo(log zap.Logger) {

	session := db.NewDBConn()
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	c := dbConn.C("validatorsSet")
	err := c.Insert(&vs)
	if err != nil {
		log.Error("ValidatorsSet insert error")
	} else {
		log.Info("ValidatorsSet insert success",zap.Int("height",vs.Height))
	}
	defer log.Sync()
}

// limit conf.go -> ValidatorsSetLimit
// 获取最新的100个区块验证着集合 limit默认是  conf.go -> ValidatorsSetLimit
func (vs ValidatorsSet) GetInfo(limit int) *[]ValidatorsSet{
	var ValidatorsSets []ValidatorsSet
	session := db.NewDBConn()
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	dbConn.C("validatorsSet").Find(nil).Sort("-height").Limit(limit).All(&ValidatorsSets)
	return &ValidatorsSets
}

//// check height 	防止重复存储
//func (vs ValidatorsSet) CheckInfo(height int ) int{
//	var session = db.NewDBConn() //db
//	defer session.Close()
//	dbConn := session.DB("hsnhub_db_dev")
//	var tempValidatorsSet ValidatorsSet
//	dbConn.C("validatorsSet").Find(bson.M{"blockheight": height}).One(&tempValidatorsSet)
//	if tempValidatorsSet.BlockHeight == 0 {
//		return 0
//	} else {
//		return 1
//	}
//}
