package routes

import (
	"net/http"
	"strconv"

	book "github.com/galigaribaldi/PersonalBooks/pkg/controller/book"
	"github.com/galigaribaldi/PersonalBooks/pkg/models"
	"github.com/gin-gonic/gin"
)

func addBookRoute(rg *gin.RouterGroup) {
	book := rg.Group("/book")
	book.GET("/Search", getBook)
	book.POST("/post", bookPost)

	book.DELETE("/delete", deleteBook)

}

//Funciones a usar para las vistas
// Creación de un nuevo libro
func bookPost(c *gin.Context) {
	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.CreateBook(newBook)
	c.JSON(http.StatusOK, newBook)
}

func deleteBook(c *gin.Context) {
	ids, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	}
	book.DeleteBook(ids)
	c.JSON(http.StatusOK, "Book Delete")
}

func getBook(c *gin.Context) {
	//fmt.Println(c.PostForm("id")) //id FormData
	nameBook := c.Query("nameBook")
	ids, err := strconv.Atoi(c.Query("id"))
	if err != nil && ids != 0 {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	//ID
	if ids != 0 {
		c.JSON(http.StatusOK, book.SelectBooksById(ids))
		return
	}
	//Name
	if nameBook != "" {
		c.JSON(http.StatusOK, book.SelectBookByName(nameBook))
		return
	}
	//All
	c.JSON(http.StatusOK, book.SelectAllBooks())
	return
}
