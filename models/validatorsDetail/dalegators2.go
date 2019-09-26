package validatorsDetail

import (
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/db"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Delegators2 struct {
	Address string    `json:"address"`
	Height  string    `json:"height"`
	DelegatorAddress string `json:"delegator_address"`
	Shares string `json:"shares"`
	Time    time.Time `json:"time"`
}

func (d *Delegators2) GetInfo(address string,page int, size int)(*[]Delegators2,int,int){
	var tempObj = make([]Delegators2,size)
	var maxSizeObj Delegators

	session := db.NewDBConn()
	dbConn := session.DB(conf.NewConfig().DBName)
	h, _ := time.ParseDuration("-24h")
	dbConn.C("delegations").Find(bson.M{"address": address}).Sort("-time").One(&maxSizeObj)  // get maxsize
	maxSize := len(maxSizeObj.Result)
	//TODO : user maxSize := len(maxSizeObj.Result)
	//maxSize := 100
	if maxSize<size{
		size = maxSize
		page = 0
	}
	if maxSize>5&& maxSize<page*size{
		pages := maxSize/size  //total pages
		if pages<page{
			page = pages-1
		}
	}
	totalDelegations := maxSize
	oneDayAgoDelegations,_:=dbConn.C("delegations2").Find(bson.M{"address": address,"time":bson.M{"$let":time.Now().Add(h)}}).Count()
	dbConn.C("delegations2").Find(bson.M{"address": address}).Sort("-time").Skip(page*size).Limit(size).All(&tempObj)
	return &tempObj,totalDelegations,oneDayAgoDelegations
}
func (d *Delegators2) SetInfo(log zap.Logger) {
	d.Time = time.Now()
	session := db.NewDBConn()

	dbConn := session.DB(conf.NewConfig().DBName)
	err := dbConn.C("delegations2").Insert(&d)
	if err != nil {
		log.Error("Insert delegations error!",zap.String("error",err.Error()))
	} else {
		log.Info("Insert delegations success")
	}
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

func (d *Delegators2) DeleteInfo() {
	session := db.NewDBConn()
	dbConn := session.DB(conf.NewConfig().DBName)
	h, _ := time.ParseDuration("-48h")
	err := dbConn.C("delegations2").Remove(bson.M{"time":bson.M{"lt":time.Now().Add(h)}})
	if err != nil {

	}
}
