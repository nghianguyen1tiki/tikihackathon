package crawlerfx

import (
	"context"

	"go.uber.org/fx"
	"gorm.io/gorm"

	"github.com/nghiant3223/tikihackathon/internal/crawler"
)

func invokeCrawler(lc fx.Lifecycle, db *gorm.DB) error {
	craw := crawler.NewCrawler(db)
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				return craw.Start(ctx)
			},
			OnStop: func(ctx context.Context) error {
				return craw.Stop(ctx)
			},
		},
	)
	return nil
}
