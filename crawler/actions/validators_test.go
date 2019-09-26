package actions

import (
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/logger"
	"testing"
)

func TestGetValidators(t *testing.T) {
	config := conf.NewConfig()   // CONFIG
	log := logger.NewLogger()    // LOG
	GetValidators(config, log)
	//OnlyTest()
}
