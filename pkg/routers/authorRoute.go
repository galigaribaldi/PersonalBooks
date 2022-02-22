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

	author.GET("/Search", getAuthor)
	author.POST("/post", authorPost)

	author.DELETE("/delete", deleteAuthor)

}

func getAuthor(c *gin.Context) {
	nameAuthor := c.Query("nameAuthor")
	ids, err := strconv.Atoi(c.Query("id"))
	if err != nil && ids != 0 {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	//ID
	if ids != 0 {
		c.JSON(http.StatusOK, author.SelectAuthorById(ids))
		return
	}
	//Name
	if nameAuthor != "" {
		c.JSON(http.StatusOK, author.SelectAuthorByName(nameAuthor))
		return
	}
	//All
	c.JSON(http.StatusOK, author.SelectAllAuthor())
	return
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
