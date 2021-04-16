package configfx

import (
	"log"

	"go.uber.org/fx"

	"github.com/nghiant3223/tikihackathon/pkg/config"
)

func Invoke(configFilename string) fx.Option {
	return fx.Invoke(func() {
		config.Load(configFilename)
		log.Println("after load config")
	})
}
