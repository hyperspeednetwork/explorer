package models

import (
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/db"
	"go.uber.org/zap"
	"time"
)

type Infomation struct {
	Price            string    `json:"price"`
	Height           int       `json:"height"`
	PledgeHsn        int       `json:"pledge_hsn"`
	TotalHsn         int       `json:"total_hsn"`
	Inflation        string    `json:"inflation"`
	TotalValidators  int       `json:"total_validators"`
	OnlineValidators int       `json:"online_validators"`
	BlockTime        float64   `json:"block_time"`
	Time             time.Time `json:"time"`
}

func NewInfomation() Infomation {
	var info Infomation
	return info
}

func (info Infomation) SetInfo(
	log zap.Logger,
	price string,
	height int,
	pledgeHsn int,
	totalHsn int,
	inflation string,
	totalValidators int,
	onlineValidators int,
	blockTime float64,
) error{
	info.Price = price
	info.Height = height
	info.PledgeHsn = pledgeHsn
	info.TotalHsn = totalHsn
	info.Inflation = inflation
	info.TotalValidators = totalValidators
	info.OnlineValidators = onlineValidators
	info.BlockTime = blockTime
	info.Time = time.Now()

	//TODO: STORE INFO
	session := db.NewDBConn()

	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	c := dbConn.C("public")
	err := c.Insert(&info)
	if err != nil {
		log.Error("insert error",zap.String("error",err.Error()))
		return err
	} else {
		log.Info("insert success")
	}
	defer log.Sync()
	return nil
}

func (info Infomation) GetInfo() Infomation{
	var public Infomation
	var session = db.NewDBConn() //db
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	_ =dbConn.C("public").Find(nil).Sort("-height").One(&public)
	return public
}
//func (info Infomation) DelInfo(){
//
//	var session = db.NewDBConn() //db
//
//	defer session.Close()
//	h, _ := time.ParseDuration("-11m")
//	fmt.Println("hello",time.Now().Add(h))
//	dbConn := session.DB(conf.NewConfig().DBName)
//	err:= dbConn.C("public").Remove(bson.M{"time":})
//	//bson.M{"$lte": time.Now().Add(h)}
//	fmt.Println(err)
//}
