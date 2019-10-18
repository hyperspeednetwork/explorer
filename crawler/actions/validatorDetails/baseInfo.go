package validatorDetails

import (
	"encoding/json"
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/db"
	"github.com/wongyinlong/hsnNet/models"
	"github.com/wongyinlong/hsnNet/models/validatorsDetail"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
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
	tempValue, _ := strconv.ParseFloat(getSelfToken(item.OperatorAddress,baseInfo.Address), 64)
	baseInfo.SelfToken = tempValue
	baseInfo.TotalToken,_=strconv.ParseFloat(item.Tokens, 64)
	baseInfo.OthersToken = baseInfo.TotalToken-tempValue
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
func getSelfToken(validatorAddress string, accountAddress string)string{
//http://172.38.8.89:1317/hsn1zqxayv6qe50w6h3ynfj6tq9pr09r7rtuq565clhsnvaloper1zqxayv6qe50w6h3ynfj6tq9pr09r7rtu4u3wgp
	var  delegator  validatorsDetail.Delegator
	config := conf.NewConfig()
	url := config.Remote.Lcd +"/staking/delegators/"+accountAddress+"/delegations/"+validatorAddress
	c:=&http.Client{
		Timeout:time.Second * config.Param.HTTPGetTimeOut,
	}
	resp,err:=c.Get(url)
	defer resp.Body.Close()
	if err != nil{

	}
	jsonStr ,_:= ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(jsonStr,&delegator)
	return delegator.Result.Balance.Amount
}