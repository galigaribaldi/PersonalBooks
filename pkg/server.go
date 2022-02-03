package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	server := gin.Default()

	server.GET("/:param", func(c *gin.Context) {
		param := c.Params.ByName("param")
		c.JSON(http.StatusOK, gin.H{
			"data":    param,
			"message": "Correct",
		})
	})

	server.POST("/post", func(c *gin.Context) {
		c.JSON(http.StatusOK, "POSTITO")
	})

	server.DELETE("/delete", func(c *gin.Context) {
		c.JSON(http.StatusOK, "POST")
	})

	server.PUT("/put", func(c *gin.Context) {
		c.JSON(http.StatusOK, "POST")
	})

	server.Run(":8080")
}
