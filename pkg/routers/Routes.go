package routes

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func Run() {
	getRouters()
	router.Run(":5000")
}

func getRouters() {
	// Group 1 Books
	v1 := router.Group("/v1")
	addBookRoute(v1)
	// Group 2 Author
	v2 := router.Group("/v2")
	addAuthorRoute(v2)
	// Group 3 CommentBook
	v3 := router.Group("/v3")
	addCommentBookRoute(v3)
}
