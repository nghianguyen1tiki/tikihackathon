package user

import (
	"context"

	"gorm.io/gorm"

	"github.com/nghiant3223/tikihackathon/internal/model"
)

type Repo interface {
	Get(ctx context.Context, id int) (model.User, error)
	Create(ctx context.Context, ingredient *model.User) error
}

type repo struct {
	db *gorm.DB
}
