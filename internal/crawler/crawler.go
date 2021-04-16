package crawler

import (
	"context"

	"gorm.io/gorm"
)

type Crawler struct {
	db *gorm.DB
}

func NewCrawler(db *gorm.DB) *Crawler {
	return &Crawler{db: db}
}

func (c *Crawler) Start(ctx context.Context) error {
	return nil
}

func (c *Crawler) Stop(ctx context.Context) error {
	return nil
}
