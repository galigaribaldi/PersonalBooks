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
	v1 := router.Group("/v1")
	addBookRoute(v1)
}
