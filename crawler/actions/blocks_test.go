package actions

import (
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/db"
	"github.com/wongyinlong/hsnNet/logger"
	"testing"
)

func TestGetBlock(t *testing.T) {
	config := conf.NewConfig()   // CONFIG
	log := logger.NewLogger()    // LOG
	//var lock sync.Mutex          // LOCK
	var session = db.NewDBConn() //db
	defer session.Close()
	//dbConn := session.DB("hsnhub_db_dev")
	GetBlock(config, log,)
}