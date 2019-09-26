package validatorsDetail

import (
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/db"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Delegators struct {
	Address string    `json:"address"`
	Height  string    `json:"height"`
	Result  []Result  `json:"result"`;
	Time    time.Time `json:"time"`
}
type Balance struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}
type Result struct {
	DelegatorAddress string  `json:"delegator_address"`
	ValidatorAddress string  `json:"validator_address"`
	Shares           string  `json:"shares"`
	Balance          Balance `json:"balance"`
}

func (d *Delegators) GetInfo(address string)( *Delegators,int,int){
	var tempObj,tempObj2 Delegators
	session := db.NewDBConn()
	dbConn := session.DB(conf.NewConfig().DBName)
	h, _ := time.ParseDuration("-24d")
	dbConn.C("delegations").Find(bson.M{"address": address}).Sort("-time").One(&tempObj)
	dbConn.C("delegations").Find(bson.M{"address": address,"time":bson.M{"$let":time.Now().Add(h)}}).One(&tempObj2)
	return &tempObj,len(tempObj.Result),len(tempObj2.Result)
}
func (d *Delegators) SetInfo(log zap.Logger) {
	d.Time = time.Now()
	session := db.NewDBConn()
	dbConn := session.DB(conf.NewConfig().DBName)
	err := dbConn.C("delegations").Insert(&d)
	if err != nil {
		log.Error("Insert delegations error!",zap.String("error",err.Error()))
	} else {
		log.Info("Insert delegations success")
	}
	defer log.Sync()
}

// no need
//func (d *Delegators) CheckInfo() int {
//	var tempDelgators Delegators
//	session := db.NewDBConn()
//	log := logger.NewLogger()
//	dbConn := session.DB("hsnhub_db_dev")
//	err := dbConn.C("delegations").Find(nil).One(&tempDelgators)
//	if err != nil {
//		log.Error("db read error")
//	}
//	if tempDelgators.Height < d.Height {
//		return 0
//	}
//	return 1
//}
func (d *Delegators) DeleteAllInfo() {
	session := db.NewDBConn()
	dbConn := session.DB(conf.NewConfig().DBName)
	h, _ := time.ParseDuration("-24d")
	err := dbConn.C("delegations").Remove(bson.M{"time":bson.M{"lt":time.Now().Add(h)}})
	if err != nil {

	}
}
