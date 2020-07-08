package main

import (
	"newsfeeder/httpd/handler"
	"newsfeeder/platform/newsfeeder"

	"github.com/gin-gonic/gin"
)

func main() {
	feeds := newsfeeder.New()
	r := gin.Default()
	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(feeds))
	r.POST("/newsfeed", handler.NewsfeedPost(feeds))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
