package actions

import (
	"github.com/wongyinlong/hsnNet/conf"
	"github.com/wongyinlong/hsnNet/logger"
	"testing"
)

func TestGetPublic(t *testing.T) {
	config := conf.NewConfig()   // CONFIG
	log := logger.NewLogger()    // LOG

	GetPublic(config, log)

}
