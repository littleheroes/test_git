package main

import (
	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello www.topgoer.com",
	})
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/topgoer", helloHandler)
	return r
}
