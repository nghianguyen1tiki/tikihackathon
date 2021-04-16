package logfx

import (
	"github.com/nghiant3223/tikihackathon/internal/config"
	"github.com/spf13/viper"

	"github.com/nghiant3223/tikihackathon/pkg/log"
)

func invokeLogger() {
	var cfg *config.Log
	err := viper.UnmarshalKey("log", &cfg)
	if err != nil {
		return
	}

	log.Init(log.WithMode(cfg.Mode), log.WithLoggerType(cfg.Type))
}
