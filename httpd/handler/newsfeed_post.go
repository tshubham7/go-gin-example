package handler

import (
	"net/http"
	"newsfeeder/platform/newsfeeder"

	"github.com/gin-gonic/gin"
)

// NewsfeedPost ...
func NewsfeedPost(feeds *newsfeeder.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := newsfeeder.Item{}
		c.Bind(&params)
		feeds.Add(params)
		c.JSON(http.StatusOK, gin.H{
			"message": "successfully created",
		})
	}
}
