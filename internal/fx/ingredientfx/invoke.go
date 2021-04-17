package ingredientfx

import (
	"github.com/gin-gonic/gin"

	"github.com/nghiant3223/tikihackathon/internal/ingredient"
)

func invokeIngredientHandler(repo ingredient.Repo, router *gin.Engine) {
	handler := ingredient.NewHandler(repo)
	recipeRouter := router.Group("/ingredients")
	handler.Register(recipeRouter)
}

