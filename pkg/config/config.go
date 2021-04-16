package config

import (
	"strings"

	"github.com/spf13/viper"

	"github.com/nghiant3223/tikihackathon/pkg/log"
)

func Load(name string) {
	envReplacer := strings.NewReplacer("_", ".")
	viper.SetEnvKeyReplacer(envReplacer)
	viper.AutomaticEnv()
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Errorw("cannot read in config, all configs are fallen back to env", "error", err)
	}
}
