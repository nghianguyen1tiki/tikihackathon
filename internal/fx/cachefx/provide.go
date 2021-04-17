package cachefx

import (
	"context"
	"github.com/nghiant3223/tikihackathon/internal/cache"
	"github.com/nghiant3223/tikihackathon/internal/config"
	gocache "github.com/patrickmn/go-cache"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func provideCache(lc fx.Lifecycle, db *gorm.DB) (cache.Cache, error) {
	var cfg *config.Cache
	err := viper.UnmarshalKey("cache", &cfg)
	if err != nil {
		return nil, err
	}

	innerCache := gocache.New(cfg.TTL, cfg.CleanupInterval)
	cacher := cache.New(cfg, db, innerCache)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return cacher.Warmup(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return cacher.Cooldown(ctx)
		},
	})

	return cacher, nil
}
