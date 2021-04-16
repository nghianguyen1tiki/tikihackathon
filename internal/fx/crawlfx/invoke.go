package crawlfx

import (
	"context"
	"github.com/nghiant3223/tikihackathon/pkg/assert"
	"github.com/spf13/viper"
	"net/http"

	"go.uber.org/fx"
	"gorm.io/gorm"

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

func provideCrawler(db *gorm.DB, httpClient *http.Client) *crawl.Crawler {
	target := viper.GetString("crawl.target")
	assert.NotEmpty(target, "crawl.target is empty")
	count := viper.GetInt("crawl.count")
	assert.NotZero(count, "crawl.count is zero")
	upperID := viper.GetInt("crawl.upperid")
	assert.NotZero(upperID, "crawl.upperid is zero")

	return crawl.NewCrawler(
		db,
		httpClient,
		crawl.WithTarget(target),
		crawl.WithCount(count),
		crawl.WithUpperID(upperID),
	)
}
