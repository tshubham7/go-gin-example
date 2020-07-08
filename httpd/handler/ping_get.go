package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingGet ...
func PingGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}
