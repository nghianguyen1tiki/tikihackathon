package feed

import (
	"context"
	"sort"

	"github.com/nghiant3223/tikihackathon/internal/cache"
	"github.com/nghiant3223/tikihackathon/internal/config"
	"github.com/nghiant3223/tikihackathon/internal/model"
	"github.com/nghiant3223/tikihackathon/internal/recipe"
	"github.com/nghiant3223/tikihackathon/pkg/math"
)

type Repo interface {
	GetPopularRecipes(ctx context.Context, offset *int, limit *int) ([]model.Recipe, error)
	GetPersonalizedRecipes(
		ctx context.Context,
		blacklistIngredientIDs,
		whitelistIngredientIDs,
		pantryIngredientIDs []int,
		offset, limit *int,
	) ([]model.Recipe, error)
}

var _ Repo = (*repo)(nil)

type repo struct {
	cfg        *config.Feed
	cache      cache.Cache
	recipeRepo recipe.Repo
}

func NewRepo(cfg *config.Feed, cache cache.Cache, recipeRepo recipe.Repo) Repo {
	return &repo{
		cfg:        cfg,
		cache:      cache,
		recipeRepo: recipeRepo,
	}
}

func (r *repo) GetPopularRecipes(ctx context.Context, offset *int, limit *int) ([]model.Recipe, error) {
	return r.recipeRepo.ListPopular(ctx, offset, limit)
}

func (r *repo) GetPersonalizedRecipes(
	ctx context.Context,
	blacklistIngredientIDs,
	whitelistIngredientIDs,
	pantryIngredientIDs []int,
	offset, limit *int,
) ([]model.Recipe, error) {
	recipeIDs, err := r.cache.GetAllRecipeIDs(ctx)
	if err != nil {
		return nil, err
	}

	sort.Ints(blacklistIngredientIDs)
	sort.Ints(whitelistIngredientIDs)
	sort.Ints(pantryIngredientIDs)

	recipeAndScoreList := make([][2]int, 0, len(recipeIDs))
	for _, recipeID := range recipeIDs {
		ingIDs, err := r.cache.GetIngIDsByRecipeID(ctx, recipeID)
		if err != nil {
			return nil, err
		}
		if isBlacklisted(blacklistIngredientIDs, ingIDs) {
			continue
		}
		whitelistScore := calculateScore(whitelistIngredientIDs, ingIDs)
		pantryScore := calculateScore(pantryIngredientIDs, ingIDs)
		score := whitelistScore + pantryScore
		recipeIDAndScore := [2]int{recipeID, score}
		recipeAndScoreList = append(recipeAndScoreList, recipeIDAndScore)
	}

	sort.Slice(recipeAndScoreList, func(i, j int) bool {
		scoreI := recipeAndScoreList[i][1]
		scoreJ := recipeAndScoreList[j][1]
		return scoreI < scoreJ
	})

	actualLimit := r.cfg.DefaultLimit
	actualOffset := r.cfg.DefaultOffset
	if limit != nil {
		actualLimit = *limit
	}
	if offset != nil {
		actualOffset = *offset
	}
	if actualOffset >= len(recipeAndScoreList) {
		return []model.Recipe{}, nil
	}
	actualUpper := math.MinInt(actualOffset+actualLimit+1, len(recipeAndScoreList))
	recipeAndScoreList = recipeAndScoreList[actualOffset:actualUpper]

	recipes := make([]model.Recipe, len(recipeAndScoreList))
	for i, idScorePair := range recipeAndScoreList {
		id := idScorePair[0]
		recipe, err := r.recipeRepo.Get(ctx, id)
		if err != nil {
			return nil, err
		}
		recipes[i] = *recipe
	}
	return recipes, nil
}
