package crawlfx

import (
	"context"

	"go.uber.org/fx"

	"github.com/nghiant3223/tikihackathon/internal/crawl"
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
