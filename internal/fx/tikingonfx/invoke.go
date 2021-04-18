package tikingonfx

import (
	"github.com/gin-gonic/gin"

	"github.com/nghiant3223/tikihackathon/internal/tikingon"
)

func invokeTikiNgonHandler(router *gin.Engine) {
	handler := tikingon.NewHandler()
	tikiNgonRouter := router.Group("/tiki-ngon")
	handler.Register(tikiNgonRouter)
}
