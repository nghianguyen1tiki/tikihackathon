package ginfx

import "github.com/gin-gonic/gin"

func provideGin() *gin.Engine {
	engine := gin.Default()
	engine.Use(gin.Recovery())
	return engine
}
