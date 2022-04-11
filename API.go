package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.1.2"})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func TestPingRoute(t *testing.T) {

	w := httptest.NewRecorder()

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

