package crawl

import (
	"context"
	"net/http"

	"gorm.io/gorm"
)

type Crawler struct {
	db         *gorm.DB
	cfg        *config
	httpClient *http.Client
}

func NewCrawler(db *gorm.DB, httpClient *http.Client, configFns ...configFn) *Crawler {
	cfg := &config{}
	for _, fn := range configFns {
		fn(cfg)
	}
	return &Crawler{
		db:         db,
		cfg:        cfg,
		httpClient: httpClient,
	}
}

func (c *Crawler) Start(ctx context.Context) error {
	return nil
}

func (c *Crawler) Stop(ctx context.Context) error {
	return nil
}
