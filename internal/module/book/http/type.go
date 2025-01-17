package book_http

import book_service "github.com/verlinof/golang-project-structure/internal/module/book/service"

type BookHandler struct {
	bookService book_service.BookService
}

func NewBookHandler(bookService book_service.BookService) BookHandler {
	return BookHandler{bookService: bookService}
}
