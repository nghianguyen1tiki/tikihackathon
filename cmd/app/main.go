package main

import (
	"github.com/nghiant3223/tikihackathon/internal/fx/cachefx"
	"github.com/nghiant3223/tikihackathon/internal/fx/configfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/dbfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/feedfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/ginfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/gracefulfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/httpfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/ingredientfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/logfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/recipefx"
	"github.com/nghiant3223/tikihackathon/internal/fx/serverfx"
	"go.uber.org/fx"
	"time"
)

const (
	configFilename = "crawler"
	startTimeout   = 30 * time.Second
)

func main() {
	fx.New(
		configfx.Invoke(configFilename),
		fx.StartTimeout(startTimeout),
		logfx.Invoke,
		dbfx.Provide,
		httpfx.Provide,
		ginfx.Provide,
		cachefx.Provide,
		recipefx.Provide,
		recipefx.Invoke,
		feedfx.Provide,
		feedfx.Invoke,
		ingredientfx.Provide,
		ingredientfx.Invoke,
		serverfx.Provide,
		gracefulfx.Invoke,
	).Run()
}
