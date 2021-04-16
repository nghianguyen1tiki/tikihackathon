package serverfx

import (
	"github.com/gin-gonic/gin"
	"github.com/nghiant3223/tikihackathon/internal/config"
	"github.com/nghiant3223/tikihackathon/internal/server"
	"github.com/spf13/viper"
)

func provideServer(router *gin.Engine) (server.Server, error) {
	var cfg *config.Server
	err := viper.UnmarshalKey("server", &cfg)
	if err != nil {
		return nil, err
	}
	return server.NewServer(cfg, router), nil
}
