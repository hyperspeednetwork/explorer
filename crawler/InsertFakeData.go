package crawler

import (
	"encoding/json"
	"fmt"
	"github.com/wongyinlong/hsnNet/models/accountDetail"
)

func InserFakeData() {
	//insertPowerEvnet()
	//inserValidatorDelegator()
	//insertValidatorDelegator()
	//MoocUnbondingData()
}

//func insertPowerEvnet(){
//	var txs models.Txs
//	address := "hsnvaloper1zqxayv6qe50w6h3ynfj6tq9pr09r7rtu4u3wgp"
//	txList, _ := txs.GetPowerEventInfo(address,0,0)
//	for _,item :=range *txList {
//		item.SetInfo()
//	}
//}
//
//func insertValidatorDelegator(){
//// get maxsize
//	var delegations validatorsDetail.Delegators2
//	address := "hsnvaloper1zqxayv6qe50w6h3ynfj6tq9pr09r7rtu4u3wgp"
//	items,_,_ :=delegations.GetInfo(address,0, 0)
//	for _,item :=range *items{
//		fmt.Println(item)
//		item.SetInfo()
//	}
//}

func MoocDelegatorData() string {
	return ""
}

func MoocUnbondingData() string {
	//	unbondingtr2 :=`
	//{
	//  "height": "332223",
	//  "result": [
	//    {
	//      "delegator_address": "hsn1aqwurs5jfu5z0z3k99tljt9csausdqrcaewjwv",
	//      "validator_address": "hsnvaloper1aqwurs5jfu5z0z3k99tljt9csausdqrcg39g7j",
	//      "entries": [
	//        {
	//          "creation_height": "110080",
	//          "completion_time": "2019-09-25T03:34:38.163529945Z",
	//          "initial_balance": "800000000",
	//          "balance": "800000000"
	//        }
	//      ]
	//    }
	//  ]
	//}`

	unbondingStr := `

{
 "height": "332223",
 "result": [
   {
     "delegator_address": "hsn1aqwurs5jfu5z0z3k99tljt9csausdqrcaewjwv",
     "validator_address": "hsnvaloper1aqwurs5jfu5z0z3k99tljt9csausdqrcg39g7j",
     "entries": [
       {
         "creation_height": "110080",
         "completion_time": "2019-09-25T03:34:38.163529945Z",
         "initial_balance": "800000000",
         "balance": "800000000"
       }
     ]
   },
   {
     "delegator_address": "hsn1aqwurs5jfu5z0z3k99tljt9csausdqrcaewjwv",
     "validator_address": "hsnvaloper1aqwurs5jfu5z0z3k99tljt9csausdqrcg39g7j",
     "entries": [
       {
         "creation_height": "110080",
         "completion_time": "2019-09-25T03:34:38.163529945Z",
         "initial_balance": "800000000",
         "balance": "800000000"
       }
     ]
   },
   {
     "delegator_address": "hsn1aqwurs5jfu5z0z3k99tljt9csausdqrcaewjwv",
     "validator_address": "hsnvaloper1aqwurs5jfu5z0z3k99tljt9csausdqrcg39g7j",
     "entries": [
       {
         "creation_height": "110080",
         "completion_time": "2019-09-25T03:34:38.163529945Z",
         "initial_balance": "800000000",
         "balance": "800000000"
       }
     ]
   },
   {
     "delegator_address": "hsn1aqwurs5jfu5z0z3k99tljt9csausdqrcaewjwv",
     "validator_address": "hsnvaloper1aqwurs5jfu5z0z3k99tljt9csausdqrcg39g7j",
     "entries": [
       {
         "creation_height": "110080",
         "completion_time": "2019-09-25T03:34:38.163529945Z",
         "initial_balance": "800000000",
         "balance": "800000000"
       }
     ]
   },
   {
     "delegator_address": "hsn1aqwurs5jfu5z0z3k99tljt9csausdqrcaewjwv",
     "validator_address": "hsnvaloper1aqwurs5jfu5z0z3k99tljt9csausdqrcg39g7j",
     "entries": [
       {
         "creation_height": "110080",
         "completion_time": "2019-09-25T03:34:38.163529945Z",
         "initial_balance": "800000000",
         "balance": "800000000"
       }
     ]
   },
   {
     "delegator_address": "hsn1aqwurs5jfu5z0z3k99tljt9csausdqrcaewjwv",
     "validator_address": "hsnvaloper1aqwurs5jfu5z0z3k99tljt9csausdqrcg39g7j",
     "entries": [
       {
         "creation_height": "110080",
         "completion_time": "2019-09-25T03:34:38.163529945Z",
         "initial_balance": "800000000",
         "balance": "800000000"
       }
     ]
   },
   {
     "delegator_address": "hsn1aqwurs5jfu5z0z3k99tljt9csausdqrcaewjwv",
     "validator_address": "hsnvaloper1aqwurs5jfu5z0z3k99tljt9csausdqrcg39g7j",
     "entries": [
       {
         "creation_height": "110080",
         "completion_time": "2019-09-25T03:34:38.163529945Z",
         "initial_balance": "800000000",
         "balance": "800000000"
       }
     ]
   },
   {
     "delegator_address": "hsn1aqwurs5jfu5z0z3k99tljt9csausdqrcaewjwv",
     "validator_address": "hsnvaloper1aqwurs5jfu5z0z3k99tljt9csausdqrcg39g7j",
     "entries": [
       {
         "creation_height": "110080",
         "completion_time": "2019-09-25T03:34:38.163529945Z",
         "initial_balance": "800000000",
         "balance": "800000000"
       }
     ]
   },
   {
     "delegator_address": "hsn1aqwurs5jfu5z0z3k99tljt9csausdqrcaewjwv",
     "validator_address": "hsnvaloper1aqwurs5jfu5z0z3k99tljt9csausdqrcg39g7j",
     "entries": [
       {
         "creation_height": "110080",
         "completion_time": "2019-09-25T03:34:38.163529945Z",
         "initial_balance": "800000000",
         "balance": "800000000"
       }
     ]
   },
   {
     "delegator_address": "hsn1aqwurs5jfu5z0z3k99tljt9csausdqrcaewjwv",
     "validator_address": "hsnvaloper1aqwurs5jfu5z0z3k99tljt9csausdqrcg39g7j",
     "entries": [
       {
         "creation_height": "110080",
         "completion_time": "2019-09-25T03:34:38.163529945Z",
         "initial_balance": "800000000",
         "balance": "800000000"
       }
     ]
   },
   {
     "delegator_address": "hsn1aqwurs5jfu5z0z3k99tljt9csausdqrcaewjwv",
     "validator_address": "hsnvaloper1aqwurs5jfu5z0z3k99tljt9csausdqrcg39g7j",
     "entries": [
       {
         "creation_height": "110080",
         "completion_time": "2019-09-25T03:34:38.163529945Z",
         "initial_balance": "800000000",
         "balance": "800000000"
       }
     ]
   },
   {
     "delegator_address": "hsn1aqwurs5jfu5z0z3k99tljt9csausdqrcaewjwv",
     "validator_address": "hsnvaloper1aqwurs5jfu5z0z3k99tljt9csausdqrcg39g7j",
     "entries": [
       {
         "creation_height": "110080",
         "completion_time": "2019-09-25T03:34:38.163529945Z",
         "initial_balance": "800000000",
         "balance": "800000000"
       }
     ]
   }
 ]
}

`
	var unbonding accountDetail.Unbonding
	bytesStr := []byte(unbondingStr)
	json.Unmarshal(bytesStr, &unbonding)
	fmt.Println(unbonding)
	return unbondingStr
}
