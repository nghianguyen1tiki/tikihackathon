package feed

import (
	"github.com/gin-gonic/gin"
	"github.com/nghiant3223/tikihackathon/internal/dto"
	"github.com/nghiant3223/tikihackathon/internal/server"
	"github.com/nghiant3223/tikihackathon/pkg/lang"
	"github.com/nghiant3223/tikihackathon/pkg/log"
	"github.com/spf13/cast"
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
		log.Errorw("cannot bind query", "error", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	limit := cast.ToInt(query.Size)
	offset := cast.ToInt(query.Page) * cast.ToInt(query.Size)
	popular, err := h.repo.GetPopularRecipes(c, &offset, &limit)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	blacklist := lang.StringSliceToIntSlice(query.Blacklist)
	whitelist := lang.StringSliceToIntSlice(query.Whitelist)
	pantry := lang.StringSliceToIntSlice(query.Pantry)
	personal, err := h.repo.GetPersonalizedRecipes(c, blacklist, whitelist, pantry, &offset, &limit)
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
