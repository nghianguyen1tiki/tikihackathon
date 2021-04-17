package main

import (
	"time"

	"go.uber.org/fx"

	"github.com/nghiant3223/tikihackathon/internal/fx/configfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/dbfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/ginfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/gracefulfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/httpfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/ingredientfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/logfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/recipefx"
	"github.com/nghiant3223/tikihackathon/internal/fx/serverfx"
)

func main() {
	fx.New(
		configfx.Invoke("app"),
		fx.StartTimeout(15 * time.Hour),
		logfx.Invoke,
		dbfx.Provide,
		httpfx.Provide,
		ginfx.Provide,
		recipefx.Provide,
		recipefx.Invoke,
		ingredientfx.Provide,
		ingredientfx.Invoke,
		serverfx.Provide,
		gracefulfx.Invoke,
	).Run()
}
