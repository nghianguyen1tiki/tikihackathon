package recipe

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/nghiant3223/tikihackathon/internal/server"
)

type handler struct {
	repo Repo
}

var _ server.Handler = (*handler)(nil)

func NewHandler(repo Repo) server.Handler {
	return &handler{
		repo: repo,
	}
}

func (h *handler) Register(router gin.IRouter) {
	router.GET("/", h.listRecipes)
	router.GET("/:id", h.getRecipeByID)
}

func (h *handler) listRecipes(c *gin.Context) {

}

func (h *handler) getRecipeByID(c *gin.Context) {
	recipeIDStr := c.Param("id")

	recipeID, err := strconv.Atoi(recipeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	recipe, err := h.repo.Get(c, recipeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    recipe,
	})
}
