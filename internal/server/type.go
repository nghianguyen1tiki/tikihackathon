package server

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Register(router gin.IRouter)
}

type Server interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}
