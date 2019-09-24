package models

import (
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/db"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
)

type ValidatorAddressAndKey struct {
	OperatorAddress string `json:"operator_address"`
	ConsensusPubkey string `json:"consensus_pubkey"`
	ProposerHash    string `json:"proposer_hash"`
}

func (vaak *ValidatorAddressAndKey) GetInfo(address string) string {
	session := db.NewDBConn()
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	dbConn.C("mapping").Find(bson.M{"operatoraddress":address}).One(&vaak)
	return vaak.ProposerHash
}

func (vaak *ValidatorAddressAndKey) CheckValidator(pubkey string)(int, string){
	session := db.NewDBConn()
	 count :=0
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	dbConn.C("mapping").Find(bson.M{"consensuspubkey":pubkey}).One(&vaak)
	if vaak.ProposerHash!=""{
		count = 1
	}
	return count,vaak.ProposerHash
}

func (vaak *ValidatorAddressAndKey) GetValidator(address string) string {
	session := db.NewDBConn()
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	dbConn.C("mapping").Find(bson.M{"proposerhash":address}).One(&vaak)
	return vaak.OperatorAddress
}
func (vaak *ValidatorAddressAndKey) SetInfo(log zap.Logger) {
	session := db.NewDBConn()
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	c :=dbConn.C("mapping")
	err := c.Insert(&vaak)
	if err == nil {
		log.Info("insert mappingRelationship success")
	} else {
		log.Error("insert mappingRelationship failed")
	}
}
