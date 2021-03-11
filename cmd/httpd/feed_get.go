package httpd

import (
	newsfeed "gin_example/cmd/data/news_feed"

	"github.com/gin-gonic/gin"
)

// FeedGet - returns the entire feed
func FeedGet(feed *newsfeed.List) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, feed.GetAll())
	}
}
