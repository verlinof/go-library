package book_http_route

import (
	"github.com/gin-gonic/gin"
	book_http "github.com/verlinof/golang-project-structure/internal/module/book/http"
)

func BookRoute(router *gin.RouterGroup, bookHandler book_http.BookHandler) {
	bookRoutes := router.Group("/books")
	bookRoutes.GET("/", bookHandler.GetAllBook)
	bookRoutes.GET("/:id", bookHandler.GetBookByID)
	bookRoutes.POST("/", bookHandler.CreateBook)
	bookRoutes.PATCH("/:id", bookHandler.UpdateBook)
	bookRoutes.DELETE("/:id", bookHandler.DeleteBook)
}
