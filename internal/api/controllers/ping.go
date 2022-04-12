package controllers

import "github.com/gin-gonic/gin"

type PingAPI struct{}

func (PingAPI) Get(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
