package gracefulfx

import (
	"context"
	"github.com/nghiant3223/tikihackathon/internal/server"
	"go.uber.org/fx"
)

func invokeGraceful(lc fx.Lifecycle, server server.Server) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			ctx := context.Background()
			return server.Start(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return server.Stop(ctx)
		},
	})
}
