package models

import (
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/db"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Validators struct {
	Height string   `json:"height"`
	Result []Result `json:"result"`
}
type Result struct {
	OperatorAddress string `json:"operator_address"`
	ConsensusPubkey string `json:"consensus_pubkey"`
	Jailed          bool   `json:"jailed"`
	Status          int    `json:"status"`
	Tokens          string `json:"tokens"`
	DelegatorShares string `json:"delegator_shares"`
	Description     struct {
		Moniker  string `json:"moniker"`
		Identity string `json:"identity"`
		Website  string `json:"website"`
		Details  string `json:"details"`
	} `json:"description"`
	UnbondingHeight string    `json:"unbonding_height"`
	UnbondingTime   time.Time `json:"unbonding_time"`
	Commission      struct {
		CommissionRates struct {
			Rate          string `json:"rate"`
			MaxRate       string `json:"max_rate"`
			MaxChangeRate string `json:"max_change_rate"`
		} `json:"commission_rates"`
		UpdateTime time.Time `json:"update_time"`
	} `json:"commission"`
	MinSelfDelegation string `json:"min_self_delegation"`
}
type ValidatorInfo struct {
	Avater           string `json:"avater"`
	AKA              string `json:"aka"`
	Jailed           bool   `json:"jailed"`
	ValidatorAddress string `json:"validator_address"`
	VotingPower      struct {
		Amount  float64 `json:"amount"`
		Percent float64 `json:"percent"`
	} `json:"voting_power"`
	Cumulative float64   `json:"cumulative"`
	Commission string    `json:"commission"`
	Uptime     int       `json:"uptime"`
	Time       time.Time `json:"time"`
}

func (vi ValidatorInfo) SetInfo(log zap.Logger) {
	session := db.NewDBConn()
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	c := dbConn.C("validators")
	err := c.Insert(&vi)
	if err == nil {
		log.Info("insert validatorInfo success")
	} else {
		log.Error("insert validatorInfo failed")
	}
	defer log.Sync()
}

func (vi ValidatorInfo) GetInfo() *[]ValidatorInfo {
	var list []ValidatorInfo
	session := db.NewDBConn()
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	dbConn.C("validators").Find(nil).Sort("-votingpower.amount").All(&list)
	return &list
}
func (vi ValidatorInfo) DeleteAllInfo() {
	session := db.NewDBConn()
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	dbConn.C("validators").DropCollection()
}
func (vi ValidatorInfo) GetOne(address string) *ValidatorInfo {
	var info ValidatorInfo
	session := db.NewDBConn()
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	dbConn.C("validators").Find(bson.M{"validatoraddress": address}).One(&info)
	return &info
}
