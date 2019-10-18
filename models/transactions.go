package models

import (

	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/db"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
	"time"
)

//http://172.38.8.89:1317/txs?message.action=send
type Txs struct {
	TxHash                  string    `json:"tx_hash"`
	Page                    int       `json:"page"`
	Type                    string    `json:"type"`
	Result                  bool      `json:"result"`
	Amount                  []float64 `json:"amount"`
	Fee                     float64   `json:"fee"`
	Height                  int       `json:"height"`
	TxTime                  string    `json:"tx_time"`
	Plus                    int       `json:"pluse"`
	ValidatorAddress        []string  `json:"validator_address"`
	DelegatorAddress        []string  `json:"delegator_address"`
	WithDrawRewardAmout     []float64 `json:"with_draw_reward_amout"`
	WithDrawCommissionAmout []float64 `json:"with_draw_commission_amout"`
	FromAddress             []string  `json:"from_address"`
	ToAddress               []string  `json:"to_address"`
	OutPutsAddress          []string  `json:"out_puts_address"`
	InputsAddress           []string  `json:"inputs_address"`
	VoterAddress            []string  `json:"voter_address"`
	Options                 []string  `json:"options"`
	Time                    time.Time `json:"time"`
}

type MultiSendMsg struct {
	Value struct {
		Inputs []struct {
			Address string `json:"address"`
			Coins   []struct {
				Denom  string `json:"denom"`
				Amount string `json:"amount"`
			} `json:"coins"`
		} `json:"inputs"`
		Outputs []struct {
			Address string `json:"address"`
			Coins   []struct {
				Denom  string `json:"denom"`
				Amount string `json:"amount"`
			} `json:"coins"`
		} `json:"outputs"`
	} `json:"value"`
}
type GetRewardCommissionMsg struct {
	Type  string `json:"type"`
	Value struct {
		DelegatorAddress string `json:"delegator_address"`
		ValidatorAddress string `json:"validator_address"`
		Amount           string `json:"amount"`
	} `json:"value"`
}
type GetRewardMsg struct {
	Type  string `json:"type"`
	Value struct {
		DelegatorAddress string `json:"delegator_address"`
		ValidatorAddress string `json:"validator_address"`
		Amount           string `json:"amount"`
	} `json:"value"`
}
type UndelegateMsg struct {
	Type  string `json:"type"`
	Value struct {
		DelegatorAddress string `json:"delegator_address"`
		ValidatorAddress string `json:"validator_address"`
		Amount           struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"amount"`
	} `json:"value"`
}
type DelegateMsg struct {
	Type  string `json:"type"`
	Value struct {
		DelegatorAddress string `json:"delegator_address"`
		ValidatorAddress string `json:"validator_address"`
		Amount           struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"amount"`
	} `json:"value"`
}
type VoteMsg struct {
	Type  string `json:"type"`
	Value struct {
		ProposalID string `json:"proposal_id"`
		Voter      string `json:"voter"`
		Option     string `json:"option"`
	} `json:"value"`
}
type Amount struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}
type Value struct {
	FromAddress string   `json:"from_address"`
	ToAddress   string   `json:"to_address"`
	Amount      []Amount `json:"amount"`
}
type Msg struct {
	Type  string `json:"type"`
	Value Value  `json:"value"`
}
type Gas struct {
	Used   string `json:"used"`
	Wanted string `json:"wanted"`
}

func (txs Txs) SetInfo(log zap.Logger) error {

	session := db.NewDBConn()
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	c := dbConn.C("Txs")
	err := c.Insert(&txs)
	if err != nil {
		log.Error("txs insert error", zap.String("error", err.Error()))
	} else {
		log.Info("txs insert success")
	}
	return nil

}

func (txs Txs) GetInfo(head int, page int, size int) ([]Txs, int) {
	if page <= 0 {
		// default page
		page = 0
	}
	if size == 0 {
		size = 5
	}
	var TxsSet = make([]Txs, size)
	var session = db.NewDBConn() //db
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	if head == 0 {
		var txs Txs
		_ = dbConn.C("Txs").Find(nil).Sort("-height").One(&txs)
		head = txs.Height
	}
	_= dbConn.C("Txs").Find(bson.M{"height": bson.M{
		"$lte": head,}}).Sort("-height").Limit(size).Skip(page * size).All(&TxsSet)
	totalTxsCount, _ := dbConn.C("Txs").Find(nil).Count()
	return TxsSet, totalTxsCount
}
func (txs Txs) GetDetail(txhash string) Txs {
	var session = db.NewDBConn() //db
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	var temptxs Txs
	_ = dbConn.C("Txs").Find(bson.M{"txhash": txhash}).One(&temptxs)
	return temptxs
}

func (txs Txs) CheckHash(txhash string) int {
	// use hash to query tx, 0 Unrecorded, 1 Recorded
	var session = db.NewDBConn() //db
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	var temptxs Txs
	_ = dbConn.C("Txs").Find(bson.M{"txhash": txhash}).One(&temptxs)
	if temptxs.TxHash == "" {
		return 0
	} else {
		return 1
	}
}
func (txs Txs) GetValidatorsTransactions(address string) {
	//use hash to query tx, 0 Unrecorded, 1 Recorded
	var session = db.NewDBConn() //db
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	var temptxs []Txs
	//type undelegate,delegate ...
	dbConn.C("Txs").Find(bson.M{"txhash": address}).One(&temptxs)
}

func (txs Txs) GetPowerEventInfo(address string, page, size int) (*[]Txs, int) {
	if page < 0 {
		// default page
		page = 0
	}
	if size == 0 {
		size = 5
	}
	var txsSet []Txs
	var session = db.NewDBConn() //db
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)

	var query []bson.M
	q1 := bson.M{"type": "unbonding"}
	q2 := bson.M{"type": "delegate"}
	q3 := bson.M{"type": "redelegate"}
	query = append(query, q1, q2, q3)
	count, _ := dbConn.C("Txs").Find(bson.M{"$or": query, "validatoraddress": bson.M{"$elemMatch": bson.M{"$eq": address}}, "result": true}).Count()
	_ = dbConn.C("Txs").Find(bson.M{"$or": query, "validatoraddress": bson.M{"$elemMatch": bson.M{"$eq": address}}, "result": true}).Sort("-height").Limit(size).Skip(page * size).All(&txsSet)
	return &txsSet, count
}
func (txs Txs) GetDelegatorTxs(address string, page, size int) (*[]Txs, int) {
	if page < 0 {
		// default page
		page = 0
	}

	var txsSet []Txs
	var session = db.NewDBConn() //db
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)

	var query []bson.M
	q1 := bson.M{"delegatoraddress": bson.M{"$elemMatch": bson.M{"$eq": address}}}
	//from_address
	//out_puts_address
	//voter_address
	q2 := bson.M{"fromaddress": bson.M{"$elemMatch": bson.M{"$eq": address}}}
	q3 := bson.M{"outputsaddress": bson.M{"$elemMatch": bson.M{"$eq": address}}}
	q5 := bson.M{"inputsaddress": bson.M{"$elemMatch": bson.M{"$eq": address}}}
	q4 := bson.M{"voteraddress": bson.M{"$elemMatch": bson.M{"$eq": address}}}

	query = append(query, q1, q2, q3, q4, q5)
	count, _ := dbConn.C("Txs").Find(bson.M{"$or": query}).Count()
	_ = dbConn.C("Txs").Find(bson.M{"$or": query,}).Sort("-height").Limit(size).Skip(page * size).All(&txsSet)

	return &txsSet, count
}
func (txs Txs) GetDelegatorCommissionTx(address string) (*[]Txs) {
	var txsSet []Txs
	var session = db.NewDBConn() //db
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	_ = dbConn.C("Txs").Find(bson.M{"type": "commission", "delegatoraddress": bson.M{"$elemMatch": bson.M{"$eq": address}}}).All(&txsSet)
	return &txsSet
}
func (txs Txs) GetDelegatorRewardTx(address string) (*[]Txs) {
	var txsSet []Txs
	var session = db.NewDBConn() //db
	defer session.Close()
	var query []bson.M
	dbConn := session.DB(conf.NewConfig().DBName)
	q1 := bson.M{"type": "commission"}
	q2 := bson.M{"type": "reward"}
	query = append(query, q1, q2)
	_ = dbConn.C("Txs").Find(bson.M{"$or": query, "delegatoraddress": bson.M{"$elemMatch": bson.M{"$eq": address}}}).All(&txsSet)
	return &txsSet
}
func (txs Txs) GetSpecifiedHeight(head int, page int, size int) ([]Txs, int) {

	if page <= 0 {
		// default page
		page = 0
	}
	if size == 0 {
		size = 5
	}
	var TxsSet = make([]Txs, size)
	var session = db.NewDBConn() //db
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	_ = dbConn.C("Txs").Find(bson.M{"height": head}).Sort("-height").Limit(size).Skip(page * size).All(&TxsSet)
	totalTxsCount, _ := dbConn.C("Txs").Find(bson.M{"height": head}).Count()
	return TxsSet, totalTxsCount
}
func (txs Txs) GetTxHeight() (int) {
	var session = db.NewDBConn() //db
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	_ = dbConn.C("Txs").Find(nil).Sort("-height").One(&txs)
	return txs.Height
}
