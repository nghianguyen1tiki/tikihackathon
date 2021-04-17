package ingredient

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
	router.GET("/", h.listIngredients)
}

func (h *handler) listIngredients(c *gin.Context) {
	ingredientName := c.DefaultQuery("name", "")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	ingredient, err := h.repo.List(
		c,
		page,
		limit,
		map[string]interface{}{
			"name": ingredientName,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    ingredient,
	})
}
