package feedfx

import (
	"github.com/gin-gonic/gin"
	"github.com/nghiant3223/tikihackathon/internal/feed"
)

func invokeFeedHandler(repo feed.Repo, router *gin.Engine) {
	handler := feed.NewHandler(repo)
	feedRouter := router.Group("/feed")
	handler.Register(feedRouter)
}
