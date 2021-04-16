package server

import (
	"context"
	"errors"
	"github.com/nghiant3223/tikihackathon/pkg/log"
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
	addr := net.JoinHostPort("", config.Http.Port)
	newServer.httpServer = &http.Server{
		Addr:    addr,
		Handler: newServer.router,
	}
	return newServer
}

func (s *server) Start(ctx context.Context) error {
	err := s.httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Errorf("failed to listen on port %d", s.config.Http.Port)
		return err
	}
	return nil
}

func (s *server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
