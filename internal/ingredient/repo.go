package ingredient

import (
	"context"
	"fmt"
	"gorm.io/gorm/clause"

	"gorm.io/gorm"

	"github.com/nghiant3223/tikihackathon/internal/model"
)

type Repo interface {
	List(ctx context.Context, page int, limit int, filter map[string]interface{}) ([]*model.Ingredient, error)
	ListByIDs(ctx context.Context, ids []int64) ([]*model.Ingredient, error)
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

func (r *repo) List(ctx context.Context, page int, limit int, filter map[string]interface{}) ([]*model.Ingredient, error) {
	var records []*model.Ingredient
	name := fmt.Sprintf("%v", filter["name"])

	offset := (page - 1) * limit
	err := r.db.WithContext(ctx).
		Preload(clause.Associations).
		Joins("INNER JOIN tiki_categories on ingredients.tiki_cate_id = tiki_categories.id").
		Where("ingredients.name LIKE ?", "%"+name+"%").
		Order("id DESC").
		Limit(limit).
		Offset(offset).
		Find(&records).Error
	return records, err
}

func (r *repo) ListByIDs(ctx context.Context, ids []int64) ([]*model.Ingredient, error) {
	var records []*model.Ingredient

	err := r.db.WithContext(ctx).
		Joins("INNER JOIN tiki_cat on ingredients.tiki_cate_id = tiki_cat.id").
		Where("ingredients.id IN ?", ids).
		Order("id DESC").Find(&records).Error
	return records, err
}
