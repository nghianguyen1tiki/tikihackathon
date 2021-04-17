package recipe

import (
	"context"
	"gorm.io/gorm/clause"

	"gorm.io/gorm"

	"github.com/nghiant3223/tikihackathon/internal/model"
)

type Repo interface {
	ListPopular(ctx context.Context, offset, limit *int) ([]model.Recipe, error)
	Get(ctx context.Context, id int) (*model.Recipe, error)
}

var _ Repo = (*repo)(nil)

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{
		db: db,
	}
}

func (r *repo) ListPopular(ctx context.Context, offset, limit *int) ([]model.Recipe, error) {
	var recipes []model.Recipe
	db := r.db.WithContext(ctx).Order("total_view DESC")
	if limit != nil {
		db = db.Limit(*limit)
	}
	if offset != nil {
		db = db.Offset(*offset)
	}
	err := db.Find(&recipes).Error
	if err != nil {
		return nil, err
	}
	return recipes, nil
}

func (r *repo) Get(ctx context.Context, id int) (*model.Recipe, error) {
	record := &model.Recipe{}
	err := r.db.WithContext(ctx).
		Preload("Ingredients.Ingredient.TikiCategory").
		Preload("Ingredients.Unit").
		Preload("Steps.StepPhotos").
		Preload(clause.Associations).
		Where("id = ?", id).
		First(record).Error
	if err != nil {
		return nil, err
	}

	return record, nil
}
