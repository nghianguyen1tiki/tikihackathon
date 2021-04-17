package feed

import (
	"github.com/gin-gonic/gin"
	"github.com/nghiant3223/tikihackathon/internal/dto"
	"github.com/nghiant3223/tikihackathon/internal/server"
	"net/http"
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
	router.GET("", h.getFeed)
}

func (h *handler) getFeed(c *gin.Context) {
	var query dto.GetFeedQuery
	err := c.ShouldBindQuery(&query)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	limit := query.Size
	offset := query.Page * query.Size
	popular, err := h.repo.GetPopularRecipes(c, &offset, &limit)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	personal, err := h.repo.GetPersonalizedRecipes(c, query.Blacklist, query.Whitelist, query.Pantry, &offset, &limit)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	resp := dto.GetFeedResponse{
		PopularRecipes:      popular,
		PersonalizedRecipes: personal,
	}
	c.JSON(http.StatusOK, resp)
}
