package recipe

import (
	"context"

	"gorm.io/gorm"

	"github.com/nghiant3223/tikihackathon/internal/model"
)

type Repo interface {
<<<<<<< HEAD
	Get(ctx context.Context, id int) (model.Recipe, error)
	ListPopular(ctx context.Context, offset, limit *int) ([]model.Recipe, error)
=======
	Get(ctx context.Context, id int) (*model.Recipe, error)
	Create(ctx context.Context, recipe *model.Recipe) error
	Upsert(ctx context.Context, recipe *model.Recipe) error
>>>>>>> ddc19fc (Implement get recipe detail)
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

<<<<<<< HEAD
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
=======
func (r *repo) Get(ctx context.Context, id int) (*model.Recipe, error) {
	record := &model.Recipe{}
	err := r.db.WithContext(ctx).
		Preload("Ingredients.Ingredient.TikiCategory").
		Preload("Ingredients.Unit").
		Preload("Steps.StepPhotos").
		Preload(clause.Associations).
		Where("id = ?", id).First(record).Error
	if err != nil {
		return nil, err
	}

	return record, nil
}

func (r *repo) Create(ctx context.Context, recipe *model.Recipe) error {
	panic("implement me")
>>>>>>> ddc19fc (Implement get recipe detail)
}

func (r *repo) Get(ctx context.Context, id int) (model.Recipe, error) {
	var recipe model.Recipe
	err := r.db.WithContext(ctx).Preload("Photo").First(&recipe, id).Error
	if err != nil {
		return model.Recipe{}, err
	}
	return recipe, nil
}
