package feedfx

import (
	"github.com/nghiant3223/tikihackathon/internal/cache"
	"github.com/nghiant3223/tikihackathon/internal/config"
	"github.com/nghiant3223/tikihackathon/internal/feed"
	"github.com/nghiant3223/tikihackathon/internal/recipe"
	"github.com/spf13/viper"
)

func provideRepo(cache cache.Cache, recipeRepo recipe.Repo) (feed.Repo, error) {
	var cfg *config.Feed
	err := viper.UnmarshalKey("feed", &cfg)
	if err != nil {
		return nil, err
	}

	return feed.NewRepo(cfg, cache, recipeRepo), nil
}
