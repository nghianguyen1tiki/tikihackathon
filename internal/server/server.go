package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nghiant3223/tikihackathon/internal/config"
)

type server struct {
	config     *config.Server
	router     *gin.Engine
	httpServer *http.Server
}

func NewServer(config *config.Server, router *gin.Engine) Server {
	newServer := &server{
		config: config,
		router: router,
	}
	newServer.httpServer = &http.Server{
		Handler: newServer.router,
		Addr:    net.JoinHostPort("", config.Http.Port),
	}
	newServer.initialize()
	return newServer
}

func (s *server) initialize() {
	s.addHealthCheckEndpoint()
}

func (s *server) Start(ctx context.Context) error {
	err := s.httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to listen on port %s: %w", s.config.Http.Port, err)
	}
	return nil
}

func (s *server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
