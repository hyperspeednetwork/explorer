package models

import (
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/db"
	"go.uber.org/zap"
)

type BlockInfo struct {
	BlockMeta BlockMeta `json:"block_meta"`
	Block     Block     `json:"block"`
	IntHeight int `json:"int_height"`
}

type BlockMeta struct {
	BlockId BlockId `json:"block_id"`
}
type Block struct {
	Header     Header     `json:"header"`
	//Data       Data       `json:"data"`
	//Evidence   Evidence   `json:"evidence"`
	//LastCommit LastCommit `json:"last_commit"`
}

type BlockId struct {
	Hash  string `json:"hash"`
	//Parts Parts  `json:"parts"`
}
type Parts struct {
	Total string `json:"total"`
	Hash  string `json:"hash"`
}

type Header struct {
	//Version            Version `json:"version"`
	//ChainId            string  `json:"chain_id"`
	Height             string  `json:"height"`
	Time               string  `json:"time"` //p
	NumTxs             string  `json:"num_txs"`
	TotalTxs           string  `json:"total_txs"`
	LastBlockId        BlockId `json:"last_block_id"`
	//LastCommitHash     string  `json:"last_commit_hash"`
	//DataHash           string  `json:"data_hash"`
	ValidatorsHash     string  `json:"validators_hash"`
	//NextValidatorsHash string  `json:"next_validators_hash"`
	//ConsensusHash      string  `json:"consensus_hash"`
	//AppHash            string  `json:"app_hash"`
	//LastResultHash     string  `json:"last_result_hash"`
	//EvidenceHash       string  `json:"evidence_hash"`
	ProposerAddress    string  `json:"proposer_address"`

}
type Version struct {
	Block string `json:"block"`
	App   string `json:"app"`
}

type Data struct {
	Txs []string `json:"txs"`
}

type Evidence struct {
	Evidence string `json:"evidence"`
}
type LastCommit struct {
	BlockId    BlockId          `json:"block_id"`
	Precommits []PreCommitsList `json:"precommits"`
}

type PreCommitsList struct {
	Type             int     `json:"type"`
	Height           string  `json:"height"`
	Round            string  `json:"round"`
	BlockID          BlockId `json:"block_id"`
	Timestamp        string  `json:"timestamp"`
	ValidatorAddress string  `json:"validator_address"`
	ValidatorIndex   string  `json:"validator_index"`
	Signature        string  `json:"signature"`
}

func (b *BlockInfo) SetInfo(log zap.Logger){
	var session = db.NewDBConn() //db
	defer session.Close()
	defer log.Sync()
	dbConn := session.DB(conf.NewConfig().DBName)
	c := dbConn.C("block")
	err :=c.Insert(&b)

	if err != nil {
		log.Error("insert block data error", zap.String("height", b.Block.Header.Height),zap.String("error",err.Error()))
	}else {
		log.Info("insert block data ", zap.String("height", b.Block.Header.Height))
	}

}
func (b *BlockInfo)GetAimHeightAndBlockHeight()(int,int){
	//var lock sync.Mutex          // LOCK
	var session = db.NewDBConn() //db
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	var block BlockInfo
	var public Infomation
	//lock.Lock()
	dbConn.C("public").Find(nil).Sort("-height").One(&public)
	dbConn.C("block").Find(nil).Sort("-intheight").One(&block)
	//lock.Unlock()
	lastBlockHeight :=block.IntHeight
	publicHeight := public.Height
	return lastBlockHeight,publicHeight
}