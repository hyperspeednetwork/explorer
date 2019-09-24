package validatorDetails

import (
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/db"
	"github.com/wongyinlong/hsnNet/models"
	"github.com/wongyinlong/hsnNet/models/validatorsDetail"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

/*
Acquire and process info
*/

func MakeBaseInfo(item models.Result, VS *[]models.ValidatorsSet,log zap.Logger) {
	var baseInfo validatorsDetail.ExtraValidatorInfo
	var vaada models.ValidatorAddressAndDelegatorAddress
	baseInfo.Validator = item.OperatorAddress
	_, baseInfo.Address = vaada.Check(item.OperatorAddress)
	baseInfo.Identity = item.Description.Identity
	tempValue, _ := strconv.ParseFloat(item.DelegatorShares, 64)
	baseInfo.OthersToken = tempValue
	tempValue, _ = strconv.ParseFloat(item.MinSelfDelegation, 64)
	baseInfo.SelfToken = tempValue
	baseInfo.TotalToken = baseInfo.SelfToken + baseInfo.OthersToken
	baseInfo.WebSite = item.Description.Website
	baseInfo.Details = item.Description.Details
	baseInfo.HsnHeight = getHsnHeight(item.ConsensusPubkey)
	baseInfo.MissedBlockList = getMissBlock(VS, item.ConsensusPubkey)
	Sign := baseInfo.Check()
	if Sign == 0 {
		//set
		baseInfo.Set(log)
	} else {
		//update
		baseInfo.Update(log)
	}
}

func getMissBlock(vs *[]models.ValidatorsSet, pbKey string) []validatorsDetail.MissBLockData {
	var blockRecords []validatorsDetail.MissBLockData //记录一百个块中该验证着参与的次数（通过公钥）
	var blockRecord validatorsDetail.MissBLockData
OUTLOOP:
	for _, Sets := range *vs {

		for _, item := range Sets.Validators {
			if item.PubKey == pbKey {
				blockRecord.Height = Sets.Height
				blockRecord.State = 1
				blockRecords = append(blockRecords, blockRecord)
				continue OUTLOOP
			}

		}
		blockRecord.Height = Sets.Height
		blockRecord.State = 0
		blockRecords = append(blockRecords, blockRecord)

	}

	return blockRecords
}

func getHsnHeight(pbkey string) string {
	var mappingRealationship models.ValidatorAddressAndKey
	var blockSimpleInfo validatorsDetail.Block
	_, hash := mappingRealationship.CheckValidator(pbkey)
	//get block info
	var session = db.NewDBConn() //db
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	dbConn.C("block").Find(bson.M{"block.header.proposeraddress": hash}).One(&blockSimpleInfo)
	return strconv.Itoa(blockSimpleInfo.IntHeight)

}
