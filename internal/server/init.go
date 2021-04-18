package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	healthCheckEndpoint = "/health"
	healthCheckMessage  = "ok"
)

func (s *server) addHealthCheckEndpoint() {
	s.router.GET(healthCheckEndpoint, func(c *gin.Context) {
		c.String(http.StatusOK, healthCheckMessage)
	})
}
