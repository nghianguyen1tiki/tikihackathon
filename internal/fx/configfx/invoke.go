package configfx

import (
	"github.com/nghiant3223/tikihackathon/internal/config"
	"go.uber.org/fx"
)

func Invoke(configFilename string) fx.Option {
	return fx.Invoke(func() {
		config.Load(configFilename)
	})
}
