package db

import (
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/logger"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
	"time"
)

func NewDBConn() *mgo.Session {

	var log = logger.NewLogger()
	config := conf.NewConfig()
	if config.DBstring == "" {
		log.Info("dbString is empty")
		return nil
	}
	defer func() {
		if p:=recover(); p!=nil{
			var ss []string
			ss = append(ss, config.DBstring,config.DBName)
			log.Error("Failed to connect to DB,Please check db settings.",zap.Strings("DBconfig",ss))
		}
	}()
	//mgo.DialWithTimeout(time.Second*60)
	//session, err := DialWithTimeout(url, 10*time.Second)
	dialInfo,_:=mgo.ParseURL(config.DBstring)
	dialInfo.PoolLimit=40960
	dialInfo.Timeout = time.Minute*20
	session,err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatal("creat db connection failed, check db state")
		return nil
	}
	defer session.Close()
	return session.Copy()
}

//
//session, err := mgo.Dial("mongodb://root:helloKitty@127.0.0.1:27017")

//if err != nil{
//t.Log("can not connect db")
//}else {
//t.Log("connect success")
//}
//
//defer session.Close()
//c := session.DB("hsnBCB").C("Test")
//stu1 := UserInfo{
//Name:  "wylllllllll",
//Age: 13,
//Phone: "329832984@qq.com",
//Addtime:   time.Now(),
//}
//err = c.Insert(&stu1)
//
//if err !=nil {
//t.Log("insert error")
//}else {
//t.Log("ok")
