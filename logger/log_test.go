package logger

import "testing"

func TestNewLogger(t *testing.T) {
	var log = NewLogger()
	log.Info("wo cuo le ")
	log.Info("wo you cuo le haha")
	log.Panic("wo you cuo le haha")
}
