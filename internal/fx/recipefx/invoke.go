package recipefx

import (
	"github.com/gin-gonic/gin"
	"github.com/nghiant3223/tikihackathon/internal/recipe"
)

func invokeRecipeHandler(repo recipe.Repo, router *gin.Engine) {
	handler := recipe.NewHandler(repo)
	recipeRouter := router.Group("/recipes")
	handler.Register(recipeRouter)
}
