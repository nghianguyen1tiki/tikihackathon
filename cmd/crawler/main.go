package main

import (
	"time"

	"go.uber.org/fx"

	"github.com/nghiant3223/tikihackathon/internal/fx/configfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/crawlfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/dbfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/httpfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/logfx"
)

const (
	configFilename = "crawler"
	startTimeout   = 5 * time.Minute
)

func main() {
	fx.New(
		configfx.Invoke(configFilename),
		fx.StartTimeout(startTimeout),
		logfx.Invoke,
		dbfx.Provide,
		httpfx.Provide,
		crawlfx.Provide,
		crawlfx.Invoke,
	).Run()
}
