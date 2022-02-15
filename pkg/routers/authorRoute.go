package routes

import (
	"net/http"
	"strconv"

	author "github.com/galigaribaldi/PersonalBooks/pkg/controller/author"
	"github.com/galigaribaldi/PersonalBooks/pkg/models"
	"github.com/gin-gonic/gin"
)

func addAuthorRoute(rg *gin.RouterGroup) {
	author := rg.Group("/author")
	author.GET("/All", getAllAuthor)

	author.GET("/Search", getAuthorById)
	author.POST("/post", authorPost)

	author.DELETE("/delete", deleteAuthor)

}

func getAllAuthor(c *gin.Context) {
	var newAuthor []models.Author
	newAuthor = author.SelectAllAuthor()
	c.JSON(http.StatusOK, newAuthor)
}

func getAuthorById(c *gin.Context) {
	ids, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	}
	c.JSON(http.StatusOK, author.SelectAuthorById(ids))
}

func authorPost(c *gin.Context) {
	var newAuthor models.Author
	if err := c.BindJSON(&newAuthor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	author.CreateAuthor(newAuthor)
	c.JSON(http.StatusOK, newAuthor)
}

func deleteAuthor(c *gin.Context) {
	ids, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	}
	author.DeleteAuthor(ids)
	c.JSON(http.StatusOK, "Author Delete")
}
