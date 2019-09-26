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
	vc.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", vc.Ctx.Request.Header.Get("Origin"))
	types := vc.GetString("type", "")
	var validatorInfo models.ValidatorInfo
	var vtl ValidatorTypeList
	var normallyValidatorList []models.ValidatorInfo
	var count int
	list := validatorInfo.GetInfo()
	if len(*list) == 0 {
		var errMsg errMsg
		errMsg.Data = nil
		errMsg.Code = "1"
		errMsg.Msg = "No Record!"
		vc.Data["json"] = errMsg
		vc.ServeJSON()
	}
	for _, item := range *list {
		if item.Jailed {
			lenJailedList := len(vtl.Jailed)
			if lenJailedList == 0 {
				item.Cumulative = item.VotingPower.Percent
			} else {
				aheadItemInJailedList := lenJailedList - 1
				item.Cumulative = vtl.Jailed[aheadItemInJailedList].Cumulative + item.VotingPower.Percent
			}
			vtl.Jailed = append(vtl.Jailed, item)
		} else {
			lenNormallyValidatorList := len(normallyValidatorList)
			if lenNormallyValidatorList == 0 {
				item.Cumulative = item.VotingPower.Percent
			} else {
				aheadItemInNormallyList := lenNormallyValidatorList - 1
				item.Cumulative = normallyValidatorList[aheadItemInNormallyList].Cumulative + item.VotingPower.Percent
			}
			normallyValidatorList = append(normallyValidatorList, item)
			if count < 100 {
				vtl.Active = append(vtl.Active, item)
				count++
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
