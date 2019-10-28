package controllers

import (
	"github.com/astaxie/beego"
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/db"
	"github.com/wongyinlong/hsnNet/models"
	//"github.com/wongyinlong/hsnNet/logger"
)

type PublicController struct {
	beego.Controller

}

type Public struct {
	Code string            `json:"code"`
	Data models.Infomation `json:"data"`
	Msg  string            `json:"msg"`
}

// @Title Get
// @Description public Item
// @Success code 0
// @Failure code 1
//@router / [get]
func (pb *PublicController) Get() {
	pb.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", pb.Ctx.Request.Header.Get("Origin"))
	var public models.Infomation
	var respJson Public
	var session = db.NewDBConn() //db
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	dbConn.C("public").Find(nil).Sort("-height").One(&public)
	respJson.Data = public
	respJson.Code = "0"
	respJson.Msg = "OK"
	pb.Data["json"] = respJson
	pb.ServeJSON()

}
