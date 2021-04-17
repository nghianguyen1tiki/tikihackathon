package gracefulfx

import (
	"context"
	"errors"
	"github.com/nghiant3223/tikihackathon/internal/server"
	"github.com/nghiant3223/tikihackathon/pkg/log"
	"go.uber.org/fx"
	"net/http"
)

func invokeGraceful(lc fx.Lifecycle, server server.Server) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				ctx := context.Background()
				err := server.Start(ctx)
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Errorw("failed to listen", "error", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Stop(ctx)
		},
	})
}
