package httpd

import (
	"github.com/gin-gonic/gin"
)

// PingGet handler for /ping
func PingGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
