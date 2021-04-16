package main

import (
	"github.com/nghiant3223/tikihackathon/internal/fx/httpfx"
	"go.uber.org/fx"

	"github.com/nghiant3223/tikihackathon/internal/fx/configfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/crawlfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/dbfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/logfx"
)

func main() {
	fx.New(
		configfx.Invoke("crawler"),
		logfx.Invoke,
		dbfx.Provide,
		httpfx.Provide,
		crawlfx.Provide,
		crawlfx.Invoke,
	).Run()
}
