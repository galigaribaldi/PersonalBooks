package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addBookRoute(rg *gin.RouterGroup) {
	book := rg.Group("/book")
	book.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Book Get")
	})

	book.GET("/:param", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data":    c.Params.ByName("param"),
			"message": "Correct",
		})
	})
	book.POST("/post", func(c *gin.Context) {
		c.JSON(http.StatusOK, "POSTITO")
	})

	book.DELETE("/delete", func(c *gin.Context) {
		c.JSON(http.StatusOK, "POST")
	})

	book.PUT("/put", func(c *gin.Context) {
		c.JSON(http.StatusOK, "POST")
	})
}
