package recipe

import (
	"context"

	"gorm.io/gorm/clause"

	"gorm.io/gorm"

	"github.com/nghiant3223/tikihackathon/internal/model"
)

type Repo interface {
	Get(ctx context.Context, id int) (model.Recipe, error)
	Create(ctx context.Context, recipe *model.Recipe) error
	Upsert(ctx context.Context, recipe *model.Recipe) error
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

func (r *repo) Get(ctx context.Context, id int) (model.Recipe, error) {
	panic("implement me")
}

func (r *repo) Create(ctx context.Context, recipe *model.Recipe) error {
	panic("implement me")
}

func (r *repo) Upsert(ctx context.Context, recipe *model.Recipe) error {
	return r.db.
		WithContext(ctx).
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(recipe).Error
}
