package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getListRecipes() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	}
}
