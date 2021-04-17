package crawlfx

import (
	"github.com/nghiant3223/tikihackathon/internal/config"
	"net/http"

	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/nghiant3223/tikihackathon/internal/crawl"
)

func provideCrawler(db *gorm.DB, httpClient *http.Client) (*crawl.Crawler, error) {
	var cfg *config.Crawl
	err := viper.UnmarshalKey("cache", &cfg)
	if err != nil {
		return nil, err
	}

	return crawl.NewCrawler(cfg, db, httpClient), nil
}
