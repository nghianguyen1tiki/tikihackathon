package configfx

import (
	"go.uber.org/fx"

	"github.com/nghiant3223/tikihackathon/pkg/config"
)

func Invoke(name string) fx.Option {
	return fx.Invoke(func() {
		config.Load(name)
	})
}
