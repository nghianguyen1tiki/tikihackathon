package ingredientfx

import (
	"gorm.io/gorm"

	"github.com/nghiant3223/tikihackathon/internal/ingredient"
)

func provideRepo(db *gorm.DB) ingredient.Repo {
	return ingredient.NewRepo(db)
}
