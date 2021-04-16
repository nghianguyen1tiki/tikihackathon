package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nghiant3223/tikihackathon/internal/configs"
	"github.com/nghiant3223/tikihackathon/internal/handlers"
)

type Server struct {
	config     *configs.HttpConfig
	router     *gin.Engine
	httpServer *http.Server
}

func NewServer(config *configs.HttpConfig) *Server {
	// init routers
	router := gin.New()
	router.Use(gin.Recovery())

	s := &Server{
		router: router,
		config: config,
	}
	s.initRouters()

	return s
}

func (s *Server) initRouters() {
	apiHandler := handlers.NewHandler(s.router)
	apiHandler.ConfigRouter()
}

func (s *Server) Start() {
	s.httpServer = &http.Server{
		Handler: s.router,
		Addr:    s.config.Server.Http.Addr,
	}
	_ = s.httpServer.ListenAndServe()
}

func (s *Server) Stop() {
	_ = s.httpServer.Close()
}
