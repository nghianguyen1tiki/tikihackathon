package recipefx

import (
	"github.com/nghiant3223/tikihackathon/internal/recipe"
	"gorm.io/gorm"
)

func provideRepo(db *gorm.DB) recipe.Repo {
	return recipe.NewRepo(db)
}
