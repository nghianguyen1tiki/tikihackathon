package ingredient

import (
	"context"

	"gorm.io/gorm"

	"github.com/nghiant3223/tikihackathon/internal/model"
)

type Repo interface {
	Get(ctx context.Context, id int) (model.Ingredient, error)
	Create(ctx context.Context, ingredient *model.Ingredient) error
}

type repo struct {
	db *gorm.DB
}
