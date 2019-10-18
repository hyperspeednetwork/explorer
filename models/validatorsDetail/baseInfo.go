package validatorsDetail

import (
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/db"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
	"time"
)

/*

 */

type BasicInfo struct {
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

//validator's website,details,identity ,hsn height, address

type ExtraValidatorInfo struct {
	WebSite         string   `json:"web_site"`
	Details         string   `json:"details"`
	Identity        string   `json:"identity"`
	Address         string   `json:"address"`
	Validator       string   `json:"validator"`
	HsnHeight       string   `json:"hsn_height"`
	TotalToken      float64   `json:"total_token"`
	SelfToken       float64   `json:"self_token"`
	OthersToken     float64   `json:"others_token"`
	MissedBlockList  []MissBLockData`json:"missed_block_list"`
}
type MissBLockData struct {
	Height int `json:"height"`
	State int `json:"state"`
}

type Block struct {
	IntHeight int `json:"int_height"`
}

func (evi *ExtraValidatorInfo) Set(	log zap.Logger) {
	/*found a new validator*/
	session := db.NewDBConn()
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	err :=dbConn.C("detailValidatorBase").Insert(&evi)
	if err !=nil{
		log.Error("Insert ValidatorDetailInfo Error.",zap.String("error",err.Error()))
	}else {
		log.Info("Insert ValidatorDetailInfo Success.")
	}
	defer log.Sync()
}
func (evi *ExtraValidatorInfo) Update(log zap.Logger) {
	/*验证者信息更改*/
	session := db.NewDBConn()
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	err :=dbConn.C("detailValidatorBase").Update(bson.M{"validator":evi.Validator},&evi)
	if err !=nil{
		log.Error("Update ValidatorDetailInfo Error.",zap.String("error",err.Error()))
	}else {
		log.Info("Update ValidatorDetailInfo Success.")
	}
}

func (evi *ExtraValidatorInfo) Check() int{
	session := db.NewDBConn()
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	count,_:=dbConn.C("detailValidatorBase").Find(bson.M{"validator":evi.Validator}).Count()
	return count
}

func(evi *ExtraValidatorInfo)GetOne(address string) *ExtraValidatorInfo{
	var info ExtraValidatorInfo
	session := db.NewDBConn()
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	dbConn.C("detailValidatorBase").Find(bson.M{"validator":address}).One(&info)
	return &info
}