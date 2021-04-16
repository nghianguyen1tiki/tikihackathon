package configs

import (
	"github.com/spf13/viper"

	"github.com/nghiant3223/tikihackathon/pkg/config"
	"github.com/nghiant3223/tikihackathon/pkg/log"
)

type HttpConfig struct {
	Server struct {
		Http struct {
			Addr string
		}
	}
}

func GetHttpConfig() *HttpConfig {
	config.Load("app")
	var httpConfig HttpConfig
	err := viper.Unmarshal(&httpConfig)
	if err != nil {
		log.Panicf("unable to decode into struct, %v", err)
	}
	return &httpConfig
}
