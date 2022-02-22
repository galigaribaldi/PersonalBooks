package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/galigaribaldi/PersonalBooks/pkg/controller/comment"
	commentBookOperation "github.com/galigaribaldi/PersonalBooks/pkg/controller/comment"
	"github.com/galigaribaldi/PersonalBooks/pkg/models"
	"github.com/gin-gonic/gin"
)

func addCommentBookRoute(rg *gin.RouterGroup) {
	author := rg.Group("/commentBook")

	author.GET("/Search", getCommentBookAll)
	author.POST("/post", postCommentBook)

	author.DELETE("/delete", commentBookDelete)

}

func getCommentBookAll(c *gin.Context) {
	ids, err := strconv.Atoi(c.Query("idComment"))
	idBook, err := strconv.Atoi(c.Query("idBook"))
	comment, err := strconv.Atoi(c.Query("comment"))
	nameBook := c.Query("nameBook")
	if err != nil && idBook != 0 && comment != 0 {
		fmt.Println(ids, idBook, comment, nameBook)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	//ID Comment book
	if ids != 0 {
		c.JSON(http.StatusOK, commentBookOperation.SelectByIdCommentBook(ids))
		return
	}
	//ID Book
	if idBook != 0 || comment != 0 {
		Book, commentBook := commentBookOperation.SelectBookAndCommentByBookId(idBook)
		c.JSON(http.StatusOK, gin.H{
			"Book":        Book,
			"CommentBook": commentBook,
		})
		return
	}
	if idBook != 0 {
		c.JSON(http.StatusOK, commentBookOperation.SelectCommentsByBookId(idBook))
		return
	}
	//Name Book
	if nameBook != "" {
		commentBook, Book := commentBookOperation.SelectBookAndCommentByBookName(nameBook)
		c.JSON(http.StatusOK, gin.H{
			"Book":        Book,
			"CommentBook": commentBook,
		})
		return
	}
	c.JSON(http.StatusOK, commentBookOperation.SelectAllCommentBook())
	return

}

func commentBookDelete(c *gin.Context) {
	ids, err := strconv.Atoi(c.Query("id"))
	req := comment.DeleteCommentBook(ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, req)
	}
	c.JSON(http.StatusOK, req)

}
func postCommentBook(c *gin.Context) {
	var newCommentBook models.CommentBook
	if err := c.BindJSON(&newCommentBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	comment.CreateCommentBook(newCommentBook)
	c.JSON(http.StatusOK, newCommentBook)
}
