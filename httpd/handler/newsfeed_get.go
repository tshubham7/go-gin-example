package handler

import (
	"net/http"
	"newsfeeder/platform/newsfeeder"

	"github.com/gin-gonic/gin"
)

// NewsfeedGet ...
func NewsfeedGet(feeds *newsfeeder.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, feeds.GetAll())
	}
}
