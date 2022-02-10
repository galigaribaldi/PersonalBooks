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
	book.GET("/All", getAllBook)

	book.GET("/Search", getBookById)
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
	c.JSON(http.StatusOK, gin.H{
		"id":     newBook.Id,
		"title":  newBook.Title,
		"author": newBook.Author,
		"Age":    newBook.Age,
	})
}

func getAllBook(c *gin.Context) {
	var newBook []models.Book
	newBook = book.SelectAllBooks()
	c.JSON(http.StatusOK, newBook)
}

func getBookById(c *gin.Context) {
	//fmt.Println(c.PostForm("id")) //id FormData
	ids, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
	}
	c.JSON(http.StatusOK, book.SelectBooksById(ids))
}

func deleteBook(c *gin.Context) {
	c.JSON(http.StatusOK, "POST")
}
