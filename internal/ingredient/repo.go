package ingredient

import (
	"context"

	"gorm.io/gorm/clause"

	"gorm.io/gorm"

	"github.com/nghiant3223/tikihackathon/internal/model"
)

type Repo interface {
	Get(ctx context.Context, id int) (model.Ingredient, error)
	Create(ctx context.Context, ingredient *model.Ingredient) error
	Upsert(ctx context.Context, ingredient *model.Ingredient) error
	List(ctx context.Context, page int, limit int, filter map[string]interface{}) ([]*model.Ingredient, error)
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

func (r *repo) Get(ctx context.Context, id int) (model.Ingredient, error) {
	panic("implement me")
}

func (r *repo) Create(ctx context.Context, ingredient *model.Ingredient) error {
	panic("implement me")
}

func (r *repo) Upsert(ctx context.Context, ingredient *model.Ingredient) error {
	return r.db.
		WithContext(ctx).
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(ingredient).Error
}

func (r *repo) List(ctx context.Context, page int, limit int, filter map[string]interface{}) ([]*model.Ingredient, error) {
	var records []*model.Ingredient
	name, _ := filter["name"].(string)

	offset := (page - 1) * limit
	err := r.db.WithContext(ctx).Where("name LIKE ?", "%"+name+"%").Order("id DESC").Limit(limit).Offset(offset).Find(&records).Error
	return records, err
}
