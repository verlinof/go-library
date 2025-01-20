package route

import (
	"github.com/gin-gonic/gin"
	"github.com/verlinof/golang-project-structure/configs/redis_config"
	book_http "github.com/verlinof/golang-project-structure/internal/module/book/http"
	book_http_route "github.com/verlinof/golang-project-structure/internal/module/book/http/route"
	book_service "github.com/verlinof/golang-project-structure/internal/module/book/service"
	pkg_redis "github.com/verlinof/golang-project-structure/pkg/redis"
)

func InitRoute(router *gin.Engine) {
	api := router.Group("/api/v1")

	//Dependencies
	redisManager := pkg_redis.NewRedisManager(redis_config.Config.Host, redis_config.Config.Password, redis_config.Config.Db)

	//Module Initialization
	//Book
	bookService := book_service.NewBookService()
	bookHandler := book_http.NewBookHandler(bookService, redisManager)
	book_http_route.BookRoute(api, bookHandler)
}
