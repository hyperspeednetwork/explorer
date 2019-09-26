package accountDetails

import (
	"github.com/astaxie/beego"
	"github.com/shopspring/decimal"
	"github.com/wongyinlong/hsnNet/models"
	"github.com/wongyinlong/hsnNet/models/accountDetail"
)

type KindsRewardController struct {
	beego.Controller
}

type KindsRewardErrMsg struct {
	Data error  `json:"data"`
	Msg  string `json:"msg"`
	Code string `json:"code"`
}
type KindsRewardMsg struct {
	Data Kinds  `json:"data"`
	Msg  string `json:"msg"`
	Code string `json:"code"`
}
type Kinds struct {
	// (amount percentage)
	Available   []decimal.Decimal `json:"available"`
	Delegated   []decimal.Decimal `json:"delegated"`
	Unbonding   []decimal.Decimal `json:"unbonding"`
	Reward      []decimal.Decimal `json:"reward"`
	Commission  []decimal.Decimal `json:"commission"`
	TotalAmount []decimal.Decimal `json:"total_amount"`
}

// @Title
// @Description
// @Success code 0
// @Failure code 1
//@router /
func (krc *KindsRewardController) Get() {
	krc.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", krc.Ctx.Request.Header.Get("Origin"))
	address := krc.GetString("address", "")
	if address == "" || len(address) != 42 || address[0:3] != "hsn" {
		var errMsg KindsRewardErrMsg
		errMsg.Data = nil
		errMsg.Code = "1"
		errMsg.Msg = "Address Error, Or empty!"
		krc.Data["json"] = errMsg
	} else {
		// get delegator's txs .计算各类奖励之和.(delegate, unbonding, Reward, CommissionReward,
		//Available=account's token sub (delegate+ unbonding+ Reward+ CommissionReward))
		//Reward http://172.38.8.89:1317/distribution/delegators/hsn190zh9q92xs8y7s4c0tpc784vass5lalt7f3n0h/rewards
		//delegate SUM http://172.38.8.89:1317/staking/delegators/hsn1502lgkad0tnc2szdww0whpxs30szz03lj6n06q/delegations
		//Unbonding Sum http://172.38.8.89:1317/staking/delegators/hsn1aqwurs5jfu5z0z3k99tljt9csausdqrcaewjwv/unbonding_delegations
		//Commission

		var msg KindsRewardMsg
		msg.Data.TotalAmount = getTotalAmount(address)
		msg.Data.Reward = getReward(address, msg.Data.TotalAmount)
		msg.Data.Delegated = getTotalDelegateAmount(address, msg.Data.TotalAmount)
		msg.Data.Unbonding = getTotalUnbondingAmount(address, msg.Data.TotalAmount)
		msg.Data.Commission = getTotalCommissionAmount(address, msg.Data.TotalAmount)
		msg.Data.Available = getAvailable(msg)
		msg.Code = "0"
		msg.Msg = "OK"
		krc.Data["json"] = msg

	}

	krc.ServeJSON()
}
func getTotalAmount(address string) []decimal.Decimal {
	var account accountDetail.Account
	var totalAmount []decimal.Decimal
	accountAddress, strAmount := account.GetInfo(address)
	if accountAddress == "" {
	}
	decimalRewardAmount := getDeciamlRewardAmount(address)
	decimalAmount, _ := decimal.NewFromString(strAmount)

	decimalPercentage, _ := decimal.NewFromString("1.0")
	totalAmount = append(totalAmount, decimalAmount.Add(decimalRewardAmount), decimalPercentage)
	return totalAmount
}

// from LCD interface.
//The amount is so large that there is a negative number in the results.
func getReward(address string, totalAmount []decimal.Decimal) []decimal.Decimal {
	var delegateReward accountDetail.DelegateRewards
	var rewards []decimal.Decimal
	amount := delegateReward.GetDelegateReward(address)
	decimalAmount, _ := decimal.NewFromString(amount)
	percentage := decimalAmount.Div(totalAmount[0])

	rewards = append(rewards, decimalAmount, percentage)
	return rewards
}

// from txs.
//The amount is so large that there is a negative number in the results.
//func getTotalRewardAmount(address string, totalAmount []decimal.Decimal) []decimal.Decimal {
//	var txs models.Txs
//	var reward []float64
//	var floatRewardAmount float64
//	commissionOrRewardTxs := txs.GetDelegatorRewardTx(address)
//	fmt.Println(commissionOrRewardTxs)
//	//if len(*commissionTxs) == 0 {
//	//	floatRewardAmount = 0.0
//	//} else {
//	//	for _, item := range *commissionTxs {
//	//		for index, delegator := range item.DelegatorAddress {
//	//			if delegator == address {
//	//				floatRewardAmount = floatRewardAmount + item.WithDrawCommissionAmout[index]
//	//			}
//	//		}
//	//	}
//	//}
//	percentage := floatRewardAmount / totalAmount[0]
//	reward = append(reward, floatRewardAmount, percentage)
//	return reward
//}
func getTotalDelegateAmount(address string, totalAmount []decimal.Decimal) []decimal.Decimal {
	var delegators accountDetail.Delegators
	var amount decimal.Decimal
	var delegate []decimal.Decimal
	infos := delegators.GetInfo(address)
	for _, item := range infos.Result {
		decimalAmount, _ := decimal.NewFromString(item.Balance.Amount)
		amount = amount.Add(decimalAmount)
	}
	percentage := amount.Div(totalAmount[0])
	delegate = append(delegate, amount, percentage)
	return delegate
}
func getTotalUnbondingAmount(address string, totalAmount []decimal.Decimal) []decimal.Decimal {
	var unbonding accountDetail.Unbonding
	var amount decimal.Decimal
	var unbond []decimal.Decimal
	infos := unbonding.GetInfo(address)
	for _, item := range infos.Result {
		for _, entrie := range item.Entries {
			decimalAmount, _ := decimal.NewFromString(entrie.Balance)
			amount = amount.Add(decimalAmount)
		}
	}
	percentage := amount.Div(totalAmount[0])
	unbond = append(unbond, amount, percentage)
	return unbond
}

func getTotalCommissionAmount(address string, totalAmount []decimal.Decimal) []decimal.Decimal {
	var txs models.Txs
	var commission []decimal.Decimal
	var decimalCommissionAmount decimal.Decimal
	commissionTxs := txs.GetDelegatorCommissionTx(address)
	if len(*commissionTxs) == 0 {
		decimalCommissionAmount, _ = decimal.NewFromString("0.0")
	} else {
		for _, item := range *commissionTxs {
			for index, delegator := range item.DelegatorAddress {
				if delegator == address {
					decimalWithDrawCommissionAmout := decimal.NewFromFloat(item.WithDrawCommissionAmout[index])
					decimalCommissionAmount = decimalCommissionAmount.Add(decimalWithDrawCommissionAmout)
				}
			}
		}
	}
	percentage := decimalCommissionAmount.Div((totalAmount[0]))
	commission = append(commission, decimalCommissionAmount, percentage)
	return commission
}
func getAvailable(msg KindsRewardMsg) []decimal.Decimal {
	var available []decimal.Decimal
	decimalAmount := (((msg.Data.TotalAmount[0].Sub(msg.Data.Commission[0])).Sub(msg.Data.Unbonding[0])).Sub(msg.Data.Delegated[0])).Sub(msg.Data.Reward[0])

	percentage := decimalAmount.Div(msg.Data.TotalAmount[0])
	available = append(available, decimalAmount, percentage)
	return available
}
