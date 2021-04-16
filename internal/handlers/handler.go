package handlers

import "github.com/gin-gonic/gin"

type Handler struct {
	router *gin.Engine
}

func NewHandler(r *gin.Engine) *Handler {
	return &Handler{router: r}
}

func (h *Handler) ConfigRouter() {
	apiGroup := h.router.Group("api")

	recipeHandler := apiGroup.Group("recipes")
	recipeHandler.GET("/", getListRecipes())
}
