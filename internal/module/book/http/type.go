package book_http

import (
	book_service "github.com/verlinof/golang-project-structure/internal/module/book/service"
	pkg_redis "github.com/verlinof/golang-project-structure/pkg/redis"
)

type BookHandler struct {
	bookService  book_service.BookService
	redisManager pkg_redis.RedisManager
}

func NewBookHandler(bookService book_service.BookService, redisManager pkg_redis.RedisManager) BookHandler {
	return BookHandler{
		bookService:  bookService,
		redisManager: redisManager,
	}
}
