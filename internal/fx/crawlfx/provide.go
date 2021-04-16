package crawlfx

import (
	"github.com/nghiant3223/tikihackathon/internal/crawl"
	"github.com/nghiant3223/tikihackathon/pkg/assert"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"net/http"
)

func provideCrawler(db *gorm.DB, httpClient *http.Client) *crawl.Crawler {
	target := viper.GetString("crawl.target")
	assert.NotEmpty(target, "crawl.target is empty")
	count := viper.GetInt("crawl.count")
	assert.NotZero(count, "crawl.count is zero")
	upperID := viper.GetInt("crawl.upperid")
	assert.NotZero(upperID, "crawl.upperid is zero")
	concurrency := viper.GetInt("crawl.concurrency")
	assert.NotZero(upperID, "crawl.concurrency is zero")

	return crawl.NewCrawler(
		db,
		httpClient,
		crawl.WithTarget(target),
		crawl.WithCount(count),
		crawl.WithUpperID(upperID),
		crawl.WithConcurrency(concurrency),
	)
}
