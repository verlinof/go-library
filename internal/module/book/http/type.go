package book_http

import (
	book_service "github.com/verlinof/golang-project-structure/internal/module/book/service"
	pkg_redis "github.com/verlinof/golang-project-structure/pkg/redis"
	pkg_validation "github.com/verlinof/golang-project-structure/pkg/validation"
)

type BookHandler struct {
	bookService  book_service.BookService
	redisManager pkg_redis.RedisManager
	xValidator   pkg_validation.XValidator
}

func NewBookHandler(bookService book_service.BookService, redisManager pkg_redis.RedisManager, xValidator pkg_validation.XValidator) BookHandler {
	return BookHandler{
		bookService:  bookService,
		redisManager: redisManager,
		xValidator:   xValidator,
	}
}
