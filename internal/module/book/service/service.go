package book_service

import (
	"context"

	"github.com/verlinof/golang-project-structure/db"
	book_model "github.com/verlinof/golang-project-structure/internal/module/book/model"
)

func (b *BookService) GetAllBook(ctx context.Context) ([]book_model.BookResponse, error) {
	var books []book_model.BookResponse
	err := db.DB.WithContext(ctx).Table("books").Find(&books).Error
	if err != nil {
		return []book_model.BookResponse{}, err
	}

	return books, nil
}

func (b *BookService) GetBookByID(ctx context.Context, id int) (book_model.BookResponse, error) {
	var book book_model.BookResponse
	err := db.DB.WithContext(ctx).Table("books").Where("id = ?", id).First(&book).Error
	if err != nil {
		return book_model.BookResponse{}, err
	}

	return book, nil
}

func (b *BookService) CreateBook(ctx context.Context, createBookRequest book_model.CreateBookRequest) (book_model.BookResponse, error) {
	var book *book_model.Book
	var bookResponse book_model.BookResponse

	book = &book_model.Book{
		Title:       createBookRequest.Title,
		Description: createBookRequest.Description,
	}

	err := db.DB.WithContext(ctx).Table("books").Create(&book).Error
	if err != nil {
		return book_model.BookResponse{}, err
	}

	bookResponse = book_model.BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
	}

	return bookResponse, nil
}

func (b *BookService) UpdateBook(ctx context.Context, id int, updateBookRequest book_model.UpdateBookRequest) (book_model.BookResponse, error) {
	return book_model.BookResponse{}, nil
}
