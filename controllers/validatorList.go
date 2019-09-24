package controllers

import (
	"github.com/astaxie/beego"
	"github.com/wongyinlong/hsnNet/models"
)

type ValidatorsController struct {
	beego.Controller
}
type MSGS struct {
	Code string            `json:"code"`
	Data ValidatorTypeList `json:"data"`
	Msg  string            `json:"msg"`
}
type errMsg struct {
	Code string `json:"code"`
	Data error  `json:"data"`
	Msg  string `json:"msg"`
}
type ValidatorTypeList struct {
	Jailed    []models.ValidatorInfo `json:"jailed"`
	Active    []models.ValidatorInfo `json:"active"`
	Candidate []models.ValidatorInfo `json:"candidate"`
}

// @Title 获取Validators List
// @Description get validators
// @Success code 0
// @Failure code 1
// @router /
func (vc *ValidatorsController) Get() {
	types := vc.GetString("type","")
	var validatorInfo models.ValidatorInfo
	var vtl ValidatorTypeList
	var aheadCumulative float64
	var count int
	var CumulativeList []float64
	list := validatorInfo.GetInfo()
	if len(*list) == 0 {
		var errMsg errMsg
		errMsg.Data = nil
		errMsg.Code = "1"
		errMsg.Msg = "No Record!"
		vc.Data["json"] = errMsg
		vc.ServeJSON()
	}
	for index, item := range *list {
		CumulativeList = append(CumulativeList, item.Cumulative)
		if index == 0 {
			aheadCumulative = 0.0
		} else {
			aheadCumulative = CumulativeList[index-1]
		}
		item.Cumulative = aheadCumulative + item.VotingPower.Percent
		if item.Jailed {
			vtl.Jailed = append(vtl.Jailed, item)
		} else {
			if count < 100 {
				vtl.Active = append(vtl.Active, item)
			} else {
				vtl.Candidate = append(vtl.Candidate, item)
			}
		}
	}

	var msgs MSGS

	msgs.Code = "0"
	msgs.Msg = "OK"
	switch types {
	case "":
		msgs.Data = vtl
	case "jailed":
		msgs.Data.Jailed = vtl.Jailed
	case "active":
		msgs.Data.Active = vtl.Active
	case "candidate":
		msgs.Data.Candidate = vtl.Candidate
	default:
		var errMsg errMsg
		errMsg.Data = nil
		errMsg.Code = "1"
		errMsg.Msg = "No Record!"
		vc.Data["json"] = errMsg
		vc.ServeJSON()
	}

	vc.Data["json"] = msgs
	vc.ServeJSON()
}
