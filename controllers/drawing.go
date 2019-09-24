package controllers

import (
	"github.com/astaxie/beego"
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/db"
	"github.com/wongyinlong/hsnNet/models"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

type DrawingDataController struct {
	beego.Controller
}

type Drawing struct {
	Code string `json:"code"`
	Data Items  `json:"data"`
	Msg  string `json:"msg"`
}

type Items struct {
	Price []float64 `json:"price"`
	Token []int     `json:"token"`
}

// @Title Get
// @Description 首页小图
// @Success code 0
// @Failure code 1
//@router /
func (ddc *DrawingDataController) Get() {

	var public models.Infomation
	var respJson Drawing
	var price []float64
	var token [] int
	var session = db.NewDBConn() //db
	defer session.Close()
	dbConn := session.DB(conf.NewConfig().DBName)
	now := time.Now()
	for i := 0; i <= 24; i++ {
		h, _ := time.ParseDuration("-1h")
		h1 := now.Add(time.Duration(i) * h)
		dbConn.C("public").Find(bson.M{"time": bson.M{"$lt": h1}}).Sort("-height").One(&public)
		tempPrice, _ := strconv.ParseFloat(public.Price, 64)
		price = append(price, tempPrice)
		token = append(token, public.PledgeHsn)
	}
	respJson.Data.Price = price
	respJson.Data.Token = token
	respJson.Code = "0"
	respJson.Msg = "OK"
	ddc.Data["json"] = respJson
	ddc.ServeJSON()

}
