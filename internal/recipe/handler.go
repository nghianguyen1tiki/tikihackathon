package recipe

import (
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
}

func (h *handler) listRecipes(c *gin.Context) {

}
