package main

import (
	"go.uber.org/fx"

	"github.com/nghiant3223/tikihackathon/internal/fx/configfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/crawlerfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/dbfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/logfx"
)

const (
	configFile = "crawler"
)

func main() {
	fx.New(
		configfx.Invoke(configFile),
		logfx.Invoke,
		dbfx.Provide,
		crawlerfx.Invoke,
	).Run()
}
