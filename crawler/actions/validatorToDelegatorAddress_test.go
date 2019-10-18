package actions

import (
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/logger"
	"testing"
)

func TestSetValidatorDelegatorAddress(t *testing.T) {
	config := conf.NewConfig()   // CONFIG
	log := logger.NewLogger()    // LOG
	SetValidatorDelegatorAddress(config, log)
}
