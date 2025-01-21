package book_http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	book_model "github.com/verlinof/golang-project-structure/internal/module/book/model"
	pkg_error "github.com/verlinof/golang-project-structure/pkg/error"
	pkg_success "github.com/verlinof/golang-project-structure/pkg/success"
	"gorm.io/gorm"
)

func (bookHandler *BookHandler) GetAllBook(c *gin.Context) {
	//Pagination
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	perPageStr := c.Query("per_page")
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		perPage = 10
	}

	books, err := bookHandler.bookService.GetAllBook(c.Request.Context(), page, perPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg_error.NewInternalServerError(err))
		return
	}

	c.JSON(http.StatusOK, books)
}

func (bookHandler *BookHandler) GetBookByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Jika konversi gagal, kembalikan error ke klien
		c.JSON(http.StatusBadRequest, pkg_error.NewBadRequest(fmt.Errorf("invalid ID: %s", idStr)))
		return
	}

	//Error Handling
	book, err := bookHandler.bookService.GetBookByID(c.Request.Context(), id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, pkg_error.NewNotFound(fmt.Errorf("book with ID %d not found", id)))
			return
		}

		c.JSON(http.StatusInternalServerError, pkg_error.NewInternalServerError(err))
		return
	}

	c.JSON(http.StatusOK, pkg_success.SuccessGetData(book))
}

func (bookHandler *BookHandler) CreateBook(c *gin.Context) {
	var createBookRequest book_model.CreateBookRequest
	if err := c.ShouldBindJSON(&createBookRequest); err != nil {
		c.JSON(http.StatusBadRequest, pkg_error.NewBadRequest(err))
		return
	}

	//Validate the struct
	err := bookHandler.xValidator.Validate(createBookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg_error.NewBadRequest(err))
		return
	}

	//Error Handling
	book, err := bookHandler.bookService.CreateBook(c.Request.Context(), createBookRequest)
	if err != nil {
		//Err Duplicated Unique Key
		if err == gorm.ErrDuplicatedKey {
			c.JSON(http.StatusBadRequest, pkg_error.NewBadRequest(err))
			return
		}
		c.JSON(http.StatusInternalServerError, pkg_error.NewInternalServerError(err))
		return
	}

	c.JSON(http.StatusCreated, pkg_success.SuccessCreateData(book))
}

func (bookHandler *BookHandler) UpdateBook(c *gin.Context) {
	var updateBookRequest book_model.UpdateBookRequest
	if err := c.ShouldBindJSON(&updateBookRequest); err != nil {
		c.JSON(http.StatusBadRequest, pkg_error.NewBadRequest(err))
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Jika konversi gagal, kembalikan error ke klien
		c.JSON(http.StatusBadRequest, pkg_error.NewBadRequest(fmt.Errorf("invalid ID: %s", idStr)))
		return
	}

	// Validate struct
	err = bookHandler.xValidator.Validate(updateBookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg_error.NewBadRequest(err))
		return
	}

	// error handling
	book, err := bookHandler.bookService.UpdateBook(c.Request.Context(), id, updateBookRequest)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, pkg_error.NewNotFound(fmt.Errorf("book with ID %d not found", id)))
			return
		} else if err == gorm.ErrDuplicatedKey {
			c.JSON(http.StatusBadRequest, pkg_error.NewBadRequest(err))
			return
		}

		c.JSON(http.StatusInternalServerError, pkg_error.NewInternalServerError(err))
		return
	}

	c.JSON(http.StatusOK, pkg_success.SuccessGetData(book))
}

func (bookHandler *BookHandler) DeleteBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Jika konversi gagal, kembalikan error ke klien
		c.JSON(http.StatusBadRequest, pkg_error.NewBadRequest(fmt.Errorf("invalid ID: %s", idStr)))
		return
	}

	err = bookHandler.bookService.DeleteBook(c.Request.Context(), id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, pkg_error.NewNotFound(fmt.Errorf("book with ID %d not found", id)))
			return
		}

		c.JSON(http.StatusInternalServerError, pkg_error.NewInternalServerError(err))
		return
	}

	c.JSON(http.StatusOK, pkg_success.SuccessDeleteData(id))
}
