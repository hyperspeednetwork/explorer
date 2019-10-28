package validatorDetails

import (
	"github.com/astaxie/beego"
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/db"
	"github.com/wongyinlong/hsnNet/models"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type ProposedBlocksController struct {
	beego.Controller
}
type msgs struct {
	Code string        `json:"code"`
	Data []BlockSimple `json:"data"`
	Msg  string        `json:"msg"`
	Total int `json:"total"`
}
type BlockSimple struct {
	Height    int    `json:"height"`
	BlockHash string `json:"block_hash"`
	Proposer  string `json:"proposer"`
	Txs       string `json:"txs"`
	Time      string `json:"time"`
}


// @Title Get
// @Description get proposedBlocks
// @Success code 0
// @Failure code 1
// @router /
func (pbc *ProposedBlocksController) Get() {
	pbc.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", pbc.Ctx.Request.Header.Get("Origin"))
	address := pbc.GetString("address")
	head, _ := pbc.GetInt("head", 0)
	page, _ := pbc.GetInt("page", 0)
	size, _ := pbc.GetInt("size", 0)
	if address == "" || strings.Index(address,conf.NewConfig().Public.Bech32PrefixValAddr)!=0  {
		var errorMessage MsgErr
		errorMessage.Code = "1"
		errorMessage.Data = nil
		errorMessage.Msg = "Validator address is empty! Or error address!"
		pbc.Data["json"] = errorMessage
		pbc.ServeJSON()
	}
	var session = db.NewDBConn() //db
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	if page == 0 {
		// default page
		page = 0
	}
	if size == 0 {
		// default last SIZE
		size = 5
	}
	if head == 0 {
		// default last height
		var public models.Infomation
		_=dbConn.C("public").Find(nil).Sort("-height").One(&public)
		head = public.Height

	}

	var blocks = make([]models.BlockInfo, size)
	var blockInfoSimples = make([]BlockSimple, size)
	var respJson msgs
	var validatorAddress models.ValidatorAddressAndKey
	hashAddress := validatorAddress.GetInfo(address)
	_=dbConn.C("block").Find(
		bson.M{
			"intheight": bson.M{"$lte": head,}, "block.header.proposeraddress": hashAddress}).Sort("-intheight").Limit(size).Skip(size * page).All(&blocks)
	for i, item := range blocks {
		blockInfoSimples[i].Height = item.IntHeight
		blockInfoSimples[i].BlockHash = item.BlockMeta.BlockId.Hash
		blockInfoSimples[i].Proposer = item.Block.Header.ProposerAddress
		blockInfoSimples[i].Txs = item.Block.Header.NumTxs
		blockInfoSimples[i].Time = item.Block.Header.Time
	}
	respJson.Total ,_= dbConn.C("block").Find(bson.M{"block.header.proposeraddress": hashAddress}).Count()
	respJson.Data = blockInfoSimples
	respJson.Code = "0"
	respJson.Msg = "OK"
	pbc.Data["json"] = respJson
	pbc.ServeJSON()
}
