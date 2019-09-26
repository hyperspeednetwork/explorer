package models

import (
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/db"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
)

type ValidatorAddressAndDelegatorAddress struct {
	ValidatorAddress string `json:"validator_address"`
	DelegatorAddress string `json:"delegator_address"`
}

func (vaada *ValidatorAddressAndDelegatorAddress) Set(log zap.Logger) {
	// store

	session := db.NewDBConn()
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	c := dbConn.C("mapping")
	err := c.Insert(&vaada)
	if err == nil {
		log.Info("insert mappingRelationship success")
	} else {
		log.Error("insert mappingRelationship failed",zap.String("error",err.Error()))
	}
}

func (vaada *ValidatorAddressAndDelegatorAddress) Get() {

}

func (vaada *ValidatorAddressAndDelegatorAddress) Check(address string) (int, string) {
	var tempValue ValidatorAddressAndDelegatorAddress
	session := db.NewDBConn()
	count := 0
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	dbConn.C("mapping").Find(bson.M{"validatoraddress": address}).One(&tempValue)
	if tempValue.DelegatorAddress != "" {
		count = 1
	}
	return count, tempValue.DelegatorAddress
}

type GenesisFile struct {
	Result struct {
		Genesis struct {
			AppState struct {
				Genutil struct {
					Gentxs []struct {
						Value struct {
							Msg []struct {
								Value struct {
									DelegatorAddress string `json:"delegator_address"`
									ValidatorAddress string `json:"validator_address"`
								} `json:"value"`
							} `json:"msg"`
						} `json:"value"`
					} `json:"gentxs"`
				} `json:"genutil"`
			} `json:"app_state"`
		} `json:"genesis"`
	} `json:"result"`
}

type CreatValidator struct {
	Txs []struct {
		Tx struct {
			Value struct {
				Msg []struct {
					Value struct {
						DelegatorAddress string `json:"delegator_address"`
						ValidatorAddress string `json:"validator_address"`
					} `json:"value"`
				} `json:"msg"`
			} `json:"value"`
		} `json:"tx"`
	} `json:"txs"`
}
