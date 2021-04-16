package logfx

import (
	"github.com/spf13/viper"

	"github.com/nghiant3223/tikihackathon/pkg/log"
)

func invokeLogger() {
	mode := viper.GetString("log.mode")
	loggerType := viper.GetString("log.type")

	log.Init(log.WithMode(mode), log.WithLoggerType(loggerType))
}
