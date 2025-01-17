package route

import (
	"github.com/gin-gonic/gin"
	book_http "github.com/verlinof/golang-project-structure/internal/module/book/http"
	book_http_route "github.com/verlinof/golang-project-structure/internal/module/book/http/route"
	book_service "github.com/verlinof/golang-project-structure/internal/module/book/service"
)

func InitRoute(router *gin.Engine) {
	api := router.Group("/api/v1")

	//Book
	bookService := book_service.NewBookService()
	bookHandler := book_http.NewBookHandler(bookService)
	book_http_route.BookRoute(api, bookHandler)
}
