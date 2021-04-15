package crawlerfx

import (
	"context"
	"github.com/nghiant3223/tikihackathon/internal/crawler"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func initializeCrawler(lc fx.Lifecycle, db *gorm.DB) error {
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
