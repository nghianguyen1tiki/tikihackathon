package main

import (
	"github.com/nghiant3223/tikihackathon/internal/fx/configfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/crawlfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/dbfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/httpfx"
	"github.com/nghiant3223/tikihackathon/internal/fx/logfx"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		configfx.Invoke("crawler"),
		fx.StartTimeout(viper.GetDuration("crawl.timeout")),
		logfx.Invoke,
		dbfx.Provide,
		httpfx.Provide,
		crawlfx.Provide,
		crawlfx.Invoke,
	).Run()
}
