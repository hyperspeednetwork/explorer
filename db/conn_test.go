package db

import (
	"fmt"
	"github.com/wongyinlong/hsnNet/logger"
	"testing"
	"time"
)

type UserInfo struct {
	Name    string    `json:"name"`
	Age     int       `json:"age"`
	Phone   string    `json:"phone"`
	Addtime time.Time `json:"addtime"`
}

func TestNewDBConn(t *testing.T) {
	log := logger.NewLogger()
	session := NewDBConn()
	defer session.Close()
	dbConn := session.DB("hsnhub_db_dev_test")
	if dbConn == nil{
		return
	}
	c := dbConn.C("hello_test")
	fmt.Println("hello there")
	stu1 := UserInfo{
		Name:    "wyl",
		Age:     13,
		Phone:   "329832984@qq.com",
		Addtime: time.Now(),
	}

	err := c.Insert(&stu1)
	if err != nil {
		log.Error("insert error")
	} else {
		log.Info("insert success")

	}
}
