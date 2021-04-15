package logfx

import (
	"github.com/nghiant3223/tikihackathon/pkg/log"
	"github.com/spf13/viper"
)

func initializeLogger() {
	mode := viper.GetString("log.mode")
	loggerType := viper.GetString("log.type")

	log.Init(log.WithMode(mode), log.WithLoggerType(loggerType))
}
