package main

import (
	"github.com/nghiant3223/tikihackathon/internal/fx/configfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/dbfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/ginfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/gracefulfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/httpfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/logfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/recipefx"
	"github.com/nghiant3223/tikihackathon/internal/fx/serverfx"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		configfx.Invoke("app"),
		logfx.Invoke,
		dbfx.Provide,
		httpfx.Provide,
		ginfx.Provide,
		recipefx.Provide,
		recipefx.Invoke,
		serverfx.Provide,
		gracefulfx.Invoke,
	).Run()
}
