package main

import (
	newsfeed "gin_example/cmd/data/news_feed"
	"gin_example/cmd/httpd"

	"github.com/gin-gonic/gin"
)

func main() {
	feed := newsfeed.New()
	feed.Add(newsfeed.Item{
		ID:    1,
		Title: "Hey",
		Post:  "You got this",
	})

	r := gin.Default()
	r.GET("/ping", httpd.PingGet)
	r.GET("/feed", httpd.FeedGet(feed))
	r.Run()
}
