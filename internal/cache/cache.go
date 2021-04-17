package cache

import (
	"context"
	"fmt"
	"github.com/nghiant3223/tikihackathon/internal/config"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"sort"
)

type Cache interface {
	Warmup(ctx context.Context) error
	Cooldown(ctx context.Context) error
	GetAllRecipeIDs(ctx context.Context) ([]int, error)
	GetIngIDsByRecipeID(ctx context.Context, recipeID int) ([]int, error)
}

var _ Cache = (*memCache)(nil)

type memCache struct {
	db      *gorm.DB
	cfg     *config.Cache
	goCache *cache.Cache
}

func New(cfg *config.Cache, db *gorm.DB, goCache *cache.Cache) Cache {
	return &memCache{
		db:      db,
		cfg:     cfg,
		goCache: goCache,
	}
}

func (c *memCache) Cooldown(ctx context.Context) error {
	c.goCache.Flush()
	return nil
}

func (c *memCache) Warmup(ctx context.Context) error {
	recipeIDs, err := c.GetAllRecipeIDs(ctx)
	if err != nil {
		return fmt.Errorf("failed to warm up all recipes: %w", err)
	}

	for _, recipeID := range recipeIDs {
		_, err = c.GetIngIDsByRecipeID(ctx, recipeID)
		if err != nil {
			return fmt.Errorf("failed to warm up ingredients: %w", err)
		}
	}

	return nil
}

func (c *memCache) GetAllRecipeIDs(ctx context.Context) ([]int, error) {
	key := formatAllRecipeIDsKey()
	recipesID, ok := c.goCache.Get(key)
	if ok {
		return recipesID.([]int), nil
	}
	return c.cacheAllRecipeIDs(ctx)
}

func (c *memCache) GetIngIDsByRecipeID(ctx context.Context, recipeID int) ([]int, error) {
	key := formatIngIDsByRecipeIDKey(recipeID)
	ingredientIDs, ok := c.goCache.Get(key)
	if ok {
		return ingredientIDs.([]int), nil
	}
	return c.cacheIngredientIDsByRecipeID(ctx, recipeID)
}

func (c *memCache) cacheAllRecipeIDs(ctx context.Context) ([]int, error) {
	var recipeIDs []int
	err := c.db.WithContext(ctx).
		Table("recipes").
		Select("id").
		Scan(&recipeIDs).Error
	if err != nil {
		return nil, err
	}
	key := formatAllRecipeIDsKey()
	c.goCache.Set(key, recipeIDs, c.cfg.TTL)
	return recipeIDs, nil
}

func (c *memCache) cacheIngredientIDsByRecipeID(ctx context.Context, recipeID int) ([]int, error) {
	var ingredientIDs []int
	err := c.db.WithContext(ctx).
		Table("includings").
		Select("ingredient_id").
		Where("recipe_id = ?", recipeID).
		Scan(&ingredientIDs).Error
	if err != nil {
		return nil, err
	}
	sort.Ints(ingredientIDs)
	key := formatIngIDsByRecipeIDKey(recipeID)
	c.goCache.Set(key, ingredientIDs, c.cfg.TTL)
	return ingredientIDs, nil
}

func formatIngIDsByRecipeIDKey(recipeID int) string {
	return fmt.Sprintf("recipes:%d:ingredients", recipeID)
}

func formatAllRecipeIDsKey() string {
	return fmt.Sprintf("recipes")
}
