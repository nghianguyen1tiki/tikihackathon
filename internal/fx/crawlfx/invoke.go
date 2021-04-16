package crawlfx

import (
	"context"
	"github.com/nghiant3223/tikihackathon/internal/crawl"
	"go.uber.org/fx"
)

func invokeCrawler(lc fx.Lifecycle, crawler *crawl.Crawler) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				return crawler.Start(ctx)
			},
			OnStop: func(ctx context.Context) error {
				return crawler.Stop(ctx)
			},
		},
	)
}
