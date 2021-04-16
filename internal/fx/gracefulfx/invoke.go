package gracefulfx

import (
	"context"
	"github.com/nghiant3223/tikihackathon/internal/server"
	"go.uber.org/fx"
	"os"
	"os/signal"
	"syscall"
)

func invokeGraceful(lc fx.Lifecycle, server server.Server) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			ctx, done, cancel := newStartContext()

			go func() {
				<-done
				cancel()
			}()

			go func() {
				server.Start(ctx)
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Stop(ctx)
		},
	})
}

func newStartContext() (context.Context, chan os.Signal, context.CancelFunc) {
	ctx := context.Background()
	done := make(chan os.Signal)
	ctx, cancel := context.WithCancel(ctx)
	signal.Notify(done, syscall.SIGINT, syscall.SIGKILL)
	return ctx, done, cancel
}